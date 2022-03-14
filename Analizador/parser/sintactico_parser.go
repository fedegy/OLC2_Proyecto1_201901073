// Code generated from Sintactico.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Sintactico

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
import "OLC2_PROYECTO1_201901073/Analizador/Ast/Expresion"
import arrayL "github.com/colegno/arrayList"
import "OLC2_PROYECTO1_201901073/Analizador/Ast/Instrucciones"
import "OLC2_PROYECTO1_201901073/Analizador/Entorno"
import "OLC2_PROYECTO1_201901073/Analizador/Ast"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 192,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 31, 10, 3, 13, 3, 14, 3, 32, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 68, 10, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 79, 10, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 114, 10, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 7, 9, 131, 10, 9, 12, 9, 14, 9, 134, 11, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 144, 10, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 156, 10,
	10, 12, 10, 14, 10, 159, 11, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 7, 11, 170, 10, 11, 12, 11, 14, 11, 173, 11, 11,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 190, 10, 13, 3, 13, 2, 5, 16, 18,
	20, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 5, 3, 2, 35, 36,
	3, 2, 33, 34, 3, 2, 21, 26, 2, 203, 2, 26, 3, 2, 2, 2, 4, 30, 3, 2, 2,
	2, 6, 36, 3, 2, 2, 2, 8, 43, 3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12, 67, 3,
	2, 2, 2, 14, 78, 3, 2, 2, 2, 16, 113, 3, 2, 2, 2, 18, 143, 3, 2, 2, 2,
	20, 160, 3, 2, 2, 2, 22, 174, 3, 2, 2, 2, 24, 189, 3, 2, 2, 2, 26, 27,
	5, 4, 3, 2, 27, 28, 8, 2, 1, 2, 28, 3, 3, 2, 2, 2, 29, 31, 5, 6, 4, 2,
	30, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3,
	2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 35, 8, 3, 1, 2, 35, 5, 3, 2, 2, 2, 36,
	37, 7, 3, 2, 2, 37, 38, 7, 16, 2, 2, 38, 39, 5, 14, 8, 2, 39, 40, 7, 17,
	2, 2, 40, 41, 7, 30, 2, 2, 41, 42, 8, 4, 1, 2, 42, 7, 3, 2, 2, 2, 43, 44,
	7, 13, 2, 2, 44, 45, 7, 8, 2, 2, 45, 46, 5, 10, 6, 2, 46, 47, 7, 29, 2,
	2, 47, 48, 5, 12, 7, 2, 48, 49, 7, 27, 2, 2, 49, 50, 5, 14, 8, 2, 50, 51,
	8, 5, 1, 2, 51, 9, 3, 2, 2, 2, 52, 53, 7, 44, 2, 2, 53, 54, 8, 6, 1, 2,
	54, 11, 3, 2, 2, 2, 55, 56, 7, 4, 2, 2, 56, 68, 8, 7, 1, 2, 57, 58, 7,
	5, 2, 2, 58, 68, 8, 7, 1, 2, 59, 60, 7, 6, 2, 2, 60, 68, 8, 7, 1, 2, 61,
	62, 7, 7, 2, 2, 62, 68, 8, 7, 1, 2, 63, 64, 7, 10, 2, 2, 64, 68, 8, 7,
	1, 2, 65, 66, 7, 9, 2, 2, 66, 68, 8, 7, 1, 2, 67, 55, 3, 2, 2, 2, 67, 57,
	3, 2, 2, 2, 67, 59, 3, 2, 2, 2, 67, 61, 3, 2, 2, 2, 67, 63, 3, 2, 2, 2,
	67, 65, 3, 2, 2, 2, 68, 13, 3, 2, 2, 2, 69, 70, 5, 18, 10, 2, 70, 71, 8,
	8, 1, 2, 71, 79, 3, 2, 2, 2, 72, 73, 5, 20, 11, 2, 73, 74, 8, 8, 1, 2,
	74, 79, 3, 2, 2, 2, 75, 76, 5, 16, 9, 2, 76, 77, 8, 8, 1, 2, 77, 79, 3,
	2, 2, 2, 78, 69, 3, 2, 2, 2, 78, 72, 3, 2, 2, 2, 78, 75, 3, 2, 2, 2, 79,
	15, 3, 2, 2, 2, 80, 81, 8, 9, 1, 2, 81, 82, 7, 34, 2, 2, 82, 83, 5, 14,
	8, 2, 83, 84, 8, 9, 1, 2, 84, 114, 3, 2, 2, 2, 85, 86, 7, 16, 2, 2, 86,
	87, 5, 14, 8, 2, 87, 88, 7, 17, 2, 2, 88, 89, 8, 9, 1, 2, 89, 114, 3, 2,
	2, 2, 90, 91, 7, 4, 2, 2, 91, 92, 7, 28, 2, 2, 92, 93, 7, 11, 2, 2, 93,
	94, 7, 16, 2, 2, 94, 95, 5, 16, 9, 2, 95, 96, 7, 31, 2, 2, 96, 97, 5, 16,
	9, 2, 97, 98, 7, 17, 2, 2, 98, 99, 8, 9, 1, 2, 99, 114, 3, 2, 2, 2, 100,
	101, 7, 5, 2, 2, 101, 102, 7, 28, 2, 2, 102, 103, 7, 11, 2, 2, 103, 104,
	7, 16, 2, 2, 104, 105, 5, 16, 9, 2, 105, 106, 7, 31, 2, 2, 106, 107, 5,
	16, 9, 2, 107, 108, 7, 17, 2, 2, 108, 109, 8, 9, 1, 2, 109, 114, 3, 2,
	2, 2, 110, 111, 5, 22, 12, 2, 111, 112, 8, 9, 1, 2, 112, 114, 3, 2, 2,
	2, 113, 80, 3, 2, 2, 2, 113, 85, 3, 2, 2, 2, 113, 90, 3, 2, 2, 2, 113,
	100, 3, 2, 2, 2, 113, 110, 3, 2, 2, 2, 114, 132, 3, 2, 2, 2, 115, 116,
	12, 9, 2, 2, 116, 117, 9, 2, 2, 2, 117, 118, 5, 16, 9, 10, 118, 119, 8,
	9, 1, 2, 119, 131, 3, 2, 2, 2, 120, 121, 12, 8, 2, 2, 121, 122, 9, 3, 2,
	2, 122, 123, 5, 16, 9, 9, 123, 124, 8, 9, 1, 2, 124, 131, 3, 2, 2, 2, 125,
	126, 12, 7, 2, 2, 126, 127, 7, 37, 2, 2, 127, 128, 5, 16, 9, 8, 128, 129,
	8, 9, 1, 2, 129, 131, 3, 2, 2, 2, 130, 115, 3, 2, 2, 2, 130, 120, 3, 2,
	2, 2, 130, 125, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2,
	132, 133, 3, 2, 2, 2, 133, 17, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 136,
	8, 10, 1, 2, 136, 137, 7, 20, 2, 2, 137, 138, 5, 18, 10, 4, 138, 139, 8,
	10, 1, 2, 139, 144, 3, 2, 2, 2, 140, 141, 5, 20, 11, 2, 141, 142, 8, 10,
	1, 2, 142, 144, 3, 2, 2, 2, 143, 135, 3, 2, 2, 2, 143, 140, 3, 2, 2, 2,
	144, 157, 3, 2, 2, 2, 145, 146, 12, 6, 2, 2, 146, 147, 7, 19, 2, 2, 147,
	148, 5, 18, 10, 7, 148, 149, 8, 10, 1, 2, 149, 156, 3, 2, 2, 2, 150, 151,
	12, 5, 2, 2, 151, 152, 7, 18, 2, 2, 152, 153, 5, 18, 10, 6, 153, 154, 8,
	10, 1, 2, 154, 156, 3, 2, 2, 2, 155, 145, 3, 2, 2, 2, 155, 150, 3, 2, 2,
	2, 156, 159, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158,
	19, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2, 160, 161, 8, 11, 1, 2, 161, 162,
	5, 16, 9, 2, 162, 163, 8, 11, 1, 2, 163, 171, 3, 2, 2, 2, 164, 165, 12,
	4, 2, 2, 165, 166, 9, 4, 2, 2, 166, 167, 5, 20, 11, 5, 167, 168, 8, 11,
	1, 2, 168, 170, 3, 2, 2, 2, 169, 164, 3, 2, 2, 2, 170, 173, 3, 2, 2, 2,
	171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 21, 3, 2, 2, 2, 173, 171,
	3, 2, 2, 2, 174, 175, 5, 24, 13, 2, 175, 176, 8, 12, 1, 2, 176, 23, 3,
	2, 2, 2, 177, 178, 7, 38, 2, 2, 178, 190, 8, 13, 1, 2, 179, 180, 7, 40,
	2, 2, 180, 190, 8, 13, 1, 2, 181, 182, 7, 39, 2, 2, 182, 190, 8, 13, 1,
	2, 183, 184, 7, 42, 2, 2, 184, 190, 8, 13, 1, 2, 185, 186, 7, 43, 2, 2,
	186, 190, 8, 13, 1, 2, 187, 188, 7, 41, 2, 2, 188, 190, 8, 13, 1, 2, 189,
	177, 3, 2, 2, 2, 189, 179, 3, 2, 2, 2, 189, 181, 3, 2, 2, 2, 189, 183,
	3, 2, 2, 2, 189, 185, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 25, 3, 2,
	2, 2, 13, 32, 67, 78, 113, 130, 132, 143, 155, 157, 171, 189,
}
var literalNames = []string{
	"", "'println!'", "'i64'", "'f64'", "'&str'", "'char'", "'mut'", "'String'",
	"'bool'", "'pow'", "'as'", "'let'", "'to_owned()'", "'to_string()'", "'('",
	"')'", "'&&'", "'||'", "'!'", "'!='", "'>='", "'<='", "'>'", "'<'", "'=='",
	"'='", "'::'", "':'", "';'", "','", "'.'", "'+'", "'-'", "'*'", "'/'",
	"'%'", "", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "PRINTLN", "R_INT", "R_FLOAT", "R_STR", "R_CHAR", "MUT", "R_STRING",
	"R_BOOL", "POW", "AS", "LET", "TOOWNED", "TOSTRING", "PARENA", "PARENC",
	"AND", "OR", "NOT", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE",
	"MENORQUE", "IGUALIGUAL", "IGUAL", "CUATROPT", "DOSPT", "PTCOMA", "COMA",
	"PUNTO", "MAS", "MENOS", "POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT",
	"CADENA", "CARACTER", "R_TRUE", "R_FALSE", "ID", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "declaracion", "listides", "tipos",
	"expr", "expr_op", "expr_log", "expr_rel", "expr_val", "primitivo",
}

