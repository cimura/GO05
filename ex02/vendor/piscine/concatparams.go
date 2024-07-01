package piscine

func ConcatParams(args []string) string {
	var constr string
	for i, _ := range args {
		if i > 0 {
			constr += "\n"
		}
		constr += args[i]
	}
	return constr
}