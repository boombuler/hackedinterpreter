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
    Sym_A_btn             gold.SymbolId = 25 // 'a_btn'
    Sym_Abs               gold.SymbolId = 26 // abs
    Sym_B_btn             gold.SymbolId = 27 // 'b_btn'
    Sym_Copy              gold.SymbolId = 28 // copy
    Sym_Down              gold.SymbolId = 29 // down
    Sym_Draw              gold.SymbolId = 30 // draw
    Sym_Draw_text         gold.SymbolId = 31 // 'draw_text'
    Sym_Else              gold.SymbolId = 32 // else
    Sym_F1                gold.SymbolId = 33 // 'f1'
    Sym_F10               gold.SymbolId = 34 // 'f10'
    Sym_F11               gold.SymbolId = 35 // 'f11'
    Sym_F12               gold.SymbolId = 36 // 'f12'
    Sym_F13               gold.SymbolId = 37 // 'f13'
    Sym_F14               gold.SymbolId = 38 // 'f14'
    Sym_F15               gold.SymbolId = 39 // 'f15'
    Sym_F16               gold.SymbolId = 40 // 'f16'
    Sym_F2                gold.SymbolId = 41 // 'f2'
    Sym_F3                gold.SymbolId = 42 // 'f3'
    Sym_F4                gold.SymbolId = 43 // 'f4'
    Sym_F5                gold.SymbolId = 44 // 'f5'
    Sym_F6                gold.SymbolId = 45 // 'f6'
    Sym_F7                gold.SymbolId = 46 // 'f7'
    Sym_F8                gold.SymbolId = 47 // 'f8'
    Sym_F9                gold.SymbolId = 48 // 'f9'
    Sym_False             gold.SymbolId = 49 // false
    Sym_Fill              gold.SymbolId = 50 // fill
    Sym_Foreach           gold.SymbolId = 51 // foreach
    Sym_Function          gold.SymbolId = 52 // function
    Sym_Height            gold.SymbolId = 53 // height
    Sym_If                gold.SymbolId = 54 // if
    Sym_In                gold.SymbolId = 55 // in
    Sym_Input             gold.SymbolId = 56 // Input
    Sym_Insert            gold.SymbolId = 57 // insert
    Sym_Is_list           gold.SymbolId = 58 // 'is_list'
    Sym_Left              gold.SymbolId = 59 // left
    Sym_Length            gold.SymbolId = 60 // length
    Sym_Map               gold.SymbolId = 61 // map
    Sym_Max               gold.SymbolId = 62 // max
    Sym_Min               gold.SymbolId = 63 // min
    Sym_Mod               gold.SymbolId = 64 // mod
    Sym_New_list          gold.SymbolId = 65 // 'new_list'
    Sym_Number            gold.SymbolId = 66 // Number
    Sym_Pop               gold.SymbolId = 67 // pop
    Sym_Pow               gold.SymbolId = 68 // pow
    Sym_Push              gold.SymbolId = 69 // push
    Sym_Random            gold.SymbolId = 70 // random
    Sym_Remove            gold.SymbolId = 71 // remove
    Sym_Return            gold.SymbolId = 72 // return
    Sym_Right             gold.SymbolId = 73 // right
    Sym_Sort              gold.SymbolId = 74 // sort
    Sym_Sort_with         gold.SymbolId = 75 // 'sort_with'
    Sym_Time              gold.SymbolId = 76 // time
    Sym_True              gold.SymbolId = 77 // true
    Sym_Up                gold.SymbolId = 78 // up
    Sym_Variable          gold.SymbolId = 79 // Variable
    Sym_While             gold.SymbolId = 80 // while
    Sym_Width             gold.SymbolId = 81 // width
    Sym_Addexp            gold.SymbolId = 82 // <Add Exp>
    Sym_Array             gold.SymbolId = 83 // <Array>
    Sym_Arrayvalues       gold.SymbolId = 84 // <ArrayValues>
    Sym_Assign            gold.SymbolId = 85 // <Assign>
    Sym_Block             gold.SymbolId = 86 // <Block>
    Sym_Bool              gold.SymbolId = 87 // <Bool>
    Sym_Boolalgexpression gold.SymbolId = 88 // <BoolAlgExpression>
    Sym_Codeblock         gold.SymbolId = 89 // <CodeBlock>
    Sym_Compexpression    gold.SymbolId = 90 // <Comp Expression>
    Sym_Custfncall        gold.SymbolId = 91 // <Cust Fn Call>
    Sym_Custfndef         gold.SymbolId = 92 // <Cust Fn Def>
    Sym_Custfnname        gold.SymbolId = 93 // <Cust Fn Name>
    Sym_Elseblock         gold.SymbolId = 94 // <ElseBlock>
    Sym_Expression        gold.SymbolId = 95 // <Expression>
    Sym_Fnparameters      gold.SymbolId = 96 // <Fn Parameters>
    Sym_Fnparametersdef   gold.SymbolId = 97 // <Fn Parameters Def>
    Sym_Foreachloop       gold.SymbolId = 98 // <ForEachLoop>
    Sym_Functioncall      gold.SymbolId = 99 // <Function Call>
    Sym_Functionname      gold.SymbolId = 100 // <Function Name>
    Sym_Getindex          gold.SymbolId = 101 // <GetIndex>
    Sym_Ifblock           gold.SymbolId = 102 // <IfBlock>
    Sym_Lambdafncall      gold.SymbolId = 103 // <Lambda Fn Call>
    Sym_Lambdafndef       gold.SymbolId = 104 // <Lambda Fn Def>
    Sym_Lambdatarget      gold.SymbolId = 105 // <Lambda Target>
    Sym_Methodcall        gold.SymbolId = 106 // <Method Call>
    Sym_Methodname        gold.SymbolId = 107 // <Method Name>
    Sym_Multexp           gold.SymbolId = 108 // <Mult Exp>
    Sym_Program           gold.SymbolId = 109 // <Program>
    Sym_Returnstmnt       gold.SymbolId = 110 // <ReturnStmnt>
    Sym_Singlestatement   gold.SymbolId = 111 // <Single Statement>
    Sym_Statement         gold.SymbolId = 112 // <Statement>
    Sym_Statements        gold.SymbolId = 113 // <Statements>
    Sym_Value             gold.SymbolId = 114 // <Value>
    Sym_Variable2         gold.SymbolId = 115 // <Variable>
    Sym_Whileloop         gold.SymbolId = 116 // <WhileLoop>
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
    Rule_FunctionnameTime             gold.RuleId = 78 // <Function Name> ::= time
    Rule_FunctionnameRandom           gold.RuleId = 79 // <Function Name> ::= random
    Rule_FunctionnameB_btn            gold.RuleId = 80 // <Function Name> ::= 'b_btn'
    Rule_FunctionnameA_btn            gold.RuleId = 81 // <Function Name> ::= 'a_btn'
    Rule_FunctionnameDown             gold.RuleId = 82 // <Function Name> ::= down
    Rule_FunctionnameUp               gold.RuleId = 83 // <Function Name> ::= up
    Rule_FunctionnameRight            gold.RuleId = 84 // <Function Name> ::= right
    Rule_FunctionnameLeft             gold.RuleId = 85 // <Function Name> ::= left
    Rule_FunctionnameHeight           gold.RuleId = 86 // <Function Name> ::= height
    Rule_FunctionnameWidth            gold.RuleId = 87 // <Function Name> ::= width
    Rule_FunctionnameDraw_text        gold.RuleId = 88 // <Function Name> ::= 'draw_text'
    Rule_FunctionnameDraw             gold.RuleId = 89 // <Function Name> ::= draw
    Rule_FunctionnameNew_list         gold.RuleId = 90 // <Function Name> ::= 'new_list'
    Rule_FunctionnamePow              gold.RuleId = 91 // <Function Name> ::= pow
    Rule_FunctionnameMax              gold.RuleId = 92 // <Function Name> ::= max
    Rule_FunctionnameMin              gold.RuleId = 93 // <Function Name> ::= min
    Rule_FunctionnameMod              gold.RuleId = 94 // <Function Name> ::= mod
    Rule_FunctionnameAbs              gold.RuleId = 95 // <Function Name> ::= abs
    Rule_FunctioncallLparenRparen     gold.RuleId = 96 // <Function Call> ::= <Function Name> '(' <ArrayValues> ')'
    Rule_MethodcallDotLparenRparen    gold.RuleId = 97 // <Method Call> ::= <Value> '.' <Method Name> '(' <ArrayValues> ')'
    Rule_MethodcallDot                gold.RuleId = 98 // <Method Call> ::= <Value> '.' <Method Name>
    Rule_LambdafndefFunctionMinusgt   gold.RuleId = 99 // <Lambda Fn Def> ::= function <Fn Parameters Def> '->' <Statement>
    Rule_LambdafncallLparenRparen     gold.RuleId = 100 // <Lambda Fn Call> ::= <Lambda Target> '(' <Fn Parameters> ')'
    Rule_Lambdatarget                 gold.RuleId = 101 // <Lambda Target> ::= <Variable>
    Rule_Lambdatarget2                gold.RuleId = 102 // <Lambda Target> ::= <GetIndex>
    Rule_Lambdatarget3                gold.RuleId = 103 // <Lambda Target> ::= <Method Call>
    Rule_Lambdatarget4                gold.RuleId = 104 // <Lambda Target> ::= <Cust Fn Call>
    Rule_Lambdatarget5                gold.RuleId = 105 // <Lambda Target> ::= <Lambda Fn Call>
    Rule_LambdatargetLparenRparen     gold.RuleId = 106 // <Lambda Target> ::= '(' <Expression> ')'
    Rule_ValueNumber                  gold.RuleId = 107 // <Value> ::= Number
    Rule_Value                        gold.RuleId = 108 // <Value> ::= <Lambda Target>
    Rule_Value2                       gold.RuleId = 109 // <Value> ::= <Function Call>
    Rule_Value3                       gold.RuleId = 110 // <Value> ::= <Bool>
    Rule_Value4                       gold.RuleId = 111 // <Value> ::= <Array>
    Rule_Value5                       gold.RuleId = 112 // <Value> ::= <Cust Fn Name>
    Rule_Value6                       gold.RuleId = 113 // <Value> ::= <Lambda Fn Def>
)
