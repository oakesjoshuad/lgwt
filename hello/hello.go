package main

import (
	"fmt"
)

const french = "French"
const spanish = "Spanish"
const en_prefix = "Hello, "
const sp_prefix = "Hola, "
const fr_prefix = "Bonjour, "

func greetingPrefix(lang string) (prefix string) {
	switch lang {
		case french:
			prefix = fr_prefix
		case spanish:
			prefix = sp_prefix
		default:
			prefix = en_prefix
	}
	return
}

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println("Joshua","English")
}
