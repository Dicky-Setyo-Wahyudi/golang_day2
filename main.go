package main

import (
	"training_golang_kedua/connection"
	"training_golang_kedua/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleReq()
}
