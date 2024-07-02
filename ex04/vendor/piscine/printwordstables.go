package piscine

import "ft"

func Printstr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	for _, str := range a {
		Printstr(str)
	}
}