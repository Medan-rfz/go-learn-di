package server

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/dig"
)

// IHandler dependency of the handler
type IHandler interface {
	Func() http.HandlerFunc
}

// Server handler
type Server struct {
	handler IHandler
}

// Params initialization params
type Params struct {
	dig.In

	Handler IHandler
}

// NewServer constructor
func NewServer(p Params) *Server {
	return &Server{
		handler: p.Handler,
	}
}

// Run starting of the server
func (s *Server) Run(host string) error {
	router := mux.NewRouter()
	attachSwagger(router)

	router.
		Name("Func").
		Methods("GET").
		Path("/func").
		Handler(s.handler.Func())

	return http.ListenAndServe(host, router)
}

func attachSwagger(router *mux.Router) {
	fs := http.FileServer(http.Dir("./docs"))
	router.PathPrefix("/swagger/docs/").Handler(http.StripPrefix("/swagger/docs", fs))

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/docs/swagger.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}
