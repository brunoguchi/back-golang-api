package main

import "github.com/brunoguchi/back-golang-api/server"

func main() {
	server := server.NewServer()
	server.Run()
}