type Sintactico struct {
	*antlr.BaseParser
}

// NewSintactico produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Sintactico instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewSintactico(input antlr.TokenStream) *Sintactico {
	this := new(Sintactico)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sintactico.g4"

	return this
}

// Sintactico tokens.
const (
	SintacticoEOF        = antlr.TokenEOF
	SintacticoPRINTLN    = 1
	SintacticoR_INT      = 2
	SintacticoR_FLOAT    = 3
	SintacticoR_STR      = 4
	SintacticoR_CHAR     = 5
	SintacticoMUT        = 6
	SintacticoR_STRING   = 7
	SintacticoR_BOOL     = 8
	SintacticoPOW        = 9
	SintacticoAS         = 10
	SintacticoLET        = 11
	SintacticoTOOWNED    = 12
	SintacticoTOSTRING   = 13
	SintacticoPARENA     = 14
	SintacticoPARENC     = 15
	SintacticoAND        = 16
	SintacticoOR         = 17
	SintacticoNOT        = 18
	SintacticoDIFERENTE  = 19
	SintacticoMAYORIGUAL = 20
	SintacticoMENORIGUAL = 21
	SintacticoMAYORQUE   = 22
	SintacticoMENORQUE   = 23
	SintacticoIGUALIGUAL = 24
	SintacticoIGUAL      = 25
	SintacticoCUATROPT   = 26
	SintacticoDOSPT      = 27
	SintacticoPTCOMA     = 28
	SintacticoCOMA       = 29
	SintacticoPUNTO      = 30
	SintacticoMAS        = 31
	SintacticoMENOS      = 32
	SintacticoPOR        = 33
	SintacticoDIVIDIDO   = 34
	SintacticoMODULO     = 35
	SintacticoENTERO     = 36
	SintacticoFLOAT      = 37
	SintacticoCADENA     = 38
	SintacticoCARACTER   = 39
	SintacticoR_TRUE     = 40
	SintacticoR_FALSE    = 41
	SintacticoID         = 42
	SintacticoWHITESPACE = 43
)

