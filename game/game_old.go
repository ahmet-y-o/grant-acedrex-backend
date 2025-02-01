package game

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"unicode"
)

func (t *Tile) NewPiece(p PieceType, c Color) {
	switch p {
	case AancaType:
		t.NewAanca(c)
	case CrocodileType:
		t.NewCrocodile(c)
	case GiraffeType:
		t.NewGiraffe(c)
	case KingType:
		t.NewKing(c)
	case LionType:
		t.NewLion(c)
	case PawnType:
		t.NewPawn(c)
	case RookType:
		t.NewRook(c)
	case UnicornoType:
		t.NewUnicornio(c)
	}
}

func CharToNum(c int) (int, error) {
	if c < 'a' && c >= 'm' {
		return -1, errors.New("out of bounds")
	}
	return c - 'a', nil
}

func NumToChar(n int) (int, error) {
	if n < 0 && n >= 12 {
		return -1, errors.New("out of bounds")
	}
	return n + 'a', nil
}

func (c Color) String() string {
	if c == White {
		return "white"
	} else {
		return "black"
	}
}

func CreateEmptyBoardGame() *Game {
	b := [][]Tile{}
	for y := 1; y <= 12; y++ {
		row := []Tile{}
		for x := 0; x < 12; x++ {
			c, err := NumToChar(x)
			if err != nil {
				log.Fatalf("Error when initilazing game %v", err)
			}
			row = append(row, Tile{X: c, Y: y})
		}
		b = append(b, row)
	}

	return &Game{
		Board: b,
		Turn:  White,
	}
}

func InitilaizeGame() *Game {
	b := [][]Tile{}

	for y := 1; y <= 12; y++ {
		row := []Tile{}
		for x := 0; x < 12; x++ {
			c, err := NumToChar(x)
			if err != nil {
				log.Fatalf("Error when initilazing game %v", err)
			}
			row = append(row, Tile{X: c, Y: y})
		}
		b = append(b, row)
	}

	// initialize white
	for y := 0; y < 12; y++ {
		b[3][y].NewPawn(White)
	}
	b[0][0].NewRook(White)
	b[0][11].NewRook(White)
	b[0][1].NewLion(White)
	b[0][10].NewLion(White)
	b[0][2].NewUnicornio(White)
	b[0][9].NewUnicornio(White)
	b[0][4].NewCrocodile(White)
	b[0][7].NewCrocodile(White)
	b[0][3].NewGiraffe(White)
	b[0][8].NewGiraffe(White)
	b[0][5].NewAanca(White)
	b[0][6].NewKing(White)

	// initialize black
	for y := 0; y < 12; y++ {
		b[8][y].NewPawn(Black)
	}
	b[11][0].NewRook(Black)
	b[11][11].NewRook(Black)
	b[11][1].NewLion(Black)
	b[11][10].NewLion(Black)
	b[11][2].NewUnicornio(Black)
	b[11][9].NewUnicornio(Black)
	b[11][4].NewCrocodile(Black)
	b[11][7].NewCrocodile(Black)
	b[11][3].NewGiraffe(Black)
	b[11][8].NewGiraffe(Black)
	b[11][5].NewAanca(Black)
	b[11][6].NewKing(Black)

	g := Game{
		Turn:  White, // white's turn is true, black's turn is false
		Board: b,
	}

	return &g
}

// x is char, y is number. standart notation
func (g *Game) GetTile(x, y int) *Tile {
	rx, err := CharToNum(x)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return &(g.Board[y-1][rx])
}

func (g *Game) changeTurn() {
	if g.Turn == White {
		g.Turn = Black
	} else {
		g.Turn = White
	}
}

// GAFENotation stands for Grand Acedrex Notation
// It extends FEN for grand acedrex variant
// https://en.wikipedia.org/wiki/Grand_Acedex
func (g *Game) GAFENotation() string {
	toReturn := ""

	// start from top
	for y := 11; y >= 0; y-- {
		acc := 0
		for x := 0; x < 12; x++ {
			p := g.Board[y][x].Piece
			if p == nil {
				acc++
			} else if acc > 0 {
				toReturn += strconv.Itoa(acc)
				acc = 0
			} else {
				toReturn += p.GAFENotation()
			}
		}
		if acc > 0 {
			toReturn += strconv.Itoa(acc)
			acc = 0
		}
		// last '/' is not printed
		if y != 0 {
			toReturn += "/"
		}
	}

	toReturn += " "
	// push turn
	if g.Turn == White {
		toReturn += "w"
	} else {
		toReturn += "b"
	}

	toReturn += " "
	// push king's ability to move 2 squares (denoted by J or j respectively)
	var whiteKing King
	var blackKing King
	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			if (whiteKing == King{}) && (blackKing == King{}) {
				break
			}
			p := g.Board[y][x].Piece
			if p == nil {
				continue
			} else if p.GetType() == KingType {
				if p.GetColor() == White {
					whiteKing = p.(King)
				} else {
					blackKing = p.(King)
				}
			}
		}
	}
	if !whiteKing.HasMoved {
		toReturn += "J"
	}
	if !blackKing.HasMoved {
		toReturn += "j"
	}
	if whiteKing.HasMoved && blackKing.HasMoved {
		toReturn += "-"
	}

	return toReturn
}

