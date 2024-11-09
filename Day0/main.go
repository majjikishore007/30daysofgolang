package main

import "fmt"

// seperating the domain code from side effects
const spanish = "Spanish"
const french = "French"
const frenchHelloPrefix="Bonjour, "
const spanishHelloPrefix = "Hola, "
const englishHelloPrefix = "Hello, "


// func Hello(name string,language string) string {
// 	if name == "" {
// 		return englishHelloPrefix+"World!"
// 	}
// 	if language == spanish {
// 		return spanishHelloPrefix+name
// 	}
// 	if language == french{
// 		return frenchHelloPrefix+name
// 	}
// 	return englishHelloPrefix+name
// }

//  Refactor the code to make it more readable

func greetingPrefix(language string)string {
	switch language {
		case french:
			return frenchHelloPrefix
		case spanish:
			return spanishHelloPrefix
		default:
			return englishHelloPrefix
	}
}

func Hello(name string,language string) string {
	if name == "" {
		name = "World!"
	}
	return greetingPrefix(language)+name
}
func main() {
	fmt.Println(Hello("world","English"))
}