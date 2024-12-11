package server

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/ekediala/expensix"
	"github.com/ekediala/expensix/httpio"
	"github.com/ekediala/expensix/sqlx"
	"github.com/ekediala/expensix/templ/pages/health"
)

//go:embed public/*
var assets embed.FS

var (
	PathLanding = fmt.Sprintf("%s %s", http.MethodGet, expensix.RouteLanding)
	PathAssets  = fmt.Sprintf("%s /public/", http.MethodGet)
)

type Server struct {
	http.Handler
}

func New(db sqlx.Querier) *Server {
	mux := http.NewServeMux()
	mux.Handle(PathLanding, Landing())
	mux.Handle(PathAssets, http.FileServer(http.FS(assets)))

	s := Server{
		Handler: mux,
	}
	return &s
}

func Landing() httpio.Handler {
	return func(w http.ResponseWriter, r *http.Request) httpio.Handler {
		return httpio.Code(http.StatusOK, httpio.HTML(health.Health(), httpio.OK))
	}
}
