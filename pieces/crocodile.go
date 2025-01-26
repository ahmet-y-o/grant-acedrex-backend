package pieces

// Same as bishop

type Crocodile struct {
	Tile  *Tile
	Color Color
}

func (t *Tile) NewCrocodile(c Color) {
	t.Piece = Crocodile{
		Tile:  t,
		Color: c,
	}
}

func (c Crocodile) GetColor() Color {
	return c.Color
}

func (c Crocodile) GetTile() *Tile {
	return c.Tile
}

func (c Crocodile) GetType() PieceType {
	return CrocodileType
}

func (c Crocodile) String() string {
	if c.Color == White {
		return "♗"
	} else {
		return "♝"
	}
}