func GAFENPieceToType(piece rune) (PieceType, Color, error) {
	switch piece {
	case 'P':
		return PawnType, White, nil
	case 'p':
		return PawnType, Black, nil
	case 'R':
		return RookType, White, nil
	case 'r':
		return RookType, Black, nil
	case 'C':
		return CrocodileType, White, nil
	case 'c':
		return CrocodileType, Black, nil
	case 'G':
		return GiraffeType, White, nil
	case 'g':
		return GiraffeType, Black, nil
	case 'L':
		return LionType, White, nil
	case 'l':
		return LionType, Black, nil
	case 'U':
		return UnicornoType, White, nil
	case 'u':
		return UnicornoType, Black, nil
	case 'A':
		return AancaType, White, nil
	case 'a':
		return AancaType, Black, nil
	case 'K':
		return KingType, White, nil
	case 'k':
		return KingType, Black, nil
	default:
		return 0, 0, fmt.Errorf("unknown piece %v", piece)
	}
}

func FromGAFENotation(notation string) *Game {
	b := [][]Tile{}
	for y := 1; y <= 12; y++ {
		row := []Tile{}
		for x := 0; x < 12; x++ {
			c, err := NumToChar(x)
			if err != nil {
				log.Fatalf("Error when initilazing game %v", err)
			}
			row = append(row, Tile{X: c, Y: y})
		}
		b = append(b, row)
	}
	g := &Game{Board: b}

	row, col := 0, 0

	for _, char := range notation {
		digitAcc := ""
		if unicode.IsDigit(char) {
			digitAcc += string(char)
			continue
		} else if digitAcc != "" {
			digit, err := strconv.Atoi(digitAcc)
			if err != nil {
				log.Fatalf("Error when initilazing game %v", err)
			}
			for i := 0; i < digit; i++ {
				b[row][col].Piece = nil
				col++
				if col == 12 {
					row++
					col = 0
				}
			}
			digitAcc = ""
			continue
		}

		if char == 'w' {
			g.Turn = White
		} else if char == 'b' {
			g.Turn = Black
		} else if char == 'J' {
			g.WhiteKingMoved = false
		} else if char == 'j' {
			g.BlackKingMoved = false
		} else if unicode.IsLetter(char) {
			pType, color, err := GAFENPieceToType(char)
			if err != nil {
				log.Fatalf("Error when initilazing game %v", err)
			}
			b[row][col].NewPiece(pType, color)
		} else if char == '/' {
			row++
			col = 0
		} else if char == '-' {
			g.WhiteKingMoved = true
			g.BlackKingMoved = true
		}

	}

	return nil
}

// does not check turn order
func (g *Game) UnrestrictedMove(sxc, sy, exc, ey int) {
	p := g.GetTile(sxc, sy).Piece
	g.GetTile(exc, ey).NewPiece(p.GetType(), p.GetColor())
	g.GetTile(sxc, sy).Piece = nil
}

func (g *Game) Move(sx, sy, ex, ey string) error {
	// check inputs
	sxi := int(sx[0])
	if sxi < 'a' || sxi > 'l' {
		return errors.New("invalid input start_X")
	}
	exi := int(ex[0])
	if exi < 'a' || exi > 'l' {
		return errors.New("invalid input end_X")
	}
	syi, err := strconv.Atoi(sy)
	if err != nil {
		return err
	}
	if syi < 1 || syi > 12 {
		return errors.New("invalid input start_Y")
	}
	eyi, err := strconv.Atoi(ey)
	if err != nil {
		return err
	}
	if eyi < 1 || eyi > 12 {
		return errors.New("invalid input end_Y")
	}

	tile := g.GetTile(sxi, syi)
	if tile == nil {
		return errors.New("start tile is nil")
	}
	if tile.Piece == nil {
		return errors.New("start tile has no piece")
	}

	g.UnrestrictedMove(sxi, syi, exi, eyi)
	return nil
}

