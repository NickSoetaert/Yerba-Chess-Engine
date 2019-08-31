package board

/*
Board represents one possible board orientation.
The a1 square is bit position 0, b2 = 1,..., g8 = 62, h8 = 63
To figure out if , AND together with White or Black
*/
type Board struct {
	Pawns, Knights, Bishops, Rooks, Queens, Kings, White, Black Piece
}

//Piece represents the position for one type of piece
type Piece uint64

/*
BinaryBoard is a 64-bit mask for each tile. Popcount of a tile is always 1.
For example:
A1 = 0000000000000000000000000000000000000000000000000000000000000001
H8 = 1000000000000000000000000000000000000000000000000000000000000000

bit-position is defined as follows:

    8: 56 57 58 59 60 61 62 63
    7: 48 49 50 51 52 53 54 55
    6: 40 41 42 43 44 45 46 47
    5: 32 33 34 35 36 37 38 39
    4: 24 25 26 27 28 29 30 31
    3: 16 17 18 19 20 21 22 23
    2: 8  9  10 11 12 13 14 15
	1: 0  1  2  3  4  5  6  7
	   -----------------------
       A  B  C  D  E  F  G  H

*/
type BinaryBoard uint64

const (
	A1 BinaryBoard = 1 << iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

const (
	AFile = A1 | A2 | A3 | A4 | A5 | A6 | A7 | A8
	BFile = B1 | B2 | B3 | B4 | B5 | B6 | B7 | B8
	CFile = C1 | C2 | C3 | C4 | C5 | C6 | C7 | C8
	DFile = D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8
	EFile = E1 | E2 | E3 | E4 | E5 | E6 | E7 | E8
	FFile = F1 | F2 | F3 | F4 | F5 | F6 | F7 | F8
	GFile = G1 | G2 | G3 | G4 | G5 | G6 | G7 | G8
	HFile = H1 | H2 | H3 | H4 | H5 | H6 | H7 | H8
)

const (
	FirstRank   = A1 | B1 | C1 | D1 | E1 | F1 | G1 | H1
	SecondRank  = A2 | B2 | C2 | D2 | E2 | F2 | G2 | H2
	ThirdRank   = A3 | B3 | C3 | D3 | E3 | F3 | G3 | H3
	FourthRank  = A4 | B4 | C4 | D4 | E4 | F4 | G4 | H4
	FifthRank   = A5 | B5 | C5 | D5 | E5 | F5 | G5 | H5
	SixthRank   = A6 | B6 | C6 | D6 | E6 | F6 | G6 | H6
	SeventhRank = A7 | B7 | C7 | D7 | E7 | F7 | G7 | H7
	EighthRank  = A8 | B8 | C8 | D8 | E8 | F8 | G8 | H8
)

//diagonals
const ()

//PieceType - pawn, knight, etc.
type PieceType uint8

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

//White if true means the piece or tile is white. False means black.
type White bool

const (
	white White = true
	black White = false
)