// Sintactico rules.
const (
	SintacticoRULE_start         = 0
	SintacticoRULE_instrucciones = 1
	SintacticoRULE_instruccion   = 2
	SintacticoRULE_declaracion   = 3
	SintacticoRULE_listides      = 4
	SintacticoRULE_tipos         = 5
	SintacticoRULE_expr          = 6
	SintacticoRULE_expr_op       = 7
	SintacticoRULE_expr_log      = 8
	SintacticoRULE_expr_rel      = 9
	SintacticoRULE_expr_val      = 10
	SintacticoRULE_primitivo     = 11
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetAst returns the ast attribute.
	GetAst() ast.Ast

	// SetAst sets the ast attribute.
	SetAst(ast.Ast)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	ast            ast.Ast
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetAst() ast.Ast { return s.ast }

func (s *StartContext) SetAst(v ast.Ast) { s.ast = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Sintactico) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SintacticoRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).ast = ast.NewAst(localctx.(*StartContext).Get_instrucciones().GetL())

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayL.List

	// SetL sets the l attribute.
	SetL(*arrayL.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayL.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetL() *arrayL.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayL.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Sintactico) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SintacticoRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayL.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SintacticoPRINTLN {
		{
			p.SetState(27)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}
	fmt.Printf("Tipo %T", localctx.(*InstruccionesContext).GetE())

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruccion)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	instr  abstract.Instruccion
	_expr  IExprContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_expr() IExprContext { return s._expr }

func (s *InstruccionContext) Set_expr(v IExprContext) { s._expr = v }

func (s *InstruccionContext) GetInstr() abstract.Instruccion { return s.instr }

func (s *InstruccionContext) SetInstr(v abstract.Instruccion) { s.instr = v }

func (s *InstruccionContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(SintacticoPRINTLN, 0)
}

func (s *InstruccionContext) PARENA() antlr.TerminalNode {
	return s.GetToken(SintacticoPARENA, 0)
}

func (s *InstruccionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InstruccionContext) PARENC() antlr.TerminalNode {
	return s.GetToken(SintacticoPARENC, 0)
}

func (s *InstruccionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(SintacticoPTCOMA, 0)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Sintactico) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SintacticoRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(SintacticoPRINTLN)
	}
	{
		p.SetState(35)
		p.Match(SintacticoPARENA)
	}
	{
		p.SetState(36)

		var _x = p.Expr()

		localctx.(*InstruccionContext)._expr = _x
	}
	{
		p.SetState(37)
		p.Match(SintacticoPARENC)
	}
	{
		p.SetState(38)
		p.Match(SintacticoPTCOMA)
	}
	localctx.(*InstruccionContext).instr = instrucciones.NewImprimir(localctx.(*InstruccionContext).Get_expr().GetP())

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_listides returns the _listides rule contexts.
	Get_listides() IListidesContext

	// Get_tipos returns the _tipos rule contexts.
	Get_tipos() ITiposContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_listides sets the _listides rule contexts.
	Set_listides(IListidesContext)

	// Set_tipos sets the _tipos rule contexts.
	Set_tipos(ITiposContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// GetInstr returns the instr attribute.
	GetInstr() abstract.Instruccion

	// SetInstr sets the instr attribute.
	SetInstr(abstract.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	instr     abstract.Instruccion
	_listides IListidesContext
	_tipos    ITiposContext
	_expr     IExprContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_listides() IListidesContext { return s._listides }

func (s *DeclaracionContext) Get_tipos() ITiposContext { return s._tipos }

func (s *DeclaracionContext) Get_expr() IExprContext { return s._expr }

func (s *DeclaracionContext) Set_listides(v IListidesContext) { s._listides = v }

func (s *DeclaracionContext) Set_tipos(v ITiposContext) { s._tipos = v }

func (s *DeclaracionContext) Set_expr(v IExprContext) { s._expr = v }

func (s *DeclaracionContext) GetInstr() abstract.Instruccion { return s.instr }

func (s *DeclaracionContext) SetInstr(v abstract.Instruccion) { s.instr = v }

func (s *DeclaracionContext) LET() antlr.TerminalNode {
	return s.GetToken(SintacticoLET, 0)
}

func (s *DeclaracionContext) MUT() antlr.TerminalNode {
	return s.GetToken(SintacticoMUT, 0)
}

func (s *DeclaracionContext) Listides() IListidesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListidesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListidesContext)
}

func (s *DeclaracionContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(SintacticoDOSPT, 0)
}

func (s *DeclaracionContext) Tipos() ITiposContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITiposContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITiposContext)
}

func (s *DeclaracionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(SintacticoIGUAL, 0)
}

func (s *DeclaracionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *Sintactico) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SintacticoRULE_declaracion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(SintacticoLET)
	}
	{
		p.SetState(42)
		p.Match(SintacticoMUT)
	}
	{
		p.SetState(43)

		var _x = p.Listides()

		localctx.(*DeclaracionContext)._listides = _x
	}
	{
		p.SetState(44)
		p.Match(SintacticoDOSPT)
	}
	{
		p.SetState(45)

		var _x = p.Tipos()

		localctx.(*DeclaracionContext)._tipos = _x
	}
	{
		p.SetState(46)
		p.Match(SintacticoIGUAL)
	}
	{
		p.SetState(47)

		var _x = p.Expr()

		localctx.(*DeclaracionContext)._expr = _x
	}
	localctx.(*DeclaracionContext).instr = instrucciones.NewDeclaracionValor(localctx.(*DeclaracionContext).Get_listides().GetLista(), localctx.(*DeclaracionContext).Get_tipos().GetTipo(), localctx.(*DeclaracionContext).Get_expr().GetP())

	return localctx
}

