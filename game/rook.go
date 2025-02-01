package game

type Rook struct {
	Tile  *Tile
	Color Color
}

func (t *Tile) NewRook(c Color) {
	t.Piece = Rook{
		Tile:  t,
		Color: c,
	}
}

func (r Rook) GetColor() Color {
	return r.Color
}

func (r Rook) GetTile() *Tile {
	return r.Tile
}

func (r Rook) GetType() PieceType {
	return RookType
}

func (r Rook) String() string {
	if r.Color == White {
		return "♖"
	} else {
		return "♜"
	}
}

func (r Rook) GetAvailableMoves(g *Game) []*Tile {
	return nil
}

func (r Rook) GAFENotation() string {
	if r.Color == White {
		return "R"
	} else {
		return "r"
	}
}
