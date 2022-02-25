package expresion

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
)

type Relacional struct {
	HIzq     abstract.Expresion
	Operador string
	HDer     abstract.Expresion
}

func NewRelacional(HIzq abstract.Expresion, Operador string, HDer abstract.Expresion) Relacional {
	env := Relacional{HIzq, Operador, HDer}
	return env
}

func (p Relacional) GetValor(env entorno.Entorno) entorno.RetornoTipo {
	var valIzq entorno.RetornoTipo
	var valDer entorno.RetornoTipo

	valIzq = p.HIzq.GetValor(env)
	valDer = p.HDer.GetValor(env)

	switch p.Operador {
	case ">":
		{
			if valIzq.Tipo == entorno.INT && valDer.Tipo == entorno.INT {
				return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(int) > valDer.Valor.(int))}
			} else if valIzq.Tipo == entorno.FLOAT && valDer.Tipo == entorno.FLOAT {
				return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(float64) > valDer.Valor.(float64))}
			} else if valIzq.Tipo == entorno.STRING && valDer.Tipo == entorno.STRING {
				return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(string) > valDer.Valor.(string))}
			}
		}

	case "<":
		{
			if valIzq.Tipo == entorno.INT && valDer.Tipo == entorno.INT {
				return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(int) < valDer.Valor.(int))}
			} else if valIzq.Tipo == entorno.FLOAT && valDer.Tipo == entorno.FLOAT {
				return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(float64) < valDer.Valor.(float64))}
			} else if valIzq.Tipo == entorno.STRING && valDer.Tipo == entorno.STRING {
				return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(string) < valDer.Valor.(string))}
			}
		}
	case ">=": {
		if valIzq.Tipo == entorno.INT && valDer.Tipo == entorno.INT {
			return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(int) >= valDer.Valor.(int))}
		} else if valIzq.Tipo == entorno.FLOAT && valDer.Tipo == entorno.FLOAT {
			return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(float64) >= valDer.Valor.(float64))}
		} else if valIzq.Tipo == entorno.STRING && valDer.Tipo == entorno.STRING {
			return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: (valIzq.Valor.(string) >= valDer.Valor.(string))}
		}
	}
	}
	return entorno.RetornoTipo{Tipo: entorno.NULL, Valor: nil}
}
