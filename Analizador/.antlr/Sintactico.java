// Generated from /home/federico/go/src/OLC2_Proyecto1_201901073/Analizador/Sintactico.g4 by ANTLR 4.8

    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Expresion"
    import arrayL "github.com/colegno/arrayList"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Instrucciones"
    import "OLC2_PROYECTO1_201901073/Analizador/Entorno"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast"

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
		PRINTLN=1, R_INT=2, R_FLOAT=3, R_STR=4, R_CHAR=5, MUT=6, R_STRING=7, R_BOOL=8, 
		POW=9, AS=10, LET=11, TOOWNED=12, TOSTRING=13, PARENA=14, PARENC=15, AND=16, 
		OR=17, NOT=18, DIFERENTE=19, MAYORIGUAL=20, MENORIGUAL=21, MAYORQUE=22, 
		MENORQUE=23, IGUALIGUAL=24, IGUAL=25, CUATROPT=26, DOSPT=27, PTCOMA=28, 
		COMA=29, PUNTO=30, MAS=31, MENOS=32, POR=33, DIVIDIDO=34, MODULO=35, ENTERO=36, 
		FLOAT=37, CADENA=38, CARACTER=39, R_TRUE=40, R_FALSE=41, ID=42, WHITESPACE=43;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_imprimir = 3, 
		RULE_declaracion = 4, RULE_listides = 5, RULE_tipos = 6, RULE_expr = 7, 
		RULE_expr_op = 8, RULE_expr_log = 9, RULE_expr_rel = 10, RULE_expr_val = 11, 
		RULE_primitivo = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "imprimir", "declaracion", "listides", 
			"tipos", "expr", "expr_op", "expr_log", "expr_rel", "expr_val", "primitivo"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println!'", "'i64'", "'f64'", "'&str'", "'char'", "'mut'", "'String'", 
			"'bool'", "'pow'", "'as'", "'let'", "'to_owned()'", "'to_string()'", 
			"'('", "')'", "'&&'", "'||'", "'!'", "'!='", "'>='", "'<='", "'>'", "'<'", 
			"'=='", "'='", "'::'", "':'", "';'", "','", "'.'", "'+'", "'-'", "'*'", 
			"'/'", "'%'", null, null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PRINTLN", "R_INT", "R_FLOAT", "R_STR", "R_CHAR", "MUT", "R_STRING", 
			"R_BOOL", "POW", "AS", "LET", "TOOWNED", "TOSTRING", "PARENA", "PARENC", 
			"AND", "OR", "NOT", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", 
			"MENORQUE", "IGUALIGUAL", "IGUAL", "CUATROPT", "DOSPT", "PTCOMA", "COMA", 
			"PUNTO", "MAS", "MENOS", "POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT", 
			"CADENA", "CARACTER", "R_TRUE", "R_FALSE", "ID", "WHITESPACE"
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
		public ast.Ast ast;
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
			setState(26);
			((StartContext)_localctx).instrucciones = instrucciones();
			_localctx.ast = ast.NewAst(((StartContext)_localctx).instrucciones.l)
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
			setState(30); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(29);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(32); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==PRINTLN || _la==LET );

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
		public ImprimirContext imprimir;
		public DeclaracionContext declaracion;
		public ImprimirContext imprimir() {
			return getRuleContext(ImprimirContext.class,0);
		}
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(42);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINTLN:
				enterOuterAlt(_localctx, 1);
				{
				setState(36);
				((InstruccionContext)_localctx).imprimir = imprimir();
				_localctx.instr = ((InstruccionContext)_localctx).imprimir.instr
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(39);
				((InstruccionContext)_localctx).declaracion = declaracion();
				_localctx.instr = ((InstruccionContext)_localctx).declaracion.instr
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

	public static class ImprimirContext extends ParserRuleContext {
		public abstract.Instruccion instr;
		public ExprContext expr;
		public TerminalNode PRINTLN() { return getToken(Sintactico.PRINTLN, 0); }
		public TerminalNode PARENA() { return getToken(Sintactico.PARENA, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PARENC() { return getToken(Sintactico.PARENC, 0); }
		public TerminalNode PTCOMA() { return getToken(Sintactico.PTCOMA, 0); }
		public ImprimirContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_imprimir; }
	}

	public final ImprimirContext imprimir() throws RecognitionException {
		ImprimirContext _localctx = new ImprimirContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_imprimir);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(PRINTLN);
			setState(45);
			match(PARENA);
			setState(46);
			((ImprimirContext)_localctx).expr = expr();
			setState(47);
			match(PARENC);
			setState(48);
			match(PTCOMA);
			_localctx.instr = instrucciones.NewImprimir(((ImprimirContext)_localctx).expr.p)
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

	public static class DeclaracionContext extends ParserRuleContext {
		public abstract.Instruccion instr;
		public ListidesContext listides;
		public TiposContext tipos;
		public ExprContext expr;
		public TerminalNode LET() { return getToken(Sintactico.LET, 0); }
		public TerminalNode MUT() { return getToken(Sintactico.MUT, 0); }
		public ListidesContext listides() {
			return getRuleContext(ListidesContext.class,0);
		}
		public TerminalNode DOSPT() { return getToken(Sintactico.DOSPT, 0); }
		public TiposContext tipos() {
			return getRuleContext(TiposContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Sintactico.IGUAL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51);
			match(LET);
			setState(52);
			match(MUT);
			setState(53);
			((DeclaracionContext)_localctx).listides = listides();
			setState(54);
			match(DOSPT);
			setState(55);
			((DeclaracionContext)_localctx).tipos = tipos();
			setState(56);
			match(IGUAL);
			setState(57);
			((DeclaracionContext)_localctx).expr = expr();
			_localctx.instr = instrucciones.NewDeclaracionValor(((DeclaracionContext)_localctx).listides.lista, ((DeclaracionContext)_localctx).tipos.tipo, ((DeclaracionContext)_localctx).expr.p)
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

	public static class ListidesContext extends ParserRuleContext {
		public *arrayL.List lista;
		public Token ID;
		public TerminalNode ID() { return getToken(Sintactico.ID, 0); }
		public ListidesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listides; }
	}

	public final ListidesContext listides() throws RecognitionException {
		ListidesContext _localctx = new ListidesContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_listides);
		 _localctx.lista = arrayL.New()
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			((ListidesContext)_localctx).ID = match(ID);
			_localctx.lista.Add(expresion.NewIdentificador((((ListidesContext)_localctx).ID!=null?((ListidesContext)_localctx).ID.getText():null)))
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

	public static class TiposContext extends ParserRuleContext {
		public entorno.TipoDato tipo;
		public TerminalNode R_INT() { return getToken(Sintactico.R_INT, 0); }
		public TerminalNode R_FLOAT() { return getToken(Sintactico.R_FLOAT, 0); }
		public TerminalNode R_STR() { return getToken(Sintactico.R_STR, 0); }
		public TerminalNode R_CHAR() { return getToken(Sintactico.R_CHAR, 0); }
		public TerminalNode R_BOOL() { return getToken(Sintactico.R_BOOL, 0); }
		public TerminalNode R_STRING() { return getToken(Sintactico.R_STRING, 0); }
		public TiposContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipos; }
	}

	public final TiposContext tipos() throws RecognitionException {
		TiposContext _localctx = new TiposContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_tipos);
		try {
			setState(75);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(63);
				match(R_INT);
				_localctx.tipo = entorno.INT
				}
				break;
			case R_FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(65);
				match(R_FLOAT);
				_localctx.tipo = entorno.FLOAT
				}
				break;
			case R_STR:
				enterOuterAlt(_localctx, 3);
				{
				setState(67);
				match(R_STR);
				_localctx.tipo = entorno.STRING
				}
				break;
			case R_CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(69);
				match(R_CHAR);
				_localctx.tipo = entorno.CHAR
				}
				break;
			case R_BOOL:
				enterOuterAlt(_localctx, 5);
				{
				setState(71);
				match(R_BOOL);
				_localctx.tipo = entorno.BOOL
				}
				break;
			case R_STRING:
				enterOuterAlt(_localctx, 6);
				{
				setState(73);
				match(R_STRING);
				_localctx.tipo = entorno.STRING
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
		enterRule(_localctx, 14, RULE_expr);
		try {
			setState(86);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(77);
				((ExprContext)_localctx).expr_log = expr_log(0);
				_localctx.p = ((ExprContext)_localctx).expr_log.p
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(80);
				((ExprContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((ExprContext)_localctx).expr_rel.p
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(83);
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
		int _startState = 16;
		enterRecursionRule(_localctx, 16, RULE_expr_op, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				setState(89);
				match(MENOS);
				setState(90);
				((Expr_opContext)_localctx).opU = ((Expr_opContext)_localctx).expr = expr();
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).opU.p, "-", nil, true)
				}
				break;
			case PARENA:
				{
				setState(93);
				match(PARENA);
				setState(94);
				((Expr_opContext)_localctx).expr = expr();
				setState(95);
				match(PARENC);
				_localctx.p = ((Expr_opContext)_localctx).expr.p
				}
				break;
			case R_INT:
				{
				setState(98);
				match(R_INT);
				setState(99);
				match(CUATROPT);
				setState(100);
				match(POW);
				setState(101);
				match(PARENA);
				setState(102);
				((Expr_opContext)_localctx).hIzq = expr_op(0);
				setState(103);
				match(COMA);
				setState(104);
				((Expr_opContext)_localctx).hDer = expr_op(0);
				setState(105);
				match(PARENC);
				_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, "**", ((Expr_opContext)_localctx).hDer.p, false)
				}
				break;
			case R_FLOAT:
				{
				setState(108);
				match(R_FLOAT);
				setState(109);
				match(CUATROPT);
				setState(110);
				match(POW);
				setState(111);
				match(PARENA);
				setState(112);
				((Expr_opContext)_localctx).hIzq = expr_op(0);
				setState(113);
				match(COMA);
				setState(114);
				((Expr_opContext)_localctx).hDer = expr_op(0);
				setState(115);
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
				setState(118);
				((Expr_opContext)_localctx).expr_val = expr_val();
				_localctx.p = ((Expr_opContext)_localctx).expr_val.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(140);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(138);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_opContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_op);
						setState(123);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(124);
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
						setState(125);
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
						setState(128);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(129);
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
						setState(130);
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
						setState(133);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(134);
						match(MODULO);
						setState(135);
						((Expr_opContext)_localctx).hDer = expr_op(6);
						_localctx.p = expresion.NewOperacion(((Expr_opContext)_localctx).hIzq.p, (((Expr_opContext)_localctx).op!=null?((Expr_opContext)_localctx).op.getText():null), ((Expr_opContext)_localctx).hDer.p, false)
						}
						break;
					}
					} 
				}
				setState(142);
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
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_expr_log, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				{
				setState(144);
				match(NOT);
				setState(145);
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
				setState(148);
				((Expr_logContext)_localctx).expr_rel = expr_rel(0);
				_localctx.p = ((Expr_logContext)_localctx).expr_rel.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(165);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(163);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_logContext(_parentctx, _parentState);
						_localctx.hIzq = _prevctx;
						_localctx.hIzq = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_log);
						setState(153);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(154);
						match(OR);
						setState(155);
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
						setState(158);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(159);
						match(AND);
						setState(160);
						((Expr_logContext)_localctx).hDer = expr_log(4);
						_localctx.p = expresion.NewLogicas(((Expr_logContext)_localctx).hIzq.p, "&&", ((Expr_logContext)_localctx).hDer.p, false)
						}
						break;
					}
					} 
				}
				setState(167);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_expr_rel, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(169);
			((Expr_relContext)_localctx).expr_op = expr_op(0);
			_localctx.p = ((Expr_relContext)_localctx).expr_op.p
			}
			_ctx.stop = _input.LT(-1);
			setState(179);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
					setState(172);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(173);
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
					setState(174);
					((Expr_relContext)_localctx).hDer = expr_rel(3);
					_localctx.p = expresion.NewRelacional(((Expr_relContext)_localctx).hIzq.p, (((Expr_relContext)_localctx).op!=null?((Expr_relContext)_localctx).op.getText():null), ((Expr_relContext)_localctx).hDer.p)
					}
					} 
				}
				setState(181);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
		enterRule(_localctx, 22, RULE_expr_val);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
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
		enterRule(_localctx, 24, RULE_primitivo);
		try {
			setState(197);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				enterOuterAlt(_localctx, 1);
				{
				setState(185);
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
				setState(187);
				((PrimitivoContext)_localctx).CADENA = match(CADENA);

				    str := (((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null)[1:len((((PrimitivoContext)_localctx).CADENA!=null?((PrimitivoContext)_localctx).CADENA.getText():null))-1]
				    _localctx.p = expresion.NewPrimitivo(str, entorno.STRING)

				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 3);
				{
				setState(189);
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
				setState(191);
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
				setState(193);
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
				setState(195);
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
		case 8:
			return expr_op_sempred((Expr_opContext)_localctx, predIndex);
		case 9:
			return expr_log_sempred((Expr_logContext)_localctx, predIndex);
		case 10:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3-\u00ca\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\2\3\3\6\3!\n\3\r\3\16\3\"\3\3"+
		"\3\3\3\4\3\4\3\4\3\4\3\4\3\4\5\4-\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\5\bN\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5"+
		"\tY\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\5\n|\n\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\7\n\u008d\n\n\f\n\16\n\u0090\13\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\3\13\5\13\u009a\n\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13"+
		"\7\13\u00a6\n\13\f\13\16\13\u00a9\13\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\7\f\u00b4\n\f\f\f\16\f\u00b7\13\f\3\r\3\r\3\r\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u00c8\n\16\3\16\2\5\22"+
		"\24\26\17\2\4\6\b\n\f\16\20\22\24\26\30\32\2\5\3\2#$\3\2!\"\3\2\25\32"+
		"\2\u00d5\2\34\3\2\2\2\4 \3\2\2\2\6,\3\2\2\2\b.\3\2\2\2\n\65\3\2\2\2\f"+
		">\3\2\2\2\16M\3\2\2\2\20X\3\2\2\2\22{\3\2\2\2\24\u0099\3\2\2\2\26\u00aa"+
		"\3\2\2\2\30\u00b8\3\2\2\2\32\u00c7\3\2\2\2\34\35\5\4\3\2\35\36\b\2\1\2"+
		"\36\3\3\2\2\2\37!\5\6\4\2 \37\3\2\2\2!\"\3\2\2\2\" \3\2\2\2\"#\3\2\2\2"+
		"#$\3\2\2\2$%\b\3\1\2%\5\3\2\2\2&\'\5\b\5\2\'(\b\4\1\2(-\3\2\2\2)*\5\n"+
		"\6\2*+\b\4\1\2+-\3\2\2\2,&\3\2\2\2,)\3\2\2\2-\7\3\2\2\2./\7\3\2\2/\60"+
		"\7\20\2\2\60\61\5\20\t\2\61\62\7\21\2\2\62\63\7\36\2\2\63\64\b\5\1\2\64"+
		"\t\3\2\2\2\65\66\7\r\2\2\66\67\7\b\2\2\678\5\f\7\289\7\35\2\29:\5\16\b"+
		"\2:;\7\33\2\2;<\5\20\t\2<=\b\6\1\2=\13\3\2\2\2>?\7,\2\2?@\b\7\1\2@\r\3"+
		"\2\2\2AB\7\4\2\2BN\b\b\1\2CD\7\5\2\2DN\b\b\1\2EF\7\6\2\2FN\b\b\1\2GH\7"+
		"\7\2\2HN\b\b\1\2IJ\7\n\2\2JN\b\b\1\2KL\7\t\2\2LN\b\b\1\2MA\3\2\2\2MC\3"+
		"\2\2\2ME\3\2\2\2MG\3\2\2\2MI\3\2\2\2MK\3\2\2\2N\17\3\2\2\2OP\5\24\13\2"+
		"PQ\b\t\1\2QY\3\2\2\2RS\5\26\f\2ST\b\t\1\2TY\3\2\2\2UV\5\22\n\2VW\b\t\1"+
		"\2WY\3\2\2\2XO\3\2\2\2XR\3\2\2\2XU\3\2\2\2Y\21\3\2\2\2Z[\b\n\1\2[\\\7"+
		"\"\2\2\\]\5\20\t\2]^\b\n\1\2^|\3\2\2\2_`\7\20\2\2`a\5\20\t\2ab\7\21\2"+
		"\2bc\b\n\1\2c|\3\2\2\2de\7\4\2\2ef\7\34\2\2fg\7\13\2\2gh\7\20\2\2hi\5"+
		"\22\n\2ij\7\37\2\2jk\5\22\n\2kl\7\21\2\2lm\b\n\1\2m|\3\2\2\2no\7\5\2\2"+
		"op\7\34\2\2pq\7\13\2\2qr\7\20\2\2rs\5\22\n\2st\7\37\2\2tu\5\22\n\2uv\7"+
		"\21\2\2vw\b\n\1\2w|\3\2\2\2xy\5\30\r\2yz\b\n\1\2z|\3\2\2\2{Z\3\2\2\2{"+
		"_\3\2\2\2{d\3\2\2\2{n\3\2\2\2{x\3\2\2\2|\u008e\3\2\2\2}~\f\t\2\2~\177"+
		"\t\2\2\2\177\u0080\5\22\n\n\u0080\u0081\b\n\1\2\u0081\u008d\3\2\2\2\u0082"+
		"\u0083\f\b\2\2\u0083\u0084\t\3\2\2\u0084\u0085\5\22\n\t\u0085\u0086\b"+
		"\n\1\2\u0086\u008d\3\2\2\2\u0087\u0088\f\7\2\2\u0088\u0089\7%\2\2\u0089"+
		"\u008a\5\22\n\b\u008a\u008b\b\n\1\2\u008b\u008d\3\2\2\2\u008c}\3\2\2\2"+
		"\u008c\u0082\3\2\2\2\u008c\u0087\3\2\2\2\u008d\u0090\3\2\2\2\u008e\u008c"+
		"\3\2\2\2\u008e\u008f\3\2\2\2\u008f\23\3\2\2\2\u0090\u008e\3\2\2\2\u0091"+
		"\u0092\b\13\1\2\u0092\u0093\7\24\2\2\u0093\u0094\5\24\13\4\u0094\u0095"+
		"\b\13\1\2\u0095\u009a\3\2\2\2\u0096\u0097\5\26\f\2\u0097\u0098\b\13\1"+
		"\2\u0098\u009a\3\2\2\2\u0099\u0091\3\2\2\2\u0099\u0096\3\2\2\2\u009a\u00a7"+
		"\3\2\2\2\u009b\u009c\f\6\2\2\u009c\u009d\7\23\2\2\u009d\u009e\5\24\13"+
		"\7\u009e\u009f\b\13\1\2\u009f\u00a6\3\2\2\2\u00a0\u00a1\f\5\2\2\u00a1"+
		"\u00a2\7\22\2\2\u00a2\u00a3\5\24\13\6\u00a3\u00a4\b\13\1\2\u00a4\u00a6"+
		"\3\2\2\2\u00a5\u009b\3\2\2\2\u00a5\u00a0\3\2\2\2\u00a6\u00a9\3\2\2\2\u00a7"+
		"\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\25\3\2\2\2\u00a9\u00a7\3\2\2"+
		"\2\u00aa\u00ab\b\f\1\2\u00ab\u00ac\5\22\n\2\u00ac\u00ad\b\f\1\2\u00ad"+
		"\u00b5\3\2\2\2\u00ae\u00af\f\4\2\2\u00af\u00b0\t\4\2\2\u00b0\u00b1\5\26"+
		"\f\5\u00b1\u00b2\b\f\1\2\u00b2\u00b4\3\2\2\2\u00b3\u00ae\3\2\2\2\u00b4"+
		"\u00b7\3\2\2\2\u00b5\u00b3\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\27\3\2\2"+
		"\2\u00b7\u00b5\3\2\2\2\u00b8\u00b9\5\32\16\2\u00b9\u00ba\b\r\1\2\u00ba"+
		"\31\3\2\2\2\u00bb\u00bc\7&\2\2\u00bc\u00c8\b\16\1\2\u00bd\u00be\7(\2\2"+
		"\u00be\u00c8\b\16\1\2\u00bf\u00c0\7\'\2\2\u00c0\u00c8\b\16\1\2\u00c1\u00c2"+
		"\7*\2\2\u00c2\u00c8\b\16\1\2\u00c3\u00c4\7+\2\2\u00c4\u00c8\b\16\1\2\u00c5"+
		"\u00c6\7)\2\2\u00c6\u00c8\b\16\1\2\u00c7\u00bb\3\2\2\2\u00c7\u00bd\3\2"+
		"\2\2\u00c7\u00bf\3\2\2\2\u00c7\u00c1\3\2\2\2\u00c7\u00c3\3\2\2\2\u00c7"+
		"\u00c5\3\2\2\2\u00c8\33\3\2\2\2\16\",MX{\u008c\u008e\u0099\u00a5\u00a7"+
		"\u00b5\u00c7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}