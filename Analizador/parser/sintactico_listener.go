// Code generated from Sintactico.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Sintactico

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SintacticoListener is a complete listener for a parse tree produced by Sintactico.
type SintacticoListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterListides is called when entering the listides production.
	EnterListides(c *ListidesContext)

	// EnterTipos is called when entering the tipos production.
	EnterTipos(c *TiposContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExpr_op is called when entering the expr_op production.
	EnterExpr_op(c *Expr_opContext)

	// EnterExpr_log is called when entering the expr_log production.
	EnterExpr_log(c *Expr_logContext)

	// EnterExpr_rel is called when entering the expr_rel production.
	EnterExpr_rel(c *Expr_relContext)

	// EnterExpr_val is called when entering the expr_val production.
	EnterExpr_val(c *Expr_valContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitListides is called when exiting the listides production.
	ExitListides(c *ListidesContext)

	// ExitTipos is called when exiting the tipos production.
	ExitTipos(c *TiposContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExpr_op is called when exiting the expr_op production.
	ExitExpr_op(c *Expr_opContext)

	// ExitExpr_log is called when exiting the expr_log production.
	ExitExpr_log(c *Expr_logContext)

	// ExitExpr_rel is called when exiting the expr_rel production.
	ExitExpr_rel(c *Expr_relContext)

	// ExitExpr_val is called when exiting the expr_val production.
	ExitExpr_val(c *Expr_valContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}
