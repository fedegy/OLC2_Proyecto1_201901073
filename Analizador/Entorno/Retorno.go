package entorno

type TipoDato int

const (
	INT TipoDato = iota
	FLOAT
	STRING
	BOOL
	CHAR
)

type RetornoType struct {
	Tipo  TipoDato
	Valor interface{}
}
