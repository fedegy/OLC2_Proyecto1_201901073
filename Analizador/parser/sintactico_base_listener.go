// Code generated from Sintactico.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Sintactico

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSintacticoListener is a complete listener for a parse tree produced by Sintactico.
type BaseSintacticoListener struct{}

var _ SintacticoListener = &BaseSintacticoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSintacticoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSintacticoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSintacticoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSintacticoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSintacticoListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSintacticoListener) ExitStart(ctx *StartContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseSintacticoListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseSintacticoListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseSintacticoListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseSintacticoListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSintacticoListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSintacticoListener) ExitExpr(ctx *ExprContext) {}

// EnterExpr_op is called when production expr_op is entered.
func (s *BaseSintacticoListener) EnterExpr_op(ctx *Expr_opContext) {}

// ExitExpr_op is called when production expr_op is exited.
func (s *BaseSintacticoListener) ExitExpr_op(ctx *Expr_opContext) {}

// EnterExpr_log is called when production expr_log is entered.
func (s *BaseSintacticoListener) EnterExpr_log(ctx *Expr_logContext) {}

// ExitExpr_log is called when production expr_log is exited.
func (s *BaseSintacticoListener) ExitExpr_log(ctx *Expr_logContext) {}

// EnterExpr_rel is called when production expr_rel is entered.
func (s *BaseSintacticoListener) EnterExpr_rel(ctx *Expr_relContext) {}

// ExitExpr_rel is called when production expr_rel is exited.
func (s *BaseSintacticoListener) ExitExpr_rel(ctx *Expr_relContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseSintacticoListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseSintacticoListener) ExitPrimitivo(ctx *PrimitivoContext) {}
