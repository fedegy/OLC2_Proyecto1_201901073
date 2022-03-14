package controllers

import (
	analizador "OLC2_PROYECTO1_201901073/Analizador"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
	"OLC2_PROYECTO1_201901073/Analizador/parser"
	complementos "OLC2_PROYECTO1_201901073/Complementos"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Solicitud struct {
	Text string `json:"text"`
}

func Inicio() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]interface{}{"ok": "Exitoso"})
	}
}

func Data() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var solicitu Solicitud

		if err := json.NewDecoder(r.Body).Decode(&solicitu); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{"Status": http.StatusBadRequest, "Message": "Error"})
			return
		}
		errores := &complementos.CustomErrorListener{}

		//errores := &complementos.CustomErrorListener{}
		//is := antlr.NewInputStream("{system.out.println( \"\nresultado: \"+( 10*1)+ \"\n\") ;system.out.println( \"\nresultado: \"+( 10*1)+ \"\n\") ;}")
		is := antlr.NewInputStream(solicitu.Text)
		// Creación de lexer
		lexer := parser.NewLexico(is)
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(errores)

		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		// Creando el parser
		p := parser.NewSintactico(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(errores)
		p.BuildParseTrees = true
		// Obteniendo la raiz
		tree := p.Start()
		fmt.Printf("\nErrores este punto %v", errores.Errores)
		// Listener para recorrer el arbol
		var listener *complementos.TreeShapeListener = complementos.NewTreeShapeListener()
		if len(errores.Errores) == 0 {
			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		}
		AST := listener.Ast

		ENTORNO_GLOBAL := entorno.NewEntorno("GLOBAL", nil)

		for i := 0; i < AST.ListaInstrucciones.Len(); i++ {

			r := AST.ListaInstrucciones.GetValue(i)
			if r != nil {
				r.(interfaces.Instruccion).Ejecutar(ENTORNO_GLOBAL)
			}

		}

		json.NewEncoder(w).Encode(map[string]interface{}{"val": analizador.Consola})
	}
}
