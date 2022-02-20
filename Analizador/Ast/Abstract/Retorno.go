package abstract

type TipoDato int

const (
	INT	TipoDato = iota
	FLOAT
	STRING
	BOOL
	CHAR
)

type RetornoTipo struct {
	Tipo TipoDato
	Valor interface{}
}