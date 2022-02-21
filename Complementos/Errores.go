package complementos

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CustomError struct {
	linea, columna int
	desc           string
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errores []CustomError
}

func (c *CustomErrorListener) CustomError(rec antlr.Recognizer, offsimbol interface{}, linea, columna int, desc string, e antlr.RecognitionException) {
	c.Errores = append(c.Errores, CustomError{
		linea:   linea,
		columna: columna,
		desc:    desc,
	})
}
