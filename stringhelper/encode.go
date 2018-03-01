package stringhelper

import "fmt"

// Encode receives an int64 from the database an return its base62
func Encode(i int64) string {
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X","Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	base := int64(len(alphabet))
	s := ""

	if i == 0 {
		return alphabet[0]
	}

	for i > 0 {
		s = fmt.Sprintf("%s%s", s, alphabet[i % base])
		i = i / base
	}

	return reverse(s)
}