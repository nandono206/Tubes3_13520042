package algorithm

import (
	"math"
)

/* Realisasi algoritma Longest Common Subsequence (lcs) */
func LCS(s1 string, s2 string) int {
	// s1 akan menjadi baris
	// s2 akan menjadi kolom

	matrix := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		matrix[i] = make([]int, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = int(math.Max(float64(matrix[i-1][j]), float64(matrix[i][j-1])))
			}
		}
	}

	return matrix[len(s1)][len(s2)]
}