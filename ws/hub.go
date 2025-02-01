package ws

import (
	"acedrex/game"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
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

// TODO: garbage collector for rooms without players
// TODO: check if the players are still connected

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
	room, exists := rooms[roomId]

	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	side, err := room.determinePlayerSide()

	if err != nil {
		// room is full
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Error upgrading to websocket", http.StatusInternalServerError)
		return
	}

	player := Player{
		Nickname: "Player " + side.String(),
		Conn:     conn,
	}
	defer func() {
		conn.Close()
		room.handlePlayerDisconnect(side)
	}()

	if side == game.White {
		room.White = player
	} else {
		room.Black = player
	}

	conn.SetCloseHandler(func(code int, text string) error {
		room.handlePlayerDisconnect(side)
		return nil
	})

	if err := conn.WriteJSON(WSMessage{
		Type: "side",
		Data: side.String(),
	}); err != nil {
		// TODO: understand how if err := conn works
		return
	}

	room.BroadcastJSON(WSMessage{
		Type: "chat",
		Data: side.String() + " has joined the room.",
	})

	// Main loop
	for {
		var payload WSMessage

		err := conn.ReadJSON(&payload)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error: %v", err)
			}
			break
		}

		// TODO: debug, delete later
		log.Println(payload)

		switch payload.Type {
		case "move":
			// TODO: check if move is valid
			tokenized_move, err := TokenizeMoveMessage(payload.Data)
			if err != nil {
				log.Println(err)
			}
			log.Println(tokenized_move)
			if side != room.Game.Turn {
				conn.WriteJSON(WSMessage{
					Type: "error",
					Data: "Not your turn",
				})
				continue
			}
			err = room.Game.Move(tokenized_move[0], tokenized_move[1], tokenized_move[2], tokenized_move[3])
			if err != nil {
				conn.WriteJSON(WSMessage{
					Type: "error",
					Data: err.Error(),
				})
				continue
			}
			room.BroadcastJSON(WSMessage{
				Type: "move",
				Data: payload.Data,
			})
			room.Game.PrintBoard(os.Stdout, false)
		case "chat":
			room.BroadcastJSON(WSMessage{
				Type: "chat",
				Data: payload.Data,
			})
		}
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

func (room *Room) determinePlayerSide() (game.Color, error) {
	switch {
	case room.Black.Conn == nil && room.White.Conn == nil:
		return game.White, nil
	case room.Black.Conn != nil && room.White.Conn == nil:
		return game.White, nil
	case room.Black.Conn == nil && room.White.Conn != nil:
		return game.Black, nil
	default:
		return game.White, fmt.Errorf("room is full")
	}
}

func (r *Room) handlePlayerDisconnect(side game.Color) {
	if side == game.White {
		r.White = Player{}
	} else {
		r.Black = Player{}
	}
	r.BroadcastJSON(WSMessage{
		Type: "chat",
		Data: side.String() + " has left the room.",
	})
}
