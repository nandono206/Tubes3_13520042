package algorithm

/* Realisasi Border Function */
// Menerima input pattern p dan mengembalikan border function untuk setiap indeks j pada p
func ComputeFail(p string) []int {
	border := make([]int, len(p))
	for j := 0; j < len(p); j++ {
		/* size of the largest prefix of P[0..j-1] that is also a suffix of P[1..j-1] */
		if j <= 1 {
			border[j] = 0
		} else {
			k := j - 1
			found := false
			for !found && k > 0 {
				if p[0:k] == p[j-k:j] {
					found = true
				} else {
					k--
				}
			}
			border[j] = k
		}
	}
	return border
}

/* Realisasi algoritma kmp */
// Menerima input teks t dan pattern p.
// Mengembalikan nilai boolean yang menyatakan apakah pattern p ada pada teks t
func KmpMatch(t string, p string) bool {
	n := len(t)
	m := len(p)

	fail := ComputeFail(p)

	i := 0
	j := 0

	for i < n {
		if p[j] == t[i] {
			if j == m-1 {
				return true // match
			}
			i++
			j++
		} else if j > 0 {
			j = fail[j-1]
		} else {
			i++
		}
	}

	return false
}