func (g *Game) AttemptMove(sxc, sy, exc, ey int) (*Game, error) {

	// TODO: implement promotion

	sx, err := CharToNum(sxc)
	if err != nil {
		log.Fatalf(err.Error())
	}
	ex, err := CharToNum(exc)
	if err != nil {
		log.Fatalf(err.Error())
	}

	attemptedBoard := make([][]Tile, len(g.Board))
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
	case AancaType: // moves 1 tile diagonally and if wanted like a rook
		// TODO: check if there are any pieces in between
		if (AbsInt(sx-ex) == 1 && AbsInt(sy-ey) >= 1) || (AbsInt(sx-ex) >= 1 && AbsInt(sy-ey) == 1) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.changeTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("aanca can't move like that")
		}
	case CrocodileType: // moves like a bishop
		// TODO: check if there are any pieces in between
		if AbsInt(sx-ex) == AbsInt(sy-ey) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.changeTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("crocodile can't move like that")
		}
	case GiraffeType: // girafe is 3-2 leaper
		if (AbsInt(sx-ex) == 3 && AbsInt(sy-ey) == 2) || (AbsInt(sx-ex) == 2 && AbsInt(sy-ey) == 3) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.changeTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("giraffe can't move like that")
		}

	case KingType: // moves like a standart king, on its first move can move 2 tiles
		k, ok := attemptedSource.Piece.(King)
		if ok {
			// TODO: tidy up
			// should always be ok
			if (AbsInt(sx-ex) == 1 && (AbsInt(sy-ey) == 1 || AbsInt(sy-ey) == 0)) ||
				(AbsInt(sy-ey) == 1 && (AbsInt(sx-ex) == 1 || AbsInt(sx-ex) == 0)) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				nk, _ := attemptedGame.GetTile(exc, ey).Piece.(King)
				nk.HasMoved = true
				attemptedGame.changeTurn()
				unrestictedMoved = true
			} else if !k.HasMoved && ((AbsInt(sx-ex) == 2 && (AbsInt(sy-ey) == 2 || AbsInt(sy-ey) == 0)) ||
				(AbsInt(sy-ey) == 2 && (AbsInt(sx-ex) == 2 || AbsInt(sx-ex) == 0))) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				nk, _ := attemptedGame.GetTile(exc, ey).Piece.(King)
				nk.HasMoved = true
				attemptedGame.changeTurn()
				unrestictedMoved = true
			}
		}

	case LionType: // 3-1 leaper and 3-0 leaper
		if (AbsInt(sx-ex) == 3 && (AbsInt(sy-ey) == 1 || AbsInt(sy-ey) == 0)) ||
			(AbsInt(sy-ey) == 3 && (AbsInt(sx-ex) == 1 || AbsInt(sx-ex) == 0)) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.changeTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("lion can't move like that")
		}

	case PawnType: // no initial jump, moves like a standart pawn
		// check if the pawn is moving foward 1 tile
		if ey == sy+1 && (g.Turn == White) {
			if (ex == sx) && (attemptedDestination.Piece == nil) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				attemptedGame.changeTurn()
				unrestictedMoved = true

			} else if AbsInt(sx-ex) == 1 && attemptedDestination.Piece != nil && attemptedDestination.Piece.GetColor() != g.Turn {
				// captures
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				attemptedGame.changeTurn()
				unrestictedMoved = true
			} else {
				return g, errors.New("pawns can't move that way")
			}
			// check if pawn is moving forward 1 tile but for black
		} else if ey == sy-1 && (g.Turn == Black) {
			if (ex == sx) && (attemptedDestination.Piece == nil) {
				attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
				attemptedGame.changeTurn()
				unrestictedMoved = true
			} else {
				return g, errors.New("pawns can't move that way")
			}
		} else {
			return g, errors.New("pawn must move forward vertically")
		}

	case RookType: // moves like a modern rook
		if (AbsInt(sx-ex) == 0 && AbsInt(sy-ey) >= 1) ||
			(AbsInt(sx-ex) >= 1 && AbsInt(sy-ey) == 0) {
			attemptedGame.UnrestrictedMove(sxc, sy, exc, ey)
			attemptedGame.changeTurn()
			unrestictedMoved = true
		} else {
			return g, errors.New("rook can't move like that")
		}
	case UnicornoType: // most complex piece

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
