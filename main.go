package main

const englishHelloPrefix = "Hello "
const Stranger = "Stranger"

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + Stranger
	}
	return englishHelloPrefix + name
}
