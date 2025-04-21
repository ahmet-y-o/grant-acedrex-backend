package game

// calculates pseudo-legal moves

func pieceMoves(start_x int, start_y int, game *Game) []Coords {
	piece := game.Board[start_y][start_x]

	switch piece.Type {
	case PawnType:
		return pawnMoves(start_x, start_y, game)
	case RookType:
		return rookMoves(start_x, start_y, game)
	case CrocodileType:
		return crocodileMoves(start_x, start_y, game)
	case GiraffeType:
		return giraffeMoves(start_x, start_y, game)
	case LionType:
		return lionMoves(start_x, start_y, game)
	case UnicornoType:
		return unicornioMoves(start_x, start_y, game)
	case AancaType:
		return aancaMoves(start_x, start_y, game)
	case KingType:
		return kingMoves(start_x, start_y, game)
	default:
		panic("Unknown piece type")
	}
}

// TODO: implement promotion
func pawnMoves(start_x int, start_y int, game *Game) []Coords {
	pawn := game.Board[start_y][start_x]
	var direction int
	if pawn.Color == White {
		direction = -1
	} else {
		direction = +1
	}
	toReturn := []Coords{}

	// can move forward
	new_x := start_x
	new_y := start_y + direction
	if InBounds(new_x, new_y) && (game.Board[new_y][new_x] == nil) {
		toReturn = append(toReturn, Coords{X: new_x, Y: new_y})
	}

	// can capture diagonally
	new_x = start_x + 1
	new_y = start_y + direction
	if InBounds(new_x, new_y) && (game.Board[new_y][new_x] != nil) && (game.Board[new_y][new_x].Color != pawn.Color) {
		toReturn = append(toReturn, Coords{X: new_x, Y: new_y})
	}

	new_x = start_x - 1
	new_y = start_y + direction
	if InBounds(new_x, new_y) && (game.Board[new_y][new_x] != nil) && (game.Board[new_y][new_x].Color != pawn.Color) {
		toReturn = append(toReturn, Coords{X: new_x, Y: new_y})
	}

	return toReturn
}

func rookMoves(start_x int, start_y int, game *Game) []Coords {
	toReturn := []Coords{}

	rook := game.Board[start_y][start_x]

	directions := []Coords{
		{1, 0}, {0, 1},
		{-1, 0}, {0, -1},
	}

	for _, direction := range directions {
		new_x := start_x
		new_y := start_y
		for {
			new_x += direction.X
			new_y += direction.Y
			if !InBounds(new_x, new_y) {
				break
			}
			piece := game.Board[new_y][new_x]
			if piece == nil {
				toReturn = append(toReturn, Coords{X: new_x, Y: new_y})
			} else if piece.Color != rook.Color {
				toReturn = append(toReturn, Coords{X: new_x, Y: new_y})
				break
			} else {
				break
			}
		}
	}

	return toReturn
}

func crocodileMoves(start_x int, start_y int, game *Game) []Coords {
	// moves like a bishop
	bishop := game.Board[start_y][start_x]

	toReturn := []Coords{}

	// up left
	for i, j := start_y+1, start_x-1; i < 12 && j >= 0; i, j = i+1, j-1 {
		// check if there is a piece in the way
		piece := game.Board[i][j]
		if piece != nil && piece.Color == bishop.Color {
			break
		} else if piece != nil && piece.Color != bishop.Color {
			toReturn = append(toReturn, Coords{j, i})
			break
		}
		toReturn = append(toReturn, Coords{j, i})
	}

	// up right
	for i, j := start_y+1, start_x+1; i < 12 && j < 12; i, j = i+1, j+1 {
		// check if there is a piece in the way
		piece := game.Board[i][j]
		if piece != nil && piece.Color == bishop.Color {
			break
		} else if piece != nil && piece.Color != bishop.Color {
			toReturn = append(toReturn, Coords{j, i})
			break
		}
		toReturn = append(toReturn, Coords{j, i})
	}

	// down left
	for i, j := start_y-1, start_x-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		// check if there is a piece in the way
		piece := game.Board[i][j]
		if piece != nil && piece.Color == bishop.Color {
			break
		} else if piece != nil && piece.Color != bishop.Color {
			toReturn = append(toReturn, Coords{j, i})
			break
		}
		toReturn = append(toReturn, Coords{j, i})
	}

	// down right
	for i, j := start_y-1, start_x+1; i >= 0 && j < 12; i, j = i-1, j+1 {
		// check if there is a piece in the way
		piece := game.Board[i][j]
		if piece != nil && piece.Color == bishop.Color {
			break
		} else if piece != nil && piece.Color != bishop.Color {
			toReturn = append(toReturn, Coords{j, i})
			break
		}
		toReturn = append(toReturn, Coords{j, i})
	}

	return toReturn
}

func giraffeMoves(start_x int, start_y int, game *Game) []Coords {
	giraffe := game.Board[start_y][start_x]
	toReturn := []Coords{}
	directions := []Coords{
		{3, 2}, {3, -2}, {-3, 2}, {-3, -2},
		{2, 3}, {2, -3}, {-2, 3}, {-2, -3},
	}

	for _, direction := range directions {
		new_x := start_x + direction.X
		new_y := start_y + direction.Y

		if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
			// check if there is a piece in the way
			piece := game.Board[new_y][new_x]
			if piece == nil {
				toReturn = append(toReturn, Coords{new_x, new_y})
				continue
			} else if piece.Color == giraffe.Color {
				continue
			} else if piece.Color != giraffe.Color {
				toReturn = append(toReturn, Coords{new_x, new_y})
				continue
			}
		}
	}
	return toReturn
}

