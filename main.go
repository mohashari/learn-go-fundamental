package main

import (
	"fmt"
)

func main() {

	var languages [2]string
	languages[0] = "go"
	languages[1] = "java"
	fmt.Println(languages)
	fmt.Println(len(languages))

	languages2 := [2]string{"kotlin", "C#"}
	fmt.Println(languages2)
	fmt.Println(len(languages2))

	languages3 := [...]string{
		"javascript",
		"html",
	}
	fmt.Println(languages3)
	fmt.Println(len(languages3))

	for index, lang := range languages {
		fmt.Println("index: ", index, "lang: ", lang)
	}

	for _, lang := range languages2 {
		fmt.Println("lang: ", lang)
	}
}
