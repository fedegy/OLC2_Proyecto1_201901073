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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 25, 78, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 3, 7, 3, 19, 10, 3, 12, 3, 14, 3, 22, 11, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 55, 10, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 67, 10, 6, 12, 6, 14, 6, 70,
	11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 76, 10, 7, 3, 7, 2, 3, 10, 8, 2, 4,
	6, 8, 10, 12, 2, 5, 3, 2, 13, 13, 3, 2, 16, 17, 3, 2, 14, 15, 2, 77, 2,
	14, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 25, 3, 2, 2, 2, 8, 32, 3, 2, 2, 2,
	10, 54, 3, 2, 2, 2, 12, 75, 3, 2, 2, 2, 14, 15, 5, 4, 3, 2, 15, 16, 8,
	2, 1, 2, 16, 3, 3, 2, 2, 2, 17, 19, 5, 6, 4, 2, 18, 17, 3, 2, 2, 2, 19,
	22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2,
	2, 22, 20, 3, 2, 2, 2, 23, 24, 8, 3, 1, 2, 24, 5, 3, 2, 2, 2, 25, 26, 7,
	5, 2, 2, 26, 27, 7, 3, 2, 2, 27, 28, 5, 8, 5, 2, 28, 29, 7, 4, 2, 2, 29,
	30, 7, 12, 2, 2, 30, 31, 8, 4, 1, 2, 31, 7, 3, 2, 2, 2, 32, 33, 5, 10,
	6, 2, 33, 34, 8, 5, 1, 2, 34, 9, 3, 2, 2, 2, 35, 36, 8, 6, 1, 2, 36, 37,
	5, 12, 7, 2, 37, 38, 8, 6, 1, 2, 38, 55, 3, 2, 2, 2, 39, 40, 7, 3, 2, 2,
	40, 41, 5, 8, 5, 2, 41, 42, 7, 4, 2, 2, 42, 43, 8, 6, 1, 2, 43, 55, 3,
	2, 2, 2, 44, 45, 7, 6, 2, 2, 45, 46, 7, 11, 2, 2, 46, 47, 7, 10, 2, 2,
	47, 48, 7, 3, 2, 2, 48, 49, 5, 10, 6, 2, 49, 50, 9, 2, 2, 2, 50, 51, 5,
	10, 6, 2, 51, 52, 7, 4, 2, 2, 52, 53, 8, 6, 1, 2, 53, 55, 3, 2, 2, 2, 54,
	35, 3, 2, 2, 2, 54, 39, 3, 2, 2, 2, 54, 44, 3, 2, 2, 2, 55, 68, 3, 2, 2,
	2, 56, 57, 12, 7, 2, 2, 57, 58, 9, 3, 2, 2, 58, 59, 5, 10, 6, 8, 59, 60,
	8, 6, 1, 2, 60, 67, 3, 2, 2, 2, 61, 62, 12, 6, 2, 2, 62, 63, 9, 4, 2, 2,
	63, 64, 5, 10, 6, 7, 64, 65, 8, 6, 1, 2, 65, 67, 3, 2, 2, 2, 66, 56, 3,
	2, 2, 2, 66, 61, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68,
	69, 3, 2, 2, 2, 69, 11, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 72, 7, 19,
	2, 2, 72, 76, 8, 7, 1, 2, 73, 74, 7, 21, 2, 2, 74, 76, 8, 7, 1, 2, 75,
	71, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 13, 3, 2, 2, 2, 7, 20, 54, 66,
	68, 75,
}
var literalNames = []string{
	"", "'('", "')'", "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", "'pow'",
	"'::'", "';'", "','", "'+'", "'-'", "'*'", "'/'", "'%'", "", "", "", "'true'",
	"'false'",
}
var symbolicNames = []string{
	"", "PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL",
	"POW", "CUATROPT", "PTCOMA", "COMA", "MAS", "MENOS", "POR", "DIVIDIDO",
	"MODULO", "ENTERO", "FLOAT", "CADENA", "R_TRUE", "R_FALSE", "ID", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "expr", "expr_op", "primitivo",
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
	SintacticoCUATROPT   = 9
	SintacticoPTCOMA     = 10
	SintacticoCOMA       = 11
	SintacticoMAS        = 12
	SintacticoMENOS      = 13
	SintacticoPOR        = 14
	SintacticoDIVIDIDO   = 15
	SintacticoMODULO     = 16
	SintacticoENTERO     = 17
	SintacticoFLOAT      = 18
	SintacticoCADENA     = 19
	SintacticoR_TRUE     = 20
	SintacticoR_FALSE    = 21
	SintacticoID         = 22
	SintacticoWHITESPACE = 23
)

