package main

import "strconv"

func main() {
	input := []string{"4", "13", "5", "/", "+"}
	output := evalRPN(input)
	println(output)
}

func evalRPN(tokens []string) int {
	tokenOperatorToOperation := map[string]func(int, int) int{
		"+": Add,
		"-": Subtract,
		"*": Multiply,
		"/": Divide,
	}

	tokensInOperation := []int{}
	for _, token := range tokens {
		number, err := strconv.Atoi(string(token))
		if err != nil {
			operation := tokenOperatorToOperation[token]
			a := tokensInOperation[len(tokensInOperation)-2]
			b := tokensInOperation[len(tokensInOperation)-1]
			tokensInOperation = tokensInOperation[:len(tokensInOperation)-2]
			tokensInOperation = append(tokensInOperation, operation(a, b))
		} else {
			tokensInOperation = append(tokensInOperation, number)
		}
	}

	return tokensInOperation[0]
}

func Add(val1 int, val2 int) int {
	return val1 + val2
}

func Subtract(val1 int, val2 int) int {
	return val1 - val2
}

func Divide(val1 int, val2 int) int {
	return val1 / val2
}

func Multiply(val1 int, val2 int) int {
	return val1 * val2
}
