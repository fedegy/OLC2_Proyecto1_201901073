package entorno

import arraylist "github.com/colegno/arrayList"

type Simbolo struct {
	Linea         int
	Columna       int
	Identificador string
	Valor         interface{}
	Tipo          TipoDato
	Constante     bool
	FlagFuncion   bool
	ListaParams   *arraylist.List
}

func NewSimboloIdentificador(linea, columna int, identificador string) *Simbolo {
	return &Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Valor:         nil,
		Constante:     false,
		FlagFuncion:   false,
	}
}

func NewSimboloIdentificadorValor(linea , columna int, identificador string, valor interface{}, dato TipoDato) Simbolo {
	e := Simbolo{
		Linea:         linea,
		Columna:       columna,
		Identificador: identificador,
		Constante:     false,
		FlagFuncion:   false,
		Valor:         valor,
		Tipo:          dato,
	}
	return e
}
