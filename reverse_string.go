package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	inRunes:= []rune(input)
	outRunes := make([]rune,0,len(inRunes))
	for i:= len(inRunes)-1; i >= 0 ; i-- {
		outRunes = append(outRunes, inRunes[i])
	}
	return string(outRunes)
}
