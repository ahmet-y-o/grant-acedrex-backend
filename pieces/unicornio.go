package pieces

type Unicornio struct {
	Tile  *Tile
	Color Color
}

func (t *Tile) NewUnicornio(c Color) {
	t.Piece = Unicornio{
		Tile:  t,
		Color: c,
	}
}

func (u Unicornio) GetColor() Color {
	return u.Color
}

func (u Unicornio) GetTile() *Tile {
	return u.Tile
}

func (u Unicornio) GetType() PieceType {
	return UnicornoType
}

func (u Unicornio) String() string {
	if u.Color == White {
		return "🩎"
	} else {
		return "🩑"
	}
}
