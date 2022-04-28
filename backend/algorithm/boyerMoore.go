package algorithm

import "math"

/* Menghitung fungsi Last(x) untuk tiap karakter pada DNA */
func BuildLast(p string) []int {
	last := make([]int, 4) // A = last[0], C = last[1], G = last[2], T = last[3]

	// Inisialisasi array last
	for i := 0; i < 4; i++ {
		last[i] = -1
	}

	for i := 0; i < len(p); i++ {
		switch p[i] {
		case 'A':
			last[0] = i
		case 'C':
			last[1] = i
		case 'G':
			last[2] = i
		case 'T':
			last[3] = i
		}
	}
	return last
}

/* Realisasi algoritma Boyer Moore */
func BmMatch(t string, p string) bool {
	last := BuildLast(p)
	n := len(t)
	m := len(p)
	i := m - 1

	if i > n-1 {
		return false // panjang pattern > panjang text
	}

	j := m - 1
	for i < n {
		if p[j] == t[i] {
			if j == 0 {
				return true // match
			} else {
				i--
				j--
			}
		} else {
			var lo int
			switch t[i] {
			case 'A':
				lo = last[0]
			case 'B':
				lo = last[1]
			case 'C':
				lo = last[2]
			case 'D':
				lo = last[3]
			}
			i = i + m - int(math.Min(float64(j), float64(1+lo)))
			j = m - 1
		}
	}
	return false
}