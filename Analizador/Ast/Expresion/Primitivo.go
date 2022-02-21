package expresion

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
)

type Primitivo struct {
	Valor interface{}
	Tipo  abstract.TipoDato
}

func (p Primitivo) GetValor(env entorno.Entorno) abstract.RetornoTipo {
	return abstract.RetornoTipo{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func NewPrimitivo(valor interface{}, tipo abstract.TipoDato) Primitivo {
	e := Primitivo{valor, tipo}
	return e
}
