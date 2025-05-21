package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/asruldev/cab/cmd/server"
	"github.com/asruldev/cab/pkg/config"
)

// Package main Cab API.
//
// @title Cab API
// @version 1.0
// @description This is a sample server for Cab API.
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.LoadEnv()
	r := server.SetupRouter()

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
