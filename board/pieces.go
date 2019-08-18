package board

//Piece is an enum of the 12 piece typeps
type Piece int

const (
	wp Piece = 1 + iota
	wn
	wb
	wr
	wq
	wk

	bp
	bn
	bb
	br
	bq
	bk
)