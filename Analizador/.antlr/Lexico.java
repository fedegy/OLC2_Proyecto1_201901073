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
		PTCOMA=8, ENTERO=9, FLOAT=10, CADENA=11, R_TRUE=12, R_FALSE=13, ID=14, 
		WHITESPACE=15;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", "R_BOOL", 
			"PTCOMA", "ENTERO", "FLOAT", "CADENA", "R_TRUE", "R_FALSE", "ID", "WHITESPACE", 
			"ESC_SEQ"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'println!'", "'i64'", "'f64'", "'&str'", "'bool'", 
			"';'", null, null, null, "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARENA", "PARENC", "PRINTLN", "R_INT", "R_FLOAT", "R_STRING", 
			"R_BOOL", "PTCOMA", "ENTERO", "FLOAT", "CADENA", "R_TRUE", "R_FALSE", 
			"ID", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\21y\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\3\2\3\2\3"+
		"\3\3\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6"+
		"\3\6\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\n\6\nF\n\n\r\n"+
		"\16\nG\3\13\6\13K\n\13\r\13\16\13L\3\13\3\13\6\13Q\n\13\r\13\16\13R\3"+
		"\f\3\f\7\fW\n\f\f\f\16\fZ\13\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\17\3\17\7\17k\n\17\f\17\16\17n\13\17\3\20\6\20q\n"+
		"\20\r\20\16\20r\3\20\3\20\3\21\3\21\3\21\2\2\22\3\3\5\4\7\5\t\6\13\7\r"+
		"\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\2\3\2\t\3\2\62;"+
		"\3\2\60\60\3\2$$\5\2C\\aac|\6\2\62;C\\aac|\5\2\13\f\17\17\"\"\t\2\"#%"+
		"%--/\60<<BB]_\2}\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13"+
		"\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2"+
		"\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\3"+
		"#\3\2\2\2\5%\3\2\2\2\7\'\3\2\2\2\t\60\3\2\2\2\13\64\3\2\2\2\r8\3\2\2\2"+
		"\17=\3\2\2\2\21B\3\2\2\2\23E\3\2\2\2\25J\3\2\2\2\27T\3\2\2\2\31]\3\2\2"+
		"\2\33b\3\2\2\2\35h\3\2\2\2\37p\3\2\2\2!v\3\2\2\2#$\7*\2\2$\4\3\2\2\2%"+
		"&\7+\2\2&\6\3\2\2\2\'(\7r\2\2()\7t\2\2)*\7k\2\2*+\7p\2\2+,\7v\2\2,-\7"+
		"n\2\2-.\7p\2\2./\7#\2\2/\b\3\2\2\2\60\61\7k\2\2\61\62\78\2\2\62\63\7\66"+
		"\2\2\63\n\3\2\2\2\64\65\7h\2\2\65\66\78\2\2\66\67\7\66\2\2\67\f\3\2\2"+
		"\289\7(\2\29:\7u\2\2:;\7v\2\2;<\7t\2\2<\16\3\2\2\2=>\7d\2\2>?\7q\2\2?"+
		"@\7q\2\2@A\7n\2\2A\20\3\2\2\2BC\7=\2\2C\22\3\2\2\2DF\t\2\2\2ED\3\2\2\2"+
		"FG\3\2\2\2GE\3\2\2\2GH\3\2\2\2H\24\3\2\2\2IK\t\2\2\2JI\3\2\2\2KL\3\2\2"+
		"\2LJ\3\2\2\2LM\3\2\2\2MN\3\2\2\2NP\t\3\2\2OQ\t\2\2\2PO\3\2\2\2QR\3\2\2"+
		"\2RP\3\2\2\2RS\3\2\2\2S\26\3\2\2\2TX\7$\2\2UW\n\4\2\2VU\3\2\2\2WZ\3\2"+
		"\2\2XV\3\2\2\2XY\3\2\2\2Y[\3\2\2\2ZX\3\2\2\2[\\\7$\2\2\\\30\3\2\2\2]^"+
		"\7v\2\2^_\7t\2\2_`\7w\2\2`a\7g\2\2a\32\3\2\2\2bc\7h\2\2cd\7c\2\2de\7n"+
		"\2\2ef\7u\2\2fg\7g\2\2g\34\3\2\2\2hl\t\5\2\2ik\t\6\2\2ji\3\2\2\2kn\3\2"+
		"\2\2lj\3\2\2\2lm\3\2\2\2m\36\3\2\2\2nl\3\2\2\2oq\t\7\2\2po\3\2\2\2qr\3"+
		"\2\2\2rp\3\2\2\2rs\3\2\2\2st\3\2\2\2tu\b\20\2\2u \3\2\2\2vw\7^\2\2wx\t"+
		"\b\2\2x\"\3\2\2\2\t\2GLRXlr\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}