package hello

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const japanese = "Japanese"
const helloPrefix = "Hello, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "
const helloPrefixJapanese = "Konnichiwa, "

func Hello(name, lang string) string {
	if name == "" {
		return helloPrefix + "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = helloPrefixFrench
	case spanish:
		prefix = helloPrefixSpanish
	case japanese:
		prefix = helloPrefixJapanese
	default:
		prefix = helloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
