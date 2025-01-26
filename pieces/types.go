package pieces

import "errors"

type Tile struct {
	X     int
	Y     int
	Piece Piece
}

func (t *Tile) NewPiece(p PieceType, c Color) {

	switch p {
	case AancaType:
		t.NewAanca(c)
	case CrocodileType:
		t.NewCrocodile(c)
	case GiraffeType:
		t.NewGiraffe(c)
	case KingType:
		t.NewKing(c)
	case LionType:
		t.NewLion(c)
	case PawnType:
		t.NewPawn(c)
	case RookType:
		t.NewRook(c)
	case UnicornoType:
		t.NewUnicornio(c)
	}
}

func CharToNum(c int) (int, error) {
	if c < 'a' && c >= 'm' {
		return -1, errors.New("out of bounds")
	}
	return c - 'a', nil
}

func NumToChar(n int) (int, error) {
	if n < 0 && n >= 12 {
		return -1, errors.New("out of bounds")
	}
	return n + 'a', nil
}

type Piece interface {
	GetColor() Color
	GetTile() *Tile
	GetType() PieceType
	String() string
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

type Color int

const (
	White Color = iota
	Black
)
