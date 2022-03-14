lexer grammar Lexico;

PRINTLN         : 'println!';
R_INT           : 'i64';
R_FLOAT         : 'f64';
R_STR           : '&str';
R_CHAR          : 'char';
MUT             : 'mut';
R_STRING        : 'String';
R_BOOL          : 'bool';
POW             : 'pow';
AS              : 'as';
LET             : 'let';
TOOWNED         : 'to_owned()';
TOSTRING        : 'to_string()';

PARENA          : '(';
PARENC          : ')';
AND             : '&&';
OR              : '||';
NOT             : '!';
DIFERENTE       : '!=';
MAYORIGUAL      : '>=';
MENORIGUAL      : '<=';
MAYORQUE        : '>';
MENORQUE        : '<';
IGUALIGUAL      : '==';

IGUAL           : '=';
CUATROPT        : '::';
DOSPT           : ':';
PTCOMA          : ';';
COMA            : ',';
PUNTO           : '.';

MAS             : '+';
MENOS           : '-';
POR             : '*';
DIVIDIDO        : '/';
MODULO          : '%';

ENTERO          : [0-9]+;
FLOAT           : [0-9]+[.][0-9]+;
CADENA          : '"'~["]*'"';
CARACTER        : '\''~[']'\'';
R_TRUE          : 'true';
R_FALSE         : 'false';
ID              : [a-zA-Z_][a-zA-Z0-9_]*;
WHITESPACE      : [ \r\n\t]+ -> skip;

fragment
ESC_SEQ : '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
;
