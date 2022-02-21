package instrucciones

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
)

type Imprimir struct {
	Expresiones abstract.Expresion
}

func NewImprimir(valor abstract.Expresion) Imprimir {
	e := Imprimir{valor}
	return e
}
