package solver

// The min and max Unicode code points for valid characters.
const (
	unicodeMinPoint = 0x61  // 'a'
	unicodeMaxPoint = unicodeMinPoint + 26 - 1  // 'z'
)

// ValidString checks whether the given string is valid in this game (true if
// all characters are valid according to ValidChar).
func ValidString(str string) bool {
	for _, char := range str {
		if !ValidChar(char) {
			return false
		}
	}
	return true
}

// ValidChar checks whether the given rune is valid in this game. Only lowercase
// a-z are valid.
func ValidChar(char rune) bool {
	v := int(char)
	return v >= unicodeMinPoint && v <= unicodeMaxPoint
}
