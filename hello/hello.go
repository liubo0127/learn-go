package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func selectLanguage(language string) string {
	prefix := ""
	switch language {
	case "english":
		prefix = englishHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	case "french":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return selectLanguage(language) + name
}

func main() {
	fmt.Println(Hello("John", "english"))
}
