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
}

start returns [*arrayL.List lista] : instrucciones {$lista = $instrucciones.l}
;

instrucciones returns [*arrayL.List l] @init {
    $l = arrayL.New()
} : e += instruccion* {
    listInt := localctx.(*InstruccionesContext).GetE()
    for _,e := range listInt {
        $l.Add(e.GetInstr())
    }
    fmt.Printf("Tipo %T", localctx.(*InstruccionesContext).GetE())
}
;

instruccion returns[abstract.Instruccion instr]
: PRINTLN PARENA expr PARENC PTCOMA {$instr = instrucciones.NewImprimir($expr.p)}
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
    | primitivo {$p = $primitivo.p}
    | PARENA expr PARENC {$p = $expr.p}
    | R_INT CUATROPT POW PARENA hIzq = expr_op ',' hDer = expr_op PARENC {$p = expresion.NewOperacion($hIzq.p, "**", $hDer.p, false)}
    | R_FLOAT CUATROPT POW PARENA hIzq = expr_op ',' hDer = expr_op PARENC {$p = expresion.NewOperacion($hIzq.p, "f64**", $hDer.p, false)}
;

expr_log returns[abstract.Expresion p]
    : hIzq = expr_rel '||' hDer = expr_rel {$p = expresion.NewLogicas($hIzq.p, "||", $hDer.p, false)}
    | hIzq = expr_rel '&&' hDer = expr_rel {$p = expresion.NewLogicas($hIzq.p, "&&", $hDer.p, false)}
    | '!' hIzq = expr_rel    {$p = expresion.NewLogicas($hIzq.p, "!", nil, true)}
    | expr_rel {$p = $expr_rel.p}
;

expr_rel returns[abstract.Expresion p]
    : hIzq = expr_rel op = ('==' | '!=' | '<' | '<=' | '>' | '>=') hDer = expr_rel {$p = expresion.NewRelacional($hIzq.p, $op.text, $hDer.p)}
    | expr_op {$p = $expr_op.p}
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
;