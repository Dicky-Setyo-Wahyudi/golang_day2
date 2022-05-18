package main

import (
	"training_golang_keempat/connection"
	"training_golang_keempat/handlers"
)

func main() {
	connection.Connect()
	handlers.HandleReq()
}
