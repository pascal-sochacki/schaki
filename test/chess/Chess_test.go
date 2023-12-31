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
