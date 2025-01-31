package ws

import (
	"errors"
	"unicode"
)

type WSMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
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
		} else if v == '-' {
			toReturn = append(toReturn, acc)
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
