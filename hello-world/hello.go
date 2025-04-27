package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
)

const spanishGrettingPrefix = "Hola, "
const frenchGrettingPrefix = "Bonjour, "
const englishGrettingPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getGrettingPrefix(language) + name
}

func getGrettingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishGrettingPrefix
	case french:
		prefix = frenchGrettingPrefix
	default:
		prefix = englishGrettingPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
