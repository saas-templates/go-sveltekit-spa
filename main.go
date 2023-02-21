package main

import (
	"embed"
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//go:embed all:build/**
var buildFS embed.FS

var (
	addr    = flag.String("addr", ":8080", "Bind addr to listen for request")
	noLogin = flag.Bool("nologin", false, "Simulate no login on /me call")
)

func main() {
	flag.Parse()

	log.Printf("starting server at %s...", *addr)
	if err := http.ListenAndServe(*addr, newRouter()); err != nil {
		log.Fatalf("server exited: %s", err)
	}
}

func newRouter() chi.Router {
	appFS := mustSub("build")

	nonAppServer := http.FileServer(http.FS(appFS))
	appServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		indexFile, _ := buildFS.ReadFile("build/app/index.html")
		w.Write(indexFile)
	})

	r := chi.NewRouter()
	r.Mount("/", nonAppServer)
	r.Mount("/app", appServer)
	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("PING received")
		})

		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			// Note: -nologin flag can be used to simulate error.
			if *noLogin {
				writeJSON(w, http.StatusUnauthorized, nil)
			} else {
				writeJSON(w, http.StatusOK, map[string]any{
					"id":       "yJIABar",
					"name":     "Bob Builder",
					"picture":  "https://www.gravatar.com/avatar?s=64",
					"email":    "bob@builder.com",
					"username": "bobthebuilder",
				})
			}
		})
	})
	return r
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func mustSub(subPath string) fs.FS {
	f, err := fs.Sub(buildFS, subPath)
	if err != nil {
		panic(err)
	}
	return f
}
