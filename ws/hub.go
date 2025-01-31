package ws

import (
	"acedrex/game"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var rooms map[string]*Room

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Player struct {
	Nickname string
	Conn     *websocket.Conn
}

type Room struct {
	RoomId string    `json:"roomId"`
	Game   game.Game `json:"-"`
	White  Player    `json:"-"`
	Black  Player    `json:"-"`
}

func InitializeRooms() {
	rooms = map[string]*Room{}
}

func generateHashID() string {
	// Combine random number and timestamp for uniqueness
	rand.New(rand.NewSource(time.Now().Unix()))
	randomComponent := rand.Intn(1000)
	timestamp := time.Now().UnixNano()

	// Concatenate the timestamp and random component into a string
	data := fmt.Sprintf("%d-%d", timestamp, randomComponent)

	// Generate a MD5 hash and return the first 8 characters
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", hash)[:8]
}

func GetRoomsList(w http.ResponseWriter, r *http.Request) {
	idsSlice := []string{}
	for key := range rooms {
		idsSlice = append(idsSlice, key)
	}
	jsonData, err := json.Marshal(idsSlice)
	if err != nil {
		log.Fatalln("Failed at attempting marhaling rooms.")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonData)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {

	// generate random Id
	uuid := generateHashID()
	rooms[uuid] = &Room{
		RoomId: uuid,
		Game:   *game.InitilaizeGame(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("{\"roomId\":\"" + uuid + "\"}"))
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	// get room id from query params
	roomId := chi.URLParam(r, "roomId")
	room := rooms[roomId]
	// TODO: check if room exists

	// TODO: check if side exists (do we have to check if the query param exists?)
	// TODO: query params?
	// side = r.URL.Query().Get("side")
	var side string
	if room.Black.Conn == nil && room.White.Conn == nil {
		// no player is present
		side = "w" // auto choose white
	} else if room.Black.Conn != nil && room.White.Conn == nil {
		// only black is present
		side = "w" // auto choose white
	} else if room.Black.Conn == nil && room.White.Conn != nil {
		// only white is present
		side = "b"
	} else {
		// both are present
		// set status code 403
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"error\":\"Room is full\"}"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("Error when upgrading to webscoket", err)
	}
	defer conn.Close()

	player := Player{
		Nickname: side,
		Conn:     conn,
	}

	if side == "w" {
		room.White = player
	} else {
		room.Black = player
	}
	conn.WriteJSON(WSMessage{
		Type: "chat",
		Data: side,
	})

	for {
		var payload WSMessage
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Println("Error in ReadMessage: ", err)
		}
		log.Println(payload)
		switch payload.Type {
		case "move":
			// TODO: check if move is valid
			tokenized_move, err := TokenizeMoveMessage(payload.Data)
			if err != nil {
				log.Println(err)
				return
			}
			room.Game.AttemptMove(tokenized_move[0], tokenized_move[1], tokenized_move[2], tokenized_move[3])
		case "chat":
			room.BroadcastJSON(WSMessage{
				Type: "chat",
				Data: payload.Data,
			})

		}

		/*
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
			}
		*/
	}

}

func (r *Room) Broadcast(messageType int, p []byte) {
	if r.White.Conn != nil {
		r.White.Conn.WriteMessage(messageType, p)
	}
	if r.Black.Conn != nil {
		r.Black.Conn.WriteMessage(messageType, p)
	}
}

func (r *Room) BroadcastJSON(m interface{}) {
	if r.White.Conn != nil {
		r.White.Conn.WriteJSON(m)
	}
	if r.Black.Conn != nil {
		r.Black.Conn.WriteJSON(m)
	}
}
