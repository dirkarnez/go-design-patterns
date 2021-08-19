package strategy

import (
	"fmt"
)

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (a Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Subtraction struct{}

func (s Subtraction) Apply(lval, rval int) int {
	return lval - rval
}

type Multiplication struct{}

func (m Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

type Division struct{}

func (d Division) Apply(lval, rval int) int {
	return lval / rval
}

func Start() {
	operationCode := "*" // +, -, *, /
	operation := DoOperation(operationCode)
	fmt.Println(operation.Operate(5, 2))
}

func DoOperation(operationCode string) Operation {
	operationMap := make(map[string]Operation)
	operationMap [ "+" ] = Operation{Addition{}}
	operationMap [ "-" ] = Operation{Subtraction{}}
	operationMap [ "*" ] = Operation{Multiplication{}}
	operationMap [ "/ " ] = Operation{Division{}}
	return operationMap[operationCode]
}
