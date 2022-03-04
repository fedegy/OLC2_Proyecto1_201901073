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

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 34, 146,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 23, 10, 3, 12, 3,
	14, 3, 26, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 46, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 81, 10, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 98, 10, 6, 12, 6, 14, 6, 101, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 120, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 7, 8, 131, 10, 8, 12, 8, 14, 8, 134, 11, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 144, 10, 9, 3, 9, 2, 4, 10,
	14, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 5, 3, 2, 25, 26, 3, 2, 23, 24, 3,
	2, 14, 19, 2, 154, 2, 18, 3, 2, 2, 2, 4, 24, 3, 2, 2, 2, 6, 29, 3, 2, 2,
	2, 8, 45, 3, 2, 2, 2, 10, 80, 3, 2, 2, 2, 12, 119, 3, 2, 2, 2, 14, 121,
	3, 2, 2, 2, 16, 143, 3, 2, 2, 2, 18, 19, 5, 4, 3, 2, 19, 20, 8, 2, 1, 2,
	20, 3, 3, 2, 2, 2, 21, 23, 5, 6, 4, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2,
	2, 2, 24, 22, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 27, 3, 2, 2, 2, 26, 24,
	3, 2, 2, 2, 27, 28, 8, 3, 1, 2, 28, 5, 3, 2, 2, 2, 29, 30, 7, 5, 2, 2,
	30, 31, 7, 3, 2, 2, 31, 32, 5, 8, 5, 2, 32, 33, 7, 4, 2, 2, 33, 34, 7,
	21, 2, 2, 34, 35, 8, 4, 1, 2, 35, 7, 3, 2, 2, 2, 36, 37, 5, 12, 7, 2, 37,
	38, 8, 5, 1, 2, 38, 46, 3, 2, 2, 2, 39, 40, 5, 14, 8, 2, 40, 41, 8, 5,
	1, 2, 41, 46, 3, 2, 2, 2, 42, 43, 5, 10, 6, 2, 43, 44, 8, 5, 1, 2, 44,
	46, 3, 2, 2, 2, 45, 36, 3, 2, 2, 2, 45, 39, 3, 2, 2, 2, 45, 42, 3, 2, 2,
	2, 46, 9, 3, 2, 2, 2, 47, 48, 8, 6, 1, 2, 48, 49, 7, 24, 2, 2, 49, 50,
	5, 8, 5, 2, 50, 51, 8, 6, 1, 2, 51, 81, 3, 2, 2, 2, 52, 53, 5, 16, 9, 2,
	53, 54, 8, 6, 1, 2, 54, 81, 3, 2, 2, 2, 55, 56, 7, 3, 2, 2, 56, 57, 5,
	8, 5, 2, 57, 58, 7, 4, 2, 2, 58, 59, 8, 6, 1, 2, 59, 81, 3, 2, 2, 2, 60,
	61, 7, 6, 2, 2, 61, 62, 7, 20, 2, 2, 62, 63, 7, 10, 2, 2, 63, 64, 7, 3,
	2, 2, 64, 65, 5, 10, 6, 2, 65, 66, 7, 22, 2, 2, 66, 67, 5, 10, 6, 2, 67,
	68, 7, 4, 2, 2, 68, 69, 8, 6, 1, 2, 69, 81, 3, 2, 2, 2, 70, 71, 7, 7, 2,
	2, 71, 72, 7, 20, 2, 2, 72, 73, 7, 10, 2, 2, 73, 74, 7, 3, 2, 2, 74, 75,
	5, 10, 6, 2, 75, 76, 7, 22, 2, 2, 76, 77, 5, 10, 6, 2, 77, 78, 7, 4, 2,
	2, 78, 79, 8, 6, 1, 2, 79, 81, 3, 2, 2, 2, 80, 47, 3, 2, 2, 2, 80, 52,
	3, 2, 2, 2, 80, 55, 3, 2, 2, 2, 80, 60, 3, 2, 2, 2, 80, 70, 3, 2, 2, 2,
	81, 99, 3, 2, 2, 2, 82, 83, 12, 9, 2, 2, 83, 84, 9, 2, 2, 2, 84, 85, 5,
	10, 6, 10, 85, 86, 8, 6, 1, 2, 86, 98, 3, 2, 2, 2, 87, 88, 12, 8, 2, 2,
	88, 89, 9, 3, 2, 2, 89, 90, 5, 10, 6, 9, 90, 91, 8, 6, 1, 2, 91, 98, 3,
	2, 2, 2, 92, 93, 12, 7, 2, 2, 93, 94, 7, 27, 2, 2, 94, 95, 5, 10, 6, 8,
	95, 96, 8, 6, 1, 2, 96, 98, 3, 2, 2, 2, 97, 82, 3, 2, 2, 2, 97, 87, 3,
	2, 2, 2, 97, 92, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99,
	100, 3, 2, 2, 2, 100, 11, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 103, 5,
	14, 8, 2, 103, 104, 7, 12, 2, 2, 104, 105, 5, 14, 8, 2, 105, 106, 8, 7,
	1, 2, 106, 120, 3, 2, 2, 2, 107, 108, 5, 14, 8, 2, 108, 109, 7, 11, 2,
	2, 109, 110, 5, 14, 8, 2, 110, 111, 8, 7, 1, 2, 111, 120, 3, 2, 2, 2, 112,
	113, 7, 13, 2, 2, 113, 114, 5, 14, 8, 2, 114, 115, 8, 7, 1, 2, 115, 120,
	3, 2, 2, 2, 116, 117, 5, 14, 8, 2, 117, 118, 8, 7, 1, 2, 118, 120, 3, 2,
	2, 2, 119, 102, 3, 2, 2, 2, 119, 107, 3, 2, 2, 2, 119, 112, 3, 2, 2, 2,
	119, 116, 3, 2, 2, 2, 120, 13, 3, 2, 2, 2, 121, 122, 8, 8, 1, 2, 122, 123,
	5, 10, 6, 2, 123, 124, 8, 8, 1, 2, 124, 132, 3, 2, 2, 2, 125, 126, 12,
	4, 2, 2, 126, 127, 9, 4, 2, 2, 127, 128, 5, 14, 8, 5, 128, 129, 8, 8, 1,
	2, 129, 131, 3, 2, 2, 2, 130, 125, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132,
	130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 15, 3, 2, 2, 2, 134, 132, 3,
	2, 2, 2, 135, 136, 7, 28, 2, 2, 136, 144, 8, 9, 1, 2, 137, 138, 7, 30,
	2, 2, 138, 144, 8, 9, 1, 2, 139, 140, 7, 29, 2, 2, 140, 144, 8, 9, 1, 2,
	141, 142, 7, 31, 2, 2, 142, 144, 8, 9, 1, 2, 143, 135, 3, 2, 2, 2, 143,
	137, 3, 2, 2, 2, 143, 139, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 144, 17, 3,
	2, 2, 2, 10, 24, 45, 80, 97, 99, 119, 132, 143,
}
var literalNames = []string{
	"", "'('", "')'", "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", "'pow'",
	"'&&'", "'||'", "'!'", "'!='", "'>='", "'<='", "'>'", "'<'", "'=='", "'::'",
	"';'", "','", "'+'", "'-'", "'*'", "'/'", "'%'", "", "", "", "'true'",
	"'false'",
}
var symbolicNames = []string{
	"", "PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL",
	"POW", "AND", "OR", "NOT", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE",
	"MENORQUE", "IGUALIGUAL", "CUATROPT", "PTCOMA", "COMA", "MAS", "MENOS",
	"POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT", "CADENA", "R_TRUE", "R_FALSE",
	"ID", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "expr", "expr_op", "expr_log",
	"expr_rel", "primitivo",
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
	SintacticoPARENA     = 1
	SintacticoPARENC     = 2
	SintacticoPRINTLN    = 3
	SintacticoR_INT      = 4
	SintacticoR_FLOAT    = 5
	SintacticoR_STRING   = 6
	SintacticoR_BOOL     = 7
	SintacticoPOW        = 8
	SintacticoAND        = 9
	SintacticoOR         = 10
	SintacticoNOT        = 11
	SintacticoDIFERENTE  = 12
	SintacticoMAYORIGUAL = 13
	SintacticoMENORIGUAL = 14
	SintacticoMAYORQUE   = 15
	SintacticoMENORQUE   = 16
	SintacticoIGUALIGUAL = 17
	SintacticoCUATROPT   = 18
	SintacticoPTCOMA     = 19
	SintacticoCOMA       = 20
	SintacticoMAS        = 21
	SintacticoMENOS      = 22
	SintacticoPOR        = 23
	SintacticoDIVIDIDO   = 24
	SintacticoMODULO     = 25
	SintacticoENTERO     = 26
	SintacticoFLOAT      = 27
	SintacticoCADENA     = 28
	SintacticoR_TRUE     = 29
	SintacticoR_FALSE    = 30
	SintacticoID         = 31
	SintacticoWHITESPACE = 32
)

