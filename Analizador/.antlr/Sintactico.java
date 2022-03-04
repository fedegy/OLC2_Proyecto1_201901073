// Generated from /home/federico/go/src/OLC2_Proyecto1_201901073/Analizador/Sintactico.g4 by ANTLR 4.8

    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Expresion"
    import arrayL "github.com/colegno/arrayList"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Instrucciones"
    import "OLC2_PROYECTO1_201901073/Analizador/Entorno"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Sintactico extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PRINTLN=1, R_INT=2, R_FLOAT=3, R_STRING=4, R_BOOL=5, POW=6, AS=7, TOOWNED=8, 
		TOSTRING=9, PARENA=10, PARENC=11, AND=12, OR=13, NOT=14, DIFERENTE=15, 
		MAYORIGUAL=16, MENORIGUAL=17, MAYORQUE=18, MENORQUE=19, IGUALIGUAL=20, 
		CUATROPT=21, PTCOMA=22, COMA=23, PUNTO=24, MAS=25, MENOS=26, POR=27, DIVIDIDO=28, 
		MODULO=29, ENTERO=30, FLOAT=31, CADENA=32, CARACTER=33, R_TRUE=34, R_FALSE=35, 
		ID=36, WHITESPACE=37;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_expr = 3, 
		RULE_expr_op = 4, RULE_expr_log = 5, RULE_expr_rel = 6, RULE_expr_val = 7, 
		RULE_primitivo = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "expr", "expr_op", "expr_log", 
			"expr_rel", "expr_val", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", "'pow'", "'as'", 
			"'to_owned()'", "'to_string()'", "'('", "')'", "'&&'", "'||'", "'!'", 
			"'!='", "'>='", "'<='", "'>'", "'<'", "'=='", "'::'", "';'", "','", "'.'", 
			"'+'", "'-'", "'*'", "'/'", "'%'", null, null, null, null, "'true'", 
			"'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL", "POW", "AS", 
			"TOOWNED", "TOSTRING", "PARENA", "PARENC", "AND", "OR", "NOT", "DIFERENTE", 
			"MAYORIGUAL", "MENORIGUAL", "MAYORQUE", "MENORQUE", "IGUALIGUAL", "CUATROPT", 
			"PTCOMA", "COMA", "PUNTO", "MAS", "MENOS", "POR", "DIVIDIDO", "MODULO", 
			"ENTERO", "FLOAT", "CADENA", "CARACTER", "R_TRUE", "R_FALSE", "ID", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Sintactico.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Sintactico(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public *arrayL.List lista;
		public InstruccionesContext instrucciones;
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(18);
			((StartContext)_localctx).instrucciones = instrucciones();
			_localctx.lista = ((StartContext)_localctx).instrucciones.l
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public *arrayL.List l;
		public InstruccionContext instruccion;
		public List<InstruccionContext> e = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);

		    _localctx.l = arrayL.New()

		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(24);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PRINTLN) {
				{
				{
				setState(21);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(26);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			    listInt := localctx.(*InstruccionesContext).GetE()
			    for _,e := range listInt {
			        _localctx.l.Add(e.GetInstr())
			    }
			    fmt.Printf("Tipo %T", localctx.(*InstruccionesContext).GetE())

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public abstract.Instruccion instr;
		public ExprContext expr;
		public TerminalNode PRINTLN() { return getToken(Sintactico.PRINTLN, 0); }
		public TerminalNode PARENA() { return getToken(Sintactico.PARENA, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARENC() { return getToken(Sintactico.PARENC, 0); }
		public TerminalNode PTCOMA() { return getToken(Sintactico.PTCOMA, 0); }
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(29);
			match(PRINTLN);
			setState(30);
			match(PARENA);
			setState(31);
			((InstruccionContext)_localctx).expr = expr();
			setState(32);
			match(PARENC);
			setState(33);
			match(PTCOMA);
			_localctx.instr = instrucciones.NewImprimir(((InstruccionContext)_localctx).expr.p)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExprContext extends ParserRuleContext {
		public abstract.Expresion p;
		public Expr_logContext expr_log;
		public Expr_relContext expr_rel;
		public Expr_opContext expr_op;
		public Expr_logContext expr_log() {
			return getRuleContext(Expr_logContext.class,0);
		}
		public Expr_relContext expr_rel() {
			return getRuleContext(Expr_relContext.class,0);
		}
		public Expr_opContext expr_op() {
			return getRuleContext(Expr_opContext.class,0);
		}
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		ExprContext _localctx = new ExprContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_expr);
		try {
			setState(45);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				((ExprContext)_localctx).expr_log = expr_log(0);
				_localctx.p = ((ExprContext)_localctx).expr_log.p
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(39);
				((ExprContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((ExprContext)_localctx).expr_rel.p
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(42);
				((ExprContext)_localctx).expr_op = expr_op(0);
				_localctx.p = ((ExprContext)_localctx).expr_op.p
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expr_opContext extends ParserRuleContext {
		public abstract.Expresion p;
		public Expr_opContext hIzq;
		public ExprContext opU;
		public ExprContext expr;
		public Expr_opContext hDer;
		public Expr_valContext expr_val;
		public Token op;
		public TerminalNode MENOS() { return getToken(Sintactico.MENOS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARENA() { return getToken(Sintactico.PARENA, 0); }
		public TerminalNode PARENC() { return getToken(Sintactico.PARENC, 0); }
		public TerminalNode R_INT() { return getToken(Sintactico.R_INT, 0); }
		public TerminalNode CUATROPT() { return getToken(Sintactico.CUATROPT, 0); }
		public TerminalNode POW() { return getToken(Sintactico.POW, 0); }
		public TerminalNode COMA() { return getToken(Sintactico.COMA, 0); }
		public List<Expr_opContext> expr_op() {
			return getRuleContexts(Expr_opContext.class);
		}
		public Expr_opContext expr_op(int i) {
			return getRuleContext(Expr_opContext.class,i);
		}
		public TerminalNode R_FLOAT() { return getToken(Sintactico.R_FLOAT, 0); }
		public Expr_valContext expr_val() {
			return getRuleContext(Expr_valContext.class,0);
		}
		public TerminalNode POR() { return getToken(Sintactico.POR, 0); }
		public TerminalNode DIVIDIDO() { return getToken(Sintactico.DIVIDIDO, 0); }
		public TerminalNode MAS() { return getToken(Sintactico.MAS, 0); }
		public TerminalNode MODULO() { return getToken(Sintactico.MODULO, 0); }
		public Expr_opContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_op; }
	}

	public final Expr_opContext expr_op() throws RecognitionException {
		return expr_op(0);
	}

	private Expr_opContext expr_op(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_opContext _localctx = new Expr_opContext(_ctx, _parentState);
		Expr_opContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_expr_op, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				setState(48);
				match(MENOS);
				setState(49);
				((Expr_opContext)_localctx).opU = ((Expr_opContext)_localctx).expr = expr();
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).opU.p, "-", nil, true)
				}
				break;
			case PARENA:
				{
				setState(52);
				match(PARENA);
				setState(53);
				((Expr_opContext)_localctx).expr = expr();
				setState(54);
				match(PARENC);
				_localctx.p = ((Expr_opContext)_localctx).expr.p
				}
				break;
			case R_INT:
				{
				setState(57);
				match(R_INT);
				setState(58);
				match(CUATROPT);
				setState(59);
				match(POW);
				setState(60);
				match(PARENA);
				setState(61);
				((Expr_opContext)_localctx).hIzq = expr_op(0);
				setState(62);
				match(COMA);
				setState(63);
				((Expr_opContext)_localctx).hDer = expr_op(0);
				setState(64);
				match(PARENC);
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, "**", ((Expr_opContext)_localctx).hDer.p, false)
				}
				break;
			case R_FLOAT:
				{
				setState(67);
				match(R_FLOAT);
				setState(68);
				match(CUATROPT);
				setState(69);
				match(POW);
				setState(70);
				match(PARENA);
				setState(71);
				((Expr_opContext)_localctx).hIzq = expr_op(0);
				setState(72);
				match(COMA);
				setState(73);
				((Expr_opContext)_localctx).hDer = expr_op(0);
				setState(74);
				match(PARENC);
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, "f64**", ((Expr_opContext)_localctx).hDer.p, false)
				}
				break;
			case ENTERO:
			case FLOAT:
			case CADENA:
			case CARACTER:
			case R_TRUE:
			case R_FALSE:
				{
				setState(77);
				((Expr_opContext)_localctx).expr_val = expr_val();
				_localctx.p = ((Expr_opContext)_localctx).expr_val.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(99);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(97);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_opContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_op);
						setState(82);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(83);
						((Expr_opContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==POR || _la==DIVIDIDO) ) {
							((Expr_opContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(84);
						((Expr_opContext)_localctx).hDer = expr_op(8);
						_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, (((Expr_opContext)_localctx).op!=null?((Expr_opContext)_localctx).op.getText():null), ((Expr_opContext)_localctx).hDer.p, false)
						}
						break;
					case 2:
						{
						_localctx = new Expr_opContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_op);
						setState(87);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(88);
						((Expr_opContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MAS || _la==MENOS) ) {
							((Expr_opContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(89);
						((Expr_opContext)_localctx).hDer = expr_op(7);
						_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, (((Expr_opContext)_localctx).op!=null?((Expr_opContext)_localctx).op.getText():null), ((Expr_opContext)_localctx).hDer.p, false)
						}
						break;
					case 3:
						{
						_localctx = new Expr_opContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_op);
						setState(92);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(93);
						match(MODULO);
						setState(94);
						((Expr_opContext)_localctx).hDer = expr_op(6);
						_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, (((Expr_opContext)_localctx).op!=null?((Expr_opContext)_localctx).op.getText():null), ((Expr_opContext)_localctx).hDer.p, false)
						}
						break;
					}
					} 
				}
				setState(101);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Expr_logContext extends ParserRuleContext {
		public abstract.Expresion p;
		public Expr_logContext hIzq;
		public Expr_relContext expr_rel;
		public Expr_logContext hDer;
		public TerminalNode NOT() { return getToken(Sintactico.NOT, 0); }
		public List<Expr_logContext> expr_log() {
			return getRuleContexts(Expr_logContext.class);
		}
		public Expr_logContext expr_log(int i) {
			return getRuleContext(Expr_logContext.class,i);
		}
		public Expr_relContext expr_rel() {
			return getRuleContext(Expr_relContext.class,0);
		}
		public TerminalNode OR() { return getToken(Sintactico.OR, 0); }
		public TerminalNode AND() { return getToken(Sintactico.AND, 0); }
		public Expr_logContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_log; }
	}

	public final Expr_logContext expr_log() throws RecognitionException {
		return expr_log(0);
	}

	private Expr_logContext expr_log(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_logContext _localctx = new Expr_logContext(_ctx, _parentState);
		Expr_logContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_expr_log, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(103);
				match(NOT);
				setState(104);
				((Expr_logContext)_localctx).hIzq = expr_log(2);
				_localctx.p = expresion.NewLogicas(((Expr_logContext)_localctx).hIzq.p, "!", nil, true)
				}
				break;
			case R_INT:
			case R_FLOAT:
			case PARENA:
			case MENOS:
			case ENTERO:
			case FLOAT:
			case CADENA:
			case CARACTER:
			case R_TRUE:
			case R_FALSE:
				{
				setState(107);
				((Expr_logContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((Expr_logContext)_localctx).expr_rel.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(124);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(122);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_logContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_log);
						setState(112);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(113);
						match(OR);
						setState(114);
						((Expr_logContext)_localctx).hDer = expr_log(5);
						_localctx.p = expresion.NewLogicas(((Expr_logContext)_localctx).hIzq.p, "||", ((Expr_logContext)_localctx).hDer.p, false)
						}
						break;
					case 2:
						{
						_localctx = new Expr_logContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_log);
						setState(117);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(118);
						match(AND);
						setState(119);
						((Expr_logContext)_localctx).hDer = expr_log(4);
						_localctx.p = expresion.NewLogicas(((Expr_logContext)_localctx).hIzq.p, "&&", ((Expr_logContext)_localctx).hDer.p, false)
						}
						break;
					}
					} 
				}
				setState(126);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Expr_relContext extends ParserRuleContext {
		public abstract.Expresion p;
		public Expr_relContext hIzq;
		public Expr_opContext expr_op;
		public Token op;
		public Expr_relContext hDer;
		public Expr_opContext expr_op() {
			return getRuleContext(Expr_opContext.class,0);
		}
		public List<Expr_relContext> expr_rel() {
			return getRuleContexts(Expr_relContext.class);
		}
		public Expr_relContext expr_rel(int i) {
			return getRuleContext(Expr_relContext.class,i);
		}
		public TerminalNode IGUALIGUAL() { return getToken(Sintactico.IGUALIGUAL, 0); }
		public TerminalNode DIFERENTE() { return getToken(Sintactico.DIFERENTE, 0); }
		public TerminalNode MENORQUE() { return getToken(Sintactico.MENORQUE, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Sintactico.MENORIGUAL, 0); }
		public TerminalNode MAYORQUE() { return getToken(Sintactico.MAYORQUE, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(Sintactico.MAYORIGUAL, 0); }
		public Expr_relContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_rel; }
	}

	public final Expr_relContext expr_rel() throws RecognitionException {
		return expr_rel(0);
	}

	private Expr_relContext expr_rel(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_relContext _localctx = new Expr_relContext(_ctx, _parentState);
		Expr_relContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_expr_rel, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(128);
			((Expr_relContext)_localctx).expr_op = expr_op(0);
			_localctx.p = ((Expr_relContext)_localctx).expr_op.p
			}
			_ctx.stop = _input.LT(-1);
			setState(138);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Expr_relContext(_parentctx, _parentState);
					_localctx.hIzq = _prevctx;
					_localctx.hIzq = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expr_rel);
					setState(131);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(132);
					((Expr_relContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DIFERENTE) | (1L << MAYORIGUAL) | (1L << MENORIGUAL) | (1L << MAYORQUE) | (1L << MENORQUE) | (1L << IGUALIGUAL))) != 0)) ) {
						((Expr_relContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(133);
					((Expr_relContext)_localctx).hDer = expr_rel(3);
					_localctx.p = expresion.NewRelacional(((Expr_relContext)_localctx).hIzq.p, (((Expr_relContext)_localctx).op!=null?((Expr_relContext)_localctx).op.getText():null), ((Expr_relContext)_localctx).hDer.p)
					}
					} 
				}
				setState(140);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class Expr_valContext extends ParserRuleContext {
		public abstract.Expresion p;
		public PrimitivoContext primitivo;
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public Expr_valContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_val; }
	}

	public final Expr_valContext expr_val() throws RecognitionException {
		Expr_valContext _localctx = new Expr_valContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_expr_val);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			((Expr_valContext)_localctx).primitivo = primitivo();
			_localctx.p = ((Expr_valContext)_localctx).primitivo.p
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrimitivoContext extends ParserRuleContext {
		public abstract.Expresion p;
		public Token ENTERO;
		public Token CADENA;
		public Token FLOAT;
		public Token R_TRUE;
		public Token R_FALSE;
		public Token CARACTER;
		public TerminalNode ENTERO() { return getToken(Sintactico.ENTERO, 0); }
		public TerminalNode CADENA() { return getToken(Sintactico.CADENA, 0); }
		public TerminalNode FLOAT() { return getToken(Sintactico.FLOAT, 0); }
		public TerminalNode R_TRUE() { return getToken(Sintactico.R_TRUE, 0); }
		public TerminalNode R_FALSE() { return getToken(Sintactico.R_FALSE, 0); }
		public TerminalNode CARACTER() { return getToken(Sintactico.CARACTER, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_primitivo);
		try {
			setState(156);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(144);
				((PrimitivoContext)_localctx).ENTERO = match(ENTERO);

				    num, err := strconv.Atoi((((PrimitivoContext)_localctx).ENTERO!=null?((PrimitivoContext)_localctx).ENTERO.getText():null))
				    if err != nil {
				        fmt.Println(err)
				    }
				    _localctx.p = expresion.NewPrimitivo(num, entorno.INT)

				}
				break;
			case CADENA:
				enterOuterAlt(_localctx, 2);
				{
				setState(146);
				((PrimitivoContext)_localctx).CADENA = match(CADENA);

				    str := (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)[1:len((((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null))-1]
				    _localctx.p = expresion.NewPrimitivo(str, entorno.STRING)

				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 3);
				{
				setState(148);
				((PrimitivoContext)_localctx).FLOAT = match(FLOAT);

				    num, err := strconv.ParseFloat((((PrimitivoContext)_localctx).FLOAT!=null?((PrimitivoContext)_localctx).FLOAT.getText():null), 64)
				    if err != nil {
				        fmt.Println(err)
				    }
				    _localctx.p = expresion.NewPrimitivo(num, entorno.FLOAT)

				}
				break;
			case R_TRUE:
				enterOuterAlt(_localctx, 4);
				{
				setState(150);
				((PrimitivoContext)_localctx).R_TRUE = match(R_TRUE);

				    bool_true, err := strconv.ParseBool((((PrimitivoContext)_localctx).R_TRUE!=null?((PrimitivoContext)_localctx).R_TRUE.getText():null))
				    if err != nil {
				        fmt.Println(err)
				    }
				    _localctx.p = expresion.NewPrimitivo(bool_true, entorno.BOOL)

				}
				break;
			case R_FALSE:
				enterOuterAlt(_localctx, 5);
				{
				setState(152);
				((PrimitivoContext)_localctx).R_FALSE = match(R_FALSE);

				    bool_false, err := strconv.ParseBool((((PrimitivoContext)_localctx).R_FALSE!=null?((PrimitivoContext)_localctx).R_FALSE.getText():null))
				    if err != nil {
				        fmt.Println(err)
				    }
				    _localctx.p = expresion.NewPrimitivo(bool_false, entorno.BOOL)

				}
				break;
			case CARACTER:
				enterOuterAlt(_localctx, 6);
				{
				setState(154);
				((PrimitivoContext)_localctx).CARACTER = match(CARACTER);

				    str := (((PrimitivoContext)_localctx).CARACTER!=null?((PrimitivoContext)_localctx).CARACTER.getText():null)[1:len((((PrimitivoContext)_localctx).CARACTER!=null?((PrimitivoContext)_localctx).CARACTER.getText():null))-1]
				    _localctx.p = expresion.NewPrimitivo(str, entorno.STRING)

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return expr_op_sempred((Expr_opContext)_localctx, predIndex);
		case 5:
			return expr_log_sempred((Expr_logContext)_localctx, predIndex);
		case 6:
			return expr_rel_sempred((Expr_relContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_op_sempred(Expr_opContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 7);
		case 1:
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean expr_log_sempred(Expr_logContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 4);
		case 4:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean expr_rel_sempred(Expr_relContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\'\u00a1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\3\2\3\2"+
		"\3\2\3\3\7\3\31\n\3\f\3\16\3\34\13\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5\60\n\5\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6S\n\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6d\n\6\f\6\16\6g\13\6\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7q\n\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\7\7}\n\7\f\7\16\7\u0080\13\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\7\b\u008b\n\b\f\b\16\b\u008e\13\b\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n"+
		"\3\n\3\n\3\n\3\n\3\n\3\n\3\n\5\n\u009f\n\n\3\n\2\5\n\f\16\13\2\4\6\b\n"+
		"\f\16\20\22\2\5\3\2\35\36\3\2\33\34\3\2\21\26\2\u00aa\2\24\3\2\2\2\4\32"+
		"\3\2\2\2\6\37\3\2\2\2\b/\3\2\2\2\nR\3\2\2\2\fp\3\2\2\2\16\u0081\3\2\2"+
		"\2\20\u008f\3\2\2\2\22\u009e\3\2\2\2\24\25\5\4\3\2\25\26\b\2\1\2\26\3"+
		"\3\2\2\2\27\31\5\6\4\2\30\27\3\2\2\2\31\34\3\2\2\2\32\30\3\2\2\2\32\33"+
		"\3\2\2\2\33\35\3\2\2\2\34\32\3\2\2\2\35\36\b\3\1\2\36\5\3\2\2\2\37 \7"+
		"\3\2\2 !\7\f\2\2!\"\5\b\5\2\"#\7\r\2\2#$\7\30\2\2$%\b\4\1\2%\7\3\2\2\2"+
		"&\'\5\f\7\2\'(\b\5\1\2(\60\3\2\2\2)*\5\16\b\2*+\b\5\1\2+\60\3\2\2\2,-"+
		"\5\n\6\2-.\b\5\1\2.\60\3\2\2\2/&\3\2\2\2/)\3\2\2\2/,\3\2\2\2\60\t\3\2"+
		"\2\2\61\62\b\6\1\2\62\63\7\34\2\2\63\64\5\b\5\2\64\65\b\6\1\2\65S\3\2"+
		"\2\2\66\67\7\f\2\2\678\5\b\5\289\7\r\2\29:\b\6\1\2:S\3\2\2\2;<\7\4\2\2"+
		"<=\7\27\2\2=>\7\b\2\2>?\7\f\2\2?@\5\n\6\2@A\7\31\2\2AB\5\n\6\2BC\7\r\2"+
		"\2CD\b\6\1\2DS\3\2\2\2EF\7\5\2\2FG\7\27\2\2GH\7\b\2\2HI\7\f\2\2IJ\5\n"+
		"\6\2JK\7\31\2\2KL\5\n\6\2LM\7\r\2\2MN\b\6\1\2NS\3\2\2\2OP\5\20\t\2PQ\b"+
		"\6\1\2QS\3\2\2\2R\61\3\2\2\2R\66\3\2\2\2R;\3\2\2\2RE\3\2\2\2RO\3\2\2\2"+
		"Se\3\2\2\2TU\f\t\2\2UV\t\2\2\2VW\5\n\6\nWX\b\6\1\2Xd\3\2\2\2YZ\f\b\2\2"+
		"Z[\t\3\2\2[\\\5\n\6\t\\]\b\6\1\2]d\3\2\2\2^_\f\7\2\2_`\7\37\2\2`a\5\n"+
		"\6\bab\b\6\1\2bd\3\2\2\2cT\3\2\2\2cY\3\2\2\2c^\3\2\2\2dg\3\2\2\2ec\3\2"+
		"\2\2ef\3\2\2\2f\13\3\2\2\2ge\3\2\2\2hi\b\7\1\2ij\7\20\2\2jk\5\f\7\4kl"+
		"\b\7\1\2lq\3\2\2\2mn\5\16\b\2no\b\7\1\2oq\3\2\2\2ph\3\2\2\2pm\3\2\2\2"+
		"q~\3\2\2\2rs\f\6\2\2st\7\17\2\2tu\5\f\7\7uv\b\7\1\2v}\3\2\2\2wx\f\5\2"+
		"\2xy\7\16\2\2yz\5\f\7\6z{\b\7\1\2{}\3\2\2\2|r\3\2\2\2|w\3\2\2\2}\u0080"+
		"\3\2\2\2~|\3\2\2\2~\177\3\2\2\2\177\r\3\2\2\2\u0080~\3\2\2\2\u0081\u0082"+
		"\b\b\1\2\u0082\u0083\5\n\6\2\u0083\u0084\b\b\1\2\u0084\u008c\3\2\2\2\u0085"+
		"\u0086\f\4\2\2\u0086\u0087\t\4\2\2\u0087\u0088\5\16\b\5\u0088\u0089\b"+
		"\b\1\2\u0089\u008b\3\2\2\2\u008a\u0085\3\2\2\2\u008b\u008e\3\2\2\2\u008c"+
		"\u008a\3\2\2\2\u008c\u008d\3\2\2\2\u008d\17\3\2\2\2\u008e\u008c\3\2\2"+
		"\2\u008f\u0090\5\22\n\2\u0090\u0091\b\t\1\2\u0091\21\3\2\2\2\u0092\u0093"+
		"\7 \2\2\u0093\u009f\b\n\1\2\u0094\u0095\7\"\2\2\u0095\u009f\b\n\1\2\u0096"+
		"\u0097\7!\2\2\u0097\u009f\b\n\1\2\u0098\u0099\7$\2\2\u0099\u009f\b\n\1"+
		"\2\u009a\u009b\7%\2\2\u009b\u009f\b\n\1\2\u009c\u009d\7#\2\2\u009d\u009f"+
		"\b\n\1\2\u009e\u0092\3\2\2\2\u009e\u0094\3\2\2\2\u009e\u0096\3\2\2\2\u009e"+
		"\u0098\3\2\2\2\u009e\u009a\3\2\2\2\u009e\u009c\3\2\2\2\u009f\23\3\2\2"+
		"\2\f\32/Rcep|~\u008c\u009e";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}