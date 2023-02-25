package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Usage: <calculator> <operation> <operand>")
		os.Exit(1)
	}

	operation := args[0]
	operands := args[1:]

	switch operation {
	case "add":
		result := add(operands)
		fmt.Println(result)
	case "sub":
		result := sub(operands)
		fmt.Println(result)
	case "mul":
		result := mul(operands)
		fmt.Println(result)
	case "div":
		result, err := div(operands)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result)
	default:
		fmt.Println("Unknown operation:", operation)
		os.Exit(1)
	}
}

func add(operands []string) int {
	result := 0
	for _, operand := range operands {
		value, _ := strconv.Atoi(operand)
		result += value
	}
	return result
}

func sub(operands []string) int {
	result, _ := strconv.Atoi(operands[0])
	for _, operand := range operands[1:] {
		value, _ := strconv.Atoi(operand)
		result -= value
	}
	return result
}

func mul(operands []string) int {
	result := 1
	for _, operand := range operands {
		value, _ := strconv.Atoi(operand)
		result *= value
	}
	return result
}

func div(operands []string) (int, error) {
	if len(operands) != 2 {
		return 0, fmt.Errorf("usage: calculator div <dividend> <divisor>")
	}
	dividend, err := strconv.Atoi(operands[0])
	if err != nil {
		return 0, fmt.Errorf("invalid dividend: %s", operands[0])
	}
	divisor, err := strconv.Atoi(operands[1])
	if err != nil {
		return 0, fmt.Errorf("invalid divisor: %s", operands[1])
	}
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}
