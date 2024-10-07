package main

import "fmt"

const (
	helloPrefix   = "hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
	hindiPrefix   = "Namaste, "

	spanish = "spanish"
	french  = "french"
	hindi   = "hindi"
)

// domain - core logic
func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

// prefix is a named return value, having the default value - ""
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	case hindi:
		prefix = hindiPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
