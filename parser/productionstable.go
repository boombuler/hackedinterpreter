
package parser

import _rt "../runtime"

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
		String: `Program : error	<<  >>`,
		Id: "Program",
		NTType: 1,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Variable : var	<<  >>`,
		Id: "Variable",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Variable : "input"	<<  >>`,
		Id: "Variable",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bool : "true"	<< _rt.NewConstBool(true, p(X[0])) >>`,
		Id: "Bool",
		NTType: 3,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewConstBool(true, p(X[0]))
		},
	},
	ProdTabEntry{
		String: `Bool : "false"	<< _rt.NewConstBool(false, p(X[0])) >>`,
		Id: "Bool",
		NTType: 3,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewConstBool(false, p(X[0]))
		},
	},
	ProdTabEntry{
		String: `Callable_Object : Assignable	<< X[0], nil >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Callable_Object : "(" Expression ")"	<< X[1], nil >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
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
		String: `Callable_Object : cust_fn_name	<< _rt.NewGetCustomFunction(str(X[0]), p(X[0])) >>`,
		Id: "Callable_Object",
		NTType: 4,
		Index: 11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewGetCustomFunction(str(X[0]), p(X[0]))
		},
	},
	ProdTabEntry{
		String: `Object : Callable_Object	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 12,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : Bool	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 13,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : int	<< _rt.NewConstInt(str(X[0]), p(X[0])) >>`,
		Id: "Object",
		NTType: 5,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewConstInt(str(X[0]), p(X[0]))
		},
	},
	ProdTabEntry{
		String: `Object : ListDef	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Object : Method_Call	<< X[0], nil >>`,
		Id: "Object",
		NTType: 5,
		Index: 16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Assignable : Variable	<< _rt.NewVariable(str(X[0]), p(X[0])), nil >>`,
		Id: "Assignable",
		NTType: 6,
		Index: 17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewVariable(str(X[0]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Assignable : Object "[" Expression "]"	<< _rt.NewListItem(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Assignable",
		NTType: 6,
		Index: 18,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewListItem(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Post_Inc_Expr : Assignable "++"	<< _rt.NewIncAssignable(assignable(X[0]), p(X[1])), nil >>`,
		Id: "Post_Inc_Expr",
		NTType: 7,
		Index: 19,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewIncAssignable(assignable(X[0]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Post_Inc_Expr : Assignable "--"	<< _rt.NewDecAssignable(assignable(X[0]), p(X[1])), nil >>`,
		Id: "Post_Inc_Expr",
		NTType: 7,
		Index: 20,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewDecAssignable(assignable(X[0]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Unary_Expr : Object	<< X[0], nil >>`,
		Id: "Unary_Expr",
		NTType: 8,
		Index: 21,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Unary_Expr : Post_Inc_Expr	<< X[0], nil >>`,
		Id: "Unary_Expr",
		NTType: 8,
		Index: 22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Unary_Expr : "-" Object	<< _rt.NewNegInt(c(X[1]), p(X[0])), nil >>`,
		Id: "Unary_Expr",
		NTType: 8,
		Index: 23,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewNegInt(c(X[1]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Unary_Expr : "!" Object	<< _rt.NewNegBool(c(X[1]), p(X[0])), nil >>`,
		Id: "Unary_Expr",
		NTType: 8,
		Index: 24,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewNegBool(c(X[1]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Mult_Expr : Mult_Expr "*" Unary_Expr	<< _rt.NewMul(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Mult_Expr",
		NTType: 9,
		Index: 25,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewMul(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Mult_Expr : Mult_Expr "/" Unary_Expr	<< _rt.NewDiv(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Mult_Expr",
		NTType: 9,
		Index: 26,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewDiv(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Mult_Expr : Unary_Expr	<< X[0], nil >>`,
		Id: "Mult_Expr",
		NTType: 9,
		Index: 27,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Add_Expr : Add_Expr "+" Mult_Expr	<< _rt.NewAdd(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Add_Expr",
		NTType: 10,
		Index: 28,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewAdd(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Add_Expr : Add_Expr "-" Mult_Expr	<< _rt.NewSub(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Add_Expr",
		NTType: 10,
		Index: 29,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewSub(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Add_Expr : Mult_Expr	<< X[0], nil >>`,
		Id: "Add_Expr",
		NTType: 10,
		Index: 30,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr ">" Add_Expr	<< _rt.NewGt(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Comp_Expr",
		NTType: 11,
		Index: 31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewGt(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr "<" Add_Expr	<< _rt.NewLt(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Comp_Expr",
		NTType: 11,
		Index: 32,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewLt(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr "==" Add_Expr	<< _rt.NewEqual(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Comp_Expr",
		NTType: 11,
		Index: 33,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewEqual(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Comp_Expr "!=" Add_Expr	<< _rt.NewNotEqual(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Comp_Expr",
		NTType: 11,
		Index: 34,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewNotEqual(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Comp_Expr : Add_Expr	<< X[0], nil >>`,
		Id: "Comp_Expr",
		NTType: 11,
		Index: 35,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bool_Expr : Bool_Expr "&&" Comp_Expr	<< _rt.NewAND(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Bool_Expr",
		NTType: 12,
		Index: 36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewAND(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Bool_Expr : Bool_Expr "||" Comp_Expr	<< _rt.NewOR(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Bool_Expr",
		NTType: 12,
		Index: 37,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewOR(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Bool_Expr : Comp_Expr	<< X[0], nil >>`,
		Id: "Bool_Expr",
		NTType: 12,
		Index: 38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Assign : Assignable "=" Expression	<< _rt.NewAssign(assignable(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Assign",
		NTType: 13,
		Index: 39,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewAssign(assignable(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Expression : Bool_Expr	<< X[0], nil >>`,
		Id: "Expression",
		NTType: 14,
		Index: 40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Assign	<< X[0], nil >>`,
		Id: "Expression",
		NTType: 14,
		Index: 41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expression : Lambda_Def	<< X[0], nil >>`,
		Id: "Expression",
		NTType: 14,
		Index: 42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Values : Values "," Expression	<< _rt.NewAddToValues(c(X[0]), c(X[2])), nil >>`,
		Id: "Values",
		NTType: 15,
		Index: 43,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewAddToValues(c(X[0]), c(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Values : Expression	<< _rt.NewValues(c(X[0])), nil >>`,
		Id: "Values",
		NTType: 15,
		Index: 44,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewValues(c(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `ListDef : "[" "]"	<< _rt.NewEmptyList(p(X[0])), nil >>`,
		Id: "ListDef",
		NTType: 16,
		Index: 45,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewEmptyList(p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `ListDef : "[" Values "]"	<< _rt.NewListValues(c(X[1]), p(X[0])), nil >>`,
		Id: "ListDef",
		NTType: 16,
		Index: 46,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewListValues(c(X[1]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Fn_Call : fn_name "(" ")"	<< _rt.NewCallFunction(str(X[0]), nil, p(X[1])) >>`,
		Id: "Fn_Call",
		NTType: 17,
		Index: 47,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewCallFunction(str(X[0]), nil, p(X[1]))
		},
	},
	ProdTabEntry{
		String: `Fn_Call : fn_name "(" Values ")"	<< _rt.NewCallFunction(str(X[0]), c(X[2]), p(X[1])) >>`,
		Id: "Fn_Call",
		NTType: 17,
		Index: 48,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewCallFunction(str(X[0]), c(X[2]), p(X[1]))
		},
	},
	ProdTabEntry{
		String: `Lambda_Call : Callable_Object "(" Values ")"	<< _rt.NewCallLambda(c(X[0]), c(X[2]), p(X[1])), nil >>`,
		Id: "Lambda_Call",
		NTType: 18,
		Index: 49,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewCallLambda(c(X[0]), c(X[2]), p(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Method_Call : Object "." fn_name	<< _rt.NewMethodCall(c(X[0]), str(X[2]), nil, p(X[2])) >>`,
		Id: "Method_Call",
		NTType: 19,
		Index: 50,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewMethodCall(c(X[0]), str(X[2]), nil, p(X[2]))
		},
	},
	ProdTabEntry{
		String: `Method_Call : Object "." fn_name "(" ")"	<< _rt.NewMethodCall(c(X[0]), str(X[2]), nil, p(X[2])) >>`,
		Id: "Method_Call",
		NTType: 19,
		Index: 51,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewMethodCall(c(X[0]), str(X[2]), nil, p(X[2]))
		},
	},
	ProdTabEntry{
		String: `Method_Call : Object "." fn_name "(" Values ")"	<< _rt.NewMethodCall(c(X[0]), str(X[2]), c(X[4]), p(X[2])) >>`,
		Id: "Method_Call",
		NTType: 19,
		Index: 52,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewMethodCall(c(X[0]), str(X[2]), c(X[4]), p(X[2]))
		},
	},
	ProdTabEntry{
		String: `Code_Block : "{" Statements "}"	<< X[1], nil >>`,
		Id: "Code_Block",
		NTType: 20,
		Index: 53,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Func_Param_Def : Func_Param_Def "," Variable	<< _rt.NewAddToParamDef(X[0].([]string), str(X[2])), nil >>`,
		Id: "Func_Param_Def",
		NTType: 21,
		Index: 54,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewAddToParamDef(X[0].([]string), str(X[2])), nil
		},
	},
	ProdTabEntry{
		String: `Func_Param_Def : Variable	<< _rt.NewParamDef(str(X[0])), nil >>`,
		Id: "Func_Param_Def",
		NTType: 21,
		Index: 55,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewParamDef(str(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Cust_Fn_def : "function" cust_fn_name ":" Func_Param_Def Code_Block	<< _rt.NewCustFuncDev(str(X[1]), X[3].([]string), c(X[4]), p(X[0])) >>`,
		Id: "Cust_Fn_def",
		NTType: 22,
		Index: 56,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewCustFuncDev(str(X[1]), X[3].([]string), c(X[4]), p(X[0]))
		},
	},
	ProdTabEntry{
		String: `Statement : Expression	<< X[0], nil >>`,
		Id: "Statement",
		NTType: 23,
		Index: 57,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statement : "return" Expression	<< _rt.NewReturn(c(X[1]), p(X[0])), nil >>`,
		Id: "Statement",
		NTType: 23,
		Index: 58,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewReturn(c(X[1]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Single_Statement : Statement ";"	<< X[0], nil >>`,
		Id: "Single_Statement",
		NTType: 24,
		Index: 59,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Single_Statement : Block	<< X[0], nil >>`,
		Id: "Single_Statement",
		NTType: 24,
		Index: 60,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Single_Statement : Cust_Fn_def	<< X[0], nil >>`,
		Id: "Single_Statement",
		NTType: 24,
		Index: 61,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : Single_Statement Statements	<< _rt.NewStatements(c(X[0]), c(X[1])), nil >>`,
		Id: "Statements",
		NTType: 25,
		Index: 62,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewStatements(c(X[0]), c(X[1])), nil
		},
	},
	ProdTabEntry{
		String: `Statements : Single_Statement	<< X[0], nil >>`,
		Id: "Statements",
		NTType: 25,
		Index: 63,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Statements : Statement	<< X[0], nil >>`,
		Id: "Statements",
		NTType: 25,
		Index: 64,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `If_Block : "if" Expression Code_Block "else" Code_Block	<< _rt.NewIfThenElse(c(X[1]), c(X[2]), c(X[4]), p(X[0])), nil >>`,
		Id: "If_Block",
		NTType: 26,
		Index: 65,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewIfThenElse(c(X[1]), c(X[2]), c(X[4]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `If_Block : "if" Expression Code_Block	<< _rt.NewIfThenElse(c(X[1]), c(X[2]), nil, p(X[0])), nil >>`,
		Id: "If_Block",
		NTType: 26,
		Index: 66,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewIfThenElse(c(X[1]), c(X[2]), nil, p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `While_Loop : "while" Expression Code_Block	<< _rt.NewWhileLoop(c(X[1]), c(X[2]), p(X[0])), nil >>`,
		Id: "While_Loop",
		NTType: 27,
		Index: 67,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewWhileLoop(c(X[1]), c(X[2]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `For_Each_Loop : "foreach" Variable "in" Expression Code_Block	<< _rt.NewForEach(str(X[1]), c(X[3]), c(X[4]), p(X[0])), nil >>`,
		Id: "For_Each_Loop",
		NTType: 28,
		Index: 68,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewForEach(str(X[1]), c(X[3]), c(X[4]), p(X[0])), nil
		},
	},
	ProdTabEntry{
		String: `Block : If_Block	<< X[0], nil >>`,
		Id: "Block",
		NTType: 29,
		Index: 69,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Block : While_Loop	<< X[0], nil >>`,
		Id: "Block",
		NTType: 29,
		Index: 70,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Block : For_Each_Loop	<< X[0], nil >>`,
		Id: "Block",
		NTType: 29,
		Index: 71,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Lambda_Def : "function" Func_Param_Def "->" Statement	<< _rt.NewLambdaDef(X[1].([]string), c(X[3]), p(X[0])), nil >>`,
		Id: "Lambda_Def",
		NTType: 30,
		Index: 72,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return _rt.NewLambdaDef(X[1].([]string), c(X[3]), p(X[0])), nil
		},
	},
	
}
