package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	vietNam            = "VietNam"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	vietNamHelloPrefix = "Xin chào, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case vietNam:
		prefix = vietNamHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