// IListidesContext is an interface to support dynamic dispatch.
type IListidesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetLista returns the lista attribute.
	GetLista() *arrayL.List

	// SetLista sets the lista attribute.
	SetLista(*arrayL.List)

	// IsListidesContext differentiates from other interfaces.
	IsListidesContext()
}

type ListidesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  *arrayL.List
	_ID    antlr.Token
}

func NewEmptyListidesContext() *ListidesContext {
	var p = new(ListidesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_listides
	return p
}

func (*ListidesContext) IsListidesContext() {}

func NewListidesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListidesContext {
	var p = new(ListidesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_listides

	return p
}

func (s *ListidesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListidesContext) Get_ID() antlr.Token { return s._ID }

func (s *ListidesContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ListidesContext) GetLista() *arrayL.List { return s.lista }

func (s *ListidesContext) SetLista(v *arrayL.List) { s.lista = v }

func (s *ListidesContext) ID() antlr.TerminalNode {
	return s.GetToken(SintacticoID, 0)
}

func (s *ListidesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListidesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListidesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterListides(s)
	}
}

func (s *ListidesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitListides(s)
	}
}

func (p *Sintactico) Listides() (localctx IListidesContext) {
	localctx = NewListidesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SintacticoRULE_listides)
	localctx.(*ListidesContext).lista = arrayL.New()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)

		var _m = p.Match(SintacticoID)

		localctx.(*ListidesContext)._ID = _m
	}
	localctx.(*ListidesContext).lista.Add(expresion.NewIdentificador((func() string {
		if localctx.(*ListidesContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ListidesContext).Get_ID().GetText()
		}
	}())))

	return localctx
}

// ITiposContext is an interface to support dynamic dispatch.
type ITiposContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTipo returns the tipo attribute.
	GetTipo() entorno.TipoDato

	// SetTipo sets the tipo attribute.
	SetTipo(entorno.TipoDato)

	// IsTiposContext differentiates from other interfaces.
	IsTiposContext()
}

type TiposContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tipo   entorno.TipoDato
}

func NewEmptyTiposContext() *TiposContext {
	var p = new(TiposContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_tipos
	return p
}

func (*TiposContext) IsTiposContext() {}

func NewTiposContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TiposContext {
	var p = new(TiposContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_tipos

	return p
}

func (s *TiposContext) GetParser() antlr.Parser { return s.parser }

func (s *TiposContext) GetTipo() entorno.TipoDato { return s.tipo }

func (s *TiposContext) SetTipo(v entorno.TipoDato) { s.tipo = v }

func (s *TiposContext) R_INT() antlr.TerminalNode {
	return s.GetToken(SintacticoR_INT, 0)
}

func (s *TiposContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(SintacticoR_FLOAT, 0)
}

func (s *TiposContext) R_STR() antlr.TerminalNode {
	return s.GetToken(SintacticoR_STR, 0)
}

func (s *TiposContext) R_CHAR() antlr.TerminalNode {
	return s.GetToken(SintacticoR_CHAR, 0)
}

func (s *TiposContext) R_BOOL() antlr.TerminalNode {
	return s.GetToken(SintacticoR_BOOL, 0)
}

func (s *TiposContext) R_STRING() antlr.TerminalNode {
	return s.GetToken(SintacticoR_STRING, 0)
}

func (s *TiposContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TiposContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TiposContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterTipos(s)
	}
}

func (s *TiposContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitTipos(s)
	}
}

func (p *Sintactico) Tipos() (localctx ITiposContext) {
	localctx = NewTiposContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SintacticoRULE_tipos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoR_INT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(SintacticoR_INT)
		}
		localctx.(*TiposContext).tipo = entorno.INT

	case SintacticoR_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Match(SintacticoR_FLOAT)
		}
		localctx.(*TiposContext).tipo = entorno.FLOAT

	case SintacticoR_STR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Match(SintacticoR_STR)
		}
		localctx.(*TiposContext).tipo = entorno.STRING

	case SintacticoR_CHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(59)
			p.Match(SintacticoR_CHAR)
		}
		localctx.(*TiposContext).tipo = entorno.CHAR

	case SintacticoR_BOOL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(61)
			p.Match(SintacticoR_BOOL)
		}
		localctx.(*TiposContext).tipo = entorno.BOOL

	case SintacticoR_STRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(63)
			p.Match(SintacticoR_STRING)
		}
		localctx.(*TiposContext).tipo = entorno.STRING

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_log returns the _expr_log rule contexts.
	Get_expr_log() IExpr_logContext

	// Get_expr_rel returns the _expr_rel rule contexts.
	Get_expr_rel() IExpr_relContext

	// Get_expr_op returns the _expr_op rule contexts.
	Get_expr_op() IExpr_opContext

	// Set_expr_log sets the _expr_log rule contexts.
	Set_expr_log(IExpr_logContext)

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// Set_expr_op sets the _expr_op rule contexts.
	Set_expr_op(IExpr_opContext)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         abstract.Expresion
	_expr_log IExpr_logContext
	_expr_rel IExpr_relContext
	_expr_op  IExpr_opContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Get_expr_log() IExpr_logContext { return s._expr_log }

func (s *ExprContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *ExprContext) Get_expr_op() IExpr_opContext { return s._expr_op }

func (s *ExprContext) Set_expr_log(v IExpr_logContext) { s._expr_log = v }

func (s *ExprContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *ExprContext) Set_expr_op(v IExpr_opContext) { s._expr_op = v }

func (s *ExprContext) GetP() abstract.Expresion { return s.p }

func (s *ExprContext) SetP(v abstract.Expresion) { s.p = v }

func (s *ExprContext) Expr_log() IExpr_logContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_logContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_logContext)
}

