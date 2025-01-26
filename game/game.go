package game

import (
	"acedrex/pieces"
	"errors"
	"io"
	"log"
)

type Game struct {
	Turn  pieces.Color
	Board [][]pieces.Tile
}

type IntPair struct {
	X int
	Y int
}

func InitilaizeGame() *Game {
	b := [][]pieces.Tile{}

	for y := 1; y <= 12; y++ {
		row := []pieces.Tile{}
		for x := 0; x < 12; x++ {
			c, err := pieces.NumToChar(x)
			if err != nil {
				log.Fatalf("Error when initilazing game", err)
			}
			row = append(row, pieces.Tile{X: c, Y: y})
		}
		b = append(b, row)
	}

	// initialize white
	for y := 0; y < 12; y++ {
		b[3][y].NewPawn(pieces.White)
	}
	b[0][0].NewRook(pieces.White)
	b[0][11].NewRook(pieces.White)
	b[0][1].NewLion(pieces.White)
	b[0][10].NewLion(pieces.White)
	b[0][2].NewUnicornio(pieces.White)
	b[0][9].NewUnicornio(pieces.White)
	b[0][4].NewCrocodile(pieces.White)
	b[0][7].NewCrocodile(pieces.White)
	b[0][3].NewGiraffe(pieces.White)
	b[0][8].NewGiraffe(pieces.White)
	b[0][5].NewAanca(pieces.White)
	b[0][6].NewKing(pieces.White)

	// initialize black
	for y := 0; y < 12; y++ {
		b[8][y].NewPawn(pieces.Black)
	}
	b[11][0].NewRook(pieces.Black)
	b[11][11].NewRook(pieces.Black)
	b[11][1].NewLion(pieces.Black)
	b[11][10].NewLion(pieces.Black)
	b[11][2].NewUnicornio(pieces.Black)
	b[11][9].NewUnicornio(pieces.Black)
	b[11][4].NewCrocodile(pieces.Black)
	b[11][7].NewCrocodile(pieces.Black)
	b[11][3].NewGiraffe(pieces.Black)
	b[11][8].NewGiraffe(pieces.Black)
	b[11][5].NewAanca(pieces.Black)
	b[11][6].NewKing(pieces.Black)

	g := Game{
		Turn:  pieces.White, // white's turn is true, black's turn is false
		Board: b,
	}

	return &g
}

// x is char, y is number. standart notation
func (g *Game) GetTile(x, y int) *pieces.Tile {
	rx, err := pieces.CharToNum(x)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &(g.Board[y-1][rx])
}

func (g *Game) PassTurn() {
	if g.Turn == pieces.White {
		g.Turn = pieces.Black
	} else {
		g.Turn = pieces.White
	}
}

// does not check turn order
func (g *Game) UnrestrictedMove(sxc, sy, exc, ey int) {
	p := g.GetTile(sxc, sy).Piece
	g.GetTile(exc, ey).NewPiece(p.GetType(), p.GetColor())
	g.GetTile(sxc, sy).Piece = nil
}

