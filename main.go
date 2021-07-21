package main

import (
	"net/http"

	"github.com/karimsalhab24/go-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
