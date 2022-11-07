package main

type FuncParams struct {
	MathExpression string  `JSON:"mathExpression"`
	A              float64 `JSON:"a"`
	B              float64 `JSON:"b"`
	Precision      float64 `JSON:"precision"`
}
