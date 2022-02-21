package expresion

import (
	abstract2 "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
	"fmt"
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

func (p Operacion) GetValor(env entorno.Entorno) abstract2.RetornoTipo {
	var valIzq abstract2.RetornoTipo
	var valDer abstract2.RetornoTipo

	if p.Unario == true {
		valIzq = p.HIzq.GetValor(env)
	} else {
		valIzq = p.HIzq.GetValor(env)
		valDer = p.HDer.GetValor(env)
	}

	var dominante abstract2.TipoDato

	switch p.Operador {
	case "+":
		{
			dominante = suma[valIzq.Tipo][valDer.Tipo]
			if dominante == abstract2.INT {
				fmt.Println(valIzq.Tipo)
				fmt.Println(valDer.Tipo)
				return abstract2.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(int) + valDer.Valor.(int))}
			} else if dominante == abstract2.FLOAT {
				return abstract2.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(float64) + valDer.Valor.(float64))}
			} else if dominante == abstract2.STRING {
				res1 := fmt.Sprintf("%v", valIzq.Valor)
				res2 := fmt.Sprintf("%v", valDer.Valor)
				return abstract2.RetornoTipo{Tipo: dominante, Valor: res1 + res2}
			}
		}
	}
	return abstract2.RetornoTipo{Tipo: abstract2.NULL, Valor: nil}
}
