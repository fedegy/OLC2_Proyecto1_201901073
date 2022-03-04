package expresion

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
)

type Logicas struct {
	HIzq 		abstract.Expresion
	Operador	string
	HDer		abstract.Expresion
	Unario		bool
}

func NewLogicas(HIzq abstract.Expresion, Operador string, HDer abstract.Expresion, Unario bool) Logicas {
	env := Logicas{HIzq, Operador, HDer, Unario}
	return env
}

func (p Logicas) GetValor(env entorno.Entorno) entorno.RetornoTipo {
	
	var valIzq entorno.RetornoTipo
	var valDer entorno.RetornoTipo

	if p.Unario == true{
		valIzq = p.HIzq.GetValor(env)
	}else{
		valIzq = p.HIzq.GetValor(env)
		valDer = p.HDer.GetValor(env)
	}

	switch p.Operador {
	case "||" : {
		if valIzq.Tipo == entorno.BOOL && valDer.Tipo == entorno.BOOL {
			return entorno.RetornoTipo{Tipo: entorno.BOOL,Valor: valIzq.Valor.(bool) || valDer.Valor.(bool)}
		}
	}

	case "&&": {
		if valIzq.Tipo == entorno.BOOL && valDer.Tipo == entorno.BOOL{
			return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor:  valIzq.Valor.(bool) && valDer.Valor.(bool)}
		}
	}

case "!": {
	if valIzq.Tipo == entorno.BOOL {
		return entorno.RetornoTipo{Tipo: entorno.BOOL, Valor: ! valIzq.Valor.(bool)}
	}
}
	}
	return entorno.RetornoTipo{Tipo: entorno.NULL, Valor: nil}
}