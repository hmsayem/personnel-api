package router

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type chiRouter struct {
	chi *chi.Mux
}

func NewChiRouter() Router {
	return &chiRouter{chi: chi.NewRouter()}
}

func (r *chiRouter) Get(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.chi.Get(uri, f)
}

func (r *chiRouter) Put(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.chi.Put(uri, f)
}

func (r *chiRouter) Post(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.chi.Post(uri, f)
}

func (r *chiRouter) Delete(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	r.chi.Delete(uri, f)
}

func (r *chiRouter) Serve(port string) {
	log.Printf("Chi HTTP server is running on port %v", port)
	err := http.ListenAndServe(port, r.chi)
	if err != nil {
		log.Fatal("Error starting HTTP server: ", err)
	}
}
