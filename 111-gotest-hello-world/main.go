package main

import "fmt"

const french = "french"
const spanish = "spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (helloPrefix string) {
	switch language {
	case french:
		helloPrefix = frenchHelloPrefix
	case spanish:
		helloPrefix = spanishHelloPrefix
	default:
		helloPrefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Lana", ""))
}
