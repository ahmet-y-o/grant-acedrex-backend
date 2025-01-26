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
	"strconv"
	"strings"
	"time"

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
	RoomId  string    `json:"roomId"`
	Game    game.Game `json:"-"`
	Player1 Player    `json:"-"`
	Player2 Player    `json:"-"`
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

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	// get room id from query params
	// TODO: upgrade connection to WS
	g := game.InitilaizeGame()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("Error when upgrading to webscoket", err)
	}

	// TODO: assign a player from the game to the connection
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error in ReadMessage: ", err)

		}

		// move y can have 1 or 2 digits
		// have "-" in between characters
		message := strings.Split(string(p), "-")
		log.Println(message)
		//echo message to client
		writer, err := conn.NextWriter(messageType)
		if err != nil {
			log.Println("Error in ReadMessage: ", err)
		}
		sxc := int(message[0][0])
		sy, _ := strconv.Atoi(string(message[1]))
		exc := int(message[2][0])
		ey, _ := strconv.Atoi(string(message[3]))

		fmt.Println(sxc, sy, exc, ey)
		g, err := g.AttemptMove(sxc, sy, exc, ey)
		if err != nil {
			log.Println("Error in ReadMessage: ", err)

		}
		g.PrintBoard(os.Stdout, true)
		// echo board to the client
		g.PrintBoard(writer, true)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)

		}
	}

}
