package iteration

// Repeat function takes a character and a repeat count, then repeats
// the character till the count
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
