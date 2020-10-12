package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ovasquezbrito/candyshop/pkg/http/rest"
)

func main() {

	fmt.Println("starting server on port 8080")
	router := rest.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", router))
}
