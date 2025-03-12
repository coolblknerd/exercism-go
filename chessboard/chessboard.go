package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (total int) {
	for _, v := range cb[file] {
		if v {
			total++
		}
	}

	return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (total int) {
	for _, v := range cb {
		if rank > 0 && rank <= len(v) {
			if v[rank-1] {
				total++
			}
		}
	}

	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (total int) {
	for _, _ = range cb {
		total += len(cb)
	}

	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (total int) {
	for _, v := range cb {
		for _, vv := range v {
			if vv {
				total++
			}
		}
	}

	return
}
