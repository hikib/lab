package main

const engPrefix = "Hello, "
const gerPrefix = "Hallo, "
const spaPrefix = "Hola, "

func Hello(s, language string) (greet string) {
	if isEmpty(s) {
		return engPrefix + "World"
	}

	switch language {
	case "Spanish":
		greet = spaPrefix + s
	case "German":
		greet = gerPrefix + s
	default:
		greet = engPrefix + s
	}
	return
}

func isEmpty(s string) bool {
	return len(s) == 0
}
