package main

import (
	"fmt"
)

func main() {
	age := 10

	if age > 10 {
		fmt.Println("umur ku di atas 10 ")
	} else if age > 20 {
		fmt.Println("umur ku di atas 20 ")
	} else {
		fmt.Println("umur ku di bawah 10 ")
	}
}
