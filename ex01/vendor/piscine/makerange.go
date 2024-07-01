package piscine

func MakeRange(min, max int) []int {
	var size int
	if max > min {
		size = max - min
	} else {
		size = 0
	}
	box := make([]int, size)
	for i, _ := range box {
		box[i] += min
		min++
	}
	return box
}