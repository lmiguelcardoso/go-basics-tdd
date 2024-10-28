package main

import "fmt"

const (
	spanish = "Spanish"
	portuguese = "Portuguese"

	englishHelloPrefix = "Hello, "
	portugueseHelloPrefix = "Olá, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == ""{
		name = "World"
	}
	
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Chris", ""))
}