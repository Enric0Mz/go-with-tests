package hello

const englishHelloPrefix = "Hello "
const spanishPrefix = "Hola "
const frentchPrefix = "Bonjour "
const Stranger = "Stranger"

func Hello(name, lang string) string {
	prefix := DefinePrefix(lang)
	if name == "" {
		return prefix + Stranger
	}
	return prefix + name
}

func DefinePrefix(lang string) string {
	switch lang {
	case "Spanish":
		return spanishPrefix
	case "French":
		return frentchPrefix
	default:
		return englishHelloPrefix
	}
}
