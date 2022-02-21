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
		PTCOMA=8, MAS=9, MENOS=10, POR=11, DIVIDIDO=12, ENTERO=13, FLOAT=14, CADENA=15, 
		R_TRUE=16, R_FALSE=17, ID=18, WHITESPACE=19;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL", 
			"PTCOMA", "MAS", "MENOS", "POR", "DIVIDIDO", "ENTERO", "FLOAT", "CADENA", 
			"R_TRUE", "R_FALSE", "ID", "WHITESPACE", "ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", 
			"';'", "'+'", "'-'", "'*'", "'/'", null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", 
			"R_BOOL", "PTCOMA", "MAS", "MENOS", "POR", "DIVIDIDO", "ENTERO", "FLOAT", 
			"CADENA", "R_TRUE", "R_FALSE", "ID", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\25\u0089\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\3\3\3\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7"+
		"\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\6"+
		"\16V\n\16\r\16\16\16W\3\17\6\17[\n\17\r\17\16\17\\\3\17\3\17\6\17a\n\17"+
		"\r\17\16\17b\3\20\3\20\7\20g\n\20\f\20\16\20j\13\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\23\3\23\7\23{\n\23\f\23"+
		"\16\23~\13\23\3\24\6\24\u0081\n\24\r\24\16\24\u0082\3\24\3\24\3\25\3\25"+
		"\3\25\2\2\26\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16"+
		"\33\17\35\20\37\21!\22#\23%\24\'\25)\2\3\2\t\3\2\62;\3\2\60\60\3\2$$\5"+
		"\2C\\aac|\6\2\62;C\\aac|\5\2\13\f\17\17\"\"\t\2\"#%%--/\60<<BB]_\2\u008d"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\3+\3\2\2\2\5-\3\2\2\2\7/\3\2\2\2\t8\3\2\2"+
		"\2\13<\3\2\2\2\r@\3\2\2\2\17E\3\2\2\2\21J\3\2\2\2\23L\3\2\2\2\25N\3\2"+
		"\2\2\27P\3\2\2\2\31R\3\2\2\2\33U\3\2\2\2\35Z\3\2\2\2\37d\3\2\2\2!m\3\2"+
		"\2\2#r\3\2\2\2%x\3\2\2\2\'\u0080\3\2\2\2)\u0086\3\2\2\2+,\7*\2\2,\4\3"+
		"\2\2\2-.\7+\2\2.\6\3\2\2\2/\60\7r\2\2\60\61\7t\2\2\61\62\7k\2\2\62\63"+
		"\7p\2\2\63\64\7v\2\2\64\65\7n\2\2\65\66\7p\2\2\66\67\7#\2\2\67\b\3\2\2"+
		"\289\7k\2\29:\78\2\2:;\7\66\2\2;\n\3\2\2\2<=\7h\2\2=>\78\2\2>?\7\66\2"+
		"\2?\f\3\2\2\2@A\7(\2\2AB\7u\2\2BC\7v\2\2CD\7t\2\2D\16\3\2\2\2EF\7d\2\2"+
		"FG\7q\2\2GH\7q\2\2HI\7n\2\2I\20\3\2\2\2JK\7=\2\2K\22\3\2\2\2LM\7-\2\2"+
		"M\24\3\2\2\2NO\7/\2\2O\26\3\2\2\2PQ\7,\2\2Q\30\3\2\2\2RS\7\61\2\2S\32"+
		"\3\2\2\2TV\t\2\2\2UT\3\2\2\2VW\3\2\2\2WU\3\2\2\2WX\3\2\2\2X\34\3\2\2\2"+
		"Y[\t\2\2\2ZY\3\2\2\2[\\\3\2\2\2\\Z\3\2\2\2\\]\3\2\2\2]^\3\2\2\2^`\t\3"+
		"\2\2_a\t\2\2\2`_\3\2\2\2ab\3\2\2\2b`\3\2\2\2bc\3\2\2\2c\36\3\2\2\2dh\7"+
		"$\2\2eg\n\4\2\2fe\3\2\2\2gj\3\2\2\2hf\3\2\2\2hi\3\2\2\2ik\3\2\2\2jh\3"+
		"\2\2\2kl\7$\2\2l \3\2\2\2mn\7v\2\2no\7t\2\2op\7w\2\2pq\7g\2\2q\"\3\2\2"+
		"\2rs\7h\2\2st\7c\2\2tu\7n\2\2uv\7u\2\2vw\7g\2\2w$\3\2\2\2x|\t\5\2\2y{"+
		"\t\6\2\2zy\3\2\2\2{~\3\2\2\2|z\3\2\2\2|}\3\2\2\2}&\3\2\2\2~|\3\2\2\2\177"+
		"\u0081\t\7\2\2\u0080\177\3\2\2\2\u0081\u0082\3\2\2\2\u0082\u0080\3\2\2"+
		"\2\u0082\u0083\3\2\2\2\u0083\u0084\3\2\2\2\u0084\u0085\b\24\2\2\u0085"+
		"(\3\2\2\2\u0086\u0087\7^\2\2\u0087\u0088\t\b\2\2\u0088*\3\2\2\2\t\2W\\"+
		"bh|\u0082\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}