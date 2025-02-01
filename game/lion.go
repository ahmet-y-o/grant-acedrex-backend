package game

type Lion struct {
	Tile  *Tile
	Color Color
}

func (t *Tile) NewLion(c Color) {
	t.Piece = Lion{
		Tile:  t,
		Color: c,
	}
}

func (l Lion) GetColor() Color {
	return l.Color
}

func (l Lion) GetTile() *Tile {
	return l.Tile
}

func (l Lion) GetType() PieceType {
	return LionType
}

func (l Lion) String() string {
	if l.Color == White {
		return "🩏"
	} else {
		return "🩒"
	}
}

func (l Lion) GetAvailableMoves(g *Game) []*Tile {
	return nil
}

func (l Lion) GAFENotation() string {
	if l.Color == White {
		return "L"
	} else {
		return "l"
	}
}
