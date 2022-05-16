package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	for i := 0; i < len(runes)/2; i++ {
		x := len(runes) - i - 1
		runes[i], runes[x] = runes[x], runes[i]
	}
	for _, v := range runes {
		output += string(v)
	}
	return output
}
