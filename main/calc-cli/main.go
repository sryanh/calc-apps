package main

import (
	"fmt"
	"os"

	handler "github.com/sryanh/calcy-apps/handlers"
	"github.com/sryanh/calcy-lib/calc"
)

func main() {
	handle := handler.NewHandler(os.Stdout, calc.Addition{})
	result, err := handle.Handle(os.Args[1:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", result)
}
