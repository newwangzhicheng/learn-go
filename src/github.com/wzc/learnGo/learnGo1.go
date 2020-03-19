package main

import "fmt"

const helloWorld = "hello world"
const holaWorld = "hola world"
const bonjourWorld = "bonjour world"
const spanish = "Spanish"
const french = "French"

//Hello returns a string
func Hello(name string, language string) string {
	if name == "" {
		name = "wsq"
	}

	return greetingprefix(language) + "," + name
}

func greetingprefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = holaWorld
	case french:
		prefix = bonjourWorld
	default:
		prefix = helloWorld
	}
	return
}

func main() {
	fmt.Println("go")
	fmt.Println("23")
}
