// Generated from /home/federico/go/src/OLC2_Proyecto1_201901073/Analizador/Lexico.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Lexico extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PRINTLN", "R_INT", "R_FLOAT", "R_STR", "R_CHAR", "MUT", "R_STRING", 
			"R_BOOL", "POW", "AS", "LET", "TOOWNED", "TOSTRING", "PARENA", "PARENC", 
			"AND", "OR", "NOT", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", 
			"MENORQUE", "IGUALIGUAL", "IGUAL", "CUATROPT", "DOSPT", "PTCOMA", "COMA", 
			"PUNTO", "MAS", "MENOS", "POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT", 
			"CADENA", "CARACTER", "R_TRUE", "R_FALSE", "ID", "WHITESPACE", "ESC_SEQ"
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


	public Lexico(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lexico.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2-\u0114\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\4\3\4"+
		"\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3"+
		"\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\20\3"+
		"\20\3\21\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3"+
		"\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\33\3"+
		"\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3 \3 \3!\3!\3\"\3\"\3"+
		"#\3#\3$\3$\3%\6%\u00dd\n%\r%\16%\u00de\3&\6&\u00e2\n&\r&\16&\u00e3\3&"+
		"\3&\6&\u00e8\n&\r&\16&\u00e9\3\'\3\'\7\'\u00ee\n\'\f\'\16\'\u00f1\13\'"+
		"\3\'\3\'\3(\3(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3+\3+\7+\u0106\n"+
		"+\f+\16+\u0109\13+\3,\6,\u010c\n,\r,\16,\u010d\3,\3,\3-\3-\3-\2\2.\3\3"+
		"\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21"+
		"!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!"+
		"A\"C#E$G%I&K\'M(O)Q*S+U,W-Y\2\3\2\n\3\2\62;\3\2\60\60\3\2$$\3\2))\5\2"+
		"C\\aac|\6\2\62;C\\aac|\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u0118"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\3[\3\2\2\2\5d\3\2\2\2\7h\3\2\2\2\tl\3\2\2\2\13q\3"+
		"\2\2\2\rv\3\2\2\2\17z\3\2\2\2\21\u0081\3\2\2\2\23\u0086\3\2\2\2\25\u008a"+
		"\3\2\2\2\27\u008d\3\2\2\2\31\u0091\3\2\2\2\33\u009c\3\2\2\2\35\u00a8\3"+
		"\2\2\2\37\u00aa\3\2\2\2!\u00ac\3\2\2\2#\u00af\3\2\2\2%\u00b2\3\2\2\2\'"+
		"\u00b4\3\2\2\2)\u00b7\3\2\2\2+\u00ba\3\2\2\2-\u00bd\3\2\2\2/\u00bf\3\2"+
		"\2\2\61\u00c1\3\2\2\2\63\u00c4\3\2\2\2\65\u00c6\3\2\2\2\67\u00c9\3\2\2"+
		"\29\u00cb\3\2\2\2;\u00cd\3\2\2\2=\u00cf\3\2\2\2?\u00d1\3\2\2\2A\u00d3"+
		"\3\2\2\2C\u00d5\3\2\2\2E\u00d7\3\2\2\2G\u00d9\3\2\2\2I\u00dc\3\2\2\2K"+
		"\u00e1\3\2\2\2M\u00eb\3\2\2\2O\u00f4\3\2\2\2Q\u00f8\3\2\2\2S\u00fd\3\2"+
		"\2\2U\u0103\3\2\2\2W\u010b\3\2\2\2Y\u0111\3\2\2\2[\\\7r\2\2\\]\7t\2\2"+
		"]^\7k\2\2^_\7p\2\2_`\7v\2\2`a\7n\2\2ab\7p\2\2bc\7#\2\2c\4\3\2\2\2de\7"+
		"k\2\2ef\78\2\2fg\7\66\2\2g\6\3\2\2\2hi\7h\2\2ij\78\2\2jk\7\66\2\2k\b\3"+
		"\2\2\2lm\7(\2\2mn\7u\2\2no\7v\2\2op\7t\2\2p\n\3\2\2\2qr\7e\2\2rs\7j\2"+
		"\2st\7c\2\2tu\7t\2\2u\f\3\2\2\2vw\7o\2\2wx\7w\2\2xy\7v\2\2y\16\3\2\2\2"+
		"z{\7U\2\2{|\7v\2\2|}\7t\2\2}~\7k\2\2~\177\7p\2\2\177\u0080\7i\2\2\u0080"+
		"\20\3\2\2\2\u0081\u0082\7d\2\2\u0082\u0083\7q\2\2\u0083\u0084\7q\2\2\u0084"+
		"\u0085\7n\2\2\u0085\22\3\2\2\2\u0086\u0087\7r\2\2\u0087\u0088\7q\2\2\u0088"+
		"\u0089\7y\2\2\u0089\24\3\2\2\2\u008a\u008b\7c\2\2\u008b\u008c\7u\2\2\u008c"+
		"\26\3\2\2\2\u008d\u008e\7n\2\2\u008e\u008f\7g\2\2\u008f\u0090\7v\2\2\u0090"+
		"\30\3\2\2\2\u0091\u0092\7v\2\2\u0092\u0093\7q\2\2\u0093\u0094\7a\2\2\u0094"+
		"\u0095\7q\2\2\u0095\u0096\7y\2\2\u0096\u0097\7p\2\2\u0097\u0098\7g\2\2"+
		"\u0098\u0099\7f\2\2\u0099\u009a\7*\2\2\u009a\u009b\7+\2\2\u009b\32\3\2"+
		"\2\2\u009c\u009d\7v\2\2\u009d\u009e\7q\2\2\u009e\u009f\7a\2\2\u009f\u00a0"+
		"\7u\2\2\u00a0\u00a1\7v\2\2\u00a1\u00a2\7t\2\2\u00a2\u00a3\7k\2\2\u00a3"+
		"\u00a4\7p\2\2\u00a4\u00a5\7i\2\2\u00a5\u00a6\7*\2\2\u00a6\u00a7\7+\2\2"+
		"\u00a7\34\3\2\2\2\u00a8\u00a9\7*\2\2\u00a9\36\3\2\2\2\u00aa\u00ab\7+\2"+
		"\2\u00ab \3\2\2\2\u00ac\u00ad\7(\2\2\u00ad\u00ae\7(\2\2\u00ae\"\3\2\2"+
		"\2\u00af\u00b0\7~\2\2\u00b0\u00b1\7~\2\2\u00b1$\3\2\2\2\u00b2\u00b3\7"+
		"#\2\2\u00b3&\3\2\2\2\u00b4\u00b5\7#\2\2\u00b5\u00b6\7?\2\2\u00b6(\3\2"+
		"\2\2\u00b7\u00b8\7@\2\2\u00b8\u00b9\7?\2\2\u00b9*\3\2\2\2\u00ba\u00bb"+
		"\7>\2\2\u00bb\u00bc\7?\2\2\u00bc,\3\2\2\2\u00bd\u00be\7@\2\2\u00be.\3"+
		"\2\2\2\u00bf\u00c0\7>\2\2\u00c0\60\3\2\2\2\u00c1\u00c2\7?\2\2\u00c2\u00c3"+
		"\7?\2\2\u00c3\62\3\2\2\2\u00c4\u00c5\7?\2\2\u00c5\64\3\2\2\2\u00c6\u00c7"+
		"\7<\2\2\u00c7\u00c8\7<\2\2\u00c8\66\3\2\2\2\u00c9\u00ca\7<\2\2\u00ca8"+
		"\3\2\2\2\u00cb\u00cc\7=\2\2\u00cc:\3\2\2\2\u00cd\u00ce\7.\2\2\u00ce<\3"+
		"\2\2\2\u00cf\u00d0\7\60\2\2\u00d0>\3\2\2\2\u00d1\u00d2\7-\2\2\u00d2@\3"+
		"\2\2\2\u00d3\u00d4\7/\2\2\u00d4B\3\2\2\2\u00d5\u00d6\7,\2\2\u00d6D\3\2"+
		"\2\2\u00d7\u00d8\7\61\2\2\u00d8F\3\2\2\2\u00d9\u00da\7\'\2\2\u00daH\3"+
		"\2\2\2\u00db\u00dd\t\2\2\2\u00dc\u00db\3\2\2\2\u00dd\u00de\3\2\2\2\u00de"+
		"\u00dc\3\2\2\2\u00de\u00df\3\2\2\2\u00dfJ\3\2\2\2\u00e0\u00e2\t\2\2\2"+
		"\u00e1\u00e0\3\2\2\2\u00e2\u00e3\3\2\2\2\u00e3\u00e1\3\2\2\2\u00e3\u00e4"+
		"\3\2\2\2\u00e4\u00e5\3\2\2\2\u00e5\u00e7\t\3\2\2\u00e6\u00e8\t\2\2\2\u00e7"+
		"\u00e6\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00e7\3\2\2\2\u00e9\u00ea\3\2"+
		"\2\2\u00eaL\3\2\2\2\u00eb\u00ef\7$\2\2\u00ec\u00ee\n\4\2\2\u00ed\u00ec"+
		"\3\2\2\2\u00ee\u00f1\3\2\2\2\u00ef\u00ed\3\2\2\2\u00ef\u00f0\3\2\2\2\u00f0"+
		"\u00f2\3\2\2\2\u00f1\u00ef\3\2\2\2\u00f2\u00f3\7$\2\2\u00f3N\3\2\2\2\u00f4"+
		"\u00f5\7)\2\2\u00f5\u00f6\n\5\2\2\u00f6\u00f7\7)\2\2\u00f7P\3\2\2\2\u00f8"+
		"\u00f9\7v\2\2\u00f9\u00fa\7t\2\2\u00fa\u00fb\7w\2\2\u00fb\u00fc\7g\2\2"+
		"\u00fcR\3\2\2\2\u00fd\u00fe\7h\2\2\u00fe\u00ff\7c\2\2\u00ff\u0100\7n\2"+
		"\2\u0100\u0101\7u\2\2\u0101\u0102\7g\2\2\u0102T\3\2\2\2\u0103\u0107\t"+
		"\6\2\2\u0104\u0106\t\7\2\2\u0105\u0104\3\2\2\2\u0106\u0109\3\2\2\2\u0107"+
		"\u0105\3\2\2\2\u0107\u0108\3\2\2\2\u0108V\3\2\2\2\u0109\u0107\3\2\2\2"+
		"\u010a\u010c\t\b\2\2\u010b\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d\u010b"+
		"\3\2\2\2\u010d\u010e\3\2\2\2\u010e\u010f\3\2\2\2\u010f\u0110\b,\2\2\u0110"+
		"X\3\2\2\2\u0111\u0112\7^\2\2\u0112\u0113\t\t\2\2\u0113Z\3\2\2\2\t\2\u00de"+
		"\u00e3\u00e9\u00ef\u0107\u010d\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}