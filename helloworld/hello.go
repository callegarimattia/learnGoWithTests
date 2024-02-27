package helloworld

const (
	french     = "French"
	spanish    = "Spanish"
	portuguese = "Portuguese"
	italian    = "Italian"
	english    = "English"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Hola, "
	portugueseHelloPrefix = "Oi, "
	frenchHelloPrefix     = "Bonjour, "
	italianHelloPrefix    = "Buongiorno, "
)

func Hello(name, language string) string {
	var prefix string

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	case italian:
		prefix = italianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	if name == "" {
		return prefix + "world!"
	}

	return prefix + name + "!"
}
