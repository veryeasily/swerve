package main

import (
  "net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func SetMIMEType(mimetype string) Middleware {
  return func(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", mimetype)
      f(w, r)
    }
  }
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
