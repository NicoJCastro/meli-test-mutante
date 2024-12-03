package models

type DNA struct {
	Sequences []string `json:"dna"`
}

func IsMutant(dna []string) bool {
	N := len(dna)
	sequences := 0

	checkSequences := func(c int) bool {
		sequences += c
		return sequences > 1
	}

	for i := 0; i < N && sequences <= 1; i++ {
		for j := 0; j < N && sequences <= 1; j++ {
			if j+3 < N && checkSequences(hasSequence(dna, i, j, 0, 1)) {
				return true
			}
			if i+3 < N && checkSequences(hasSequence(dna, i, j, 1, 0)) {
				return true
			}
			if i+3 < N && j+3 < N && checkSequences(hasSequence(dna, i, j, 1, 1)) {
				return true
			}
			if i+3 < N && j-3 >= 0 && checkSequences(hasSequence(dna, i, j, 1, -1)) {
				return true
			}
		}
	}
	return false
}

func hasSequence(dna []string, row, col, dRow, dCol int) int {
	base := dna[row][col]
	for k := 1; k < 4; k++ {
		if dna[row+k*dRow][col+k*dCol] != base {
			return 0
		}
	}
	return 1
}
