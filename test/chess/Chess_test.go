package chess

import (
	"testing"

	"github.com/pascal-sochacki/schaki/pkg/chess"
)

func TestNewBoard(t *testing.T) {
	chessBoard := chess.NewChessBoard()

	want := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
	is := chessBoard.String()

	if is != want {
		t.Fatalf("chessBoard: String()  = \n\t%q,\n want = \n\t%q", is, want)
	}
}

func TestConstructFromFenString(t *testing.T) {
	boards := []string{
		"rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBN1",
		"rnbqkbnr/pp1ppppp/8/2p5/4P3/5N2/PPPP1PPP/RNBQKB1R",
		"r1b1kbnr/pppp1ppp/2n5/8/3P4/8/PPP2PPP/R1BQKBNR",
		"rnbqkbnr/ppppp3/8/4P3/3p4/8/PPP2PPP/RNBQKBNR",
		"r1bqkb1r/pp1ppppp/2n5/3P4/3p4/8/PPP2PPP/R1BQKBNR",
		"r1b1kbnr/ppp1pppp/2n5/8/3P4/8/PPP2PPP/R1BQKBNR",
		"r1bqkbnr/pp1pp1pp/2n5/3Pp3/8/8/PPP2PPP/R1BQKBNR",
		"r1bqkb1r/pppp1ppp/2n5/4P3/3p4/8/PPP2PPP/R1BQKBNR",
		"rnbqkbnr/pp1pp1pp/8/3Pp3/8/8/PPP2PPP/R1BQKBNR",
		"r1bqkb1r/pp1ppppp/2n5/3P4/3p4/8/PPP2PPP/R1BQKBNR",
		"rnbqkb1r/pp1ppppp/8/3P4/3p4/8/PPP2PPP/RNBQKBNR",
		"r1bqkbnr/ppppp1pp/8/4P3/3p4/8/PPP2PPP/R1BQKBNR",
	}

	for _, test := range boards {
		board := chess.FromFenString(test)
		is := board.String()

		if is != test {
			t.Fatalf("chessBoard: FromFenString()  = \n\t%q,\n want = \n\t%q", is, test)
		}
	}
}

type moveTest struct {
	input  string
	amount int
}

func TestChessMoves(t *testing.T) {

	tests := []moveTest{
		{amount: 20, input: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"},
		{amount: 1, input: "8/8/8/4P3/8/8/8/8"},
		{amount: 8, input: "8/8/8/4N3/8/8/8/8"},
		{amount: 14, input: "8/8/8/8/8/8/8/R7"},
	}

	for _, test := range tests {
		board := chess.FromFenString(test.input)
		moves := board.GetMoves()
		is := len(moves)

		if is != test.amount {
			t.Fatalf("chessBoard: GetMoves()  = %d, want = %d", is, test.amount)
		}

	}
}