func (s *ExprContext) Expr_rel() IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *ExprContext) Expr_op() IExpr_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_opContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *Sintactico) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SintacticoRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(67)

			var _x = p.expr_log(0)

			localctx.(*ExprContext)._expr_log = _x
		}
		localctx.(*ExprContext).p = localctx.(*ExprContext).Get_expr_log().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)

			var _x = p.expr_rel(0)

			localctx.(*ExprContext)._expr_rel = _x
		}
		localctx.(*ExprContext).p = localctx.(*ExprContext).Get_expr_rel().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)

			var _x = p.expr_op(0)

			localctx.(*ExprContext)._expr_op = _x
		}
		localctx.(*ExprContext).p = localctx.(*ExprContext).Get_expr_op().GetP()

	}

	return localctx
}

// IExpr_opContext is an interface to support dynamic dispatch.
type IExpr_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetHIzq returns the hIzq rule contexts.
	GetHIzq() IExpr_opContext

	// GetOpU returns the opU rule contexts.
	GetOpU() IExprContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetHDer returns the hDer rule contexts.
	GetHDer() IExpr_opContext

	// Get_expr_val returns the _expr_val rule contexts.
	Get_expr_val() IExpr_valContext

	// SetHIzq sets the hIzq rule contexts.
	SetHIzq(IExpr_opContext)

	// SetOpU sets the opU rule contexts.
	SetOpU(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// SetHDer sets the hDer rule contexts.
	SetHDer(IExpr_opContext)

	// Set_expr_val sets the _expr_val rule contexts.
	Set_expr_val(IExpr_valContext)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsExpr_opContext differentiates from other interfaces.
	IsExpr_opContext()
}

type Expr_opContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         abstract.Expresion
	hIzq      IExpr_opContext
	opU       IExprContext
	_expr     IExprContext
	hDer      IExpr_opContext
	_expr_val IExpr_valContext
	op        antlr.Token
}

func NewEmptyExpr_opContext() *Expr_opContext {
	var p = new(Expr_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_expr_op
	return p
}

func (*Expr_opContext) IsExpr_opContext() {}

func NewExpr_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_opContext {
	var p = new(Expr_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_expr_op

	return p
}

func (s *Expr_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_opContext) GetOp() antlr.Token { return s.op }

func (s *Expr_opContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_opContext) GetHIzq() IExpr_opContext { return s.hIzq }

func (s *Expr_opContext) GetOpU() IExprContext { return s.opU }

func (s *Expr_opContext) Get_expr() IExprContext { return s._expr }

func (s *Expr_opContext) GetHDer() IExpr_opContext { return s.hDer }

func (s *Expr_opContext) Get_expr_val() IExpr_valContext { return s._expr_val }

func (s *Expr_opContext) SetHIzq(v IExpr_opContext) { s.hIzq = v }

func (s *Expr_opContext) SetOpU(v IExprContext) { s.opU = v }

func (s *Expr_opContext) Set_expr(v IExprContext) { s._expr = v }

func (s *Expr_opContext) SetHDer(v IExpr_opContext) { s.hDer = v }

func (s *Expr_opContext) Set_expr_val(v IExpr_valContext) { s._expr_val = v }

func (s *Expr_opContext) GetP() abstract.Expresion { return s.p }

func (s *Expr_opContext) SetP(v abstract.Expresion) { s.p = v }

func (s *Expr_opContext) MENOS() antlr.TerminalNode {
	return s.GetToken(SintacticoMENOS, 0)
}

func (s *Expr_opContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_opContext) PARENA() antlr.TerminalNode {
	return s.GetToken(SintacticoPARENA, 0)
}

func (s *Expr_opContext) PARENC() antlr.TerminalNode {
	return s.GetToken(SintacticoPARENC, 0)
}

func (s *Expr_opContext) R_INT() antlr.TerminalNode {
	return s.GetToken(SintacticoR_INT, 0)
}

func (s *Expr_opContext) CUATROPT() antlr.TerminalNode {
	return s.GetToken(SintacticoCUATROPT, 0)
}

func (s *Expr_opContext) POW() antlr.TerminalNode {
	return s.GetToken(SintacticoPOW, 0)
}

func (s *Expr_opContext) COMA() antlr.TerminalNode {
	return s.GetToken(SintacticoCOMA, 0)
}

func (s *Expr_opContext) AllExpr_op() []IExpr_opContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_opContext)(nil)).Elem())
	var tst = make([]IExpr_opContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_opContext)
		}
	}

	return tst
}

func (s *Expr_opContext) Expr_op(i int) IExpr_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_opContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_opContext)
}

func (s *Expr_opContext) R_FLOAT() antlr.TerminalNode {
	return s.GetToken(SintacticoR_FLOAT, 0)
}

func (s *Expr_opContext) Expr_val() IExpr_valContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_valContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_valContext)
}

func (s *Expr_opContext) POR() antlr.TerminalNode {
	return s.GetToken(SintacticoPOR, 0)
}

func (s *Expr_opContext) DIVIDIDO() antlr.TerminalNode {
	return s.GetToken(SintacticoDIVIDIDO, 0)
}

func (s *Expr_opContext) MAS() antlr.TerminalNode {
	return s.GetToken(SintacticoMAS, 0)
}

func (s *Expr_opContext) MODULO() antlr.TerminalNode {
	return s.GetToken(SintacticoMODULO, 0)
}

func (s *Expr_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterExpr_op(s)
	}
}

func (s *Expr_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitExpr_op(s)
	}
}

func (p *Sintactico) Expr_op() (localctx IExpr_opContext) {
	return p.expr_op(0)
}

