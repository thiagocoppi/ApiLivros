package main

import (
	"LibraryApi/database"
	"LibraryApi/server"
)

func main() {
	database.StartDb()

	server := server.NewServer()
	server.Run()
}
