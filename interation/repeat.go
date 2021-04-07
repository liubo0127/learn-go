package interation

func Repeat(character string, cnt int) string {
	repeated := ""
	for i := 0; i< cnt; i++ {
		repeated += character
	}
	return repeated
}