func (p *Sintactico) expr_op(_p int) (localctx IExpr_opContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_opContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_opContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, SintacticoRULE_expr_op, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoMENOS:
		{
			p.SetState(79)
			p.Match(SintacticoMENOS)
		}
		{
			p.SetState(80)

			var _x = p.Expr()

			localctx.(*Expr_opContext).opU = _x
			localctx.(*Expr_opContext)._expr = _x
		}
		localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetOpU().GetP(), "-", nil, true)

	case SintacticoPARENA:
		{
			p.SetState(83)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(84)

			var _x = p.Expr()

			localctx.(*Expr_opContext)._expr = _x
		}
		{
			p.SetState(85)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = localctx.(*Expr_opContext).Get_expr().GetP()

	case SintacticoR_INT:
		{
			p.SetState(88)
			p.Match(SintacticoR_INT)
		}
		{
			p.SetState(89)
			p.Match(SintacticoCUATROPT)
		}
		{
			p.SetState(90)
			p.Match(SintacticoPOW)
		}
		{
			p.SetState(91)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(92)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hIzq = _x
		}
		{
			p.SetState(93)
			p.Match(SintacticoCOMA)
		}
		{
			p.SetState(94)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hDer = _x
		}
		{
			p.SetState(95)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), "**", localctx.(*Expr_opContext).GetHDer().GetP(), false)

	case SintacticoR_FLOAT:
		{
			p.SetState(98)
			p.Match(SintacticoR_FLOAT)
		}
		{
			p.SetState(99)
			p.Match(SintacticoCUATROPT)
		}
		{
			p.SetState(100)
			p.Match(SintacticoPOW)
		}
		{
			p.SetState(101)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(102)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hIzq = _x
		}
		{
			p.SetState(103)
			p.Match(SintacticoCOMA)
		}
		{
			p.SetState(104)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hDer = _x
		}
		{
			p.SetState(105)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), "f64**", localctx.(*Expr_opContext).GetHDer().GetP(), false)

	case SintacticoENTERO, SintacticoFLOAT, SintacticoCADENA, SintacticoCARACTER, SintacticoR_TRUE, SintacticoR_FALSE:
		{
			p.SetState(108)

			var _x = p.Expr_val()

			localctx.(*Expr_opContext)._expr_val = _x
		}
		localctx.(*Expr_opContext).p = localctx.(*Expr_opContext).Get_expr_val().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_opContext(p, _parentctx, _parentState)
				localctx.(*Expr_opContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_op)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(114)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_opContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SintacticoPOR || _la == SintacticoDIVIDIDO) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_opContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(115)

					var _x = p.expr_op(8)

					localctx.(*Expr_opContext).hDer = _x
				}
				localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), (func() string {
					if localctx.(*Expr_opContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_opContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_opContext).GetHDer().GetP(), false)

			case 2:
				localctx = NewExpr_opContext(p, _parentctx, _parentState)
				localctx.(*Expr_opContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_op)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(119)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_opContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SintacticoMAS || _la == SintacticoMENOS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_opContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(120)

					var _x = p.expr_op(7)

					localctx.(*Expr_opContext).hDer = _x
				}
				localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), (func() string {
					if localctx.(*Expr_opContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_opContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_opContext).GetHDer().GetP(), false)

			case 3:
				localctx = NewExpr_opContext(p, _parentctx, _parentState)
				localctx.(*Expr_opContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_op)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(124)
					p.Match(SintacticoMODULO)
				}
				{
					p.SetState(125)

					var _x = p.expr_op(6)

					localctx.(*Expr_opContext).hDer = _x
				}
				localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), (func() string {
					if localctx.(*Expr_opContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_opContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_opContext).GetHDer().GetP(), false)

			}

		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_logContext is an interface to support dynamic dispatch.
type IExpr_logContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHIzq returns the hIzq rule contexts.
	GetHIzq() IExpr_logContext

	// Get_expr_rel returns the _expr_rel rule contexts.
	Get_expr_rel() IExpr_relContext

	// GetHDer returns the hDer rule contexts.
	GetHDer() IExpr_logContext

	// SetHIzq sets the hIzq rule contexts.
	SetHIzq(IExpr_logContext)

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

	// SetHDer sets the hDer rule contexts.
	SetHDer(IExpr_logContext)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsExpr_logContext differentiates from other interfaces.
	IsExpr_logContext()
}

type Expr_logContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         abstract.Expresion
	hIzq      IExpr_logContext
	_expr_rel IExpr_relContext
	hDer      IExpr_logContext
}

func NewEmptyExpr_logContext() *Expr_logContext {
	var p = new(Expr_logContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_expr_log
	return p
}

func (*Expr_logContext) IsExpr_logContext() {}

func NewExpr_logContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_logContext {
	var p = new(Expr_logContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_expr_log

	return p
}

func (s *Expr_logContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_logContext) GetHIzq() IExpr_logContext { return s.hIzq }

func (s *Expr_logContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *Expr_logContext) GetHDer() IExpr_logContext { return s.hDer }

func (s *Expr_logContext) SetHIzq(v IExpr_logContext) { s.hIzq = v }

func (s *Expr_logContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *Expr_logContext) SetHDer(v IExpr_logContext) { s.hDer = v }

func (s *Expr_logContext) GetP() abstract.Expresion { return s.p }

func (s *Expr_logContext) SetP(v abstract.Expresion) { s.p = v }

func (s *Expr_logContext) NOT() antlr.TerminalNode {
	return s.GetToken(SintacticoNOT, 0)
}

func (s *Expr_logContext) AllExpr_log() []IExpr_logContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_logContext)(nil)).Elem())
	var tst = make([]IExpr_logContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_logContext)
		}
	}

	return tst
}

func (s *Expr_logContext) Expr_log(i int) IExpr_logContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_logContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_logContext)
}

func (s *Expr_logContext) Expr_rel() IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *Expr_logContext) OR() antlr.TerminalNode {
	return s.GetToken(SintacticoOR, 0)
}

func (s *Expr_logContext) AND() antlr.TerminalNode {
	return s.GetToken(SintacticoAND, 0)
}

func (s *Expr_logContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_logContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_logContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterExpr_log(s)
	}
}

func (s *Expr_logContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitExpr_log(s)
	}
}

func (p *Sintactico) Expr_log() (localctx IExpr_logContext) {
	return p.expr_log(0)
}

