package instrucciones

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	expresion "OLC2_PROYECTO1_201901073/Analizador/Ast/Expresion"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
	"fmt"

	arraylist "github.com/colegno/arrayList"
)

var tipoDec = [6][6]entorno.TipoDato{
	{entorno.INT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.CHAR},
}

type Declaracion struct {
	ValorInicial abstract.Expresion
	TipoVariable entorno.TipoDato
	LVariables   *arraylist.List
}

func NewDeclaracion(lVariables *arraylist.List, tipoVar entorno.TipoDato) *Declaracion {
	return &Declaracion{
		TipoVariable: tipoVar,
		LVariables:   lVariables,
	}
}

func NewDeclaracionValor(lVariables *arraylist.List, tipoVar entorno.TipoDato, valor abstract.Expresion) *Declaracion {
	return &Declaracion{
		TipoVariable: tipoVar,
		LVariables:   lVariables,
		ValorInicial: valor,
	}
}

func (d *Declaracion) Ejecutar(env entorno.Entorno) interface{} {
	if d.FlagDeclarado() {
		if d.LVariables.Len() > 1 {
			return nil
		}

		retornoExpresion := d.ValorInicial.GetValor(env)
		tipoExpresion := retornoExpresion.Tipo
		tipoDeclaracion := d.TipoVariable

		TipoRes := tipoDec[tipoDeclaracion][tipoExpresion]

		if TipoRes == entorno.NULL {
			return nil
		}

		for i := 0; i < d.LVariables.Len(); i++ {
			varDeclarar := d.LVariables.GetValue(i).(expresion.Identificador)
			if env.ExisteSimbolo(varDeclarar.Identificador) {
				fmt.Printf("Error, la variable %s ya se declaro", varDeclarar.Identificador)
			} else {
				simboloTabla := entorno.NewSimboloIdentificadorValor(
					0,
					0,
					varDeclarar.Identificador,
					retornoExpresion.Valor,
					TipoRes,
				)
				env.SetSimbolo(varDeclarar.Identificador, simboloTabla)
			}
		}
	}
	return nil
}

func (d *Declaracion) FlagDeclarado() bool {
	return d.ValorInicial != nil
}
