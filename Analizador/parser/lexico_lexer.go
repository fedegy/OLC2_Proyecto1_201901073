// Code generated from Lexico.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 156,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 18, 6, 18, 105, 10, 18, 13, 18, 14, 18, 106, 3, 19, 6, 19, 110,
	10, 19, 13, 19, 14, 19, 111, 3, 19, 3, 19, 6, 19, 116, 10, 19, 13, 19,
	14, 19, 117, 3, 20, 3, 20, 7, 20, 122, 10, 20, 12, 20, 14, 20, 125, 11,
	20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 7, 23, 142, 10, 23, 12, 23, 14, 23,
	145, 11, 23, 3, 24, 6, 24, 148, 10, 24, 13, 24, 14, 24, 149, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 25, 2, 2, 26, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 2,
	3, 2, 9, 3, 2, 50, 59, 3, 2, 48, 48, 3, 2, 36, 36, 5, 2, 67, 92, 97, 97,
	99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34,
	34, 9, 2, 34, 35, 37, 37, 45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 160,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3,
	51, 3, 2, 2, 2, 5, 53, 3, 2, 2, 2, 7, 55, 3, 2, 2, 2, 9, 64, 3, 2, 2, 2,
	11, 68, 3, 2, 2, 2, 13, 72, 3, 2, 2, 2, 15, 77, 3, 2, 2, 2, 17, 82, 3,
	2, 2, 2, 19, 86, 3, 2, 2, 2, 21, 89, 3, 2, 2, 2, 23, 91, 3, 2, 2, 2, 25,
	93, 3, 2, 2, 2, 27, 95, 3, 2, 2, 2, 29, 97, 3, 2, 2, 2, 31, 99, 3, 2, 2,
	2, 33, 101, 3, 2, 2, 2, 35, 104, 3, 2, 2, 2, 37, 109, 3, 2, 2, 2, 39, 119,
	3, 2, 2, 2, 41, 128, 3, 2, 2, 2, 43, 133, 3, 2, 2, 2, 45, 139, 3, 2, 2,
	2, 47, 147, 3, 2, 2, 2, 49, 153, 3, 2, 2, 2, 51, 52, 7, 42, 2, 2, 52, 4,
	3, 2, 2, 2, 53, 54, 7, 43, 2, 2, 54, 6, 3, 2, 2, 2, 55, 56, 7, 114, 2,
	2, 56, 57, 7, 116, 2, 2, 57, 58, 7, 107, 2, 2, 58, 59, 7, 112, 2, 2, 59,
	60, 7, 118, 2, 2, 60, 61, 7, 110, 2, 2, 61, 62, 7, 112, 2, 2, 62, 63, 7,
	35, 2, 2, 63, 8, 3, 2, 2, 2, 64, 65, 7, 107, 2, 2, 65, 66, 7, 56, 2, 2,
	66, 67, 7, 54, 2, 2, 67, 10, 3, 2, 2, 2, 68, 69, 7, 104, 2, 2, 69, 70,
	7, 56, 2, 2, 70, 71, 7, 54, 2, 2, 71, 12, 3, 2, 2, 2, 72, 73, 7, 40, 2,
	2, 73, 74, 7, 117, 2, 2, 74, 75, 7, 118, 2, 2, 75, 76, 7, 116, 2, 2, 76,
	14, 3, 2, 2, 2, 77, 78, 7, 100, 2, 2, 78, 79, 7, 113, 2, 2, 79, 80, 7,
	113, 2, 2, 80, 81, 7, 110, 2, 2, 81, 16, 3, 2, 2, 2, 82, 83, 7, 114, 2,
	2, 83, 84, 7, 113, 2, 2, 84, 85, 7, 121, 2, 2, 85, 18, 3, 2, 2, 2, 86,
	87, 7, 60, 2, 2, 87, 88, 7, 60, 2, 2, 88, 20, 3, 2, 2, 2, 89, 90, 7, 61,
	2, 2, 90, 22, 3, 2, 2, 2, 91, 92, 7, 46, 2, 2, 92, 24, 3, 2, 2, 2, 93,
	94, 7, 45, 2, 2, 94, 26, 3, 2, 2, 2, 95, 96, 7, 47, 2, 2, 96, 28, 3, 2,
	2, 2, 97, 98, 7, 44, 2, 2, 98, 30, 3, 2, 2, 2, 99, 100, 7, 49, 2, 2, 100,
	32, 3, 2, 2, 2, 101, 102, 7, 39, 2, 2, 102, 34, 3, 2, 2, 2, 103, 105, 9,
	2, 2, 2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 104, 3, 2, 2,
	2, 106, 107, 3, 2, 2, 2, 107, 36, 3, 2, 2, 2, 108, 110, 9, 2, 2, 2, 109,
	108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 115, 9, 3, 2, 2, 114, 116, 9, 2,
	2, 2, 115, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2,
	117, 118, 3, 2, 2, 2, 118, 38, 3, 2, 2, 2, 119, 123, 7, 36, 2, 2, 120,
	122, 10, 4, 2, 2, 121, 120, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121,
	3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2, 2, 2, 125, 123, 3, 2,
	2, 2, 126, 127, 7, 36, 2, 2, 127, 40, 3, 2, 2, 2, 128, 129, 7, 118, 2,
	2, 129, 130, 7, 116, 2, 2, 130, 131, 7, 119, 2, 2, 131, 132, 7, 103, 2,
	2, 132, 42, 3, 2, 2, 2, 133, 134, 7, 104, 2, 2, 134, 135, 7, 99, 2, 2,
	135, 136, 7, 110, 2, 2, 136, 137, 7, 117, 2, 2, 137, 138, 7, 103, 2, 2,
	138, 44, 3, 2, 2, 2, 139, 143, 9, 5, 2, 2, 140, 142, 9, 6, 2, 2, 141, 140,
	3, 2, 2, 2, 142, 145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2,
	2, 2, 144, 46, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 148, 9, 7, 2, 2,
	147, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 149,
	150, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 152, 8, 24, 2, 2, 152, 48,
	3, 2, 2, 2, 153, 154, 7, 94, 2, 2, 154, 155, 9, 8, 2, 2, 155, 50, 3, 2,
	2, 2, 9, 2, 106, 111, 117, 123, 143, 149, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", "'pow'",
	"'::'", "';'", "','", "'+'", "'-'", "'*'", "'/'", "'%'", "", "", "", "'true'",
	"'false'",
}

