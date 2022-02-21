package controllers

import (
	parser "OLC2_PROYECTO1_201901073/Analizador/parser"
	"OLC2_PROYECTO1_201901073/Complementos"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Peticion struct {
	Text string `json:"text"`
}

func Inicio() http.HandlerFunc {
	return func (wr http.ResponseWriter, req *http.Request)  {
		wr.Header().Set("Content-Type", "application/json")
		json.NewEncoder(wr).Encode(map[string]interface{}{"Run":"200"})
	}
}

func Data() http.HandlerFunc {
	return func (w http.ResponseWriter, rq *http.Request)  {
		w.Header().Set("Content-Type", "application/json")
		var pet Peticion
		if err := json.NewDecoder(rq.Body).Decode(&pet); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Estado": http.StatusBadRequest, "Descripcion": "Error"})
			return
		}

		errores := &complementos.CustomErrorListener{}
		is := antlr.NewInputStream(pet.Text)
		lexer := parser.NewLexico(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errores)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

		pars := parser.NewSintactico(stream)
		pars.RemoveErrorListeners()
		pars.AddErrorListener(errores)

		pars.BuildParseTrees = true
		tree := pars.Start()

		fmt.Printf("\nErrores %v", errores.Errores)

		var listener *complementos.TreeShapeListener = complementos.NewTreeShapeListener()
		if len(errores.Errores) == 0 {
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		}
		fmt.Println(listener.Consola)
		json.NewEncoder(w).Encode(map[string]interface{}{"consola": listener.Consola})
	}
}