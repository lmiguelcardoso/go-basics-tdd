package main

import "fmt"

const spanish = "Spanish"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == ""{
		name = "World"
	}
	
	if language == portuguese{
		return portugueseHelloPrefix + name
	}

	if language == spanish{
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}