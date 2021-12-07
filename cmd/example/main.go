package main

import (
	"flag"
	"fmt"
	"strings"
	lab2 "github.com/CapralDavid/Architecture2"
)


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



func main(){
	fmt.Println("hello there")

	var stack []string
	//var value float64

    st:="* + 5 3 - 81 15"
	var symbols = strings.Fields(st)

	for i:=len(symbols)-1; i>=0; i--{
		fmt.Println(symbols[i])

		if isOperand(symbols[i]){
			var one, two string
			fmt.Println("OPERAND")
			//НЕТ ПРОВЕРКИ НА ОШИБКУ
			/*
			нужно что бы когда попадается операнда в стаке уже должно быть 
			2 числа.а то если будет - и только одно число то тогда как
			считать?
			что то в этом роде 				
			if len(stack)< 2{
			print
			}

			*/
			stack, two, one = popTwoFromArray(stack)
			fmt.Println("(" + one + symbols[i] + two + ")")
			otvet:="(" + one + symbols[i] + two + ")"
			stack = append(stack, otvet)



		
		}else{
			stack = append(stack, symbols[i])
		}

		fmt.Println("________THIS IS STACK NOW_______")
		fmt.Println(stack)
		fmt.Println("____THIS IS STACK NOW_____end_____")

	}
	flag.Parse()
	fmt.Println(len(symbols))
	fmt.Println(stack)
	flag.Parse()
	res, _ := lab2.PrefixToPostfix("+ 2 2")
	fmt.Println(res)
}

