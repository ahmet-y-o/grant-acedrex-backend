package game

type Piece interface {
	GetColor() Color
	GetX() int
	GetY() int
	GetType() PieceType
	GAFEN() string
	GAFEN_Rune() rune
	// TODO: clean
	GetAvailableMoves(*Game) []*Tile
}

type PieceType int

const (
	PawnType PieceType = iota
	RookType
	CrocodileType
	GiraffeType
	LionType
	UnicornoType
	AancaType
	KingType
)

type Color bool

const (
	White Color = true
	Black Color = false
)

type Tile struct {
	X     int
	Y     int
	Piece Piece
}

type Game struct {
	Turn           Color
	Board          [][]Tile
	WhiteKingMoved bool
	BlackKingMoved bool
}

type IntPair struct {
	X int
	Y int
}
