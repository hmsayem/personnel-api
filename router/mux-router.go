package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) Get(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) Put(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("PUT")
}

func (*muxRouter) Post(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) Serve(port string) {
	log.Printf("Mux HTTP server is running on port %v", port)
	err := http.ListenAndServe(port, muxDispatcher)
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}
