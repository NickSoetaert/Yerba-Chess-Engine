package main

/*
Board represents one possible board orientation.
The a1 square is bit position 0, b2 = 1,..., g8 = 62, h8 = 63
To figure out if , AND together with White or Black
*/
type Board struct {
	Pawns, Knights, Bishops, Rooks, Queens, Kings, White, Black uint64
	Move                                                        IsWhite
}

/*
bit-position is defined as follows, by order of magnitude:

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

ex:
A1 = 0000000000000000000000000000000000000000000000000000000000000001
H8 = 1000000000000000000000000000000000000000000000000000000000000000

*/


//Binary representation of single game tiles
const (
	A1 uint64 = 1 << iota
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

//Vertical lines
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

//Horizontal lines
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

const (
	QueenSide = AFile | BFile | CFile | DFile
	KingSide  = EFile | FFile | GFile | HFile
)

const (
	emptyBoard = uint64(0)
	wholeBoard = ^uint64(0)
)

// Each index represents every square that you can possibly be stopped by a blocker at, from the given index.
// Because a piece on the edge of the board does stop you from moving further, the edges of the board are not included.
// Used with magic bitboards.
var rookBlockerMask = [64]uint64{
	//A       		 	 B       			 C       			 D					E		 			F					 G					H
	0x000101010101017E, 0x000202020202027C, 0x000404040404047A, 0x0008080808080876, 0x001010101010106E, 0x002020202020205E, 0x004040404040403E, 0x008080808080807E, //1
	0x0001010101017E00, 0x0002020202027C00, 0x0004040404047A00, 0x0008080808087600, 0x0010101010106E00, 0x0020202020205E00, 0x0040404040403E00, 0x0080808080807E00, //2
	0x00010101017E0100, 0x00020202027C0200, 0x00040404047A0400, 0x0008080808760800, 0x00101010106E1000, 0x00202020205E2000, 0x00404040403E4000, 0x00808080807E8000, //3
	0x000101017E010100, 0x000202027C020200, 0x000404047A040400, 0x0008080876080800, 0x001010106E101000, 0x002020205E202000, 0x004040403E404000, 0x008080807E808000, //4
	0x0001017E01010100, 0x0002027C02020200, 0x0004047A04040400, 0x0008087608080800, 0x0010106E10101000, 0x0020205E20202000, 0x0040403E40404000, 0x0080807E80808000, //5
	0x00017E0101010100, 0x00027C0202020200, 0x00047A0404040400, 0x0008760808080800, 0x00106E1010101000, 0x00205E2020202000, 0x00403E4040404000, 0x00807E8080808000, //6
	0x007E010101010100, 0x007C020202020200, 0x007A040404040400, 0x0076080808080800, 0x006E101010101000, 0x005E202020202000, 0x003E404040404000, 0x007E808080808000, //7
	0x7E01010101010100, 0x7C02020202020200, 0x7A04040404040400, 0x7608080808080800, 0x6E10101010101000, 0x5E20202020202000, 0x3E40404040404000, 0x7E80808080808000, //8
}

var bishopBlockerMask = [64]uint64{
	//A					B					C					D					E					F					G					H
	0x0040201008040200, 0x0000402010080400, 0x0000004020100A00, 0x0000000040221400, 0x0000000002442800, 0x0000000204085000, 0x0000020408102000, 0x0002040810204000, //1
	0x0020100804020000, 0x0040201008040000, 0x00004020100A0000, 0x0000004022140000, 0x0000000244280000, 0x0000020408500000, 0x0002040810200000, 0x0004081020400000, //2
	0x0010080402000200, 0x0020100804000400, 0x004020100A000A00, 0x0000402214001400, 0x0000024428002800, 0x0002040850005000, 0x0004081020002000, 0x0008102040004000, //3
	0x0008040200020400, 0x0010080400040800, 0x0020100A000A1000, 0x0040221400142200, 0x0002442800284400, 0x0004085000500800, 0x0008102000201000, 0x0010204000402000, //4
	0x0004020002040800, 0x0008040004081000, 0x00100A000A102000, 0x0022140014224000, 0x0044280028440200, 0x0008500050080400, 0x0010200020100800, 0x0020400040201000, //5
	0x0002000204081000, 0x0004000408102000, 0x000A000A10204000, 0x0014001422400000, 0x0028002844020000, 0x0050005008040200, 0x0020002010080400, 0x0040004020100800, //6
	0x0000020408102000, 0x0000040810204000, 0x00000A1020400000, 0x0000142240000000, 0x0000284402000000, 0x0000500804020000, 0x0000201008040200, 0x0000402010080400, //7
	0x0002040810204000, 0x0004081020400000, 0x000A102040000000, 0x0014224000000000, 0x0028440200000000, 0x0050080402000000, 0x0020100804020000, 0x0040201008040200, //8
}

// Magic number for use in bitboards.
var rookMagic = [64]uint64{
	//A       		 	 B       			 C       			 D					E		 			F					 G					H
	0x0080001020400080, 0x0040001000200040, 0x0080081000200080, 0x0080040800100080, 0x0080020400080080, 0x0080010200040080, 0x0080008001000200, 0x0080002040800100, //1
	0x0000800020400080, 0x0000400020005000, 0x0000801000200080, 0x0000800800100080, 0x0000800400080080, 0x0000800200040080, 0x0000800100020080, 0x0000800040800100, //2
	0x0000208000400080, 0x0000404000201000, 0x0000808010002000, 0x0000808008001000, 0x0000808004000800, 0x0000808002000400, 0x0000010100020004, 0x0000020000408104, //3
	0x0000208080004000, 0x0000200040005000, 0x0000100080200080, 0x0000080080100080, 0x0000040080080080, 0x0000020080040080, 0x0000010080800200, 0x0000800080004100, //4
	0x0000204000800080, 0x0000200040401000, 0x0000100080802000, 0x0000080080801000, 0x0000040080800800, 0x0000020080800400, 0x0000020001010004, 0x0000800040800100, //5
	0x0000204000808000, 0x0000200040008080, 0x0000100020008080, 0x0000080010008080, 0x0000040008008080, 0x0000020004008080, 0x0000010002008080, 0x0000004081020004, //6
	0x0000204000800080, 0x0000200040008080, 0x0000100020008080, 0x0000080010008080, 0x0000040008008080, 0x0000020004008080, 0x0000800100020080, 0x0000800041000080, //7
	0x00FFFCDDFCED714A, 0x007FFCDDFCED714A, 0x003FFFCDFFD88096, 0x0000040810002101, 0x0001000204080011, 0x0001000204000801, 0x0001000082000401, 0x0001FFFAABFAD1A2, //8
}


var bishopMagic = [64]uint64{
	//A       		 	 B       			 C       			 D					E		 			F					 G					H
	0x0002020202020200, 0x0002020202020000, 0x0004010202000000, 0x0004040080000000, 0x0001104000000000, 0x0000821040000000, 0x0000410410400000, 0x0000104104104000, //1
	0x0000040404040400, 0x0000020202020200, 0x0000040102020000, 0x0000040400800000, 0x0000011040000000, 0x0000008210400000, 0x0000004104104000, 0x0000002082082000, //2
	0x0004000808080800, 0x0002000404040400, 0x0001000202020200, 0x0000800802004000, 0x0000800400A00000, 0x0000200100884000, 0x0000400082082000, 0x0000200041041000, //3
	0x0002080010101000, 0x0001040008080800, 0x0000208004010400, 0x0000404004010200, 0x0000840000802000, 0x0000404002011000, 0x0000808001041000, 0x0000404000820800, //4
	0x0001041000202000, 0x0000820800101000, 0x0000104400080800, 0x0000020080080080, 0x0000404040040100, 0x0000808100020100, 0x0001010100020800, 0x0000808080010400, //5
	0x0000820820004000, 0x0000410410002000, 0x0000082088001000, 0x0000002011000800, 0x0000080100400400, 0x0001010101000200, 0x0002020202000400, 0x0001010101000200, //6
	0x0000410410400000, 0x0000208208200000, 0x0000002084100000, 0x0000000020880000, 0x0000001002020000, 0x0000040408020000, 0x0004040404040000, 0x0002020202020000, //7
	0x0000104104104000, 0x0000002082082000, 0x0000000020841000, 0x0000000000208800, 0x0000000010020200, 0x0000000404080200, 0x0000040404040400, 0x0002020202020200, //8
}


var knightAttacks = [64]uint64{
	//A       		 	 B       			 C       			 D					E		 			F					 G					H
	0x0000000000020400, 0x0000000000050800, 0x00000000000a1100, 0x0000000000142200, 0x0000000000284400, 0x0000000000508800, 0x0000000000a01000, 0x0000000000402000, //1
	0x0000000002040004, 0x0000000005080008, 0x000000000a110011, 0x0000000014220022, 0x0000000028440044, 0x0000000050880088, 0x00000000a0100010, 0x0000000040200020, //2
	0x0000000204000402, 0x0000000508000805, 0x0000000a1100110a, 0x0000001422002214, 0x0000002844004428, 0x0000005088008850, 0x000000a0100010a0, 0x0000004020002040, //3
	0x0000020400040200, 0x0000050800080500, 0x0000000a1100110a, 0x0000142200221400, 0x0000284400442800, 0x0000508800885000, 0x0000a0100010a000, 0x0000402000204000, //4
	0x0002040004020000, 0x0005080008050000, 0x000a1100110a0000, 0x0014220022140000, 0x0028440044280000, 0x0050880088500000, 0x00a0100010a00000, 0x0040200020400000, //5
	0x0204000402000000, 0x0508000805000000, 0x0a1100110a000000, 0x1422002214000000, 0x2844004428000000, 0x5088008850000000, 0xa0100010a0000000, 0x4020002040000000, //6
	0x0400040200000000, 0x0800080500000000, 0x1100110a00000000, 0x2200221400000000, 0x4400442800000000, 0x8800885000000000, 0x100010a000000000, 0x2000204000000000, //7
	0x0004020000000000, 0x0008050000000000, 0x00110a0000000000, 0x0022140000000000, 0x0044280000000000, 0x0088500000000000, 0x0010a00000000000, 0x0020400000000000, //8
}
