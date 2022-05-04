package main

import (
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
}

func stratWebServer() {
	fmt.Println("Starting server....")
	//do smth
	fmt.Println("Server started")
}
