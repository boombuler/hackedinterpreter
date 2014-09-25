package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/boombuler/gold"
	"io"
	"log"
	"os"
)

type ExecFn func(t *gold.Token, c *Context) (Value, error)

func ExecWantList(t *gold.Token, c *Context) (*List, error) {
	val, err := Exec(t, c)
	if err != nil {
		return nil, err
	}
	valLst, ok := val.(*List)
	if !ok {
		return nil, fmt.Errorf("List expected got %v", getTypeName(val))
	}
	return valLst, nil
}

func ExecWantInt(t *gold.Token, c *Context) (int, error) {
	val, err := Exec(t, c)
	if err != nil {
		return 0, err
	}
	valInt, ok := val.(int)
	if !ok {
		return 0, fmt.Errorf("Integer expected got %v", getTypeName(val))
	}
	return valInt, nil
}

func ExecWantBool(t *gold.Token, c *Context) (bool, error) {
	val, err := Exec(t, c)
	if err != nil {
		return false, err
	}
	valBool, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("Boolean expected got %v", getTypeName(val))
	}
	return valBool, nil
}

func Exec(t *gold.Token, c *Context) (Value, error) {
	if c.result != nil || c.err != nil {
		return c.result, c.err
	}

	var execMap map[gold.RuleId]ExecFn = map[gold.RuleId]ExecFn{
		Rule_AddexpPlus:                   AddValues,
		Rule_AddexpMinus:                  SubValues,
		Rule_ValueNumber:                  ConstIntValue,
		Rule_MultexpTimes:                 MultiplyValues,
		Rule_MultexpDiv:                   DivideValues,
		Rule_LambdatargetLparenRparen:     ParenValue,
		Rule_AssignEq:                     AssignstmntVariableEq,
		Rule_Statement:                    ExecStatements,
		Rule_Statement2:                   AssignstmntVariableEq,
		Rule_VariableInput:                ValueVariable,
		Rule_VariableVariable:             ValueVariable,
		Rule_Statements:                   ExecStatements,
		Rule_Statements2:                  ExecStatements,
		Rule_Statements3:                  ExecStatements,
		Rule_SinglestatementSemi:          ExecStatements,
		Rule_CodeblockLbraceRbrace:        ExecCodeBlock,
		Rule_CompexpressionGt:             ExpressionGt,
		Rule_CompexpressionLt:             ExpressionLt,
		Rule_CompexpressionEqeq:           ExpressionEqual,
		Rule_CompexpressionExclameq:       ExpressionNotEqual,
		Rule_BoolalgexpressionAmpamp:      ExpressionLogicalAND,
		Rule_BoolalgexpressionPipepipe:    ExpressionLogicalOR,
		Rule_Value2:                       ConstBoolValue,
		Rule_BoolFalse:                    ConstBoolValue,
		Rule_BoolTrue:                     ConstBoolValue,
		Rule_ArrayLbracketRbracket:        ConstList,
		Rule_GetindexLbracketRbracket:     ListGetIndex,
		Rule_IfblockIf:                    ExecIfBlock,
		Rule_WhileloopWhile:               ExecWhileLoop,
		Rule_MethodcallDot:                CallMethodNoParam,
		Rule_MethodcallDotLparenRparen:    CallMethodParam,
		Rule_ForeachloopForeachVariableIn: ExecForEach,
		Rule_CustfndefFunctionColon:       DefineFunction,
		Rule_CustfncallLparenRparen:       CallFunction,
		Rule_ReturnstmntReturn:            Return,
		Rule_LambdafndefFunctionMinusgt:   DefineLambda,
		Rule_LambdafncallLparenRparen:     CallLambda,

		Rule_CustfnnameF1:  GetFunction,
		Rule_CustfnnameF2:  GetFunction,
		Rule_CustfnnameF3:  GetFunction,
		Rule_CustfnnameF4:  GetFunction,
		Rule_CustfnnameF5:  GetFunction,
		Rule_CustfnnameF6:  GetFunction,
		Rule_CustfnnameF7:  GetFunction,
		Rule_CustfnnameF8:  GetFunction,
		Rule_CustfnnameF9:  GetFunction,
		Rule_CustfnnameF10: GetFunction,
		Rule_CustfnnameF11: GetFunction,
		Rule_CustfnnameF12: GetFunction,
		Rule_CustfnnameF13: GetFunction,
		Rule_CustfnnameF14: GetFunction,
		Rule_CustfnnameF15: GetFunction,
		Rule_CustfnnameF16: GetFunction,
	}

	method, ok := execMap[t.Rule]
	if ok {
		return method(t, c)
	} else {
		fmt.Println("Failed to exec:", t.Text, "rule:", t.Rule, "symbol:", t.Symbol)
		panic("Invalid Operation")
		return nil, errors.New("Invalid Operation")
	}
}

func logAST(t *gold.Token, i int) {
	for x := 0; x < i; x++ {
		fmt.Print(" ")
	}
	fmt.Println(t.Text)
	for _, token := range t.Tokens {
		if !token.IsTerminal {
			logAST(token, i+2)
		}
	}
}

func PrintAST(code string) {
	reader := bytes.NewBuffer([]byte(code))
	res, err := NewParser().Parse(reader, true)
	if err != nil {
		return
	}
	logAST(res, 0)
}

func ExecuteString(code string) (Value, error) {
	return ExecuteReader(bytes.NewBuffer([]byte(code)))
}

func ExecuteReader(reader io.Reader) (Value, error) {
	res, err := NewParser().Parse(reader, true)
	if err != nil {
		return nil, err
	}

	return Exec(res, NewContext())
}

func NewParser() gold.Parser {
	f, err := os.Open("hacked.egt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	parser, err := gold.NewParser(f)
	if err != nil {
		log.Fatal(err)
	}
	return parser
}
