package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	muxDispatcher *mux.Router
}

//NewMuxRouter instancead a new Router
func NewMuxRouter() IRouter {
	muxDispatcher := mux.NewRouter()
	return &muxRouter{muxDispatcher}
}

func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v \n", port)
	server := http.ListenAndServe(port, m.muxDispatcher)
	log.Fatal(server)
}

func (m *muxRouter) GET(uri string, f funcHandler) {
	dispatcher(m, uri, f, http.MethodGet)
}

func (m *muxRouter) POST(uri string, f funcHandler) {
	dispatcher(m, uri, f, http.MethodPost)
}

func (m *muxRouter) PUT(uri string, f funcHandler) {
	dispatcher(m, uri, f, http.MethodPut)
}

func (m *muxRouter) PATCH(uri string, f funcHandler) {
	dispatcher(m, uri, f, http.MethodPatch)
}

func dispatcher(m *muxRouter, uri string, f funcHandler, method string) {
	m.muxDispatcher.
		Path(uri).
		Handler(http.HandlerFunc(f)).
		Methods(method)
}
