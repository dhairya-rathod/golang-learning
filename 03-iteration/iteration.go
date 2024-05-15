package iteration

func Repeat(chr string) string {
	output := ""
	for i := 0; i < 5; i++ {
		output += chr
	}
	return output
}
