# Learn Go Fundamental

## Learn Variable go

ada beberapa cara untuk mendeklarasikan variable di go contoh dengan *var* atau dengan menulis nama variable nya dulu beru di beri value dengan cara beri tanda *:=* 

```golang

package main

import (
	"fmt"
)

func main() {
	fmt.Println("learn go fundamental")

	var namaVariable string = "ini variable mengguanakan var"
	angkaVariable := 34

	fmt.Println(namaVariable)
	fmt.Println(angkaVariable)
}

```
## Learn if condition 
if condition di golang `if kondisi { terminalisasi }` contoh code
```golang 
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

```

## Learn Array

membuat array `var nama-array [isi-panjangnya]type-data` contoh **var languages [2]string**

fungsi dari :
- len (untuk mengetahui panjangnya array)
- range (digunakan untuk mengetahui panjang array dan biasanya digunakan saat for each )


```golang
package main

import (
	"fmt"
)

func main() {

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

```