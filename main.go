package main

import (
	"net/http"

	"github.com/smbgray/plura-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
