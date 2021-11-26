package iteration

func Repeat(char string) string {
	var repeat string
	for i := 0; i < 5; i++ {
		repeat += char
	}
	return repeat
}
