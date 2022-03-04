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
		PARENA=1, PARENC=2, PRINTLN=3, R_INT=4, R_FLOAT=5, R_STRING=6, R_BOOL=7, 
		POW=8, AND=9, OR=10, NOT=11, DIFERENTE=12, MAYORIGUAL=13, MENORIGUAL=14, 
		MAYORQUE=15, MENORQUE=16, IGUALIGUAL=17, CUATROPT=18, PTCOMA=19, COMA=20, 
		MAS=21, MENOS=22, POR=23, DIVIDIDO=24, MODULO=25, ENTERO=26, FLOAT=27, 
		CADENA=28, R_TRUE=29, R_FALSE=30, ID=31, WHITESPACE=32;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_expr = 3, 
		RULE_expr_op = 4, RULE_expr_log = 5, RULE_expr_rel = 6, RULE_primitivo = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "expr", "expr_op", "expr_log", 
			"expr_rel", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", 
			"'pow'", "'&&'", "'||'", "'!'", "'!='", "'>='", "'<='", "'>'", "'<'", 
			"'=='", "'::'", "';'", "','", "'+'", "'-'", "'*'", "'/'", "'%'", null, 
			null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", 
			"R_BOOL", "POW", "AND", "OR", "NOT", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", 
			"MAYORQUE", "MENORQUE", "IGUALIGUAL", "CUATROPT", "PTCOMA", "COMA", "MAS", 
			"MENOS", "POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT", "CADENA", "R_TRUE", 
			"R_FALSE", "ID", "WHITESPACE"
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
			setState(16);
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
			setState(22);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==PRINTLN) {
				{
				{
				setState(19);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(24);
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
			setState(27);
			match(PRINTLN);
			setState(28);
			match(PARENA);
			setState(29);
			((InstruccionContext)_localctx).expr = expr();
			setState(30);
			match(PARENC);
			setState(31);
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
			setState(43);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(34);
				((ExprContext)_localctx).expr_log = expr_log();
				_localctx.p = ((ExprContext)_localctx).expr_log.p
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(37);
				((ExprContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((ExprContext)_localctx).expr_rel.p
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(40);
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
		public PrimitivoContext primitivo;
		public Expr_opContext hDer;
		public Token op;
		public TerminalNode MENOS() { return getToken(Sintactico.MENOS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
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
			setState(78);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				setState(46);
				match(MENOS);
				setState(47);
				((Expr_opContext)_localctx).opU = ((Expr_opContext)_localctx).expr = expr();
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).opU.p, "-", nil, true)
				}
				break;
			case ENTERO:
			case FLOAT:
			case CADENA:
			case R_TRUE:
				{
				setState(50);
				((Expr_opContext)_localctx).primitivo = primitivo();
				_localctx.p = ((Expr_opContext)_localctx).primitivo.p
				}
				break;
			case PARENA:
				{
				setState(53);
				match(PARENA);
				setState(54);
				((Expr_opContext)_localctx).expr = expr();
				setState(55);
				match(PARENC);
				_localctx.p = ((Expr_opContext)_localctx).expr.p
				}
				break;
			case R_INT:
				{
				setState(58);
				match(R_INT);
				setState(59);
				match(CUATROPT);
				setState(60);
				match(POW);
				setState(61);
				match(PARENA);
				setState(62);
				((Expr_opContext)_localctx).hIzq = expr_op(0);
				setState(63);
				match(COMA);
				setState(64);
				((Expr_opContext)_localctx).hDer = expr_op(0);
				setState(65);
				match(PARENC);
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, "**", ((Expr_opContext)_localctx).hDer.p, false)
				}
				break;
			case R_FLOAT:
				{
				setState(68);
				match(R_FLOAT);
				setState(69);
				match(CUATROPT);
				setState(70);
				match(POW);
				setState(71);
				match(PARENA);
				setState(72);
				((Expr_opContext)_localctx).hIzq = expr_op(0);
				setState(73);
				match(COMA);
				setState(74);
				((Expr_opContext)_localctx).hDer = expr_op(0);
				setState(75);
				match(PARENC);
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, "f64**", ((Expr_opContext)_localctx).hDer.p, false)
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(97);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(95);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_opContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_op);
						setState(80);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(81);
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
						setState(82);
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
						setState(85);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(86);
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
						setState(87);
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
						setState(90);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(91);
						match(MODULO);
						setState(92);
						((Expr_opContext)_localctx).hDer = expr_op(6);
						_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, (((Expr_opContext)_localctx).op!=null?((Expr_opContext)_localctx).op.getText():null), ((Expr_opContext)_localctx).hDer.p, false)
						}
						break;
					}
					} 
				}
				setState(99);
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
		public Expr_relContext hIzq;
		public Expr_relContext hDer;
		public Expr_relContext expr_rel;
		public TerminalNode OR() { return getToken(Sintactico.OR, 0); }
		public List<Expr_relContext> expr_rel() {
			return getRuleContexts(Expr_relContext.class);
		}
		public Expr_relContext expr_rel(int i) {
			return getRuleContext(Expr_relContext.class,i);
		}
		public TerminalNode AND() { return getToken(Sintactico.AND, 0); }
		public TerminalNode NOT() { return getToken(Sintactico.NOT, 0); }
		public Expr_logContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_log; }
	}

	public final Expr_logContext expr_log() throws RecognitionException {
		Expr_logContext _localctx = new Expr_logContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_expr_log);
		try {
			setState(117);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(100);
				((Expr_logContext)_localctx).hIzq = expr_rel(0);
				setState(101);
				match(OR);
				setState(102);
				((Expr_logContext)_localctx).hDer = expr_rel(0);
				_localctx.p = expresion.NewLogicas(((Expr_logContext)_localctx).hIzq.p, "||", ((Expr_logContext)_localctx).hDer.p, false)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(105);
				((Expr_logContext)_localctx).hIzq = expr_rel(0);
				setState(106);
				match(AND);
				setState(107);
				((Expr_logContext)_localctx).hDer = expr_rel(0);
				_localctx.p = expresion.NewLogicas(((Expr_logContext)_localctx).hIzq.p, "&&", ((Expr_logContext)_localctx).hDer.p, false)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(110);
				match(NOT);
				setState(111);
				((Expr_logContext)_localctx).hIzq = expr_rel(0);
				_localctx.p = expresion.NewLogicas(((Expr_logContext)_localctx).hIzq.p, "!", nil, true)
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(114);
				((Expr_logContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((Expr_logContext)_localctx).expr_rel.p
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
			setState(120);
			((Expr_relContext)_localctx).expr_op = expr_op(0);
			_localctx.p = ((Expr_relContext)_localctx).expr_op.p
			}
			_ctx.stop = _input.LT(-1);
			setState(130);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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
					setState(123);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(124);
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
					setState(125);
					((Expr_relContext)_localctx).hDer = expr_rel(3);
					_localctx.p = expresion.NewRelacional(((Expr_relContext)_localctx).hIzq.p, (((Expr_relContext)_localctx).op!=null?((Expr_relContext)_localctx).op.getText():null), ((Expr_relContext)_localctx).hDer.p)
					}
					} 
				}
				setState(132);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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

	public static class PrimitivoContext extends ParserRuleContext {
		public abstract.Expresion p;
		public Token ENTERO;
		public Token CADENA;
		public Token FLOAT;
		public Token R_TRUE;
		public TerminalNode ENTERO() { return getToken(Sintactico.ENTERO, 0); }
		public TerminalNode CADENA() { return getToken(Sintactico.CADENA, 0); }
		public TerminalNode FLOAT() { return getToken(Sintactico.FLOAT, 0); }
		public TerminalNode R_TRUE() { return getToken(Sintactico.R_TRUE, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_primitivo);
		try {
			setState(141);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(133);
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
				setState(135);
				((PrimitivoContext)_localctx).CADENA = match(CADENA);

				    str := (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)[1:len((((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null))-1]
				    _localctx.p = expresion.NewPrimitivo(str, entorno.STRING)

				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 3);
				{
				setState(137);
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
				setState(139);
				((PrimitivoContext)_localctx).R_TRUE = match(R_TRUE);

				    bool_true, err := strconv.ParseBool((((PrimitivoContext)_localctx).R_TRUE!=null?((PrimitivoContext)_localctx).R_TRUE.getText():null))
				    if err != nil {
				        fmt.Println(err)
				    }
				    _localctx.p = expresion.NewPrimitivo(bool_true, entorno.BOOL)

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
	private boolean expr_rel_sempred(Expr_relContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\"\u0092\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\3\2\3\3"+
		"\7\3\27\n\3\f\3\16\3\32\13\3\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5.\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\5\6Q\n\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\7\6b\n\6\f\6\16\6e\13\6\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\5\7x\n\7\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\b\u0083\n\b\f\b\16\b\u0086\13\b\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\5\t\u0090\n\t\3\t\2\4\n\16\n\2\4\6\b\n\f\16"+
		"\20\2\5\3\2\31\32\3\2\27\30\3\2\16\23\2\u009a\2\22\3\2\2\2\4\30\3\2\2"+
		"\2\6\35\3\2\2\2\b-\3\2\2\2\nP\3\2\2\2\fw\3\2\2\2\16y\3\2\2\2\20\u008f"+
		"\3\2\2\2\22\23\5\4\3\2\23\24\b\2\1\2\24\3\3\2\2\2\25\27\5\6\4\2\26\25"+
		"\3\2\2\2\27\32\3\2\2\2\30\26\3\2\2\2\30\31\3\2\2\2\31\33\3\2\2\2\32\30"+
		"\3\2\2\2\33\34\b\3\1\2\34\5\3\2\2\2\35\36\7\5\2\2\36\37\7\3\2\2\37 \5"+
		"\b\5\2 !\7\4\2\2!\"\7\25\2\2\"#\b\4\1\2#\7\3\2\2\2$%\5\f\7\2%&\b\5\1\2"+
		"&.\3\2\2\2\'(\5\16\b\2()\b\5\1\2).\3\2\2\2*+\5\n\6\2+,\b\5\1\2,.\3\2\2"+
		"\2-$\3\2\2\2-\'\3\2\2\2-*\3\2\2\2.\t\3\2\2\2/\60\b\6\1\2\60\61\7\30\2"+
		"\2\61\62\5\b\5\2\62\63\b\6\1\2\63Q\3\2\2\2\64\65\5\20\t\2\65\66\b\6\1"+
		"\2\66Q\3\2\2\2\678\7\3\2\289\5\b\5\29:\7\4\2\2:;\b\6\1\2;Q\3\2\2\2<=\7"+
		"\6\2\2=>\7\24\2\2>?\7\n\2\2?@\7\3\2\2@A\5\n\6\2AB\7\26\2\2BC\5\n\6\2C"+
		"D\7\4\2\2DE\b\6\1\2EQ\3\2\2\2FG\7\7\2\2GH\7\24\2\2HI\7\n\2\2IJ\7\3\2\2"+
		"JK\5\n\6\2KL\7\26\2\2LM\5\n\6\2MN\7\4\2\2NO\b\6\1\2OQ\3\2\2\2P/\3\2\2"+
		"\2P\64\3\2\2\2P\67\3\2\2\2P<\3\2\2\2PF\3\2\2\2Qc\3\2\2\2RS\f\t\2\2ST\t"+
		"\2\2\2TU\5\n\6\nUV\b\6\1\2Vb\3\2\2\2WX\f\b\2\2XY\t\3\2\2YZ\5\n\6\tZ[\b"+
		"\6\1\2[b\3\2\2\2\\]\f\7\2\2]^\7\33\2\2^_\5\n\6\b_`\b\6\1\2`b\3\2\2\2a"+
		"R\3\2\2\2aW\3\2\2\2a\\\3\2\2\2be\3\2\2\2ca\3\2\2\2cd\3\2\2\2d\13\3\2\2"+
		"\2ec\3\2\2\2fg\5\16\b\2gh\7\f\2\2hi\5\16\b\2ij\b\7\1\2jx\3\2\2\2kl\5\16"+
		"\b\2lm\7\13\2\2mn\5\16\b\2no\b\7\1\2ox\3\2\2\2pq\7\r\2\2qr\5\16\b\2rs"+
		"\b\7\1\2sx\3\2\2\2tu\5\16\b\2uv\b\7\1\2vx\3\2\2\2wf\3\2\2\2wk\3\2\2\2"+
		"wp\3\2\2\2wt\3\2\2\2x\r\3\2\2\2yz\b\b\1\2z{\5\n\6\2{|\b\b\1\2|\u0084\3"+
		"\2\2\2}~\f\4\2\2~\177\t\4\2\2\177\u0080\5\16\b\5\u0080\u0081\b\b\1\2\u0081"+
		"\u0083\3\2\2\2\u0082}\3\2\2\2\u0083\u0086\3\2\2\2\u0084\u0082\3\2\2\2"+
		"\u0084\u0085\3\2\2\2\u0085\17\3\2\2\2\u0086\u0084\3\2\2\2\u0087\u0088"+
		"\7\34\2\2\u0088\u0090\b\t\1\2\u0089\u008a\7\36\2\2\u008a\u0090\b\t\1\2"+
		"\u008b\u008c\7\35\2\2\u008c\u0090\b\t\1\2\u008d\u008e\7\37\2\2\u008e\u0090"+
		"\b\t\1\2\u008f\u0087\3\2\2\2\u008f\u0089\3\2\2\2\u008f\u008b\3\2\2\2\u008f"+
		"\u008d\3\2\2\2\u0090\21\3\2\2\2\n\30-Pacw\u0084\u008f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}