grammar CalcPlus;

program :   (funcDef)+ EOF;

funcDef :   'func' 'int'? IDENT '(' paramList? ')' block ;

stmt    :   'int' IDENT (',' IDENT)* ';'        # Declare
        |   IDENT '=' expr ';'                  # ExprAssign
        |   expr ';'                            # ExprStmt
//      |   IDENT '=' 'read' '(' ')' ';'        # ReadAssign
//      |   'write' '(' expr ')' ';'            # Write
        |   'if' '(' cond ')' thenBlock=block
            ('else' elseBlock=block)?           # IfElse
        |   block                               # StmtBlock
        |   'return' expr? ';'                  # Return
        ;

expr    :   expr ('*'|'/') expr     # MulDiv
        |   expr ('+'|'-') expr     # AddSub
        |   INT                     # Int
        |   IDENT                   # Var
        |   '(' expr ')'            # Parens
        |   IDENT '(' argList? ')'  # FuncCall
        ;

paramList   : 'int' IDENT (',' 'int' IDENT)* ;
argList     : expr (',' expr)* ;

cond    :   expr ('=='|'!='|'>'|'>='|'<'|'<=') expr ;
block   :   '{' (stmt)* '}' ;

WS  : [ \t\r\n]+ -> skip;
INT : [0-9]+ ;
IDENT : [A-Za-z][A-Za-z0-9]*;