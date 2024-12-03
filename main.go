package main

func isMutant(dna []string) bool {
	N := len(dna)
	if N == 0 || len(dna[0]) != N {
		return false
	}

	sequences := 0

	checkSequence := func(c int) bool {
		sequences += c
		return sequences > 1
	}

	//Iterate over rows and columns

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j+3 < N && checkSequence(checkHorizontal(dna, i, j)) {
				return true
			}
			if i+3 < N && checkSequence(checkVertical(dna, i, j)) {
				return true
			}
			if i+3 < N && j+3 < N && checkSequence(checkDiagonalRight(dna, i, j)) {
				return true
			}
			if i+3 < N && j-3 >= 0 && checkSequence(checkDiagonalLeft(dna, i, j)) {
				return true
			}
		}

	}
	return false
}

func checkHorizontal(dna []string, row, col int) int {
	if dna[row][col] == dna[row][col+1] &&
		dna[row][col] == dna[row][col+2] &&
		dna[row][col] == dna[row][col+3] {
		return 1
	}
	return 0
}

func checkVertical(dna []string, row, col int) int {
	if dna[row][col] == dna[row+1][col] &&
		dna[row][col] == dna[row+2][col] &&
		dna[row][col] == dna[row+3][col] {
		return 1
	}
	return 0
}

func checkDiagonalRight(dna []string, row, col int) int {
	if dna[row][col] == dna[row+1][col+1] &&
		dna[row][col] == dna[row+2][col+2] &&
		dna[row][col] == dna[row+3][col+3] {
		return 1
	}
	return 0
}

func checkDiagonalLeft(dna []string, row, col int) int {
	if dna[row][col] == dna[row+1][col-1] &&
		dna[row][col] == dna[row+2][col-2] &&
		dna[row][col] == dna[row+3][col-3] {
		return 1
	}
	return 0
}

func main() {
	dnaMutant := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	dnaHuman := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}

	println("Is Mutant: ", isMutant(dnaMutant)) // true
	println("Is Human: ", isMutant(dnaHuman))   // false
}
