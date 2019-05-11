package main

import "net/http"

const USERNAME = "batman"
const PASSWORD = "secret"

func MiddlewareAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			w.Write([]byte("something went wrong"))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte("wrong username/password"))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func MiddlewareAllowOnlyGet(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET is allowed"))
			return
		}

		h.ServeHTTP(w, r)
	})
}
