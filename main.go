package main

import (
	routes "OLC2_PROYECTO1_201901073/Routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	rt := mux.NewRouter()
	routes.Interprete(rt)
	corsW := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"COntent-Type", "Origin", "Accept", "*"},
	})

	s := &http.Server{
		Handler:      corsW.Handler(rt),
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
