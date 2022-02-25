lexer grammar Lexico;

PARENA          : '(';
PARENC          : ')';
PRINTLN         : 'println!';
R_INT           : 'i64';
R_FLOAT         : 'f64';
R_STRING        : '&str';
R_BOOL          : 'bool';
POW             : 'pow';

CUATROPT        : '::';
PTCOMA          : ';';
COMA            : ',';

MAS             : '+';
MENOS           : '-';
POR             : '*';
DIVIDIDO        : '/';
MODULO          : '%';

ENTERO          : [0-9]+;
FLOAT           : [0-9]+[.][0-9]+;
CADENA          : '"'~["]*'"';
R_TRUE          : 'true';
R_FALSE         : 'false';
ID              : [a-zA-Z_][a-zA-Z0-9_]*;
WHITESPACE      : [ \r\n\t]+ -> skip;

fragment
ESC_SEQ : '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
;