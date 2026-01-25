package main

import (
	"fmt"
	"net/http"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Logging:", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		println("Finished handling request")
	})
}

func chainMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home Page")
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "User Path: %s\n", r.URL.Path)
	})

	mux.HandleFunc("/users/list", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Listing Users")
	})

	logMux := logging(mux)

	// finalHandler := chainMiddleware(mux, loggingMiddleware, authMiddleware)

	http.ListenAndServe(":8080", logMux)
}
