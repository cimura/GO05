package piscine

func Len(sRune []rune) int {
	count := 0
	for range sRune {
		count++
	}
	return count
}

func SplitWhiteSpaces(s string) []string {
	var strarray []string
	sRune := []rune(s)
	var strRune []rune

	for _, r := range sRune {
		if r != ' ' {
			strRune = append(strRune, r)
		} else if Len(strRune) > 0 {
			strarray = append(strarray, string(strRune))
			strRune = []rune{}
		}
	}
	if Len(strRune) > 0 {
		strarray = append(strarray, string(strRune))
	}
	return strarray
}