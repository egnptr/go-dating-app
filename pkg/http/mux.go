package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	Router *mux.Router
}

func NewMuxRouter() Router {
	mux := mux.NewRouter()

	return &muxRouter{
		Router: mux,
	}
}

func (m *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router.HandleFunc(uri, f).Methods("GET")
}

func (m *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router.HandleFunc(uri, f).Methods("POST")
}

func (m *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router.HandleFunc(uri, f).Methods("DELETE")
}

func (m *muxRouter) SERVE(port string) {
	fmt.Printf("HTTP server running on port %v\n", port)
	http.ListenAndServe(port, m.Router)
}
