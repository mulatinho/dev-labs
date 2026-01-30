package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Backend struct {
	Path           string
	isAuthRequired bool
	Handle         http.HandlerFunc
}

var backendList []Backend = []Backend{
	{Path: "/", isAuthRequired: false, Handle: HandleHome},
	{Path: "/admin", isAuthRequired: true, Handle: HandleAdmin},
}

func backendByPath(path string) *Backend {
	for i := range backendList {
		if backendList[i].Path == path {
			return &backendList[i]
		}
	}
	return nil
}

func middleware(handle http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backendNow := backendByPath(r.URL.Path)

		if backendNow == nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, `{"message": "Not Found"}`)
			return
		}

		if backendNow.isAuthRequired {
			auth := r.Header.Get("api-key")
			if auth != "secret123" {
				w.WriteHeader(401)
				fmt.Fprintf(w, `{"message": "Unauthorized"}`)
				return
			}
		}

		start := time.Now()
		w.Header().Set("Content-Type", "application/json")
		ending := time.Since(start)

		log.Printf(`request to "%s" executed in "%s".`, r.URL.Path, &ending)

		handle.ServeHTTP(w, r)
	})
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, `{"message": "Hello, World to Home!"}`)
}

func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, `{"message": "Hello, World to Admin"}`)
		w.WriteHeader(200)
		return
	case "POST":
		fmt.Fprintf(w, `{"message": "Create the World to Admin"}`)
		w.WriteHeader(201)
		return
	default:
		fmt.Fprintf(w, `{"message": "Method Not Allowed"}`)
		w.WriteHeader(405)
		return
	}
}

func main() {
	for _, backend := range backendList {
		log.Printf("Registering backend auth required: %5t, at path: %s", backend.isAuthRequired, backend.Path)
		http.HandleFunc(backend.Path, middleware(backend.Handle))
	}
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
