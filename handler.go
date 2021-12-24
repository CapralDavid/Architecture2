package lab2

import (
	"bufio"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.

//описываем computeHandler
//в main.go нет упоминания TurnPrefixIntoInfix,как и указано в задании
type ComputeHandler struct {
	// TODO: Add necessary fields.
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	fileScanner := bufio.NewScanner(ch.Input)

	for fileScanner.Scan() {
		input := fileScanner.Text()
		res, err := TurnPrefixIntoInfix(input)
		if err != nil {
			return err
		}
		ch.Output.Write([]byte(res))
	}

	return nil
}