// Sintactico rules.
const (
	SintacticoRULE_start         = 0
	SintacticoRULE_instrucciones = 1
	SintacticoRULE_instruccion   = 2
	SintacticoRULE_expr          = 3
	SintacticoRULE_expr_op       = 4
	SintacticoRULE_expr_log      = 5
	SintacticoRULE_expr_rel      = 6
	SintacticoRULE_primitivo     = 7
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

	// GetLista returns the lista attribute.
	GetLista() *arrayL.List

	// SetLista sets the lista attribute.
	SetLista(*arrayL.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayL.List
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

func (s *StartContext) GetLista() *arrayL.List { return s.lista }

func (s *StartContext) SetLista(v *arrayL.List) { s.lista = v }

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
		p.SetState(16)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

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
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SintacticoPRINTLN {
		{
			p.SetState(19)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(24)
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
		p.SetState(27)
		p.Match(SintacticoPRINTLN)
	}
	{
		p.SetState(28)
		p.Match(SintacticoPARENA)
	}
	{
		p.SetState(29)

		var _x = p.Expr()

		localctx.(*InstruccionContext)._expr = _x
	}
	{
		p.SetState(30)
		p.Match(SintacticoPARENC)
	}
	{
		p.SetState(31)
		p.Match(SintacticoPTCOMA)
	}
	localctx.(*InstruccionContext).instr = instrucciones.NewImprimir(localctx.(*InstruccionContext).Get_expr().GetP())

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
	p.EnterRule(localctx, 6, SintacticoRULE_expr)

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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)

			var _x = p.Expr_log()

			localctx.(*ExprContext)._expr_log = _x
		}
		localctx.(*ExprContext).p = localctx.(*ExprContext).Get_expr_log().GetP()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)

			var _x = p.expr_rel(0)

			localctx.(*ExprContext)._expr_rel = _x
		}
		localctx.(*ExprContext).p = localctx.(*ExprContext).Get_expr_rel().GetP()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)

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

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// GetHDer returns the hDer rule contexts.
	GetHDer() IExpr_opContext

	// SetHIzq sets the hIzq rule contexts.
	SetHIzq(IExpr_opContext)

	// SetOpU sets the opU rule contexts.
	SetOpU(IExprContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// SetHDer sets the hDer rule contexts.
	SetHDer(IExpr_opContext)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsExpr_opContext differentiates from other interfaces.
	IsExpr_opContext()
}

