package expresion

import (
	abstract2 "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
)

var suma = [5][5]abstract2.TipoDato{
	{abstract2.INT, abstract2.FLOAT, abstract2.STRING, abstract2.NULL, abstract2.NULL},
	{abstract2.FLOAT, abstract2.FLOAT, abstract2.STRING, abstract2.NULL, abstract2.NULL},
	{abstract2.STRING, abstract2.STRING, abstract2.STRING, abstract2.STRING, abstract2.NULL},
	{abstract2.NULL, abstract2.NULL, abstract2.STRING, abstract2.NULL, abstract2.NULL},
	{abstract2.NULL, abstract2.NULL, abstract2.NULL, abstract2.NULL, abstract2.NULL},
}

type Operacion struct {
	HIzq     abstract2.Expresion
	Operador string
	HDer     abstract2.Expresion
	Unario   bool
}

func NewOperacion(HIzq abstract2.Expresion, Operador string, HDer abstract2.Expresion, Unario bool) Operacion {
	env := Operacion{HIzq, Operador, HDer, Unario}
	return env
}
