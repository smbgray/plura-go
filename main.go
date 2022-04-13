package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.141

	fmt.Println(f)

	firstName := "Alex"
	fmt.Println(firstName)
	c := complex(3, 4)
	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
