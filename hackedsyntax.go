package main

import "github.com/boombuler/gold"

const (
    Sym_Eof               gold.SymbolId = 0 // (EOF)
    Sym_Error             gold.SymbolId = 1 // (Error)
    Sym_Whitespace        gold.SymbolId = 2 // Whitespace
    Sym_Minus             gold.SymbolId = 3 // '-'
    Sym_Exclameq          gold.SymbolId = 4 // '!='
    Sym_Ampamp            gold.SymbolId = 5 // '&&'
    Sym_Lparen            gold.SymbolId = 6 // '('
    Sym_Rparen            gold.SymbolId = 7 // ')'
    Sym_Times             gold.SymbolId = 8 // '*'
    Sym_Comma             gold.SymbolId = 9 // ','
    Sym_Dot               gold.SymbolId = 10 // '.'
    Sym_Div               gold.SymbolId = 11 // '/'
    Sym_Colon             gold.SymbolId = 12 // ':'
    Sym_Semi              gold.SymbolId = 13 // ';'
    Sym_Lbracket          gold.SymbolId = 14 // '['
    Sym_Rbracket          gold.SymbolId = 15 // ']'
    Sym_Lbrace            gold.SymbolId = 16 // '{'
    Sym_Pipepipe          gold.SymbolId = 17 // '||'
    Sym_Rbrace            gold.SymbolId = 18 // '}'
    Sym_Plus              gold.SymbolId = 19 // '+'
    Sym_Lt                gold.SymbolId = 20 // '<'
    Sym_Eq                gold.SymbolId = 21 // '='
    Sym_Eqeq              gold.SymbolId = 22 // '=='
    Sym_Gt                gold.SymbolId = 23 // '>'
    Sym_Minusgt           gold.SymbolId = 24 // '->'
    Sym_Copy              gold.SymbolId = 25 // copy
    Sym_Else              gold.SymbolId = 26 // else
    Sym_F1                gold.SymbolId = 27 // 'f1'
    Sym_F10               gold.SymbolId = 28 // 'f10'
    Sym_F11               gold.SymbolId = 29 // 'f11'
    Sym_F12               gold.SymbolId = 30 // 'f12'
    Sym_F13               gold.SymbolId = 31 // 'f13'
    Sym_F14               gold.SymbolId = 32 // 'f14'
    Sym_F15               gold.SymbolId = 33 // 'f15'
    Sym_F16               gold.SymbolId = 34 // 'f16'
    Sym_F2                gold.SymbolId = 35 // 'f2'
    Sym_F3                gold.SymbolId = 36 // 'f3'
    Sym_F4                gold.SymbolId = 37 // 'f4'
    Sym_F5                gold.SymbolId = 38 // 'f5'
    Sym_F6                gold.SymbolId = 39 // 'f6'
    Sym_F7                gold.SymbolId = 40 // 'f7'
    Sym_F8                gold.SymbolId = 41 // 'f8'
    Sym_F9                gold.SymbolId = 42 // 'f9'
    Sym_False             gold.SymbolId = 43 // false
    Sym_Fill              gold.SymbolId = 44 // fill
    Sym_Foreach           gold.SymbolId = 45 // foreach
    Sym_Function          gold.SymbolId = 46 // function
    Sym_If                gold.SymbolId = 47 // if
    Sym_In                gold.SymbolId = 48 // in
    Sym_Input             gold.SymbolId = 49 // Input
    Sym_Insert            gold.SymbolId = 50 // insert
    Sym_Is_list           gold.SymbolId = 51 // 'is_list'
    Sym_Length            gold.SymbolId = 52 // length
    Sym_Map               gold.SymbolId = 53 // map
    Sym_Number            gold.SymbolId = 54 // Number
    Sym_Pop               gold.SymbolId = 55 // pop
    Sym_Push              gold.SymbolId = 56 // push
    Sym_Remove            gold.SymbolId = 57 // remove
    Sym_Return            gold.SymbolId = 58 // return
    Sym_Sort              gold.SymbolId = 59 // sort
    Sym_Sort_with         gold.SymbolId = 60 // 'sort_with'
    Sym_True              gold.SymbolId = 61 // true
    Sym_Variable          gold.SymbolId = 62 // Variable
    Sym_While             gold.SymbolId = 63 // while
    Sym_Addexp            gold.SymbolId = 64 // <Add Exp>
    Sym_Array             gold.SymbolId = 65 // <Array>
    Sym_Arrayvalues       gold.SymbolId = 66 // <ArrayValues>
    Sym_Assign            gold.SymbolId = 67 // <Assign>
    Sym_Block             gold.SymbolId = 68 // <Block>
    Sym_Bool              gold.SymbolId = 69 // <Bool>
    Sym_Boolalgexpression gold.SymbolId = 70 // <BoolAlgExpression>
    Sym_Codeblock         gold.SymbolId = 71 // <CodeBlock>
    Sym_Compexpression    gold.SymbolId = 72 // <Comp Expression>
    Sym_Custfncall        gold.SymbolId = 73 // <Cust Fn Call>
    Sym_Custfndef         gold.SymbolId = 74 // <Cust Fn Def>
    Sym_Custfnname        gold.SymbolId = 75 // <Cust Fn Name>
    Sym_Elseblock         gold.SymbolId = 76 // <ElseBlock>
    Sym_Expression        gold.SymbolId = 77 // <Expression>
    Sym_Fnparameters      gold.SymbolId = 78 // <Fn Parameters>
    Sym_Fnparametersdef   gold.SymbolId = 79 // <Fn Parameters Def>
    Sym_Foreachloop       gold.SymbolId = 80 // <ForEachLoop>
    Sym_Getindex          gold.SymbolId = 81 // <GetIndex>
    Sym_Ifblock           gold.SymbolId = 82 // <IfBlock>
    Sym_Lambdafncall      gold.SymbolId = 83 // <Lambda Fn Call>
    Sym_Lambdafndef       gold.SymbolId = 84 // <Lambda Fn Def>
    Sym_Lambdatarget      gold.SymbolId = 85 // <Lambda Target>
    Sym_Methodcall        gold.SymbolId = 86 // <Method Call>
    Sym_Methodname        gold.SymbolId = 87 // <Method Name>
    Sym_Multexp           gold.SymbolId = 88 // <Mult Exp>
    Sym_Program           gold.SymbolId = 89 // <Program>
    Sym_Returnstmnt       gold.SymbolId = 90 // <ReturnStmnt>
    Sym_Singlestatement   gold.SymbolId = 91 // <Single Statement>
    Sym_Statement         gold.SymbolId = 92 // <Statement>
    Sym_Statements        gold.SymbolId = 93 // <Statements>
    Sym_Value             gold.SymbolId = 94 // <Value>
    Sym_Variable2         gold.SymbolId = 95 // <Variable>
    Sym_Whileloop         gold.SymbolId = 96 // <WhileLoop>
)

