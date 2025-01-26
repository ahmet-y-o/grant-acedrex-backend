package pieces

type King struct {
	Tile     *Tile
	Color    Color
	HasMoved bool
}

func (t *Tile) NewKing(c Color) {
	t.Piece = King{
		Tile:     t,
		Color:    c,
		HasMoved: false,
	}
}

func (k King) GetColor() Color {
	return k.Color
}

func (k King) GetTile() *Tile {
	return k.Tile
}

func (k King) GetType() PieceType {
	return KingType
}

func (k King) String() string {
	if k.Color == White {
		return "♔"
	} else {
		return "♚"
	}
}
