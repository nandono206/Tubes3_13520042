package algorithm

import (
	"regexp"
)

/* sanitasi input menggunakan regex untuk memastikan bahwa masukan
merupakan sequence DNA yang valid (tidak boleh ada huruf kecil, tidak boleh
ada huruf selain AGCT, dan tidak ada spasi) */
func IsDNAValid(dna string) bool {
	regularExpression := regexp.MustCompile(`[^ACGT]`)
	result := !regularExpression.Match([]byte(dna))
	return result
}