const (
    Rule_Program                      gold.RuleId = 0 // <Program> ::= <Statements>
    Rule_CodeblockLbraceRbrace        gold.RuleId = 1 // <CodeBlock> ::= '{' <Statements> '}'
    Rule_IfblockIf                    gold.RuleId = 2 // <IfBlock> ::= if <Expression> <CodeBlock> <ElseBlock>
    Rule_ElseblockElse                gold.RuleId = 3 // <ElseBlock> ::= else <CodeBlock>
    Rule_Elseblock                    gold.RuleId = 4 // <ElseBlock> ::= 
    Rule_WhileloopWhile               gold.RuleId = 5 // <WhileLoop> ::= while <Expression> <CodeBlock>
    Rule_ForeachloopForeachVariableIn gold.RuleId = 6 // <ForEachLoop> ::= foreach Variable in <Expression> <CodeBlock>
    Rule_Block                        gold.RuleId = 7 // <Block> ::= <IfBlock>
    Rule_Block2                       gold.RuleId = 8 // <Block> ::= <WhileLoop>
    Rule_Block3                       gold.RuleId = 9 // <Block> ::= <ForEachLoop>
    Rule_SinglestatementSemi          gold.RuleId = 10 // <Single Statement> ::= <Statement> ';'
    Rule_Singlestatement              gold.RuleId = 11 // <Single Statement> ::= <Block>
    Rule_Singlestatement2             gold.RuleId = 12 // <Single Statement> ::= <Cust Fn Def>
    Rule_Statements                   gold.RuleId = 13 // <Statements> ::= <Single Statement> <Statements>
    Rule_Statements2                  gold.RuleId = 14 // <Statements> ::= <Single Statement>
    Rule_Statements3                  gold.RuleId = 15 // <Statements> ::= <Statement>
    Rule_Statement                    gold.RuleId = 16 // <Statement> ::= <Expression>
    Rule_Statement2                   gold.RuleId = 17 // <Statement> ::= <ReturnStmnt>
    Rule_ReturnstmntReturn            gold.RuleId = 18 // <ReturnStmnt> ::= return <Expression>
    Rule_VariableVariable             gold.RuleId = 19 // <Variable> ::= Variable
    Rule_VariableInput                gold.RuleId = 20 // <Variable> ::= Input
    Rule_AssignEq                     gold.RuleId = 21 // <Assign> ::= <Variable> '=' <Expression>
    Rule_Expression                   gold.RuleId = 22 // <Expression> ::= <BoolAlgExpression>
    Rule_Expression2                  gold.RuleId = 23 // <Expression> ::= <Assign>
    Rule_BoolalgexpressionAmpamp      gold.RuleId = 24 // <BoolAlgExpression> ::= <BoolAlgExpression> '&&' <Comp Expression>
    Rule_BoolalgexpressionPipepipe    gold.RuleId = 25 // <BoolAlgExpression> ::= <BoolAlgExpression> '||' <Comp Expression>
    Rule_Boolalgexpression            gold.RuleId = 26 // <BoolAlgExpression> ::= <Comp Expression>
    Rule_CompexpressionGt             gold.RuleId = 27 // <Comp Expression> ::= <Comp Expression> '>' <Add Exp>
    Rule_CompexpressionLt             gold.RuleId = 28 // <Comp Expression> ::= <Comp Expression> '<' <Add Exp>
    Rule_CompexpressionEqeq           gold.RuleId = 29 // <Comp Expression> ::= <Comp Expression> '==' <Add Exp>
    Rule_CompexpressionExclameq       gold.RuleId = 30 // <Comp Expression> ::= <Comp Expression> '!=' <Add Exp>
    Rule_Compexpression               gold.RuleId = 31 // <Comp Expression> ::= <Add Exp>
    Rule_AddexpPlus                   gold.RuleId = 32 // <Add Exp> ::= <Add Exp> '+' <Mult Exp>
    Rule_AddexpMinus                  gold.RuleId = 33 // <Add Exp> ::= <Add Exp> '-' <Mult Exp>
    Rule_Addexp                       gold.RuleId = 34 // <Add Exp> ::= <Mult Exp>
    Rule_MultexpTimes                 gold.RuleId = 35 // <Mult Exp> ::= <Mult Exp> '*' <Value>
    Rule_MultexpDiv                   gold.RuleId = 36 // <Mult Exp> ::= <Mult Exp> '/' <Value>
    Rule_Multexp                      gold.RuleId = 37 // <Mult Exp> ::= <Value>
    Rule_BoolTrue                     gold.RuleId = 38 // <Bool> ::= true
    Rule_BoolFalse                    gold.RuleId = 39 // <Bool> ::= false
    Rule_GetindexLbracketRbracket     gold.RuleId = 40 // <GetIndex> ::= <Value> '[' <Expression> ']'
    Rule_ArrayvaluesComma             gold.RuleId = 41 // <ArrayValues> ::= <ArrayValues> ',' <Expression>
    Rule_Arrayvalues                  gold.RuleId = 42 // <ArrayValues> ::= <Expression>
    Rule_Arrayvalues2                 gold.RuleId = 43 // <ArrayValues> ::= 
    Rule_ArrayLbracketRbracket        gold.RuleId = 44 // <Array> ::= '[' <ArrayValues> ']'
    Rule_CustfnnameF1                 gold.RuleId = 45 // <Cust Fn Name> ::= 'f1'
    Rule_CustfnnameF2                 gold.RuleId = 46 // <Cust Fn Name> ::= 'f2'
    Rule_CustfnnameF3                 gold.RuleId = 47 // <Cust Fn Name> ::= 'f3'
    Rule_CustfnnameF4                 gold.RuleId = 48 // <Cust Fn Name> ::= 'f4'
    Rule_CustfnnameF5                 gold.RuleId = 49 // <Cust Fn Name> ::= 'f5'
    Rule_CustfnnameF6                 gold.RuleId = 50 // <Cust Fn Name> ::= 'f6'
    Rule_CustfnnameF7                 gold.RuleId = 51 // <Cust Fn Name> ::= 'f7'
    Rule_CustfnnameF8                 gold.RuleId = 52 // <Cust Fn Name> ::= 'f8'
    Rule_CustfnnameF9                 gold.RuleId = 53 // <Cust Fn Name> ::= 'f9'
    Rule_CustfnnameF10                gold.RuleId = 54 // <Cust Fn Name> ::= 'f10'
    Rule_CustfnnameF11                gold.RuleId = 55 // <Cust Fn Name> ::= 'f11'
    Rule_CustfnnameF12                gold.RuleId = 56 // <Cust Fn Name> ::= 'f12'
    Rule_CustfnnameF13                gold.RuleId = 57 // <Cust Fn Name> ::= 'f13'
    Rule_CustfnnameF14                gold.RuleId = 58 // <Cust Fn Name> ::= 'f14'
    Rule_CustfnnameF15                gold.RuleId = 59 // <Cust Fn Name> ::= 'f15'
    Rule_CustfnnameF16                gold.RuleId = 60 // <Cust Fn Name> ::= 'f16'
    Rule_FnparametersdefComma         gold.RuleId = 61 // <Fn Parameters Def> ::= <Fn Parameters Def> ',' <Variable>
    Rule_Fnparametersdef              gold.RuleId = 62 // <Fn Parameters Def> ::= <Variable>
    Rule_CustfndefFunctionColon       gold.RuleId = 63 // <Cust Fn Def> ::= function <Cust Fn Name> ':' <Fn Parameters Def> <CodeBlock>
    Rule_FnparametersComma            gold.RuleId = 64 // <Fn Parameters> ::= <Fn Parameters> ',' <Expression>
    Rule_Fnparameters                 gold.RuleId = 65 // <Fn Parameters> ::= <Expression>
    Rule_CustfncallLparenRparen       gold.RuleId = 66 // <Cust Fn Call> ::= <Cust Fn Name> '(' <Fn Parameters> ')'
    Rule_MethodnameIs_list            gold.RuleId = 67 // <Method Name> ::= 'is_list'
    Rule_MethodnameLength             gold.RuleId = 68 // <Method Name> ::= length
    Rule_MethodnamePush               gold.RuleId = 69 // <Method Name> ::= push
    Rule_MethodnamePop                gold.RuleId = 70 // <Method Name> ::= pop
    Rule_MethodnameInsert             gold.RuleId = 71 // <Method Name> ::= insert
    Rule_MethodnameRemove             gold.RuleId = 72 // <Method Name> ::= remove
    Rule_MethodnameSort               gold.RuleId = 73 // <Method Name> ::= sort
    Rule_MethodnameSort_with          gold.RuleId = 74 // <Method Name> ::= 'sort_with'
    Rule_MethodnameMap                gold.RuleId = 75 // <Method Name> ::= map
    Rule_MethodnameCopy               gold.RuleId = 76 // <Method Name> ::= copy
    Rule_MethodnameFill               gold.RuleId = 77 // <Method Name> ::= fill
    Rule_MethodcallDotLparenRparen    gold.RuleId = 78 // <Method Call> ::= <Value> '.' <Method Name> '(' <ArrayValues> ')'
    Rule_MethodcallDot                gold.RuleId = 79 // <Method Call> ::= <Value> '.' <Method Name>
    Rule_LambdafndefFunctionMinusgt   gold.RuleId = 80 // <Lambda Fn Def> ::= function <Fn Parameters Def> '->' <Statement>
    Rule_LambdafncallLparenRparen     gold.RuleId = 81 // <Lambda Fn Call> ::= <Lambda Target> '(' <Fn Parameters> ')'
    Rule_Lambdatarget                 gold.RuleId = 82 // <Lambda Target> ::= <Variable>
    Rule_Lambdatarget2                gold.RuleId = 83 // <Lambda Target> ::= <GetIndex>
    Rule_Lambdatarget3                gold.RuleId = 84 // <Lambda Target> ::= <Method Call>
    Rule_Lambdatarget4                gold.RuleId = 85 // <Lambda Target> ::= <Cust Fn Call>
    Rule_Lambdatarget5                gold.RuleId = 86 // <Lambda Target> ::= <Lambda Fn Call>
    Rule_LambdatargetLparenRparen     gold.RuleId = 87 // <Lambda Target> ::= '(' <Expression> ')'
    Rule_ValueNumber                  gold.RuleId = 88 // <Value> ::= Number
    Rule_Value                        gold.RuleId = 89 // <Value> ::= <Lambda Target>
    Rule_Value2                       gold.RuleId = 90 // <Value> ::= <Bool>
    Rule_Value3                       gold.RuleId = 91 // <Value> ::= <Array>
    Rule_Value4                       gold.RuleId = 92 // <Value> ::= <Cust Fn Name>
    Rule_Value5                       gold.RuleId = 93 // <Value> ::= <Lambda Fn Def>
)
