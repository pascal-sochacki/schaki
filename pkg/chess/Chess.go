package chess

import (
	"strconv"
	"strings"
)

type ChessBoard struct {
	whitePawns   uint64
	whiteRooks   uint64
	whiteKnights uint64
	whiteBishop  uint64
	whiteQueen   uint64
	whiteKing    uint64

	blackPawns   uint64
	blackRooks   uint64
	blackKnights uint64
	blackBishop  uint64
	blackQueen   uint64
	blackKing    uint64
}

type Move struct {
}

func NewChessBoard() *ChessBoard {
	return &ChessBoard{
		whitePawns:   0b0000000000000000000000000000000000000000000000001111111100000000,
		whiteRooks:   0b0000000000000000000000000000000000000000000000000000000010000001,
		whiteKnights: 0b0000000000000000000000000000000000000000000000000000000001000010,
		whiteBishop:  0b0000000000000000000000000000000000000000000000000000000000100100,
		whiteQueen:   0b0000000000000000000000000000000000000000000000000000000000010000,
		whiteKing:    0b0000000000000000000000000000000000000000000000000000000000001000,

		blackPawns:   0b0000000011111111000000000000000000000000000000000000000000000000,
		blackRooks:   0b1000000100000000000000000000000000000000000000000000000000000000,
		blackKnights: 0b0100001000000000000000000000000000000000000000000000000000000000,
		blackBishop:  0b0010010000000000000000000000000000000000000000000000000000000000,
		blackQueen:   0b0001000000000000000000000000000000000000000000000000000000000000,
		blackKing:    0b0000100000000000000000000000000000000000000000000000000000000000,
	}
}

func FromFenString(input string) *ChessBoard {

	for i := 2; i < 9; i++ {
		input = strings.ReplaceAll(input, strconv.Itoa(i), strings.Repeat("1", i))
	}
	result := ChessBoard{}

	var currentPoint uint64
	currentPoint = 1 << 63
	for i := 0; i < len(input); i++ {

		current := rune(input[i])

		if current == '/' {
			continue
		}

		switch {
		case current == '1':

		case current == 'P':
			result.whitePawns |= currentPoint
		case current == 'R':
			result.whiteRooks |= currentPoint
		case current == 'N':
			result.whiteKnights |= currentPoint
		case current == 'B':
			result.whiteBishop |= currentPoint
		case current == 'Q':
			result.whiteQueen |= currentPoint
		case current == 'K':
			result.whiteKing |= currentPoint

		case current == 'p':
			result.blackPawns |= currentPoint
		case current == 'r':
			result.blackRooks |= currentPoint
		case current == 'n':
			result.blackKnights |= currentPoint
		case current == 'b':
			result.blackBishop |= currentPoint
		case current == 'q':
			result.blackQueen |= currentPoint
		case current == 'k':
			result.blackKing |= currentPoint
		}

		currentPoint = currentPoint >> 1

	}

	return &result
}

func (receiver *ChessBoard) white() uint64 {
	return receiver.whitePawns | receiver.whiteRooks | receiver.whiteKnights | receiver.whiteBishop | receiver.whiteQueen | receiver.whiteKing
}

func (receiver *ChessBoard) black() uint64 {
	return receiver.blackPawns | receiver.blackRooks | receiver.blackKnights | receiver.blackBishop | receiver.blackQueen | receiver.blackKing
}

func (receiver *ChessBoard) all() uint64 {
	return receiver.white() | receiver.black()
}

func (receiver *ChessBoard) String() string {

	all := receiver.all()

	builder := strings.Builder{}

	var current uint64
	current = 1 << 63
	for y := 0; y < 8; y++ {

		empty := 0
		for x := 0; x < 8; x++ {

			if current&all > 0 && empty > 0 {
				builder.WriteString(strconv.Itoa(empty))
			}

			empty += 1
			if current&all > 0 {
				empty = 0
			}

			switch {

			case current&receiver.whitePawns > 0:
				builder.WriteString("P")
			case current&receiver.whiteRooks > 0:
				builder.WriteString("R")
			case current&receiver.whiteKnights > 0:
				builder.WriteString("N")
			case current&receiver.whiteBishop > 0:
				builder.WriteString("B")
			case current&receiver.whiteQueen > 0:
				builder.WriteString("Q")
			case current&receiver.whiteKing > 0:
				builder.WriteString("K")

			case current&receiver.blackPawns > 0:
				builder.WriteString("p")
			case current&receiver.blackRooks > 0:
				builder.WriteString("r")
			case current&receiver.blackKnights > 0:
				builder.WriteString("n")
			case current&receiver.blackKnights > 0:
				builder.WriteString("n")
			case current&receiver.blackBishop > 0:
				builder.WriteString("b")
			case current&receiver.blackQueen > 0:
				builder.WriteString("q")
			case current&receiver.blackKing > 0:
				builder.WriteString("k")
			}

			current = current >> 1

		}
		if empty != 0 {
			builder.WriteString(strconv.Itoa(empty))
		}
		if y != 7 {
			builder.WriteString("/")
		}
	}

	return builder.String()
}

func (receiver *ChessBoard) GetMoves() []Move {
	return []Move{}

}
