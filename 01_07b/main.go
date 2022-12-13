package main

import (
	"flag"
	"log"
)

const (
	Brace = iota
	Bracket
	Paren
)

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	stack := []int{}
	for _, c := range expr {
		switch c {
		case '{':
			stack = append(stack, Brace)
		case '}':
			if peek(stack) == Brace {
				stack = pop(stack)
			} else {
				return false
			}

		case '[':
			stack = append(stack, Bracket)
		case ']':
			if peek(stack) == Bracket {
				stack = pop(stack)
			} else {
				return false
			}

		case '(':
			stack = append(stack, Paren)
		case ')':
			if peek(stack) == Paren {
				stack = pop(stack)
			} else {
				return false
			}
		default:
			continue
		}
	}
	return len(stack) == 0
}

func peek(stack []int) int {
	if len(stack) == 0 {
		return -1
	}
	return stack[len(stack)-1]
}

func pop(stack []int) []int {
	return stack[:len(stack)-1]
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
