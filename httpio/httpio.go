package httpio

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

type Header struct {
	Key   string
	Value string
}

type Handler func(w http.ResponseWriter, r *http.Request) Handler

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if next := h(w, r); next != nil {
		next(w, r)
	}
}

func OK(w http.ResponseWriter, r *http.Request) Handler {
	return nil
}

func Text(s string) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, s)
		return OK
	}
}

func Code(code int, next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		w.WriteHeader(code)
		return next
	}
}

func HTML(fragment templ.Component, code int, headers ...Header) Handler {
	return func(w http.ResponseWriter, r *http.Request) Handler {
		for _, h := range headers {
			w.Header().Set(h.Key, h.Value)
		}
		templ.Handler(fragment, templ.WithStatus(code)).ServeHTTP(w, r)
		return OK
	}
}
