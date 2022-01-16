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

//идея 
// 1.простое выражение
// 2.цикл который после каждого теста добавляет к концу переменной что тестируем еще одно простое выражение
//  у тебя есть хороший туториал на англ Beautifully Simple Benchmarking with Go

// t = 1+1
// t1 = 1+1
// t2 = 1+1+1+1
// t1000 = ...

func BenchmarkTurnPrefixIntoInfix(b *testing.B) {
	var inpStr = " * + 5 3 - 81 15 "
	//var inpStrConc = " * + 5 3 - 81 15 "

	for i := 0; i < 20; i++ {
		inpStr += inpStr
		//fmt.Println(inpStr, i)

		b.Run(fmt.Sprintf("len=%d", 7+i*7), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TurnPrefixIntoInfix(inpStr)
			}
		})

	}



}

// for i := 0; i < b.N; i++ {
// 	TurnPrefixIntoInfix(inpStr)
// }


// for i := 0; i < 20; i++ {

// }

