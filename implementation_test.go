package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTurnPrefixIntoInfixOne(t *testing.T) {
	res, err := TurnPrefixIntoInfix("* + 5 3 - 81 15")
	if assert.Nil(t, err) {
		fmt.Println("НЕТ ОШИБОЧКИ")
		assert.Equal(t, "((5+3)*(81-15))", res)
	}
}


func TestTurnPrefixIntoInfixTwo(t *testing.T) {
	res, err := TurnPrefixIntoInfix("+ * 16 35 * 22 77")
	if assert.Nil(t, err) {
		fmt.Println("НЕТ ОШИБОЧКИ")
		assert.Equal(t, "((16*35)+(22*77))", res)
	}
}

//тест ошибок
// !!!!! не готово !!!!!
// должно проверить сообщение ошибки,тип правильно ли програма выбрала ошибку

//&errors.errorStringPs:"ОШИБОЧКА, неверное выражение.Короткое"
func TestTurnPrefixIntoInfixThreeERROR(t *testing.T) {
	res, err := TurnPrefixIntoInfix("+")

	//assert.Equal(t, "", res)

	//assert.Equal(t, "ОШИБОЧКА, неверное выражение.Короткое", err)

	if err ==nil {
		//fmt.Println("НЕТ ОШИБОЧКИ")
		t.Fatal("error supposed to not be nil")
	} else{
		assert.Equal(t, "", res)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := TurnPrefixIntoInfix("* + 5 3 - 81 15")
	fmt.Println(res)

	// Output:
	// ((5+3)*(81-15))
}
