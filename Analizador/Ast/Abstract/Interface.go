package abstract

import (
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
)

type Expresion interface {
	GetValor(env entorno.Entorno) entorno.RetornoTipo
}

type Instruccion interface {
	Ejecutar(env entorno.Entorno) interface{}
}
