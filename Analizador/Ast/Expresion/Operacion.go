package expresion

import (
	abstract "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
	entorno "OLC2_PROYECTO1_201901073/Analizador/Entorno"
	"fmt"
	"math"
	"strconv"
)

var suma = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.STRING, entorno.STRING, entorno.STRING, entorno.STRING, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var resta = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var multi = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var division = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var potencia = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

var modulo = [5][5]entorno.TipoDato{
	{entorno.INT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

type Operacion struct {
	HIzq     abstract.Expresion
	Operador string
	HDer     abstract.Expresion
	Unario   bool
}

func NewOperacion(HIzq abstract.Expresion, Operador string, HDer abstract.Expresion, Unario bool) Operacion {
	env := Operacion{HIzq, Operador, HDer, Unario}
	return env
}

func (p Operacion) GetValor(env entorno.Entorno) entorno.RetornoTipo {
	var valIzq entorno.RetornoTipo
	var valDer entorno.RetornoTipo

	if p.Unario == true {
		valIzq = p.HIzq.GetValor(env)
	} else {
		valIzq = p.HIzq.GetValor(env)
		valDer = p.HDer.GetValor(env)
	}

	var dominante entorno.TipoDato

	switch p.Operador {
	case "+":
		{
			dominante = suma[valIzq.Tipo][valDer.Tipo]
			if dominante == entorno.INT {
				fmt.Println(valIzq.Tipo)
				fmt.Println(valDer.Tipo)
				return entorno.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(int) + valDer.Valor.(int))}
			} else if dominante == entorno.FLOAT {
				return entorno.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(float64) + valDer.Valor.(float64))}
			} else if dominante == entorno.STRING {
				res1 := fmt.Sprintf("%v", valIzq.Valor)
				res2 := fmt.Sprintf("%v", valDer.Valor)
				return entorno.RetornoTipo{Tipo: dominante, Valor: res1 + res2}
			}
		}
	case "-":
		{
			if p.Unario {
				if valIzq.Tipo == entorno.INT {
					return entorno.RetornoTipo{Tipo: valIzq.Tipo, Valor: -valIzq.Valor.(int)}
				} else if valIzq.Tipo == entorno.FLOAT {
					return entorno.RetornoTipo{Tipo: valIzq.Tipo, Valor: -valDer.Valor.(float64)}
				}
			}

			dominante = resta[valIzq.Tipo][valDer.Tipo]
			if dominante == entorno.INT {
				return entorno.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(int) - valDer.Valor.(int))}
			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				return entorno.RetornoTipo{Tipo: dominante, Valor: val1 - val2}
			}
		}
	case "*":
		{
			dominante = multi[valIzq.Tipo][valDer.Tipo]
			if dominante == entorno.INT {
				return entorno.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(int) * valDer.Valor.(int))}
			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				return entorno.RetornoTipo{Tipo: dominante, Valor: val1 * val2}
			}
		}
	case "/":
		{
			dominante = division[valIzq.Tipo][valDer.Tipo]
			if dominante == entorno.INT {
				return entorno.RetornoTipo{Tipo: dominante, Valor: (valIzq.Valor.(int) / valDer.Valor.(int))}
			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				return entorno.RetornoTipo{Tipo: dominante, Valor: val1 / val2}
			}
		}

	case "**":
		{
			dominante = potencia[valIzq.Tipo][valDer.Tipo]
			if dominante == entorno.INT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				var res = int(math.Pow(val1, val2))
				return entorno.RetornoTipo{Tipo: dominante, Valor: res}
			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				return entorno.RetornoTipo{Tipo: dominante, Valor: math.Pow(val1, val2)}
			}
		}
	case "f64**":
		{
			dominante = potencia[valIzq.Tipo][valDer.Tipo]
			if dominante == entorno.INT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				var res = int(math.Pow(val1, val2))
				return entorno.RetornoTipo{Tipo: dominante, Valor: res}
			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				return entorno.RetornoTipo{Tipo: dominante, Valor: math.Pow(val1, val2)}
			}
		}
	case "%":
		{
			dominante = modulo[valIzq.Tipo][valDer.Tipo]
			if dominante == entorno.INT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				var res = int(math.Mod(val1, val2))
				return entorno.RetornoTipo{Tipo: dominante, Valor: res}
			} else if dominante == entorno.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", valIzq.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valDer.Valor), 64)
				var res = math.Mod(val1, val2)
				return entorno.RetornoTipo{Tipo: dominante, Valor: res}
			}
		}
	}
	return entorno.RetornoTipo{Tipo: entorno.NULL, Valor: nil}
}
