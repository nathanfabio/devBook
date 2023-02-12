package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":7000", r))
}
