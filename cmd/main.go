package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asruldev/cab/cmd/server"
)

func main() {
	r := server.SetupRouter()

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
