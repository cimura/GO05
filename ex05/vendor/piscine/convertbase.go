package piscine

func Len(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func Power(nbr int, power int) int {
	if power < 0 {
		return -1
	} else if power == 0 {
		return 1
	} else {
		return nbr * Power(nbr, power-1)
	}
	
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var nbrRune = []rune(nbr)
	nbrLen := Len(nbr)
	baseFromLen := Len(baseFrom)
	baseToLen := Len(baseTo)
	base10 := 0
	count := 0

	for i := nbrLen-1; i >= 0; i-- {
		base10 += int(nbrRune[i] - '0') * Power(baseFromLen, count)
		count++
	}
	result := ""
	Revresult := ""
	for base10 > 0 {
		result += string(base10%baseToLen + '0')
		base10 /= baseToLen
	}
	resultLen := Len(result)
	var resultRune = []rune(result)
	for i := resultLen-1; i >= 0; i-- {
		Revresult += string(resultRune[i])
	}
	return Revresult
}