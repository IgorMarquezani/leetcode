package main

func isValid(s string) bool {
	stack := []rune{}

	for _, v := range s {
		stack = append(stack, v)
		if len(stack) < 2 {
			continue
		}

		if stack[len(stack)-2] == '{' && stack[len(stack)-1] == '}' {
			stack = stack[:len(stack)-2]
		} else if stack[len(stack)-2] == '(' && stack[len(stack)-1] == ')' {
			stack = stack[:len(stack)-2]
		} else if stack[len(stack)-2] == '[' && stack[len(stack)-1] == ']' {
			stack = stack[:len(stack)-2]
		}
	}

	return len(stack) == 0
}

func main() {
	s := "[({})]"
	println(isValid(s))
	s = "[{})]"
	println(isValid(s))
}
