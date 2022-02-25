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
		CADENA=28, R_TRUE=29, R_FALSE=30, ID=31, WHITESPACE=32;
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
			"POR", "DIVIDIDO", "MODULO", "ENTERO", "FLOAT", "CADENA", "R_TRUE", "R_FALSE", 
			"ID", "WHITESPACE", "ESC_SEQ"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\"\u00c6\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3"+
		"\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t"+
		"\3\t\3\t\3\t\3\n\3\n\3\n\3\13\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\16\3\16"+
		"\3\16\3\17\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23"+
		"\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32"+
		"\3\33\6\33\u0093\n\33\r\33\16\33\u0094\3\34\6\34\u0098\n\34\r\34\16\34"+
		"\u0099\3\34\3\34\6\34\u009e\n\34\r\34\16\34\u009f\3\35\3\35\7\35\u00a4"+
		"\n\35\f\35\16\35\u00a7\13\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\37\3"+
		"\37\3\37\3\37\3\37\3\37\3 \3 \7 \u00b8\n \f \16 \u00bb\13 \3!\6!\u00be"+
		"\n!\r!\16!\u00bf\3!\3!\3\"\3\"\3\"\2\2#\3\3\5\4\7\5\t\6\13\7\r\b\17\t"+
		"\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27"+
		"-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C\2\3\2\t\3\2\62;\3\2\60"+
		"\60\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\5\2\13\f\17\17\"\"\t\2\"#%%--/\60"+
		"<<BB]_\2\u00ca\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3"+
		"\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2"+
		"\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3"+
		"\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2"+
		"\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\2"+
		"9\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\3E\3\2\2\2\5G\3"+
		"\2\2\2\7I\3\2\2\2\tR\3\2\2\2\13V\3\2\2\2\rZ\3\2\2\2\17_\3\2\2\2\21d\3"+
		"\2\2\2\23h\3\2\2\2\25k\3\2\2\2\27n\3\2\2\2\31p\3\2\2\2\33s\3\2\2\2\35"+
		"v\3\2\2\2\37y\3\2\2\2!{\3\2\2\2#}\3\2\2\2%\u0080\3\2\2\2\'\u0083\3\2\2"+
		"\2)\u0085\3\2\2\2+\u0087\3\2\2\2-\u0089\3\2\2\2/\u008b\3\2\2\2\61\u008d"+
		"\3\2\2\2\63\u008f\3\2\2\2\65\u0092\3\2\2\2\67\u0097\3\2\2\29\u00a1\3\2"+
		"\2\2;\u00aa\3\2\2\2=\u00af\3\2\2\2?\u00b5\3\2\2\2A\u00bd\3\2\2\2C\u00c3"+
		"\3\2\2\2EF\7*\2\2F\4\3\2\2\2GH\7+\2\2H\6\3\2\2\2IJ\7r\2\2JK\7t\2\2KL\7"+
		"k\2\2LM\7p\2\2MN\7v\2\2NO\7n\2\2OP\7p\2\2PQ\7#\2\2Q\b\3\2\2\2RS\7k\2\2"+
		"ST\78\2\2TU\7\66\2\2U\n\3\2\2\2VW\7h\2\2WX\78\2\2XY\7\66\2\2Y\f\3\2\2"+
		"\2Z[\7(\2\2[\\\7u\2\2\\]\7v\2\2]^\7t\2\2^\16\3\2\2\2_`\7d\2\2`a\7q\2\2"+
		"ab\7q\2\2bc\7n\2\2c\20\3\2\2\2de\7r\2\2ef\7q\2\2fg\7y\2\2g\22\3\2\2\2"+
		"hi\7(\2\2ij\7(\2\2j\24\3\2\2\2kl\7~\2\2lm\7~\2\2m\26\3\2\2\2no\7#\2\2"+
		"o\30\3\2\2\2pq\7#\2\2qr\7?\2\2r\32\3\2\2\2st\7@\2\2tu\7?\2\2u\34\3\2\2"+
		"\2vw\7>\2\2wx\7?\2\2x\36\3\2\2\2yz\7@\2\2z \3\2\2\2{|\7>\2\2|\"\3\2\2"+
		"\2}~\7?\2\2~\177\7?\2\2\177$\3\2\2\2\u0080\u0081\7<\2\2\u0081\u0082\7"+
		"<\2\2\u0082&\3\2\2\2\u0083\u0084\7=\2\2\u0084(\3\2\2\2\u0085\u0086\7."+
		"\2\2\u0086*\3\2\2\2\u0087\u0088\7-\2\2\u0088,\3\2\2\2\u0089\u008a\7/\2"+
		"\2\u008a.\3\2\2\2\u008b\u008c\7,\2\2\u008c\60\3\2\2\2\u008d\u008e\7\61"+
		"\2\2\u008e\62\3\2\2\2\u008f\u0090\7\'\2\2\u0090\64\3\2\2\2\u0091\u0093"+
		"\t\2\2\2\u0092\u0091\3\2\2\2\u0093\u0094\3\2\2\2\u0094\u0092\3\2\2\2\u0094"+
		"\u0095\3\2\2\2\u0095\66\3\2\2\2\u0096\u0098\t\2\2\2\u0097\u0096\3\2\2"+
		"\2\u0098\u0099\3\2\2\2\u0099\u0097\3\2\2\2\u0099\u009a\3\2\2\2\u009a\u009b"+
		"\3\2\2\2\u009b\u009d\t\3\2\2\u009c\u009e\t\2\2\2\u009d\u009c\3\2\2\2\u009e"+
		"\u009f\3\2\2\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a08\3\2\2\2"+
		"\u00a1\u00a5\7$\2\2\u00a2\u00a4\n\4\2\2\u00a3\u00a2\3\2\2\2\u00a4\u00a7"+
		"\3\2\2\2\u00a5\u00a3\3\2\2\2\u00a5\u00a6\3\2\2\2\u00a6\u00a8\3\2\2\2\u00a7"+
		"\u00a5\3\2\2\2\u00a8\u00a9\7$\2\2\u00a9:\3\2\2\2\u00aa\u00ab\7v\2\2\u00ab"+
		"\u00ac\7t\2\2\u00ac\u00ad\7w\2\2\u00ad\u00ae\7g\2\2\u00ae<\3\2\2\2\u00af"+
		"\u00b0\7h\2\2\u00b0\u00b1\7c\2\2\u00b1\u00b2\7n\2\2\u00b2\u00b3\7u\2\2"+
		"\u00b3\u00b4\7g\2\2\u00b4>\3\2\2\2\u00b5\u00b9\t\5\2\2\u00b6\u00b8\t\6"+
		"\2\2\u00b7\u00b6\3\2\2\2\u00b8\u00bb\3\2\2\2\u00b9\u00b7\3\2\2\2\u00b9"+
		"\u00ba\3\2\2\2\u00ba@\3\2\2\2\u00bb\u00b9\3\2\2\2\u00bc\u00be\t\7\2\2"+
		"\u00bd\u00bc\3\2\2\2\u00be\u00bf\3\2\2\2\u00bf\u00bd\3\2\2\2\u00bf\u00c0"+
		"\3\2\2\2\u00c0\u00c1\3\2\2\2\u00c1\u00c2\b!\2\2\u00c2B\3\2\2\2\u00c3\u00c4"+
		"\7^\2\2\u00c4\u00c5\t\b\2\2\u00c5D\3\2\2\2\t\2\u0094\u0099\u009f\u00a5"+
		"\u00b9\u00bf\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}