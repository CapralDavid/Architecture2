package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	res, err := TurnPrefixIntoInfix("* + 5 3 - 81 15")
	if assert.Nil(t, err) {
		assert.Equal(t, "((5+3)*(81-15))", res)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := TurnPrefixIntoInfix("* + 5 3 - 81 15")
	fmt.Println(res)

	// Output:
	// 2 2 +
}
