package game

import (
	"errors"
	"fmt"
	"io"
)

type Game struct {
	Board            [][]*Piece
	Turn             Color
	WhiteKing        *Piece
	BlackKing        *Piece
	WhiteKingCoords  Coords
	BlackKingCoords  Coords
	WhiteKingCanJump bool
	BlackKingCanJump bool
}

func (g *Game) ToGAFEN() string {
	part_1, part_2, part_3 := "", "", ""

	// part 1 is the board
	// TODO: implement empty space counting
	empty_count := 0
	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			if g.Board[y][x] == nil {
				empty_count++
			} else {
				if empty_count > 0 {
					to_add, err := IntToHex(empty_count)
					if err != nil {
						panic(err)
					}
					part_1 += to_add
				}
				part_1 += (*g.Board[y][x]).GAFEN()
				empty_count = 0
			}
		}
		if empty_count > 0 {
			to_add, err := IntToHex(empty_count)
			if err != nil {
				panic(err)
			}
			part_1 += to_add
		}
		empty_count = 0
		part_1 += "/"
	}
	// remove trailing slash
	part_1 = part_1[:len(part_1)-1]

	// part 2 is the turn
	if g.Turn == White {
		part_2 += "w"
	} else {
		part_2 += "b"
	}

	// part 3 is king's ability to jump
	if !g.WhiteKingCanJump && !g.BlackKingCanJump {
		part_3 += "-"
	} else {
		if g.WhiteKingCanJump {
			part_3 += "J"
		}
		if g.BlackKingCanJump {
			part_3 += "j"
		}
	}

	return part_1 + " " + part_2 + " " + part_3
}

/*
* 1. get the opponent king
* 2. check if the current player can capture the opponent king
 */
func (g *Game) IsLegal() bool {
	var kingCoords Coords
	if g.Turn == White {
		kingCoords = g.BlackKingCoords
	} else {
		kingCoords = g.WhiteKingCoords
	}

	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			piece := g.Board[y][x]
			if piece == nil || piece.Color != g.Turn {
				// means either there are is no piece or the piece is not ours
				continue
			} else {
				// means piece is ours
				x_str, err := IntToNotation(x)
				if err != nil {
					panic(err)
				}
				moves, err := g.GetAvailableMoves(x_str, y)
				if err != nil {
					panic(err)
				}
				for _, move := range moves {
					if move.X == kingCoords.X && move.Y == kingCoords.Y {
						return false
					}
				}

			}
		}
	}
	return true
}

/*
*
 1. Get attempted piece
 2. Check if piece movement is possible
 3. Store position as gafen
 4. Force the move
 5. Check if move is legal
 6. a- if legal, return nil
    b- if not, return error and restore position
*/
func (g *Game) Move(start_x string, start_y string, end_x string, end_y string) error {
	start_pos, err := NotationToCoords(start_x + start_y)
	if err != nil {
		return err
	}
	end_pos, err := NotationToCoords(end_x + end_y)
	if err != nil {
		return err
	}
	fmt.Println("start:", start_pos, "end:", end_pos)

	end_piece_address := g.Board[end_pos.Y][end_pos.X]

	start_piece := g.Board[start_pos.Y][start_pos.X]
	if start_piece == nil {
		return errors.New("no piece")
	}

	if start_piece.Color != g.Turn {
		return errors.New("not your piece")
	}

	moves := pieceMoves(start_pos.X, start_pos.Y, g)
	for _, move := range moves {
		if move.X == end_pos.X && move.Y == end_pos.Y {
			g.Board[start_pos.Y][start_pos.X] = nil
			g.Board[end_pos.Y][end_pos.X] = start_piece
			g.Turn = !g.Turn
			if !g.IsLegal() {
				// restore position
				g.Board[start_pos.Y][start_pos.X] = start_piece
				g.Board[end_pos.Y][end_pos.X] = end_piece_address
				return errors.New("illegal move")
			} else {
				return nil
			}
		}
	}
	return errors.New("piece cannot move like that")
}
func (g *Game) GetTurn() Color {
	return g.Turn
}

func (g *Game) CheckStatus() int {
	// TODO: implement an enum iota
	return 0
}

func (g *Game) GetAvailableMoves(x string, y int) ([]Coords, error) {
	x_int, err := NotationToInt(x)
	if err != nil {
		return nil, err
	}
	piece := g.Board[y][x_int]
	if piece == nil {
		return nil, errors.New("no piece")
	}

	return pieceMoves(x_int, y, g), nil
}

func (g *Game) PrintBoard(writer io.Writer) {
	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			if g.Board[y][x] != nil {
				fmt.Fprintf(writer, "%s ", g.Board[y][x].GAFEN())
			} else {
				fmt.Fprintf(writer, "- ")
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
