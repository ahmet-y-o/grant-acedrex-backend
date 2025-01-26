package pieces

type Aanca struct {
	Tile  *Tile
	Color Color
}

func (t *Tile) NewAanca(c Color) {
	t.Piece = Aanca{
		Tile:  t,
		Color: c,
	}
}

func (a Aanca) GetColor() Color {
	return a.Color
}

func (a Aanca) GetTile() *Tile {
	return a.Tile
}

func (a Aanca) GetType() PieceType {
	return AancaType
}

func (a Aanca) String() string {
	if a.Color == White {
		return "🨠"
	} else {
		return "🨦"
	}
}
