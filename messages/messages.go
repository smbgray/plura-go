package messages

import "fmt"

//Greet export
func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!\n", name)
}

func depart(name string) string {
	return fmt.Sprintf("Goodbye, %v.\n", name)
}
