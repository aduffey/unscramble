package solver

// The number of valid characters
const numValidChars = 26

// The code point at which valid characters start (they are consecutive)
const unicodeBasePoint = 0x61

// Check whether the given string is valid in this game (true if all characters
// are valid, i.e. all characters are lowercase a-z).
func ValidString(str string) bool {
	for _, char := range str {
		if !ValidChar(char) {
			return false
		}
	}
	return true
}

// Check whether the given rune is valid in this game. Only lowercase a-z are
// valid.
func ValidChar(char rune) bool {
	v := translate(char)
	return v >= 0 && v < numValidChars
}

// Translate the given rune to an int between 0 and numValidChars-1 (inclusive).
// Use for indexing, etc.
func translate(char rune) int {
	return int(char) - unicodeBasePoint
}
