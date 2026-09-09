package server

import (
	"http/handlers"
	"log"
	"net/http"
)

type HTTPServer struct {
	handlers *handlers.Handler
}

func NewHTTPServer(hh *handlers.Handler) *HTTPServer {
	return &HTTPServer{
		handlers: hh,
	}
}

func (hs *HTTPServer) StartServer() {

	router := http.NewServeMux()

	router.HandleFunc("POST /tasks", hs.handlers.HandleCreateTask)
	router.HandleFunc("GET /tasks", hs.handlers.HandleShowTasks)
	router.HandleFunc("PATCH /tasks/{title}", hs.handlers.HandleCompliteTask)
	router.HandleFunc("PUT /tasks/{title}", hs.handlers.HandleUnCompliteTask)
	router.HandleFunc("DELETE /tasks/{title}", hs.handlers.HandleDeleteTask)

	log.Fatal(http.ListenAndServe(":8080", router))
}
