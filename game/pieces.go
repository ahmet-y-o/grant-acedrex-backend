package game

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
	toReturn := []Coords{}

	if pawn.Color == Black {
		// black pawns can only move up
		for i := start_y + 1; i < 12; i++ {
			// check if there is a piece in the way
			piece := game.Board[i][start_x]
			if piece != nil && piece.Color == pawn.Color {
				break
			} else if piece != nil && piece.Color != pawn.Color {
				toReturn = append(toReturn, Coords{start_x, i})
				break
			}
			toReturn = append(toReturn, Coords{start_x, i})
		}
		// black pawns can only attack upleft and upright
		// check if there are any friendly piece there
		if start_x-1 >= 0 && game.Board[start_y+1][start_x-1] != nil && game.Board[start_y+1][start_x-1].Color == White {
			toReturn = append(toReturn, Coords{start_x - 1, start_y + 1})
		}
		if start_x+1 < 12 && game.Board[start_y+1][start_x+1] != nil && game.Board[start_y+1][start_x+1].Color == White {
			toReturn = append(toReturn, Coords{start_x + 1, start_y + 1})
		}
	} else {
		// white pawns can only move down
		for i := start_y - 1; i >= 0; i-- {
			// check if there is a piece in the way
			piece := game.Board[i][start_x]
			if piece != nil && piece.Color == pawn.Color {
				break
			} else if piece != nil && piece.Color != pawn.Color {
				toReturn = append(toReturn, Coords{start_x, i})
				break
			}
			toReturn = append(toReturn, Coords{start_x, i})
		}
		// white pawns can only attack downleft and downright
		// check if there are any friendly piece there
		if start_x-1 >= 0 && game.Board[start_y-1][start_x-1] != nil && game.Board[start_y-1][start_x-1].Color == Black {
			toReturn = append(toReturn, Coords{start_x - 1, start_y - 1})
		}
		if start_x+1 < 12 && game.Board[start_y-1][start_x+1] != nil && game.Board[start_y-1][start_x+1].Color == Black {
			toReturn = append(toReturn, Coords{start_x + 1, start_y - 1})
		}
	}

	return toReturn
}

func rookMoves(start_x int, start_y int, game *Game) []Coords {
	rook := game.Board[start_y][start_x]
	toReturn := []Coords{}

	// up
	for i := start_y + 1; i < 12; i++ {
		// check if there is a piece in the way
		piece := game.Board[i][start_x]
		if piece != nil && piece.Color == rook.Color {
			break
		} else if piece != nil && piece.Color != rook.Color {
			toReturn = append(toReturn, Coords{start_x, i})
			break
		}
		toReturn = append(toReturn, Coords{start_x, i})
	}

	// down
	for i := start_y - 1; i >= 0; i-- {
		// check if there is a piece in the way
		piece := game.Board[i][start_x]
		if piece != nil && piece.Color == rook.Color {
			break
		} else if piece != nil && piece.Color != rook.Color {
			toReturn = append(toReturn, Coords{start_x, i})
			break
		}
		toReturn = append(toReturn, Coords{start_x, i})
	}

	// left
	for i := start_x - 1; i >= 0; i-- {
		// check if there is a piece in the way
		piece := game.Board[start_y][i]
		if piece != nil && piece.Color == rook.Color {
			break
		} else if piece != nil && piece.Color != rook.Color {
			toReturn = append(toReturn, Coords{i, start_y})
			break
		}
		toReturn = append(toReturn, Coords{i, start_y})
	}

	// right
	for i := start_x + 1; i < 12; i++ {
		// check if there is a piece in the way
		piece := game.Board[start_y][i]
		if piece != nil && piece.Color == rook.Color {
			break
		} else if piece != nil && piece.Color != rook.Color {
			toReturn = append(toReturn, Coords{i, start_y})
			break
		}
		toReturn = append(toReturn, Coords{i, start_y})
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
			if piece != nil && piece.Color == giraffe.Color {
				continue
			} else if piece != nil && piece.Color != giraffe.Color {
				toReturn = append(toReturn, Coords{new_x, new_y})
				break
			}
			toReturn = append(toReturn, Coords{new_x, new_y})
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
	// first leaps like a classic knight (3,2 leaper)
	// then optionally moves diagonnaly in the leap direction
	leap_directions := []Coords{
		{3, 2}, {3, -2}, {-3, 2}, {-3, -2},
		{2, 3}, {2, -3}, {-2, 3}, {-2, -3},
	}
	for _, direction := range leap_directions {
		new_x := start_x + direction.X
		new_y := start_y + direction.Y
		if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
			// check if there is a piece in the way
			piece := game.Board[new_y][new_x]
			if piece != nil && piece.Color == unicornio.Color {
				continue
			} else if piece != nil && piece.Color != unicornio.Color {
				toReturn = append(toReturn, Coords{new_x, new_y})
				break
			} else {
				toReturn = append(toReturn, Coords{new_x, new_y})
				// then optional moves diagonnaly in the leap direction
				new_x += GetSign(direction.X)
				new_y += GetSign(direction.Y)
				if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
					// check if there is a piece in the way
					piece := game.Board[new_y][new_x]
					if piece != nil && piece.Color == unicornio.Color {
						break
					} else if piece != nil && piece.Color != unicornio.Color {
						toReturn = append(toReturn, Coords{new_x, new_y})
						break
					} else {
						toReturn = append(toReturn, Coords{new_x, new_y})
					}
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
		if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
			// check if there is a piece in the way
			piece := game.Board[new_y][new_x]
			if piece != nil && piece.Color == aanca.Color {
				continue
			} else if piece != nil && piece.Color != aanca.Color {
				toReturn = append(toReturn, Coords{new_x, new_y})
				break
			} else {
				toReturn = append(toReturn, Coords{new_x, new_y})
				// then optionally continuing orthogonally outward any number of squares
				// horizontal
				new_x += GetSign(direction.X)
				if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
					// check if there is a piece in the way
					piece := game.Board[new_y][new_x]
					if piece != nil && piece.Color == aanca.Color {
						break
					} else if piece != nil && piece.Color != aanca.Color {
						toReturn = append(toReturn, Coords{new_x, new_y})
						break
					} else {
						toReturn = append(toReturn, Coords{new_x, new_y})
					}
				}
				// vertical
				new_x = start_x + direction.X
				new_y += GetSign(direction.Y)
				if new_x >= 0 && new_x < 12 && new_y >= 0 && new_y < 12 {
					// check if there is a piece in the way
					piece := game.Board[new_y][new_x]
					if piece != nil && piece.Color == aanca.Color {
						break
					} else if piece != nil && piece.Color != aanca.Color {
						toReturn = append(toReturn, Coords{new_x, new_y})
						break
					} else {
						toReturn = append(toReturn, Coords{new_x, new_y})
					}
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