func (p *Sintactico) expr_log(_p int) (localctx IExpr_logContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_logContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_logContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, SintacticoRULE_expr_log, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoNOT:
		{
			p.SetState(134)
			p.Match(SintacticoNOT)
		}
		{
			p.SetState(135)

			var _x = p.expr_log(2)

			localctx.(*Expr_logContext).hIzq = _x
		}
		localctx.(*Expr_logContext).p = expresion.NewLogicas(localctx.(*Expr_logContext).GetHIzq().GetP(), "!", nil, true)

	case SintacticoR_INT, SintacticoR_FLOAT, SintacticoPARENA, SintacticoMENOS, SintacticoENTERO, SintacticoFLOAT, SintacticoCADENA, SintacticoCARACTER, SintacticoR_TRUE, SintacticoR_FALSE:
		{
			p.SetState(138)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logContext)._expr_rel = _x
		}
		localctx.(*Expr_logContext).p = localctx.(*Expr_logContext).Get_expr_rel().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_logContext(p, _parentctx, _parentState)
				localctx.(*Expr_logContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_log)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(144)
					p.Match(SintacticoOR)
				}
				{
					p.SetState(145)

					var _x = p.expr_log(5)

					localctx.(*Expr_logContext).hDer = _x
				}
				localctx.(*Expr_logContext).p = expresion.NewLogicas(localctx.(*Expr_logContext).GetHIzq().GetP(), "||", localctx.(*Expr_logContext).GetHDer().GetP(), false)

			case 2:
				localctx = NewExpr_logContext(p, _parentctx, _parentState)
				localctx.(*Expr_logContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_log)
				p.SetState(148)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(149)
					p.Match(SintacticoAND)
				}
				{
					p.SetState(150)

					var _x = p.expr_log(4)

					localctx.(*Expr_logContext).hDer = _x
				}
				localctx.(*Expr_logContext).p = expresion.NewLogicas(localctx.(*Expr_logContext).GetHIzq().GetP(), "&&", localctx.(*Expr_logContext).GetHDer().GetP(), false)

			}

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_relContext is an interface to support dynamic dispatch.
type IExpr_relContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetHIzq returns the hIzq rule contexts.
	GetHIzq() IExpr_relContext

	// Get_expr_op returns the _expr_op rule contexts.
	Get_expr_op() IExpr_opContext

	// GetHDer returns the hDer rule contexts.
	GetHDer() IExpr_relContext

	// SetHIzq sets the hIzq rule contexts.
	SetHIzq(IExpr_relContext)

	// Set_expr_op sets the _expr_op rule contexts.
	Set_expr_op(IExpr_opContext)

	// SetHDer sets the hDer rule contexts.
	SetHDer(IExpr_relContext)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsExpr_relContext differentiates from other interfaces.
	IsExpr_relContext()
}

type Expr_relContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	p        abstract.Expresion
	hIzq     IExpr_relContext
	_expr_op IExpr_opContext
	op       antlr.Token
	hDer     IExpr_relContext
}

func NewEmptyExpr_relContext() *Expr_relContext {
	var p = new(Expr_relContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_expr_rel
	return p
}

func (*Expr_relContext) IsExpr_relContext() {}

func NewExpr_relContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_relContext {
	var p = new(Expr_relContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_expr_rel

	return p
}

func (s *Expr_relContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_relContext) GetOp() antlr.Token { return s.op }

func (s *Expr_relContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_relContext) GetHIzq() IExpr_relContext { return s.hIzq }

func (s *Expr_relContext) Get_expr_op() IExpr_opContext { return s._expr_op }

func (s *Expr_relContext) GetHDer() IExpr_relContext { return s.hDer }

func (s *Expr_relContext) SetHIzq(v IExpr_relContext) { s.hIzq = v }

func (s *Expr_relContext) Set_expr_op(v IExpr_opContext) { s._expr_op = v }

func (s *Expr_relContext) SetHDer(v IExpr_relContext) { s.hDer = v }

func (s *Expr_relContext) GetP() abstract.Expresion { return s.p }

func (s *Expr_relContext) SetP(v abstract.Expresion) { s.p = v }

func (s *Expr_relContext) Expr_op() IExpr_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_opContext)
}

func (s *Expr_relContext) AllExpr_rel() []IExpr_relContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_relContext)(nil)).Elem())
	var tst = make([]IExpr_relContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_relContext)
		}
	}

	return tst
}

func (s *Expr_relContext) Expr_rel(i int) IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *Expr_relContext) IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(SintacticoIGUALIGUAL, 0)
}

func (s *Expr_relContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(SintacticoDIFERENTE, 0)
}

func (s *Expr_relContext) MENORQUE() antlr.TerminalNode {
	return s.GetToken(SintacticoMENORQUE, 0)
}

func (s *Expr_relContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(SintacticoMENORIGUAL, 0)
}

func (s *Expr_relContext) MAYORQUE() antlr.TerminalNode {
	return s.GetToken(SintacticoMAYORQUE, 0)
}

func (s *Expr_relContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(SintacticoMAYORIGUAL, 0)
}

func (s *Expr_relContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_relContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_relContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterExpr_rel(s)
	}
}

func (s *Expr_relContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitExpr_rel(s)
	}
}

func (p *Sintactico) Expr_rel() (localctx IExpr_relContext) {
	return p.expr_rel(0)
}

func (p *Sintactico) expr_rel(_p int) (localctx IExpr_relContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_relContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_relContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, SintacticoRULE_expr_rel, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)

		var _x = p.expr_op(0)

		localctx.(*Expr_relContext)._expr_op = _x
	}
	localctx.(*Expr_relContext).p = localctx.(*Expr_relContext).Get_expr_op().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_relContext(p, _parentctx, _parentState)
			localctx.(*Expr_relContext).hIzq = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_rel)
			p.SetState(162)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(163)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expr_relContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SintacticoDIFERENTE)|(1<<SintacticoMAYORIGUAL)|(1<<SintacticoMENORIGUAL)|(1<<SintacticoMAYORQUE)|(1<<SintacticoMENORQUE)|(1<<SintacticoIGUALIGUAL))) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expr_relContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(164)

				var _x = p.expr_rel(3)

				localctx.(*Expr_relContext).hDer = _x
			}
			localctx.(*Expr_relContext).p = expresion.NewRelacional(localctx.(*Expr_relContext).GetHIzq().GetP(), (func() string {
				if localctx.(*Expr_relContext).GetOp() == nil {
					return ""
				} else {
					return localctx.(*Expr_relContext).GetOp().GetText()
				}
			}()), localctx.(*Expr_relContext).GetHDer().GetP())

		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_valContext is an interface to support dynamic dispatch.
