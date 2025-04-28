package iteration

import "strings"

var repeatedCount = 5

func Repeat(character string, count int) string {
	if count > 0 {
		repeatedCount = count
	}
	var repeatedCharacter strings.Builder

	for i := 0; i < repeatedCount; i++ {
		repeatedCharacter.WriteString(character)
	}
	return repeatedCharacter.String()
}
