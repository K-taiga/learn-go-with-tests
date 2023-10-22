package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bondule, "
const spanish = "Spanish"
const french = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string){
    prefix := englishHelloPrefix

    switch language {
    case french:
        prefix = frenchHelloPrefix
    case spanish:
        prefix = spanishHelloPrefix
    default:
        prefix = englishHelloPrefix
    }
    // 返り値にprefixが設定されているのでreturnで返せる
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
