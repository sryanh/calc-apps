package handlers

import (
	"fmt"
	"io"
	"strconv"

	"github.com/sryanh/calcy-lib/calc"
)

type Calculator interface{ Calculate(a, b int) int }

type Handler struct {
	output     io.Writer
	calculator calc.Calculator
}

func NewHandler(output io.Writer, calculator calc.Calculator) *Handler {
	return &Handler{output: output, calculator: calculator}
}

func (this *Handler) Handle(values []string) (int, error) {
	if this.calculator == nil {
		return 0, errNilPointerReference
	}
	if len(values) != 2 {
		return 0, errTooFewArgs
	}
	a, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, errMalformedInput
	}

	b, err := strconv.Atoi(values[1])
	if err != nil {
		return 0, errMalformedInput
	}
	return this.calculator.Calculate(a, b), err
}

var (
	errTooFewArgs          = fmt.Errorf("Too few arguments provided.")
	errMalformedInput      = fmt.Errorf("Invalid argument type provided.")
	errNilPointerReference = fmt.Errorf("Malformed input led to nil pointer reference.")
)