type IExpr_valContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsExpr_valContext differentiates from other interfaces.
	IsExpr_valContext()
}

type Expr_valContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          abstract.Expresion
	_primitivo IPrimitivoContext
}

func NewEmptyExpr_valContext() *Expr_valContext {
	var p = new(Expr_valContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_expr_val
	return p
}

func (*Expr_valContext) IsExpr_valContext() {}

func NewExpr_valContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_valContext {
	var p = new(Expr_valContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_expr_val

	return p
}

func (s *Expr_valContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_valContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expr_valContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expr_valContext) GetP() abstract.Expresion { return s.p }

func (s *Expr_valContext) SetP(v abstract.Expresion) { s.p = v }

func (s *Expr_valContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Expr_valContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_valContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_valContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterExpr_val(s)
	}
}

func (s *Expr_valContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitExpr_val(s)
	}
}

func (p *Sintactico) Expr_val() (localctx IExpr_valContext) {
	localctx = NewExpr_valContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SintacticoRULE_expr_val)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)

		var _x = p.Primitivo()

		localctx.(*Expr_valContext)._primitivo = _x
	}
	localctx.(*Expr_valContext).p = localctx.(*Expr_valContext).Get_primitivo().GetP()

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_R_TRUE returns the _R_TRUE token.
	Get_R_TRUE() antlr.Token

	// Get_R_FALSE returns the _R_FALSE token.
	Get_R_FALSE() antlr.Token

	// Get_CARACTER returns the _CARACTER token.
	Get_CARACTER() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_R_TRUE sets the _R_TRUE token.
	Set_R_TRUE(antlr.Token)

	// Set_R_FALSE sets the _R_FALSE token.
	Set_R_FALSE(antlr.Token)

	// Set_CARACTER sets the _CARACTER token.
	Set_CARACTER(antlr.Token)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         abstract.Expresion
	_ENTERO   antlr.Token
	_CADENA   antlr.Token
	_FLOAT    antlr.Token
	_R_TRUE   antlr.Token
	_R_FALSE  antlr.Token
	_CARACTER antlr.Token
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SintacticoRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SintacticoRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *PrimitivoContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrimitivoContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *PrimitivoContext) Get_R_TRUE() antlr.Token { return s._R_TRUE }

func (s *PrimitivoContext) Get_R_FALSE() antlr.Token { return s._R_FALSE }

func (s *PrimitivoContext) Get_CARACTER() antlr.Token { return s._CARACTER }

func (s *PrimitivoContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *PrimitivoContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrimitivoContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *PrimitivoContext) Set_R_TRUE(v antlr.Token) { s._R_TRUE = v }

func (s *PrimitivoContext) Set_R_FALSE(v antlr.Token) { s._R_FALSE = v }

func (s *PrimitivoContext) Set_CARACTER(v antlr.Token) { s._CARACTER = v }

func (s *PrimitivoContext) GetP() abstract.Expresion { return s.p }

func (s *PrimitivoContext) SetP(v abstract.Expresion) { s.p = v }

func (s *PrimitivoContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(SintacticoENTERO, 0)
}

func (s *PrimitivoContext) CADENA() antlr.TerminalNode {
	return s.GetToken(SintacticoCADENA, 0)
}

func (s *PrimitivoContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(SintacticoFLOAT, 0)
}

func (s *PrimitivoContext) R_TRUE() antlr.TerminalNode {
	return s.GetToken(SintacticoR_TRUE, 0)
}

func (s *PrimitivoContext) R_FALSE() antlr.TerminalNode {
	return s.GetToken(SintacticoR_FALSE, 0)
}

func (s *PrimitivoContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(SintacticoCARACTER, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SintacticoListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Sintactico) Primitivo() (localctx IPrimitivoContext) {
	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SintacticoRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)

			var _m = p.Match(SintacticoENTERO)

			localctx.(*PrimitivoContext)._ENTERO = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_ENTERO() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_ENTERO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(num, entorno.INT)

	case SintacticoCADENA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)

			var _m = p.Match(SintacticoCADENA)

			localctx.(*PrimitivoContext)._CADENA = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CADENA().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(str, entorno.STRING)

	case SintacticoFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(179)

			var _m = p.Match(SintacticoFLOAT)

			localctx.(*PrimitivoContext)._FLOAT = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_FLOAT().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(num, entorno.FLOAT)

	case SintacticoR_TRUE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(181)

			var _m = p.Match(SintacticoR_TRUE)

			localctx.(*PrimitivoContext)._R_TRUE = _m
		}

		bool_true, err := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_R_TRUE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_R_TRUE().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(bool_true, entorno.BOOL)

	case SintacticoR_FALSE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(183)

			var _m = p.Match(SintacticoR_FALSE)

			localctx.(*PrimitivoContext)._R_FALSE = _m
		}

		bool_false, err := strconv.ParseBool((func() string {
			if localctx.(*PrimitivoContext).Get_R_FALSE() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_R_FALSE().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(bool_false, entorno.BOOL)

	case SintacticoCARACTER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(185)

			var _m = p.Match(SintacticoCARACTER)

			localctx.(*PrimitivoContext)._CARACTER = _m
		}

		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CARACTER().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CARACTER().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(str, entorno.STRING)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Sintactico) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *Expr_opContext = nil
		if localctx != nil {
			t = localctx.(*Expr_opContext)
		}
		return p.Expr_op_Sempred(t, predIndex)

	case 8:
		var t *Expr_logContext = nil
		if localctx != nil {
			t = localctx.(*Expr_logContext)
		}
		return p.Expr_log_Sempred(t, predIndex)

	case 9:
		var t *Expr_relContext = nil
		if localctx != nil {
			t = localctx.(*Expr_relContext)
		}
		return p.Expr_rel_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Sintactico) Expr_op_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Sintactico) Expr_log_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Sintactico) Expr_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
