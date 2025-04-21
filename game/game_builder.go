package game

type GameBuilder struct {
	game Game
}

func NewGameBuilder() GameBuilder {
	return GameBuilder{
		game: Game{},
	}
}

func InitStandardGame() Game {
	g := Game{
		Board:            make([][]*Piece, 12),
		Turn:             White,
		WhiteKing:        nil,
		BlackKing:        nil,
		WhiteKingCoords:  Coords{},
		BlackKingCoords:  Coords{},
		WhiteKingCanJump: true,
		BlackKingCanJump: true,
	}
	for y := 0; y < 12; y++ {
		g.Board[y] = make([]*Piece, 12)
		for x := 0; x < 12; x++ {
			g.Board[y][x] = nil
		}
	}
	// place black pieces
	g.Board[0][0] = &Piece{Black, RookType}
	g.Board[0][1] = &Piece{Black, LionType}
	g.Board[0][2] = &Piece{Black, UnicornoType}
	g.Board[0][3] = &Piece{Black, GiraffeType}
	g.Board[0][4] = &Piece{Black, CrocodileType}
	g.Board[0][5] = &Piece{Black, AancaType}
	g.Board[0][6] = &Piece{Black, KingType}
	g.BlackKing = g.Board[0][6]
	g.BlackKingCoords = Coords{6, 0}
	g.Board[0][7] = &Piece{Black, CrocodileType}
	g.Board[0][8] = &Piece{Black, GiraffeType}
	g.Board[0][9] = &Piece{Black, UnicornoType}
	g.Board[0][10] = &Piece{Black, LionType}
	g.Board[0][11] = &Piece{Black, RookType}

	// place white pieces
	g.Board[11][0] = &Piece{White, RookType}
	g.Board[11][1] = &Piece{White, LionType}
	g.Board[11][2] = &Piece{White, UnicornoType}
	g.Board[11][3] = &Piece{White, GiraffeType}
	g.Board[11][4] = &Piece{White, CrocodileType}
	g.Board[11][5] = &Piece{White, AancaType}
	g.Board[11][6] = &Piece{White, KingType}
	g.WhiteKing = g.Board[11][6]
	g.WhiteKingCoords = Coords{6, 11}
	g.Board[11][7] = &Piece{White, CrocodileType}
	g.Board[11][8] = &Piece{White, GiraffeType}
	g.Board[11][9] = &Piece{White, UnicornoType}
	g.Board[11][10] = &Piece{White, LionType}
	g.Board[11][11] = &Piece{White, RookType}

	// place pawns
	for x := 0; x < 12; x++ {
		g.Board[3][x] = &Piece{Black, PawnType}
		g.Board[8][x] = &Piece{White, PawnType}
	}

	return g
}

// TODO: implement
func FromGAFEN(gafen string) Game {
	return Game{}
}

func (gb *GameBuilder) SetTurn(turn Color) *GameBuilder {
	gb.game.Turn = turn
	return gb
}

func (gb *GameBuilder) EmptyBoard() *GameBuilder {
	gb.game.Board = make([][]*Piece, 12)
	for y := 0; y < 12; y++ {
		gb.game.Board[y] = make([]*Piece, 12)
		for x := 0; x < 12; x++ {
			gb.game.Board[y][x] = nil
		}
	}
	return gb
}

func (gb *GameBuilder) Finish() Game {
	return gb.game
}
