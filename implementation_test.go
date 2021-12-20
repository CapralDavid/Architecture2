package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnPrefixIntoInfix(t *testing.T) {
	res, err := TurnPrefixIntoInfix("* + 5 3 - 81 15")
	if assert.Nil(t, err) {
		fmt.Println("НЕТ ОШИБОЧКИ")
		assert.Equal(t, "((5+3)*(81-15))", res)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := TurnPrefixIntoInfix("* + 5 3 - 81 15")
	fmt.Println(res)

	// Output:
	// ((5+3)*(81-15))
}
