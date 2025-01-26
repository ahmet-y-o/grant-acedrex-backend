package pieces

type Pawn struct {
	Tile  *Tile
	Color Color
	// en passent
	// two moves
}

func (t *Tile) NewPawn(c Color) {
	t.Piece = Pawn{
		Tile:  t,
		Color: c,
	}
}

func (p Pawn) GetColor() Color {
	return p.Color
}

func (p Pawn) GetTile() *Tile {
	return p.Tile
}

func (p Pawn) GetType() PieceType {
	return PawnType
}

func (p Pawn) String() string {
	if p.Color == White {
		return "♙"
	} else {
		return "♟"
	}

}
