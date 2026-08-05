package main

const englishHelloPrefix = "Hello "
const spanishPrefix = "Hola "
const Stranger = "Stranger"

func Hello(name, lang string) string {
	prefix := DefinePrefix(lang)
	if name == "" {
		return prefix + Stranger
	}
	return prefix + name
}

func DefinePrefix(lang string) string {
	if lang == "Spanish" {
		return spanishPrefix
	}
	return englishHelloPrefix
}
