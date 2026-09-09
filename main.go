package main

import (
	"http/handlers"
	todo "http/list"
	"http/server"
)

func main() {

	lists := todo.NewList()
	handler := handlers.NewHandler(lists)
	servers := server.NewHTTPServer(handler)

	servers.StartServer()
}
