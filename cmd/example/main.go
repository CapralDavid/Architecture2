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
	inputFile       = flag.String("f", "", "add input file")
	outputFile      = flag.String("o", "", "Output file")
	// TODO: Add other flags support for input and output configuration.
)

//КОНСОЛЬ.если не хватает прообела между символами,то работает не до конца

func  main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	//консоль
	//если inputExpression не равен "" то ...

	if *inputExpression !="" {
		input= strings.NewReader(*inputExpression)
	}

	//отдельный файл откуда еберм инфу
	//если inputFile не равен "" то ...

	if *inputFile != "" {
		res, err := os.Open(*inputFile)
		if err != nil {
			fmt.Println("Error with input file option")
		}
		input = res
	}

	if *outputFile != "" {
		res, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error with output file option")
		}
		output = res
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



