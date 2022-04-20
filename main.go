package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.141

	fmt.Println(f)

	firstName := "Alex2"
	fmt.Println(firstName)
	c := complex(3, 4)
	r, im := real(c), imag(c)
	fmt.Println(r, im)
	// pointers

	name := "Alex"
	ptr := &name
	ptr1 := &firstName
	*ptr = *ptr1
	fmt.Println(ptr, *ptr)
}
