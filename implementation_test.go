package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnPrefixIntoInfixOne(t *testing.T) {
	res, err := TurnPrefixIntoInfix("- 81 15")
	if assert.Nil(t, err) {
		//fmt.Println("НЕТ ОШИБОЧКИ")
		assert.Equal(t, "(81-15)", res)
	}
}


func TestTurnPrefixIntoInfixTwo(t *testing.T) {
	res, err := TurnPrefixIntoInfix("+ * 16 35 * 22 77")
	if assert.Nil(t, err) {
		//fmt.Println("НЕТ ОШИБОЧКИ")
		assert.Equal(t, "((16*35)+(22*77))", res)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := TurnPrefixIntoInfix("* + 5 3 - 81 15")
	fmt.Println(res)

	// Output:
	// ((5+3)*(81-15))
}
