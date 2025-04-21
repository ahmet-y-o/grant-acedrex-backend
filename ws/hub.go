package ws

import (
	"acedrex/game"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var roomService RoomService

func InitRoomService() {
	roomService = *NewRoomService()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Player struct {
	Nickname  string
	Conn      *websocket.Conn
	SessionID string
}

type Room struct {
	RoomId           string    `json:"roomId"`
	Game             game.Game `json:"-"`
	White            Player    `json:"-"`
	Black            Player    `json:"-"`
	Started          bool      `json:"started"`
	WhiteSessionID   string
	BlackSessionID   string
	DisconnectTimers map[string]*time.Timer // Key is sessionID
	mu               sync.Mutex
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
	idsSlice := roomService.GetRoomsList()
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
	uuid := roomService.NewRoom()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("{\"roomId\":\"" + uuid + "\"}"))

}

func IsRoomFull(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	fmt.Println(roomId)
	room, exists := roomService.GetRoom(roomId)
	fmt.Println(room, exists)
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}
	if room.IsFull() {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	// get side from query params
	sideString := r.URL.Query().Get("s")
	var side game.Color
	if sideString == "white" {
		side = game.White
	} else if sideString == "black" {
		side = game.Black
	} else {
		side = game.White
	}
	// get room id from url params
	roomId := chi.URLParam(r, "roomId")
	room, exists := roomService.GetRoom(roomId)

	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	if room.Started {
		http.Error(w, "Game already started", http.StatusForbidden)
		return
	}

	side, err := room.determinePlayerSide(side)
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
		From: "server",
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
			if !room.IsFull() {
				conn.WriteJSON(WSMessage{
					Type: "error",
					Data: "Room is not full",
				})
				continue
			}
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
			room.BroadcastJSON(WSMessage{
				Type: "debug",
				Data: room.Game.AllAvailableMoves(),
			})
			room.Started = true

		case "chat":
			room.BroadcastJSON(WSMessage{
				Type: "chat",
				From: player.Nickname,
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

func (room *Room) determinePlayerSide(c game.Color) (game.Color, error) {
	if c == game.White && room.White.Conn == nil {
		return game.White, nil
	} else if c == game.Black && room.Black.Conn == nil {
		return game.Black, nil
	} else {
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
}

func (r *Room) handlePlayerDisconnect(side game.Color) {
	if side == game.White {
		r.White = Player{}
	} else {
		r.Black = Player{}
	}
	if r.IsEmpty() && !r.Started {
		roomService.DeleteRoom(r.RoomId)
	} else {
		r.BroadcastJSON(WSMessage{
			Type: "chat",
			From: "server",
			Data: side.String() + " has left the room.",
		})
	}
}

func (r *Room) IsEmpty() bool {
	return r.White.Conn == nil && r.Black.Conn == nil
}

func (r *Room) IsFull() bool {
	return r.White.Conn != nil && r.Black.Conn != nil
}
