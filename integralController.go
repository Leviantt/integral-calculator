package main

import (
	"math"

	"github.com/Knetic/govaluate"
	"github.com/benpate/convert"
)

// creates functions, that calculates math expression depending on x parameter
func MakeCalculate(expression string) func(x float64) float64 {
	_expression, _ := govaluate.NewEvaluableExpression(expression)
	parameters := make(map[string]interface{}, 8)

	return func(x float64) float64 {
		parameters["x"] = x
		result, err := _expression.Evaluate(parameters)
		if err != nil {
			return convert.Float(0)
		}
		return convert.Float(result)
	}
}

// calculates the integral by the trapezoid method, dividing the segment into n segments
func CalculateIntegral(funcParams *FuncParams, n int) float64 {
	var a = funcParams.A
	var b = funcParams.B
	var mathExpression = funcParams.MathExpression

	calculate := MakeCalculate(mathExpression)

	var step = (b - a) / float64(n)

	var sum = (calculate(a) + calculate(b)) / 2

	for x := a + step; x < b; x += step {
		sum += calculate(x)
	}
	return sum * step
}

const PARTITION_NUMBER = 25
//calculates the integral, estimating the precision by the Runge method
func CalculateIntegralWithPrecision(funcParams *FuncParams) float64 {

	var n = PARTITION_NUMBER
	var firstResult = CalculateIntegral(funcParams, n)
	var secondResult = CalculateIntegral(funcParams, n*2)

	for math.Abs(secondResult-firstResult)/float64(3) >= funcParams.Precision {
		n *= 2
		firstResult = CalculateIntegral(funcParams, n)
		secondResult = CalculateIntegral(funcParams, n*2)
	}

	return secondResult
}
