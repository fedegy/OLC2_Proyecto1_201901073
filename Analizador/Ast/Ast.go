package ast

import arrayL "github.com/colegno/arrayList"

type Ast struct {
	Instrucciones *arrayL.List
}

func NewAst(lista *arrayL.List) Ast {
	ast := Ast{Instrucciones: lista}
	return ast
}
