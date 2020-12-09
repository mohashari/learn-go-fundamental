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
## Lean if condition 
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