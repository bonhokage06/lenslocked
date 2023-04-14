package main

import (
	"fmt"
	"net/http"

	"github.com/bonhokage06/lenslocked/router"
)

func main() {
	var router router.Router
	r := router.New()
	fmt.Println("Starting the server on :3000")
	http.ListenAndServe(":3000", r)
}
