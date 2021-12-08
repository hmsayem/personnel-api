package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type muxRouter struct {
}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) GET(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	log.Printf("Mux HTTP server is running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
