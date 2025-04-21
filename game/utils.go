package game

import (
	"strconv"
)

// used for GAFEN conversion
func IntToHex(i int) (string, error) {
	if i <= 0 || i >= 12 {
		return "", strconv.ErrRange
	}
	return strconv.FormatInt(int64(i), 16), nil
}

// used for GAFEN conversion
func HexToInt(s string) (int, error) {
	i, err := strconv.ParseInt(s, 16, 64)
	if i < 0 || i > 12 {
		return 0, strconv.ErrRange
	}
	return int(i), err
}

func NotationToCoords(notation string) (Coords, error) {
	toReturn := Coords{}
	notation_x := notation[0:1] // single character string from a to l
	notation_y := notation[1:]  // number from 1 to 12 in string format

	x, err := NotationToInt(notation_x)
	if err != nil {
		return toReturn, err
	}
	y, err := strconv.Atoi(notation_y)
	if err != nil {
		return toReturn, err
	}
	if y < 1 || y > 12 {
		return toReturn, strconv.ErrRange
	}
	toReturn.X = x
	toReturn.Y = 12 - y

	return toReturn, nil
}

func CoordsToNotation(coords Coords) (string, error) {
	toReturn := ""
	x, err := IntToNotation(coords.X)
	if err != nil {
		return "", err
	}
	toReturn += x

	y := strconv.Itoa(12 - coords.Y)
	toReturn += y
	return toReturn, nil
}

func NotationToInt(notation string) (int, error) {
	switch notation {
	case "a":
		return 0, nil
	case "b":
		return 1, nil
	case "c":
		return 2, nil
	case "d":
		return 3, nil
	case "e":
		return 4, nil
	case "f":
		return 5, nil
	case "g":
		return 6, nil
	case "h":
		return 7, nil
	case "i":
		return 8, nil
	case "j":
		return 9, nil
	case "k":
		return 10, nil
	case "l":
		return 11, nil
	default:
		return 0, strconv.ErrRange
	}
}

func IntToNotation(i int) (string, error) {
	switch i {
	case 0:
		return "a", nil
	case 1:
		return "b", nil
	case 2:
		return "c", nil
	case 3:
		return "d", nil
	case 4:
		return "e", nil
	case 5:
		return "f", nil
	case 6:
		return "g", nil
	case 7:
		return "h", nil
	case 8:
		return "i", nil
	case 9:
		return "j", nil
	case 10:
		return "k", nil
	case 11:
		return "l", nil
	default:
		return "", strconv.ErrRange
	}
}

func InBounds(x int, y int) bool {
	return x >= 0 && x < 12 && y >= 0 && y < 12
}

func GetSign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}
