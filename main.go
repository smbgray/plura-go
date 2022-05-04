package main

import (
	"errors"
	"fmt"

	"github.com/smbgray/plura-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Alex",
		LastName:  "Biii",
	}
	fmt.Println(u)
	port := 3000
	_, err := stratWebServer(port)
	fmt.Println(err)
}

func stratWebServer(port int) (int, error) {
	fmt.Println("Starting server....")
	//do smth
	fmt.Println("Server started on port: ", port)
	return port, errors.New("Something went wrong")
}
