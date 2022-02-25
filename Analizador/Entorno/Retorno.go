package entorno

type TipoDato int

const (
	INT TipoDato = iota
	FLOAT
	STRING
	BOOL
	CHAR
	NULL
)

type RetornoTipo struct {
	Tipo  TipoDato
	Valor interface{}
}
