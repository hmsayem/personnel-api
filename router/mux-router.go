package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type muxRouter struct {
	mux *mux.Router
}

func NewMuxRouter() Router {
	return &muxRouter{
		mux: mux.NewRouter(),
	}
}

func (r *muxRouter) Get(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.mux.HandleFunc(uri, f).Methods("GET")
}

func (r *muxRouter) Put(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.mux.HandleFunc(uri, f).Methods("PUT")
}

func (r *muxRouter) Post(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.mux.HandleFunc(uri, f).Methods("POST")
}

func (r *muxRouter) Delete(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.mux.HandleFunc(uri, f).Methods("DELETE")
}

func (r *muxRouter) Serve(port string) {
	log.Printf("Mux HTTP server is running on port %v", port)
	err := http.ListenAndServe(port, r.mux)
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}
