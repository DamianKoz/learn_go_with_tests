package iteration

func Repeat(character string, rounds int) (result string) {
	for i := 0; i < rounds; i++ {
		result += character
	}
	return result
}
