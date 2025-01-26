package main

import (
	"acedrex/ws"
	"net/http"
)

func main() {
	/*
		g := game.InitilaizeGame()
		g, err := g.AttemptMove('a', 4, 'a', 5)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println('a', 4, 'a', 5)
		g.PrintBoard(os.Stdout, true)
	*/

	ws.InitializeRooms()

	// Register the handler function for the root URL
	http.HandleFunc("/", ws.GetRoomsList)
	http.HandleFunc("/create-room", ws.CreateRoom)
	http.HandleFunc("/join-room", ws.JoinRoom)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)

}