func lionMoves(start_x int, start_y int, game *Game) []Coords {
	lion := game.Board[start_y][start_x]
	toReturn := []Coords{}
	directions := []Coords{
		{3, 1}, {3, 0}, {3, -1},
		{-3, 1}, {-3, 0}, {-3, -1},
		{1, 3}, {0, 3}, {-1, 3},
		{1, -3}, {0, -3}, {-1, -3},
	}

	for _, direction := range directions {
		new_x := start_x + direction.X
		new_y := start_y + direction.Y
		if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
			// check if there is a piece in the way
			piece := game.Board[new_y][new_x]
			if piece != nil && piece.Color == lion.Color {
				continue
			} else if piece != nil && piece.Color != lion.Color {
				toReturn = append(toReturn, Coords{new_x, new_y})
				break
			}
			toReturn = append(toReturn, Coords{new_x, new_y})
		}
	}

	return toReturn
}

func unicornioMoves(start_x int, start_y int, game *Game) []Coords {
	toReturn := []Coords{}
	unicornio := game.Board[start_y][start_x]
	// first leaps like a classic knight (2,1 leaper)
	// then optionally moves diagonnaly in the leap direction
	leap_directions := []Coords{
		{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
		{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
	}
	for _, direction := range leap_directions {
		new_x := start_x + direction.X
		new_y := start_y + direction.Y
		if InBounds(new_x, new_y) {
			// check if there is a piece on the tile
			piece := game.Board[new_y][new_x]
			if piece == nil {
				toReturn = append(toReturn, Coords{new_x, new_y})
			} else if piece.Color != unicornio.Color {
				// capture the piece, dont check diagonal
				toReturn = append(toReturn, Coords{new_x, new_y})
				continue
			} else {
				continue
			}
			// optinonally move diagonally
			for {
				new_x += GetSign(direction.X)
				new_y += GetSign(direction.Y)
				if !InBounds(new_x, new_y) {
					break
				}
				piece = game.Board[new_y][new_x]
				if piece == nil {
					toReturn = append(toReturn, Coords{new_x, new_y})
				} else if piece.Color != unicornio.Color {
					toReturn = append(toReturn, Coords{new_x, new_y})
					break
				} else {
					break
				}
			}

		}
	}

	return toReturn
}

func aancaMoves(start_x int, start_y int, game *Game) []Coords {
	toReturn := []Coords{}
	aanca := game.Board[start_y][start_x]
	// first move 1 square diagonally
	// then optionally continuing orthogonally outward any number of squares
	diagonal_directions := []Coords{
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}
	for _, direction := range diagonal_directions {
		new_x := start_x + direction.X
		new_y := start_y + direction.Y
		if InBounds(new_x, new_y) {
			// check if there is a piece on the tile
			piece := game.Board[new_y][new_x]
			if piece == nil {
				toReturn = append(toReturn, Coords{new_x, new_y})
			} else if piece.Color != aanca.Color {
				// capture the piece, dont check diagonal
				toReturn = append(toReturn, Coords{new_x, new_y})
				continue
			} else {
				continue
			}
			//  optinonally move orthogonally
			// first on row
			for {
				new_x += GetSign(direction.X)
				if !InBounds(new_x, new_y) {
					break
				}
				piece := game.Board[new_y][new_x]
				if piece == nil {
					toReturn = append(toReturn, Coords{new_x, new_y})
				} else if piece.Color != aanca.Color {
					toReturn = append(toReturn, Coords{new_x, new_y})
					break
				} else {
					break
				}
			}
			// then on column
			// reset x position
			new_x = start_x + direction.X
			for {
				new_y += GetSign(direction.Y)
				if !InBounds(new_x, new_y) {
					break
				}
				piece := game.Board[new_y][new_x]
				if piece == nil {
					toReturn = append(toReturn, Coords{new_x, new_y})
				} else if piece.Color != aanca.Color {
					toReturn = append(toReturn, Coords{new_x, new_y})
					break
				} else {
					break
				}
			}

		}

	}
	return toReturn
}

func kingMoves(start_x int, start_y int, game *Game) []Coords {
	king := game.Board[start_y][start_x]
	toReturn := []Coords{}
	var king_can_jump bool
	if king.Color == White {
		king_can_jump = game.WhiteKingCanJump
	} else {
		king_can_jump = game.BlackKingCanJump
	}

	directions := []Coords{
		{1, 1}, {1, 0}, {1, -1},
		{0, 1}, {0, -1},
		{-1, 1}, {-1, 0}, {-1, -1},
	}
	jump_directions := []Coords{
		{2, 2}, {2, -2}, {-2, 2}, {-2, -2},
		{2, 0}, {-2, 0},
		{0, 2}, {0, -2},
	}

	for _, direction := range directions {
		new_x := start_x + direction.X
		new_y := start_y + direction.Y
		if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
			// check if there is a piece in the way
			piece := game.Board[new_y][new_x]
			if piece != nil && piece.Color == king.Color {
				continue
			} else if piece != nil && piece.Color != king.Color {
				toReturn = append(toReturn, Coords{new_x, new_y})
				break
			}
			toReturn = append(toReturn, Coords{new_x, new_y})
		}
	}

	if king_can_jump {
		for _, direction := range jump_directions {
			new_x := start_x + direction.X
			new_y := start_y + direction.Y
			if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
				// check if there is a piece in the way
				piece := game.Board[new_y][new_x]
				if piece != nil && piece.Color == king.Color {
					continue
				} else if piece != nil && piece.Color != king.Color {
					toReturn = append(toReturn, Coords{new_x, new_y})
					break
				}
				toReturn = append(toReturn, Coords{new_x, new_y})
			}
		}
	}

	return toReturn
}