type Expr_opContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          abstract.Expresion
	hIzq       IExpr_opContext
	opU        IExprContext
	_expr      IExprContext
	_primitivo IPrimitivoContext
	hDer       IExpr_opContext
	op         antlr.Token
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

func (s *Expr_opContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expr_opContext) GetHDer() IExpr_opContext { return s.hDer }

func (s *Expr_opContext) SetHIzq(v IExpr_opContext) { s.hIzq = v }

func (s *Expr_opContext) SetOpU(v IExprContext) { s.opU = v }

func (s *Expr_opContext) Set_expr(v IExprContext) { s._expr = v }

func (s *Expr_opContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expr_opContext) SetHDer(v IExpr_opContext) { s.hDer = v }

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

func (s *Expr_opContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, SintacticoRULE_expr_op, _p)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoMENOS:
		{
			p.SetState(46)
			p.Match(SintacticoMENOS)
		}
		{
			p.SetState(47)

			var _x = p.Expr()

			localctx.(*Expr_opContext).opU = _x
			localctx.(*Expr_opContext)._expr = _x
		}
		localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetOpU().GetP(), "-", nil, true)

	case SintacticoENTERO, SintacticoFLOAT, SintacticoCADENA, SintacticoR_TRUE:
		{
			p.SetState(50)

			var _x = p.Primitivo()

			localctx.(*Expr_opContext)._primitivo = _x
		}
		localctx.(*Expr_opContext).p = localctx.(*Expr_opContext).Get_primitivo().GetP()

	case SintacticoPARENA:
		{
			p.SetState(53)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(54)

			var _x = p.Expr()

			localctx.(*Expr_opContext)._expr = _x
		}
		{
			p.SetState(55)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = localctx.(*Expr_opContext).Get_expr().GetP()

	case SintacticoR_INT:
		{
			p.SetState(58)
			p.Match(SintacticoR_INT)
		}
		{
			p.SetState(59)
			p.Match(SintacticoCUATROPT)
		}
		{
			p.SetState(60)
			p.Match(SintacticoPOW)
		}
		{
			p.SetState(61)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(62)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hIzq = _x
		}
		{
			p.SetState(63)
			p.Match(SintacticoCOMA)
		}
		{
			p.SetState(64)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hDer = _x
		}
		{
			p.SetState(65)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), "**", localctx.(*Expr_opContext).GetHDer().GetP(), false)

	case SintacticoR_FLOAT:
		{
			p.SetState(68)
			p.Match(SintacticoR_FLOAT)
		}
		{
			p.SetState(69)
			p.Match(SintacticoCUATROPT)
		}
		{
			p.SetState(70)
			p.Match(SintacticoPOW)
		}
		{
			p.SetState(71)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(72)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hIzq = _x
		}
		{
			p.SetState(73)
			p.Match(SintacticoCOMA)
		}
		{
			p.SetState(74)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hDer = _x
		}
		{
			p.SetState(75)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), "f64**", localctx.(*Expr_opContext).GetHDer().GetP(), false)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_opContext(p, _parentctx, _parentState)
				localctx.(*Expr_opContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_op)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(81)

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
					p.SetState(82)

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
				p.SetState(85)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(86)

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
					p.SetState(87)

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
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(91)
					p.Match(SintacticoMODULO)
				}
				{
					p.SetState(92)

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
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IExpr_logContext is an interface to support dynamic dispatch.
type IExpr_logContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHIzq returns the hIzq rule contexts.
	GetHIzq() IExpr_relContext

	// GetHDer returns the hDer rule contexts.
	GetHDer() IExpr_relContext

	// Get_expr_rel returns the _expr_rel rule contexts.
	Get_expr_rel() IExpr_relContext

	// SetHIzq sets the hIzq rule contexts.
	SetHIzq(IExpr_relContext)

	// SetHDer sets the hDer rule contexts.
	SetHDer(IExpr_relContext)

	// Set_expr_rel sets the _expr_rel rule contexts.
	Set_expr_rel(IExpr_relContext)

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
	hIzq      IExpr_relContext
	hDer      IExpr_relContext
	_expr_rel IExpr_relContext
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

func (s *Expr_logContext) GetHIzq() IExpr_relContext { return s.hIzq }

func (s *Expr_logContext) GetHDer() IExpr_relContext { return s.hDer }

func (s *Expr_logContext) Get_expr_rel() IExpr_relContext { return s._expr_rel }

func (s *Expr_logContext) SetHIzq(v IExpr_relContext) { s.hIzq = v }

func (s *Expr_logContext) SetHDer(v IExpr_relContext) { s.hDer = v }

func (s *Expr_logContext) Set_expr_rel(v IExpr_relContext) { s._expr_rel = v }

func (s *Expr_logContext) GetP() abstract.Expresion { return s.p }

func (s *Expr_logContext) SetP(v abstract.Expresion) { s.p = v }

func (s *Expr_logContext) OR() antlr.TerminalNode {
	return s.GetToken(SintacticoOR, 0)
}

func (s *Expr_logContext) AllExpr_rel() []IExpr_relContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_relContext)(nil)).Elem())
	var tst = make([]IExpr_relContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_relContext)
		}
	}

	return tst
}

