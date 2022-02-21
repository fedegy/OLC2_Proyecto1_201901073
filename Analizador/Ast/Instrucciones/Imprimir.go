package instrucciones

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
)

type Imprimir struct {
	Expresiones abstract.Expresion
}

func NewImprimir(valor abstract.Expresion) Imprimir {
	e := Imprimir{valor}
	return e
}

func (p Imprimir) Ejecutar(env entorno.Entorno) interface{} {
	expr := p.Expresiones.GetValor(env)
	return expr.Valor
}
