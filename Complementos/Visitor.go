package complementos

import (
	parser "OLC2_PROYECTO1_201901073/Analizador/parser"
	"OLC2_PROYECTO1_201901073/Analizador/Ast"
)

type TreeShapeListener struct {
	*parser.BaseSintacticoListener
	Ast ast.Ast
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	this.Ast = ctx.GetAst()

}