// Sintactico rules.
const (
	SintacticoRULE_start         = 0
	SintacticoRULE_instrucciones = 1
	SintacticoRULE_instruccion   = 2
	SintacticoRULE_expr          = 3
	SintacticoRULE_expr_op       = 4
	SintacticoRULE_primitivo     = 5
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
		p.SetState(12)

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
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SintacticoPRINTLN {
		{
			p.SetState(15)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(20)
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
		p.SetState(23)
		p.Match(SintacticoPRINTLN)
	}
	{
		p.SetState(24)
		p.Match(SintacticoPARENA)
	}
	{
		p.SetState(25)

		var _x = p.Expr()

		localctx.(*InstruccionContext)._expr = _x
	}
	{
		p.SetState(26)
		p.Match(SintacticoPARENC)
	}
	{
		p.SetState(27)
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

	// Get_expr_op returns the _expr_op rule contexts.
	Get_expr_op() IExpr_opContext

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
	parser   antlr.Parser
	p        abstract.Expresion
	_expr_op IExpr_opContext
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

func (s *ExprContext) Get_expr_op() IExpr_opContext { return s._expr_op }

func (s *ExprContext) Set_expr_op(v IExpr_opContext) { s._expr_op = v }

func (s *ExprContext) GetP() abstract.Expresion { return s.p }

func (s *ExprContext) SetP(v abstract.Expresion) { s.p = v }

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)

		var _x = p.expr_op(0)

		localctx.(*ExprContext)._expr_op = _x
	}
	localctx.(*ExprContext).p = localctx.(*ExprContext).Get_expr_op().GetP()

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

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// GetHDer returns the hDer rule contexts.
	GetHDer() IExpr_opContext

	// SetHIzq sets the hIzq rule contexts.
	SetHIzq(IExpr_opContext)

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

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
	_primitivo IPrimitivoContext
	_expr      IExprContext
	op         antlr.Token
	hDer       IExpr_opContext
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

func (s *Expr_opContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expr_opContext) Get_expr() IExprContext { return s._expr }

func (s *Expr_opContext) GetHDer() IExpr_opContext { return s.hDer }

func (s *Expr_opContext) SetHIzq(v IExpr_opContext) { s.hIzq = v }

func (s *Expr_opContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expr_opContext) Set_expr(v IExprContext) { s._expr = v }

func (s *Expr_opContext) SetHDer(v IExpr_opContext) { s.hDer = v }

func (s *Expr_opContext) GetP() abstract.Expresion { return s.p }

func (s *Expr_opContext) SetP(v abstract.Expresion) { s.p = v }

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

func (s *Expr_opContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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

func (s *Expr_opContext) COMA() antlr.TerminalNode {
	return s.GetToken(SintacticoCOMA, 0)
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

func (s *Expr_opContext) MENOS() antlr.TerminalNode {
	return s.GetToken(SintacticoMENOS, 0)
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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoENTERO, SintacticoCADENA:
		{
			p.SetState(34)

			var _x = p.Primitivo()

			localctx.(*Expr_opContext)._primitivo = _x
		}
		localctx.(*Expr_opContext).p = localctx.(*Expr_opContext).Get_primitivo().GetP()

	case SintacticoPARENA:
		{
			p.SetState(37)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(38)

			var _x = p.Expr()

			localctx.(*Expr_opContext)._expr = _x
		}
		{
			p.SetState(39)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = localctx.(*Expr_opContext).Get_expr().GetP()

	case SintacticoR_INT:
		{
			p.SetState(42)
			p.Match(SintacticoR_INT)
		}
		{
			p.SetState(43)
			p.Match(SintacticoCUATROPT)
		}
		{
			p.SetState(44)
			p.Match(SintacticoPOW)
		}
		{
			p.SetState(45)
			p.Match(SintacticoPARENA)
		}
		{
			p.SetState(46)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hIzq = _x
		}
		{
			p.SetState(47)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_opContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SintacticoCOMA) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_opContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(48)

			var _x = p.expr_op(0)

			localctx.(*Expr_opContext).hDer = _x
		}
		{
			p.SetState(49)
			p.Match(SintacticoPARENC)
		}
		localctx.(*Expr_opContext).p = expresion.NewOperacion(localctx.(*Expr_opContext).GetHIzq().GetP(), "**", localctx.(*Expr_opContext).GetHDer().GetP(), false)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_opContext(p, _parentctx, _parentState)
				localctx.(*Expr_opContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_op)
				p.SetState(54)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(55)

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
					p.SetState(56)

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

			case 2:
				localctx = NewExpr_opContext(p, _parentctx, _parentState)
				localctx.(*Expr_opContext).hIzq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SintacticoRULE_expr_op)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(60)

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
					p.SetState(61)

					var _x = p.expr_op(5)

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
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

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

func (s *PrimitivoContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *PrimitivoContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrimitivoContext) GetP() abstract.Expresion { return s.p }

func (s *PrimitivoContext) SetP(v abstract.Expresion) { s.p = v }

func (s *PrimitivoContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(SintacticoENTERO, 0)
}

func (s *PrimitivoContext) CADENA() antlr.TerminalNode {
	return s.GetToken(SintacticoCADENA, 0)
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
	p.EnterRule(localctx, 10, SintacticoRULE_primitivo)

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

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SintacticoENTERO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)

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
			p.SetState(71)

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

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Sintactico) Expr_op_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
