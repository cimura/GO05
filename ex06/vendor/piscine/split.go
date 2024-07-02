package piscine

func Len(sRune []rune) int {
	count := 0
	for range sRune {
		count++
	}
	return count
}

func Split(s, sep string) []string {
	var strarray []string
	sRune := []rune(s)
	sepRune := []rune(sep)
	var retRune []rune
	flag := false

	for indx, sr := range sRune {
		retRune = append(retRune, sRune)	
		if sRune[indx] == sepRune[indx] {
			for _, sp := range sepRune {
				sRune[indx] = sp
				indx++
			}
		}
		strarray = append(strarray, string(sRune))
		sRune = []rune{}
	}
	return strarray
}