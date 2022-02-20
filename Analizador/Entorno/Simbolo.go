package entorno

type Simbolo struct {
	Linea			int
	Columna			int
	Id				string
	Valor			interface{}
	Tipo 			TipoDato
	Constante		bool
	FlagFuncion		bool
}

func SetSimbolo(linea int, columa int, id string)*Simbolo {
	return &Simbolo{
		Linea: linea,
		Columna: columa,
		Id: id,
		Constante: false,
		FlagFuncion: false,
		Valor: nil,
	}
}

func SetSimboloValor(linea int, columna int, id string, valor interface{}, tipo TipoDato) Simbolo{
	e := Simbolo {
		Linea: linea,
		Columna: columna,
		Id: id,
		Constante: false,
		FlagFuncion: false,
		Valor: valor,
		Tipo: tipo,
	}
	return e
}