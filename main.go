package main

import (
	"acedrex/game"
	"bufio"
	"errors"
	"fmt"
	"os"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	g := game.InitStandartGame()
	fmt.Println("Welcome to grand acedrex!")
	for {
		g.PrintBoard(os.Stdout)
		fmt.Print("Enter your move: ")
		scanner.Scan()
		move := scanner.Text()
		tokenized_move, err := TokenizeMoveMessage(move)
		fmt.Println("You entered:", tokenized_move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(tokenized_move) != 4 {
			fmt.Println("Invalid move format")
			continue
		}
		start_x := tokenized_move[0]
		start_y := tokenized_move[1]
		end_x := tokenized_move[2]
		end_y := tokenized_move[3]
		fmt.Println("Attempting move:", start_x, start_y, end_x, end_y)
		err = g.Move(start_x, start_y, end_x, end_y)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}

	/*
		ws.InitializeRooms()

		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(cors.Handler(cors.Options{
			// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"https://*", "http://*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		// Register the handler function for the root URL
		r.HandleFunc("/", ws.GetRoomsList)
		r.Get("/create-room", ws.CreateRoom)
		r.Options("/create-room", ws.CreateRoom)
		r.HandleFunc("/join-room/{roomId}", ws.JoinRoom)

		// Start the server on port 8080
		http.ListenAndServe(":8080", r)
	*/
}

func TokenizeMoveMessage(msg string) ([]string, error) {
	toReturn := []string{}

	var acc string = ""
	for _, v := range msg {
		if len(toReturn) > 4 {
			return nil, errors.New("invalid message")
		}

		// if v is letter, it is a token
		if unicode.IsLetter(v) {
			if acc != "" {
				toReturn = append(toReturn, acc)
				acc = ""
			}
			toReturn = append(toReturn, string(v))
			continue
		} else if unicode.IsDigit(v) {
			acc += string(v)
			continue
		} else {
			return nil, errors.New("invalid message")
		}
	}
	if acc != "" {
		toReturn = append(toReturn, acc)
	}
	if len(toReturn) != 4 {
		return nil, errors.New("invalid message")
	}

	return toReturn, nil
}
