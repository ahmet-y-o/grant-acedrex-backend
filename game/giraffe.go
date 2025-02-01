package game

type Giraffe struct {
	Tile  *Tile
	Color Color
}

func (t *Tile) NewGiraffe(c Color) {
	t.Piece = Giraffe{
		Tile:  t,
		Color: c,
	}
}

func (g Giraffe) GetColor() Color {
	return g.Color
}

func (g Giraffe) GetTile() *Tile {
	return g.Tile
}

func (g Giraffe) GetType() PieceType {
	return GiraffeType
}

func (g Giraffe) String() string {
	if g.Color == White {
		return "♘"
	} else {
		return "♞"
	}

}

func (giraffe Giraffe) GetAvailableMoves(g *Game) []*Tile {
	return nil
}

func (g Giraffe) GAFENotation() string {
	if g.Color == White {
		return "G"
	} else {
		return "g"
	}
}
