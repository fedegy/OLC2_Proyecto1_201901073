package expresion

import (
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
)

type Primitivo struct {
	Valor interface{}
	Tipo  entorno.TipoDato
}

func (p Primitivo) GetValor(env entorno.Entorno) entorno.RetornoTipo {
	return entorno.RetornoTipo{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func NewPrimitivo(valor interface{}, tipo entorno.TipoDato) Primitivo {
	e := Primitivo{valor, tipo}
	return e
}
