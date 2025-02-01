package game

type Aanca struct {
	X     int
	Y     int
	Color Color
}

func (a *Aanca) GetColor() Color {
	return a.Color
}

func (a *Aanca) GetX() int {
	return a.X
}

func (a *Aanca) GetY() int {
	return a.Y
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

func (a Aanca) GetAvailableMoves(g *Game) []*Tile {
	return []*Tile{}
}

func (a Aanca) GAFENotation() string {
	if a.Color == White {
		return "A"
	} else {
		return "a"
	}
}