func (s *Expr_logContext) Expr_rel(i int) IExpr_relContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_relContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_relContext)
}

func (s *Expr_logContext) AND() antlr.TerminalNode {
	return s.GetToken(SintacticoAND, 0)
}

func (s *Expr_logContext) NOT() antlr.TerminalNode {
	return s.GetToken(SintacticoNOT, 0)
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
	localctx = NewExpr_logContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SintacticoRULE_expr_log)

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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logContext).hIzq = _x
		}
		{
			p.SetState(101)
			p.Match(SintacticoOR)
		}
		{
			p.SetState(102)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logContext).hDer = _x
		}
		localctx.(*Expr_logContext).p = expresion.NewLogicas(localctx.(*Expr_logContext).GetHIzq().GetP(), "||", localctx.(*Expr_logContext).GetHDer().GetP(), false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logContext).hIzq = _x
		}
		{
			p.SetState(106)
			p.Match(SintacticoAND)
		}
		{
			p.SetState(107)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logContext).hDer = _x
		}
		localctx.(*Expr_logContext).p = expresion.NewLogicas(localctx.(*Expr_logContext).GetHIzq().GetP(), "&&", localctx.(*Expr_logContext).GetHDer().GetP(), false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(SintacticoNOT)
		}
		{
			p.SetState(111)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logContext).hIzq = _x
		}
		localctx.(*Expr_logContext).p = expresion.NewLogicas(localctx.(*Expr_logContext).GetHIzq().GetP(), "!", nil, true)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)

			var _x = p.expr_rel(0)

			localctx.(*Expr_logContext)._expr_rel = _x
		}
		localctx.(*Expr_logContext).p = localctx.(*Expr_logContext).Get_expr_rel().GetP()

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
	_startState := 12
	p.EnterRecursionRule(localctx, 12, SintacticoRULE_expr_rel, _p)
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
		p.SetState(120)

		var _x = p.expr_op(0)

		localctx.(*Expr_relContext)._expr_op = _x
	}
	localctx.(*Expr_relContext).p = localctx.(*Expr_relContext).Get_expr_op().GetP()

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_relContext(p, _parentctx, _parentState)
			localctx.(*Expr_relContext).hIzq = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_rel)
			p.SetState(123)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(124)

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
				p.SetState(125)

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
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

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

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_R_TRUE sets the _R_TRUE token.
	Set_R_TRUE(antlr.Token)

	// GetP returns the p attribute.
	GetP() abstract.Expresion

	// SetP sets the p attribute.
	SetP(abstract.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	p       abstract.Expresion
	_ENTERO antlr.Token
	_CADENA antlr.Token
	_FLOAT  antlr.Token
	_R_TRUE antlr.Token
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

func (s *PrimitivoContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *PrimitivoContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrimitivoContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *PrimitivoContext) Set_R_TRUE(v antlr.Token) { s._R_TRUE = v }

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
	p.EnterRule(localctx, 14, SintacticoRULE_primitivo)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)

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
			p.SetState(135)

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
			p.SetState(137)

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
			p.SetState(139)

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

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Sintactico) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *Expr_opContext = nil
		if localctx != nil {
			t = localctx.(*Expr_opContext)
		}
		return p.Expr_op_Sempred(t, predIndex)

	case 6:
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

func (p *Sintactico) Expr_rel_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
