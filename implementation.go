package lab2

import (
	"fmt"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}




/////////////////////////////////////////////////


func isOperand(x string) (bool){
	switch x{
	case "+", "-", "*", "/", "^":
		return true

	default:
		return false

	}

}

//вынимаем из стака два последних элемента
func popTwoFromArray(stack []string) ([]string, string, string) {
	var ab []string
	stack, ab = stack[:len(stack)-2], stack[len(stack)-2:]
	return stack, ab[0], ab[1]
}


func TurnPrefixIntoInfix(s string) (string, error){
	/*
	стак куда складываем наше ификсное выражение
	список символов которые будем рассматривать СПРАВО НАЛЕВО

	*/
	var stack []string
	var symbols = strings.Fields(s)
	//var value float64


//!!!!!!!!!!!!!!!!
	for i:=len(symbols)-1; i>=0; i--{
		//fmt.Println(symbols[i])


		if isOperand(symbols[i]){
			var one, two string
			//fmt.Println("OPERAND")
			//проверяем на ошибку
			if len(stack)< 2{
				return "0", fmt.Errorf("ОШИБОЧКА")
			}

			stack, two, one = popTwoFromArray(stack)
			//fmt.Println("(" + one + symbols[i] + two + ")")
			otvet:="(" + one + symbols[i] + two + ")"
			stack = append(stack, otvet)
		
		}else{
			stack = append(stack, symbols[i])
		}

	}

	//fmt.Println("_______FINAL STACK")
	//fmt.Println(stack[0])

	return stack[0], nil



}