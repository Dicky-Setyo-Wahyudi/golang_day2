package main

import (
	"training_golang_ketiga/connection"
	"training_golang_ketiga/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleReq()
}
