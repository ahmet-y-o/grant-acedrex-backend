package game

type GameNew struct {
	Board     [][]Piece
	Turn      Color
	WhiteKing *King
	BlackKing *King
}
