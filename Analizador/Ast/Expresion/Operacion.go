package expresion

import (
	abstract2 "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
	"fmt"
)

var suma = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.STRING, entorno.STRING, entorno.STRING, entorno.STRING, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var resta = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var multi = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var division = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var potencia = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var modulo = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
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

func (p Operacion) GetValor(env entorno.Entorno) entorno.RetornoTipo {
	var valIzq entorno.RetornoTipo
	var valDer entorno.RetornoTipo

	if p.Unario == true {
		valIzq = p.HIzq.GetValor(env)
	} else {
		valIzq = p.HIzq.GetValor(env)
		valDer = p.HDer.GetValor(env)
	}

	var dominante entorno.TipoDato

	switch p.Operador {
		case "+": {
				dominante = suma[valIzq.Tipo][valDer.Tipo]
				if dominante == entorno.INT {
					fmt.Println(valIzq.Tipo)
					fmt.Println(valDer.Tipo)
					return entorno.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(int) + valDer.Valor.(int))}
				} else if dominante == entorno.FLOAT {
					return entorno.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(float64) + valDer.Valor.(float64))}
				} else if dominante == entorno.STRING {
					res1 := fmt.Sprintf("%v", valIzq.Valor)
					res2 := fmt.Sprintf("%v", valDer.Valor)
					return entorno.RetornoTipo{Tipo: dominante, Valor: res1 + res2}
				}
			}
		case "-": {
			dominante = resta[valIzq.Tipo][valDer.Tipo]
		}

	}
	return entorno.RetornoTipo{Tipo: entorno.NULL, Valor: nil}
}
