package game

type Piece struct {
	Color Color
	Type  PieceType
}

func (p *Piece) GAFEN() string {
	if p == nil {
		return "."
	}

	if p.Color == White {
		switch p.Type {
		case PawnType:
			return "P"
		case RookType:
			return "R"
		case CrocodileType:
			return "D"
		case GiraffeType:
			return "G"
		case LionType:
			return "L"
		case UnicornoType:
			return "U"
		case AancaType:
			return "X"
		case KingType:
			return "K"
		default:
			return "."
		}
	} else {
		switch p.Type {
		case PawnType:
			return "p"
		case RookType:
			return "r"
		case CrocodileType:
			return "d"
		case GiraffeType:
			return "g"
		case LionType:
			return "l"
		case UnicornoType:
			return "u"
		case AancaType:
			return "x"
		case KingType:
			return "k"
		default:
			return "."
		}
	}
}

type Coords struct {
	X int
	Y int
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
