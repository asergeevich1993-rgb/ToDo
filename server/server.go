package server

import (
	"http/handlers"
	"http/handlers/middleware"
	"net/http"
)

type HTTPServer struct {
	handlers *handlers.Handler
	svr      *http.Server
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

	logger := middleware.Logger(router)

	hs.svr = &http.Server{
		Addr:    ":8080",
		Handler: logger,
	}

	hs.svr.ListenAndServe()

}
