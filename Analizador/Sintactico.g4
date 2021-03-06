parser grammar Sintactico;

options {
    tokenVocab = Lexico;
}

@header {
    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Abstract"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Expresion"
    import arrayL "github.com/colegno/arrayList"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast/Instrucciones"
    import "OLC2_PROYECTO1_201901073/Analizador/Entorno"
    import "OLC2_PROYECTO1_201901073/Analizador/Ast"
}

start returns[ast.Ast ast]
: instrucciones                     {$ast = ast.NewAst($instrucciones.l)}
;

instrucciones returns [*arrayL.List l] @init {
    $l = arrayL.New()
} : e += instruccion+ {
    listInt := localctx.(*InstruccionesContext).GetE()
    for _,e := range listInt {
        $l.Add(e.GetInstr())
    }
    fmt.Printf("Tipo %T", localctx.(*InstruccionesContext).GetE())
}
;

instruccion returns[abstract.Instruccion instr]
: imprimir          {$instr = $imprimir.instr}
| declaracion       {$instr = $declaracion.instr}
;

imprimir returns[abstract.Instruccion instr]
    : PRINTLN PARENA expr PARENC PTCOMA {$instr = instrucciones.NewImprimir($expr.p)}
;

declaracion returns[abstract.Instruccion instr]
    : LET MUT listides DOSPT tipos IGUAL expr {$instr = instrucciones.NewDeclaracionValor($listides.lista, $tipos.tipo, $expr.p)}
;

listides returns[*arrayL.List lista]
@init { $lista = arrayL.New()}
    : ID            {$lista.Add(expresion.NewIdentificador($ID.text))}                      
;

tipos returns[entorno.TipoDato tipo]
    : R_INT                     {$tipo = entorno.INT}
    | R_FLOAT                   {$tipo = entorno.FLOAT}
    | R_STR                     {$tipo = entorno.STRING}
    | R_CHAR                    {$tipo = entorno.CHAR}
    | R_BOOL                    {$tipo = entorno.BOOL}
    | R_STRING                  {$tipo = entorno.STRING}
;

expr returns[abstract.Expresion p]
    : expr_log          {$p = $expr_log.p}
    | expr_rel          {$p = $expr_rel.p}
    | expr_op           {$p = $expr_op.p}
;

expr_op returns[abstract.Expresion p]
    : '-' opU = expr                             {$p = expresion.NewOperacion($opU.p, "-", nil, true)}
    | hIzq = expr_op op=('*'|'/') hDer = expr_op {$p = expresion.NewOperacion($hIzq.p, $op.text, $hDer.p, false)}
    | hIzq = expr_op op=('+'|'-') hDer = expr_op {$p = expresion.NewOperacion($hIzq.p, $op.text, $hDer.p, false)}
    | hIzq = expr_op '%' hDer = expr_op {$p = expresion.NewOperacion($hIzq.p, $op.text, $hDer.p, false)}
    | PARENA expr PARENC {$p = $expr.p}
    | R_INT CUATROPT POW PARENA hIzq = expr_op ',' hDer = expr_op PARENC {$p = expresion.NewOperacion($hIzq.p, "**", $hDer.p, false)}
    | R_FLOAT CUATROPT POW PARENA hIzq = expr_op ',' hDer = expr_op PARENC {$p = expresion.NewOperacion($hIzq.p, "f64**", $hDer.p, false)}
    | expr_val {$p = $expr_val.p}
;

expr_log returns[abstract.Expresion p]
    : hIzq = expr_log '||' hDer = expr_log {$p = expresion.NewLogicas($hIzq.p, "||", $hDer.p, false)}
    | hIzq = expr_log '&&' hDer = expr_log {$p = expresion.NewLogicas($hIzq.p, "&&", $hDer.p, false)}
    | '!' hIzq = expr_log   {$p = expresion.NewLogicas($hIzq.p, "!", nil, true)}
    | expr_rel {$p = $expr_rel.p}
;

expr_rel returns[abstract.Expresion p]
    : hIzq = expr_rel op = ('==' | '!=' | '<' | '<=' | '>' | '>=') hDer = expr_rel {$p = expresion.NewRelacional($hIzq.p, $op.text, $hDer.p)}
    | expr_op {$p = $expr_op.p}
;

expr_val returns[abstract.Expresion p]
: primitivo         {$p = $primitivo.p}
;

primitivo returns[abstract.Expresion p]
: ENTERO {
    num, err := strconv.Atoi($ENTERO.text)
    if err != nil {
        fmt.Println(err)
    }
    $p = expresion.NewPrimitivo(num, entorno.INT)
}
| CADENA {
    str := $CADENA.text[1:len($CADENA.text)-1]
    $p = expresion.NewPrimitivo(str, entorno.STRING)
}
| FLOAT {
    num, err := strconv.ParseFloat($FLOAT.text, 64)
    if err != nil {
        fmt.Println(err)
    }
    $p = expresion.NewPrimitivo(num, entorno.FLOAT)
}
| R_TRUE {
    bool_true, err := strconv.ParseBool($R_TRUE.text)
    if err != nil {
        fmt.Println(err)
    }
    $p = expresion.NewPrimitivo(bool_true, entorno.BOOL)
}
| R_FALSE {
    bool_false, err := strconv.ParseBool($R_FALSE.text)
    if err != nil {
        fmt.Println(err)
    }
    $p = expresion.NewPrimitivo(bool_false, entorno.BOOL)
}
| CARACTER {
    str := $CARACTER.text[1:len($CARACTER.text)-1]
    $p = expresion.NewPrimitivo(str, entorno.STRING)
}
;