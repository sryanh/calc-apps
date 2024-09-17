package main

import (
	"flag"
	"fmt"
	"os"

	handler "github.com/sryanh/calcy-apps/handlers"
	"github.com/sryanh/calcy-lib/calc"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "+", "Valid operations: +, - , *, /")
	flag.Parse()
	vals := flag.Args()
	operation, _ := calculatorOperations[op] // Could handle error here...

	handle := handler.NewHandler(os.Stdout, operation)
	result, err := handle.Handle(vals)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", result)
}

// I feel like there is a better place for this but I don't know where yet.
var calculatorOperations = map[string]handler.Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"*": calc.Multiplication{},
	"/": calc.Division{},
	//"^" : calc.Power{}
}
