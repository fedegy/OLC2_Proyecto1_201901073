package expresion

import (
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
	"reflect"
)

type Identificador struct {
	Identificador string
}

func NewIdentificador(identificador string) Identificador{
	return Identificador{Identificador: identificador}
}

func (id Identificador) GetValor(env entorno.Entorno) entorno.RetornoTipo {
	var existe bool = env.ExisteSimbolo(id.Identificador)
	if existe == false {
		return entorno.RetornoTipo{Valor: nil, Tipo: entorno.NULL}
	}

	sim := env.GetSimbolo(id.Identificador)
	if reflect.TypeOf(sim) == reflect.TypeOf(entorno.Simbolo{}){
		dato := sim.(entorno.Simbolo)
		return entorno.RetornoTipo{Valor: dato.Valor, Tipo: dato.Tipo}
	}
	return entorno.RetornoTipo{Valor: -1, Tipo: entorno.NULL}
}