func (g *Game) AttemptMove(sxc, sy, exc, ey int) (*Game, error) {

	// TODO: implement promotion

	sx, err := pieces.CharToNum(sxc)
	if err != nil {
		log.Fatalf(err.Error())
	}
	ex, err := pieces.CharToNum(exc)
	if err != nil {
		log.Fatalf(err.Error())
	}

	attemptedBoard := make([][]pieces.Tile, len(g.Board))
	copy(attemptedBoard, g.Board)
	unrestictedMoved := false

	attemptedGame := Game{
		Turn:  g.Turn,
		Board: attemptedBoard, // does this copy or is it reference?
	}

	if sx == ex && sy == ey {
		return g, errors.New("can't move in place")
	}

	if ex >= 12 || ey >= 12 {
		return g, errors.New("destination is out of the board")
	}

	attemptedSource := g.Board[sy-1][sx]
	attemptedDestination := g.Board[ey-1][ex]

	if attemptedSource.Piece == nil {
		return g, errors.New("no piece at source")
	}

	if g.Turn != attemptedSource.Piece.GetColor() {
		return g, errors.New("not your piece")
	}

	if attemptedDestination.Piece != nil && g.Turn == attemptedDestination.Piece.GetColor() {
		return g, errors.New("you can't capture your own piece")
	}

	// check if piece's movement is such
	switch attemptedSource.Piece.GetType() {
	case pieces.AancaType: // moves 1 tile diagonally and if wanted like a rook
		// TODO: check if there are any pieces in between
		if (AbsInt(sx-ex) == 1 && AbsInt(sy-ey) >= 1) || (AbsInt(sx-ex) >= 1 && AbsInt(sy-ey) == 1) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.PassTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("aanca can't move like that")
		}
	case pieces.CrocodileType: // moves like a bishop
		// TODO: check if there are any pieces in between
		if AbsInt(sx-ex) == AbsInt(sy-ey) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.PassTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("crocodile can't move like that")
		}
	case pieces.GiraffeType: // girafe is 3-2 leaper
		if (AbsInt(sx-ex) == 3 && AbsInt(sy-ey) == 2) || (AbsInt(sx-ex) == 2 && AbsInt(sy-ey) == 3) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.PassTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("giraffe can't move like that")
		}

	case pieces.KingType: // moves like a standart king, on its first move can move 2 tiles
		k, ok := attemptedSource.Piece.(pieces.King)
		if ok {
			// TODO: tidy up
			// should always be ok
			if (AbsInt(sx-ex) == 1 && (AbsInt(sy-ey) == 1 || AbsInt(sy-ey) == 0)) ||
				(AbsInt(sy-ey) == 1 && (AbsInt(sx-ex) == 1 || AbsInt(sx-ex) == 0)) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				nk, _ := attemptedGame.GetTile(exc, ey).Piece.(pieces.King)
				nk.HasMoved = true
				attemptedGame.PassTurn()
				unrestictedMoved = true
			} else if !k.HasMoved && ((AbsInt(sx-ex) == 2 && (AbsInt(sy-ey) == 2 || AbsInt(sy-ey) == 0)) ||
				(AbsInt(sy-ey) == 2 && (AbsInt(sx-ex) == 2 || AbsInt(sx-ex) == 0))) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				nk, _ := attemptedGame.GetTile(exc, ey).Piece.(pieces.King)
				nk.HasMoved = true
				attemptedGame.PassTurn()
				unrestictedMoved = true
			}
		}

	case pieces.LionType: // 3-1 leaper and 3-0 leaper
		if (AbsInt(sx-ex) == 3 && (AbsInt(sy-ey) == 1 || AbsInt(sy-ey) == 0)) ||
			(AbsInt(sy-ey) == 3 && (AbsInt(sx-ex) == 1 || AbsInt(sx-ex) == 0)) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.PassTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("lion can't move like that")
		}

	case pieces.PawnType: // no initial jump, moves like a standart pawn
		// check if the pawn is moving foward 1 tile
		if ey == sy+1 && (g.Turn == pieces.White) {
			if (ex == sx) && (attemptedDestination.Piece == nil) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				attemptedGame.PassTurn()
				unrestictedMoved = true

			} else if AbsInt(sx-ex) == 1 && attemptedDestination.Piece != nil && attemptedDestination.Piece.GetColor() != g.Turn {
				// captures
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				attemptedGame.PassTurn()
				unrestictedMoved = true
			} else {
				return g, errors.New("pawns can't move that way")
			}
			// check if pawn is moving forward 1 tile but for black
		} else if ey == sy-1 && (g.Turn == pieces.Black) {
			if (ex == sx) && (attemptedDestination.Piece == nil) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				attemptedGame.PassTurn()
				unrestictedMoved = true
			} else {
				return g, errors.New("pawns can't move that way")
			}
		} else {
			return g, errors.New("pawn must move forward vertically")
		}

	case pieces.RookType: // moves like a modern rook
		if (AbsInt(sx-ex) == 0 && AbsInt(sy-ey) >= 1) ||
			(AbsInt(sx-ex) >= 1 && AbsInt(sy-ey) == 0) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.PassTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("rook can't move like that")
		}
	case pieces.UnicornoType: // most complex piece

	}

	if !unrestictedMoved {
		return g, errors.New("no such movement. (shouldn't get this error, means we have edge cases in movement checks)")
	}

	// check if your king is in check after the move
	if !attemptedGame.IsLegal() {
		return g, errors.New("illegal move")
	}

	return &attemptedGame, nil
}

func (g *Game) IsLegal() bool {

	return true
}

func (g *Game) PrintBoard(w io.Writer, turned bool) {

	if !turned {
		w.Write([]byte("|"))
		for i := 1; i <= 23; i++ {
			w.Write([]byte("-"))
		}
		w.Write([]byte("|\n"))

		for y := 0; y < 12; y++ {
			// each row
			w.Write([]byte("|"))
			for x := 0; x < 12; x++ {
				if g.Board[y][x].Piece == nil {
					w.Write([]byte(" "))
					//w.Write([]byte(strconv.Itoa(x) + " " + strconv.Itoa(y)))
				} else {
					w.Write([]byte(g.Board[y][x].Piece.String()))
				}
				w.Write([]byte("|"))
			}
			w.Write([]byte("\n"))

			w.Write([]byte("|"))
			for i := 1; i <= 23; i++ {
				w.Write([]byte("-"))
			}
			w.Write([]byte("|\n"))
		}
	} else {
		w.Write([]byte("|"))
		for i := 1; i <= 23; i++ {
			w.Write([]byte("-"))
		}
		w.Write([]byte("|\n"))

		for y := 11; y >= 0; y-- {
			// each row
			w.Write([]byte("|"))
			for x := 0; x < 12; x++ {
				if g.Board[y][x].Piece == nil {
					w.Write([]byte(" "))
					//w.Write([]byte(strconv.Itoa(x) + " " + strconv.Itoa(y)))
				} else {
					w.Write([]byte(g.Board[y][x].Piece.String()))
				}
				w.Write([]byte("|"))
			}
			w.Write([]byte("\n"))

			w.Write([]byte("|"))
			for i := 1; i <= 23; i++ {
				w.Write([]byte("-"))
			}
			w.Write([]byte("|\n"))
		}
	}
}

func AbsInt(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
