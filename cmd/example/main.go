package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	lab2 "github.com/CapralDavid/Architecture2"
)


var (
	inputExpression = flag.String("e", "", "Expression to evaluate")
	// TODO: Add other flags support for input and output configuration.
)

//КОНСОЛЬ.если не хватает прообела между символами,то работает не до конца

func  main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *inputExpression !="" {
		input= strings.NewReader(*inputExpression)
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Println(err.Error())
	}

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	//res, _ := lab2.TurnPrefixIntoInfix("* + 5 3 - 81 15")
	//fmt.Println(res)
}



