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
