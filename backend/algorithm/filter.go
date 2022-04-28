package algorithm

import (
	"regexp"
	"strings"
)

func reverse(str []string) []string {
	for i := 0; i < len(str)/2; i++ {
		j := len(str) - i - 1
		str[i], str[j] = str[j], str[i]
	}
	return str
}

/* regrex for filter data history */
func Filter(s string) map[string]string {
	m := make(map[string]string)
	var tanggal string
	var namaPenyakit string
	bulan := map[string]string{
		"JANUARI" : "01", 
		"FEBRUARI" : "02",
		"MARET" : "03",
		"APRIL" : "04",
		"MEI" : "05",
		"JUNI" : "06",
		"JULI" : "07",
		"AGUSTUS" : "08",
		"SEPTEMBER" : "09",
		"OKTOBER" : "10",
		"NOVEMBER" : "11",
		"DESEMBER" : "12",
	}
	regrexpBulan := `(JANUARI|FEBRUARI|MARET|APRIL|MEI|JUNI|JULI|AGUSTUS|SEPTEMBER|OKTOBER|NOVEMBER|DESEMBER)`
	regrexpPenyakit := regexp.MustCompile(`\b\w+\b$`)
	formatTanggal1 := regexp.MustCompile(`(3[01]|[12][0-9]|0?[1-9]) ` + regrexpBulan + ` \d{4}`) // 1 JANUARI 2020
	formatTanggal2 := regexp.MustCompile(`(3[01]|[12][0-9]|0?[1-9])[\/\-]((0?[0-9])|((1)[0-2]))[\/\-]\d{4}`) // 31/12/2019 or 31-12-2019

	// Fix string format
	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(s, " ") // Mengganti spasi yang lebih dari satu dengan satu spasi
	s = strings.ToUpper(s) // Membuat string menjadi Upper Case

	// Mengambil string tanggal
	if (formatTanggal1.FindString(s) != "") {
		tanggal = formatTanggal1.FindString(s)
		formatBulan := regexp.MustCompile(regrexpBulan)
		namaBulan := formatBulan.FindString(tanggal)
		tanggal = formatBulan.ReplaceAllString(tanggal, bulan[namaBulan])
		tanggal = strings.Replace(tanggal, " ", "-", -1)
	} else if (formatTanggal2.FindString(s) != "") {
		tanggal = formatTanggal2.FindString(s)
		formatSplit := regexp.MustCompile(`[\/\-]`)
		tanggal = strings.Join(formatSplit.Split(tanggal, -1), "-")
	}

	s = strings.Replace(s, tanggal, "", -1)
	s = strings.Trim(s, " ")

	namaPenyakit = regrexpPenyakit.FindString(s)
	m["tanggal"] = tanggal
	m["penyakit"] = namaPenyakit
	return m
}