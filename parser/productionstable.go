
package parser

import _rt "../runtime"
import "../token"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Program	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Program : Statements	<<  >>`,
		Id: "Program",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Variable : var	<<  >>`,
		Id: "Variable",
		NTType: 2,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Variable : "input"	<<  >>`,
		Id: "Variable",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< _rt.Bool(true), nil >>`,
		Id: "Bool",
		NTType: 3,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.Bool(true), nil
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< _rt.Bool(false), nil >>`,
		Id: "Bool",
		NTType: 3,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.Bool(false), nil
		},
	},
	ProdTabEntry{
		String: `Callable_Object : Variable	<< _rt.MakeGetVariable(X[0].(*token.Token).Lit), nil >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeGetVariable(X[0].(*token.Token).Lit), nil
		},
	},
	ProdTabEntry{
		String: `Callable_Object : "(" Expression ")"	<< X[1], nil >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Callable_Object : Get_Index	<< X[0], nil >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Callable_Object : Fn_Call	<< X[0], nil >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Callable_Object : Lambda_Call	<< X[0], nil >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : Callable_Object	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : Bool	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : int	<< _rt.NewConstInt(X[0].(*token.Token).Lit) >>`,
		Id: "Object",
		NTType: 5,
		Index: 13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewConstInt(X[0].(*token.Token).Lit)
		},
	},
	ProdTabEntry{
		String: `Object : ListDef	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : Method_Call	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Mult_Expr : Mult_Expr "*" Object	<< _rt.MakeMul(c(X[0]), c(X[2])), nil >>`,
		Id: "Mult_Expr",
		NTType: 6,
		Index: 16,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeMul(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Mult_Expr : Mult_Expr "/" Object	<< _rt.MakeDiv(c(X[0]), c(X[2])), nil >>`,
		Id: "Mult_Expr",
		NTType: 6,
		Index: 17,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeDiv(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Mult_Expr : Object	<< X[0], nil >>`,
		Id: "Mult_Expr",
		NTType: 6,
		Index: 18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Add_Expr : Add_Expr "+" Mult_Expr	<< _rt.MakeAdd(c(X[0]), c(X[2])), nil >>`,
		Id: "Add_Expr",
		NTType: 7,
		Index: 19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeAdd(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Add_Expr : Add_Expr "-" Mult_Expr	<< _rt.MakeSub(c(X[0]), c(X[2])), nil >>`,
		Id: "Add_Expr",
		NTType: 7,
		Index: 20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeSub(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Add_Expr : Mult_Expr	<< X[0], nil >>`,
		Id: "Add_Expr",
		NTType: 7,
		Index: 21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr ">" Add_Expr	<< _rt.MakeGt(c(X[0]), c(X[2])), nil >>`,
		Id: "Comp_Expr",
		NTType: 8,
		Index: 22,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeGt(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr "<" Add_Expr	<< _rt.MakeLt(c(X[0]), c(X[2])), nil >>`,
		Id: "Comp_Expr",
		NTType: 8,
		Index: 23,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeLt(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr "==" Add_Expr	<< _rt.MakeEqual(c(X[0]), c(X[2])), nil >>`,
		Id: "Comp_Expr",
		NTType: 8,
		Index: 24,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeEqual(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr "!=" Add_Expr	<< _rt.MakeNotEqual(c(X[0]), c(X[2])), nil >>`,
		Id: "Comp_Expr",
		NTType: 8,
		Index: 25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeNotEqual(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Add_Expr	<< X[0], nil >>`,
		Id: "Comp_Expr",
		NTType: 8,
		Index: 26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bool_Expr : Bool_Expr "&&" Comp_Expr	<< _rt.MakeAND(c(X[0]), c(X[2])), nil >>`,
		Id: "Bool_Expr",
		NTType: 9,
		Index: 27,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeAND(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Bool_Expr : Bool_Expr "||" Comp_Expr	<< _rt.MakeOR(c(X[0]), c(X[2])), nil >>`,
		Id: "Bool_Expr",
		NTType: 9,
		Index: 28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeOR(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Bool_Expr : Comp_Expr	<< X[0], nil >>`,
		Id: "Bool_Expr",
		NTType: 9,
		Index: 29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Get_Index : Object "[" Expression "]"	<< _rt.MakeGetListItem(c(X[0]), c(X[2])), nil >>`,
		Id: "Get_Index",
		NTType: 10,
		Index: 30,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeGetListItem(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Assign : Variable "=" Expression	<< _rt.MakeSetVariable(X[0].(*token.Token).Lit, c(X[2])), nil >>`,
		Id: "Assign",
		NTType: 11,
		Index: 31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeSetVariable(X[0].(*token.Token).Lit, c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Assign : Object "[" Expression "]" "=" Expression	<< _rt.MakeSetListItem(c(X[0]), c(X[2]), c(X[5])), nil >>`,
		Id: "Assign",
		NTType: 11,
		Index: 32,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeSetListItem(c(X[0]), c(X[2]), c(X[5])), nil
		},
	},
	ProdTabEntry{
		String: `Expression : Bool_Expr	<< X[0], nil >>`,
		Id: "Expression",
		NTType: 12,
		Index: 33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Assign	<< X[0], nil >>`,
		Id: "Expression",
		NTType: 12,
		Index: 34,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Lambda_Def	<< X[0], nil >>`,
		Id: "Expression",
		NTType: 12,
		Index: 35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Values : Values "," Expression	<< _rt.AddToValues(c(X[0]), c(X[2])), nil >>`,
		Id: "Values",
		NTType: 13,
		Index: 36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.AddToValues(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Values : Expression	<< _rt.NewValues(c(X[0])), nil >>`,
		Id: "Values",
		NTType: 13,
		Index: 37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewValues(c(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `ListDef : "[]"	<< _rt.MakeEmptyList(), nil >>`,
		Id: "ListDef",
		NTType: 14,
		Index: 38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeEmptyList(), nil
		},
	},
	ProdTabEntry{
		String: `ListDef : "[" Values "]"	<< _rt.MakeListValues(c(X[1])), nil >>`,
		Id: "ListDef",
		NTType: 14,
		Index: 39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeListValues(c(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Fn_Name : fn_name	<< X[0], nil >>`,
		Id: "Fn_Name",
		NTType: 15,
		Index: 40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Fn_Name : cust_fn_name	<< X[0], nil >>`,
		Id: "Fn_Name",
		NTType: 15,
		Index: 41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Fn_Call : Fn_Name "()"	<< _rt.CallFunction(X[0].(*token.Token).Lit, nil) >>`,
		Id: "Fn_Call",
		NTType: 16,
		Index: 42,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.CallFunction(X[0].(*token.Token).Lit, nil)
		},
	},
	ProdTabEntry{
		String: `Fn_Call : Fn_Name "(" Values ")"	<< _rt.CallFunction(X[0].(*token.Token).Lit, c(X[2])) >>`,
		Id: "Fn_Call",
		NTType: 16,
		Index: 43,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.CallFunction(X[0].(*token.Token).Lit, c(X[2]))
		},
	},
	ProdTabEntry{
		String: `Lambda_Call : Callable_Object "(" Values ")"	<< _rt.MakeCallLambda(c(X[0]), c(X[2])), nil >>`,
		Id: "Lambda_Call",
		NTType: 17,
		Index: 44,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeCallLambda(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Method_Call : Object "." fn_name	<<  >>`,
		Id: "Method_Call",
		NTType: 18,
		Index: 45,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Method_Call : Object "." Fn_Call	<<  >>`,
		Id: "Method_Call",
		NTType: 18,
		Index: 46,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `CodeBlock : "{" Statements "}"	<< X[1], nil >>`,
		Id: "CodeBlock",
		NTType: 19,
		Index: 47,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Func_Param_Def : Func_Param_Def "," Variable	<< _rt.AddToParamDef(X[0].([]string), X[2].(*token.Token).Lit), nil >>`,
		Id: "Func_Param_Def",
		NTType: 20,
		Index: 48,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.AddToParamDef(X[0].([]string), X[2].(*token.Token).Lit), nil
		},
	},
	ProdTabEntry{
		String: `Func_Param_Def : Variable	<< _rt.NewParamDef(X[0].(*token.Token).Lit), nil >>`,
		Id: "Func_Param_Def",
		NTType: 20,
		Index: 49,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewParamDef(X[0].(*token.Token).Lit), nil
		},
	},
	ProdTabEntry{
		String: `Cust_Fn_def : "function" cust_fn_name ":" Func_Param_Def CodeBlock	<< _rt.MakeCustFuncDev(X[1].(*token.Token).Lit, X[3].([]string), c(X[4])) >>`,
		Id: "Cust_Fn_def",
		NTType: 21,
		Index: 50,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeCustFuncDev(X[1].(*token.Token).Lit, X[3].([]string), c(X[4]))
		},
	},
	ProdTabEntry{
		String: `Statement : Expression	<< X[0], nil >>`,
		Id: "Statement",
		NTType: 22,
		Index: 51,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : "return" Expression	<<  >>`,
		Id: "Statement",
		NTType: 22,
		Index: 52,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Single_Statement : Statement ";"	<< X[0], nil >>`,
		Id: "Single_Statement",
		NTType: 23,
		Index: 53,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Single_Statement : Block	<< X[0], nil >>`,
		Id: "Single_Statement",
		NTType: 23,
		Index: 54,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Single_Statement : Cust_Fn_def	<< X[0], nil >>`,
		Id: "Single_Statement",
		NTType: 23,
		Index: 55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : Single_Statement Statements	<< _rt.MakeStatements(c(X[0]), c(X[1])), nil >>`,
		Id: "Statements",
		NTType: 24,
		Index: 56,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeStatements(c(X[0]), c(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Statements : Single_Statement	<< X[0], nil >>`,
		Id: "Statements",
		NTType: 24,
		Index: 57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : Statement	<< X[0], nil >>`,
		Id: "Statements",
		NTType: 24,
		Index: 58,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IfBlock : "if" Expression CodeBlock "else" CodeBlock	<< _rt.MakeIfThenElse(c(X[1]), c(X[2]), c(X[4])), nil >>`,
		Id: "IfBlock",
		NTType: 25,
		Index: 59,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeIfThenElse(c(X[1]), c(X[2]), c(X[4])), nil
		},
	},
	ProdTabEntry{
		String: `IfBlock : "if" Expression CodeBlock	<< _rt.MakeIfThenElse(c(X[1]), c(X[2]), nil), nil >>`,
		Id: "IfBlock",
		NTType: 25,
		Index: 60,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeIfThenElse(c(X[1]), c(X[2]), nil), nil
		},
	},
	ProdTabEntry{
		String: `WhileLoop : "while" Expression CodeBlock	<< _rt.MakeWhileLoop(c(X[1]), c(X[2])), nil >>`,
		Id: "WhileLoop",
		NTType: 26,
		Index: 61,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeWhileLoop(c(X[1]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `ForEachLoop : "foreach" Variable "in" Expression CodeBlock	<< _rt.MakeForEach(X[1].(*token.Token).Lit, c(X[3]), c(X[4])), nil >>`,
		Id: "ForEachLoop",
		NTType: 27,
		Index: 62,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeForEach(X[1].(*token.Token).Lit, c(X[3]), c(X[4])), nil
		},
	},
	ProdTabEntry{
		String: `Block : IfBlock	<<  >>`,
		Id: "Block",
		NTType: 28,
		Index: 63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Block : WhileLoop	<<  >>`,
		Id: "Block",
		NTType: 28,
		Index: 64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Block : ForEachLoop	<<  >>`,
		Id: "Block",
		NTType: 28,
		Index: 65,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Lambda_Def : "function" Func_Param_Def "->" Statement	<< _rt.MakeLambda(X[1].([]string), c(X[3])), nil >>`,
		Id: "Lambda_Def",
		NTType: 29,
		Index: 66,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.MakeLambda(X[1].([]string), c(X[3])), nil
		},
	},
	
}