var lexerSymbolicNames = []string{
	"", "PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL",
	"POW", "CUATROPT", "PTCOMA", "COMA", "MAS", "MENOS", "POR", "DIVIDIDO",
	"MODULO", "ENTERO", "FLOAT", "CADENA", "R_TRUE", "R_FALSE", "ID", "WHITESPACE",
}

var lexerRuleNames = []string{
	"PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL",
	"POW", "CUATROPT", "PTCOMA", "COMA", "MAS", "MENOS", "POR", "DIVIDIDO",
	"MODULO", "ENTERO", "FLOAT", "CADENA", "R_TRUE", "R_FALSE", "ID", "WHITESPACE",
	"ESC_SEQ",
}

type Lexico struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewLexico produces a new lexer instance for the optional input antlr.CharStream.
//
// The *Lexico instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewLexico(input antlr.CharStream) *Lexico {
	l := new(Lexico)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Lexico.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Lexico tokens.
const (
	LexicoPARENA     = 1
	LexicoPARENC     = 2
	LexicoPRINTLN    = 3
	LexicoR_INT      = 4
	LexicoR_FLOAT    = 5
	LexicoR_STRING   = 6
	LexicoR_BOOL     = 7
	LexicoPOW        = 8
	LexicoCUATROPT   = 9
	LexicoPTCOMA     = 10
	LexicoCOMA       = 11
	LexicoMAS        = 12
	LexicoMENOS      = 13
	LexicoPOR        = 14
	LexicoDIVIDIDO   = 15
	LexicoMODULO     = 16
	LexicoENTERO     = 17
	LexicoFLOAT      = 18
	LexicoCADENA     = 19
	LexicoR_TRUE     = 20
	LexicoR_FALSE    = 21
	LexicoID         = 22
	LexicoWHITESPACE = 23
)
