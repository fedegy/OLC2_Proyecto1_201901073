package routes

import (
	controllers "OLC2_PROYECTO1_201901073/Controllers"

	"github.com/gorilla/mux"
)

func Interprete(r *mux.Router) {
	r.HandleFunc("/", controllers.Inicio()).Methods("GET")
	r.HandleFunc("/compilar", controllers.Data()).Methods("POST")
}
