package chess

import (
	"math/bits"
	"strconv"
	"strings"
)

const (
	WHITE_PAWN_START_POSITION BitMap = 0b0000000000000000000000000000000000000000000000001111111100000000
	BLACK_PAWN_START_POSITION BitMap = 0b0000000011111111000000000000000000000000000000000000000000000000
	MAX_POSITION              BitMap = 0b1000000000000000000000000000000000000000000000000000000000000000
	MIN_POSITION              BitMap = 0b0000000000000000000000000000000000000000000000000000000000000001
)

type BitMap uint64

func (receiver *BitMap) String() string {
	builder := strings.Builder{}
	point := MIN_POSITION
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if point&*receiver > 0 {
				builder.WriteString("X")
			} else {
				builder.WriteString(".")
			}
			point = point << 1

		}
		builder.WriteString("\n")
	}
	return builder.String()
}

type ChessBoard struct {
	whitePawns   BitMap
	whiteRooks   BitMap
	whiteKnights BitMap
	whiteBishop  BitMap
	whiteQueen   BitMap
	whiteKing    BitMap

	blackPawns   BitMap
	blackRooks   BitMap
	blackKnights BitMap
	blackBishop  BitMap
	blackQueen   BitMap
	blackKing    BitMap
}

var knightJumps [64]BitMap

// 2  1  0  1  2
// 6  7  8  9 10
//14 15 16 17 18

func init() {
	var i BitMap
	i = 1
	position := 0

	var seventeen BitMap
	seventeen = 0b1111111011111110111111101111111011111110111111101111111011111110

	var ten BitMap
	ten = 0b1111110011111100111111001111110011111100111111001111110011111100

	var fiveteen BitMap
	fiveteen = 0b0111111101111111011111110111111101111111011111110111111101111111

	var six BitMap
	six = 0b0011111100111111001111110011111100111111001111110011111100111111

	for {
		var jumps BitMap

		jumps |= ((i << 17) & seventeen)
		jumps |= (i << 10 & ten)
		jumps |= ((i << 15) & fiveteen)
		jumps |= ((i << 6) & six)

		jumps |= ((i >> 6) & ten)
		jumps |= ((i >> 10) & six)
		jumps |= ((i >> 15) & seventeen)
		jumps |= ((i >> 17) & fiveteen)

		if i == BitMap(MAX_POSITION) {
			break
		}
		knightJumps[position] = jumps
		i = i << 1
		position += 1
	}

}

type Move struct {
}

func NewChessBoard() *ChessBoard {
	return &ChessBoard{
		whitePawns:   WHITE_PAWN_START_POSITION,
		whiteRooks:   0b0000000000000000000000000000000000000000000000000000000010000001,
		whiteKnights: 0b0000000000000000000000000000000000000000000000000000000001000010,
		whiteBishop:  0b0000000000000000000000000000000000000000000000000000000000100100,
		whiteQueen:   0b0000000000000000000000000000000000000000000000000000000000010000,
		whiteKing:    0b0000000000000000000000000000000000000000000000000000000000001000,

		blackPawns:   BLACK_PAWN_START_POSITION,
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

	var currentPoint BitMap
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

func (receiver *ChessBoard) white() BitMap {
	return receiver.whitePawns | receiver.whiteRooks | receiver.whiteKnights | receiver.whiteBishop | receiver.whiteQueen | receiver.whiteKing
}

func (receiver *ChessBoard) black() BitMap {
	return receiver.blackPawns | receiver.blackRooks | receiver.blackKnights | receiver.blackBishop | receiver.blackQueen | receiver.blackKing
}

func (receiver *ChessBoard) all() BitMap {
	return receiver.white() | receiver.black()
}

func (receiver *ChessBoard) String() string {

	all := receiver.all()

	builder := strings.Builder{}

	var current BitMap
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

func findLSBSetBit(num BitMap) uint8 {
	var bitIndex uint8
	bitIndex = 0
	for num > 0 {
		if num&1 == 1 {
			return bitIndex
		}
		num >>= 1
		bitIndex++
	}
	return 0 // If no bit is set
}

func (receiver ChessBoard) GetMoves() []Move {

	result := []Move{}

	pawn_at_starting_pos := receiver.whitePawns & WHITE_PAWN_START_POSITION
	if pawn_at_starting_pos > 0 {
		amount := bits.OnesCount64(uint64(pawn_at_starting_pos))
		for i := 0; i < amount; i++ {
			result = append(result, Move{})
		}
	}

	amount := bits.OnesCount64(uint64(receiver.whitePawns))
	for i := 0; i < amount; i++ {
		result = append(result, Move{})
	}

    tdb(receiver.whiteKnights, func(u uint8) {
		jumps := knightJumps[u] & ^receiver.white()
		amount := bits.OnesCount64(uint64(jumps))
		for i := 0; i < amount; i++ {
			result = append(result, Move{})
		}
    })

    tdb(receiver.whiteRooks, func(u uint8) {

    })


	return result

}

func tdb(pieces BitMap, do func(uint8) ) {

    var current uint8
    for {
		current = findLSBSetBit(pieces)
		if current == 0 {
			break
		}
        do(current)
		pieces = pieces ^ 1<<current
    }
}
