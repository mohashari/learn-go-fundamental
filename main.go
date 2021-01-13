package main

import "fmt"

func main() {
	learnVariable()
	learnConditionIF()
	learnArray()
	learnSlice()

}

func learnVariable() {
	fmt.Println("-----------------> Learn Variable")
	fmt.Println("learn go fundamental")
	var namaVariable string = "ini variable mengguanakan var"
	angkaVariable := 34
	fmt.Println(namaVariable)
	fmt.Println(angkaVariable)
}

func learnConditionIF() {

	fmt.Println("-----------------> Learn Conditional IF")
	age := 10

	if age > 10 {
		fmt.Println("umur ku di atas 10 ")
	} else if age > 20 {
		fmt.Println("umur ku di atas 20 ")
	} else {
		fmt.Println("umur ku di bawah 10 ")
	}
}

func learnArray() {
	fmt.Println("-----------------> Learn Array")

	//cara pertama
	var languages [2]string
	languages[0] = "go"
	languages[1] = "java"
	fmt.Println(languages)
	fmt.Println(len(languages))

	//cara kedua untuk mendefinisikan array
	languages2 := [2]string{"kotlin", "C#"}
	fmt.Println(languages2)
	fmt.Println(len(languages2))

	//cara ketiga untuk mendefinisikan array
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

func learnSlice() {
	fmt.Println("-----------------> Learn Slice")
	//cara pertama
	var gamingConsole []string
	gamingConsole = append(gamingConsole, "PS1")
	fmt.Println(gamingConsole)

	//cara kedua
	gamingConsole2 := []string{"PS2"}
	fmt.Println(gamingConsole2)

	//cara ketiga
	gamingConsole3 := make([]string, 0)
	gamingConsole3 = append(gamingConsole3, "Nitendo", "PSP")
	fmt.Println(gamingConsole3)

	for _, game := range gamingConsole3 {
		fmt.Println("Game Console: ", game)
	}
}
