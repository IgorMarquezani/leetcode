package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for _, v := range tokens {
		switch v {
		case "+":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b-a)
		case "*":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a, b := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, b/a)
		default:
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func main() {
	println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
