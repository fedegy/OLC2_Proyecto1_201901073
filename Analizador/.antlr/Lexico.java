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
		PARENA=1, PARENC=2, PRINTLN=3, R_INT=4, R_FLOAT=5, R_STRING=6, R_BOOL=7, 
		POW=8, AND=9, OR=10, NOT=11, DIFERENTE=12, MAYORIGUAL=13, MENORIGUAL=14, 
		MAYORQUE=15, MENORQUE=16, IGUALIGUAL=17, CUATROPT=18, PTCOMA=19, COMA=20, 
		MAS=21, MENOS=22, POR=23, DIVIDIDO=24, MODULO=25, ENTERO=26, FLOAT=27, 
		CADENA=28, CARACTER=29, R_TRUE=30, R_FALSE=31, ID=32, WHITESPACE=33;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL", 
			"POW", "AND", "OR", "NOT", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYORQUE", 
			"MENORQUE", "IGUALIGUAL", "CUATROPT", "PTCOMA", "COMA", "MAS", "MENOS", 
			"POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT", "CADENA", "CARACTER", 
			"R_TRUE", "R_FALSE", "ID", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", 
			"'pow'", "'&&'", "'||'", "'!'", "'!='", "'>='", "'<='", "'>'", "'<'", 
			"'=='", "'::'", "';'", "','", "'+'", "'-'", "'*'", "'/'", "'%'", null, 
			null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", 
			"R_BOOL", "POW", "AND", "OR", "NOT", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", 
			"MAYORQUE", "MENORQUE", "IGUALIGUAL", "CUATROPT", "PTCOMA", "COMA", "MAS", 
			"MENOS", "POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT", "CADENA", "CARACTER", 
			"R_TRUE", "R_FALSE", "ID", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2#\u00cc\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3"+
		"\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\16"+
		"\3\16\3\16\3\17\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\22\3\23\3\23"+
		"\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32"+
		"\3\32\3\33\6\33\u0095\n\33\r\33\16\33\u0096\3\34\6\34\u009a\n\34\r\34"+
		"\16\34\u009b\3\34\3\34\6\34\u00a0\n\34\r\34\16\34\u00a1\3\35\3\35\7\35"+
		"\u00a6\n\35\f\35\16\35\u00a9\13\35\3\35\3\35\3\36\3\36\3\36\3\36\3\37"+
		"\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3!\3!\7!\u00be\n!\f!\16!\u00c1"+
		"\13!\3\"\6\"\u00c4\n\"\r\"\16\"\u00c5\3\"\3\"\3#\3#\3#\2\2$\3\3\5\4\7"+
		"\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C"+
		"#E\2\3\2\n\3\2\62;\3\2\60\60\3\2$$\3\2))\5\2C\\aac|\6\2\62;C\\aac|\5\2"+
		"\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u00d0\2\3\3\2\2\2\2\5\3\2\2\2\2"+
		"\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2"+
		"\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2"+
		"\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2"+
		"\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2"+
		"\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2"+
		"\2\2A\3\2\2\2\2C\3\2\2\2\3G\3\2\2\2\5I\3\2\2\2\7K\3\2\2\2\tT\3\2\2\2\13"+
		"X\3\2\2\2\r\\\3\2\2\2\17a\3\2\2\2\21f\3\2\2\2\23j\3\2\2\2\25m\3\2\2\2"+
		"\27p\3\2\2\2\31r\3\2\2\2\33u\3\2\2\2\35x\3\2\2\2\37{\3\2\2\2!}\3\2\2\2"+
		"#\177\3\2\2\2%\u0082\3\2\2\2\'\u0085\3\2\2\2)\u0087\3\2\2\2+\u0089\3\2"+
		"\2\2-\u008b\3\2\2\2/\u008d\3\2\2\2\61\u008f\3\2\2\2\63\u0091\3\2\2\2\65"+
		"\u0094\3\2\2\2\67\u0099\3\2\2\29\u00a3\3\2\2\2;\u00ac\3\2\2\2=\u00b0\3"+
		"\2\2\2?\u00b5\3\2\2\2A\u00bb\3\2\2\2C\u00c3\3\2\2\2E\u00c9\3\2\2\2GH\7"+
		"*\2\2H\4\3\2\2\2IJ\7+\2\2J\6\3\2\2\2KL\7r\2\2LM\7t\2\2MN\7k\2\2NO\7p\2"+
		"\2OP\7v\2\2PQ\7n\2\2QR\7p\2\2RS\7#\2\2S\b\3\2\2\2TU\7k\2\2UV\78\2\2VW"+
		"\7\66\2\2W\n\3\2\2\2XY\7h\2\2YZ\78\2\2Z[\7\66\2\2[\f\3\2\2\2\\]\7(\2\2"+
		"]^\7u\2\2^_\7v\2\2_`\7t\2\2`\16\3\2\2\2ab\7d\2\2bc\7q\2\2cd\7q\2\2de\7"+
		"n\2\2e\20\3\2\2\2fg\7r\2\2gh\7q\2\2hi\7y\2\2i\22\3\2\2\2jk\7(\2\2kl\7"+
		"(\2\2l\24\3\2\2\2mn\7~\2\2no\7~\2\2o\26\3\2\2\2pq\7#\2\2q\30\3\2\2\2r"+
		"s\7#\2\2st\7?\2\2t\32\3\2\2\2uv\7@\2\2vw\7?\2\2w\34\3\2\2\2xy\7>\2\2y"+
		"z\7?\2\2z\36\3\2\2\2{|\7@\2\2| \3\2\2\2}~\7>\2\2~\"\3\2\2\2\177\u0080"+
		"\7?\2\2\u0080\u0081\7?\2\2\u0081$\3\2\2\2\u0082\u0083\7<\2\2\u0083\u0084"+
		"\7<\2\2\u0084&\3\2\2\2\u0085\u0086\7=\2\2\u0086(\3\2\2\2\u0087\u0088\7"+
		".\2\2\u0088*\3\2\2\2\u0089\u008a\7-\2\2\u008a,\3\2\2\2\u008b\u008c\7/"+
		"\2\2\u008c.\3\2\2\2\u008d\u008e\7,\2\2\u008e\60\3\2\2\2\u008f\u0090\7"+
		"\61\2\2\u0090\62\3\2\2\2\u0091\u0092\7\'\2\2\u0092\64\3\2\2\2\u0093\u0095"+
		"\t\2\2\2\u0094\u0093\3\2\2\2\u0095\u0096\3\2\2\2\u0096\u0094\3\2\2\2\u0096"+
		"\u0097\3\2\2\2\u0097\66\3\2\2\2\u0098\u009a\t\2\2\2\u0099\u0098\3\2\2"+
		"\2\u009a\u009b\3\2\2\2\u009b\u0099\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u009d"+
		"\3\2\2\2\u009d\u009f\t\3\2\2\u009e\u00a0\t\2\2\2\u009f\u009e\3\2\2\2\u00a0"+
		"\u00a1\3\2\2\2\u00a1\u009f\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a28\3\2\2\2"+
		"\u00a3\u00a7\7$\2\2\u00a4\u00a6\n\4\2\2\u00a5\u00a4\3\2\2\2\u00a6\u00a9"+
		"\3\2\2\2\u00a7\u00a5\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\u00aa\3\2\2\2\u00a9"+
		"\u00a7\3\2\2\2\u00aa\u00ab\7$\2\2\u00ab:\3\2\2\2\u00ac\u00ad\7)\2\2\u00ad"+
		"\u00ae\n\5\2\2\u00ae\u00af\7)\2\2\u00af<\3\2\2\2\u00b0\u00b1\7v\2\2\u00b1"+
		"\u00b2\7t\2\2\u00b2\u00b3\7w\2\2\u00b3\u00b4\7g\2\2\u00b4>\3\2\2\2\u00b5"+
		"\u00b6\7h\2\2\u00b6\u00b7\7c\2\2\u00b7\u00b8\7n\2\2\u00b8\u00b9\7u\2\2"+
		"\u00b9\u00ba\7g\2\2\u00ba@\3\2\2\2\u00bb\u00bf\t\6\2\2\u00bc\u00be\t\7"+
		"\2\2\u00bd\u00bc\3\2\2\2\u00be\u00c1\3\2\2\2\u00bf\u00bd\3\2\2\2\u00bf"+
		"\u00c0\3\2\2\2\u00c0B\3\2\2\2\u00c1\u00bf\3\2\2\2\u00c2\u00c4\t\b\2\2"+
		"\u00c3\u00c2\3\2\2\2\u00c4\u00c5\3\2\2\2\u00c5\u00c3\3\2\2\2\u00c5\u00c6"+
		"\3\2\2\2\u00c6\u00c7\3\2\2\2\u00c7\u00c8\b\"\2\2\u00c8D\3\2\2\2\u00c9"+
		"\u00ca\7^\2\2\u00ca\u00cb\t\t\2\2\u00cbF\3\2\2\2\t\2\u0096\u009b\u00a1"+
		"\u00a7\u00bf\u00c5\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}