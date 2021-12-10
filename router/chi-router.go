package router

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

var (
	chiDispatcher = chi.NewRouter()
)

func (*chiRouter) Get(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) Post(uri string, f func(writer http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) Serve(port string) {
	log.Printf("Chi HTTP server is running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
