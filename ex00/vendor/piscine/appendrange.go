package piscine

func AppendRange(min, max int) []int {
	var box []int
	for i := 0; i < max - min; i++ {
		box = append(box, min + i)
	}
	return box
}