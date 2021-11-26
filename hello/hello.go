package hello

import "fmt"

const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {

	if name == "" {
		name = "world"
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case "fr":
		prefix = frenchHelloPrefix
	case "zh":
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("jiejie", ""))
}
