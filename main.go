package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/time", func(w http.ResponseWriter, r *http.Request) {
		y, m, d := time.Now().Date()
		fmt.Fprint(w, fmt.Sprintf("%d/%d/%d", y, m, d))

	})
	r.Get("/fun", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("just have fun "))
	})
	http.ListenAndServe(":6969", r)
}
