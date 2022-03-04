package complementos

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
	parser "OLC2_PROYECTO1_201901073/Analizador/parser"
	"fmt"
)

type TreeShapeListener struct {
	parser.BaseSintacticoListener
	Consola string
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	res := ctx.GetLista()
	for i := 0; i < res.Len(); i++ {
		r := res.GetValue(i)
		if r != nil {
			valor := res.GetValue(i).(abstract.Instruccion).Ejecutar(entorno.NewEntorno("hl", nil))
			this.Consola += fmt.Sprintf("%v", valor)
		}
	}
}
