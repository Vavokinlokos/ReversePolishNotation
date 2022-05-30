package main

import (
	"fmt"
	"github.com/Vavokinlokos/ReversePolishNotation/stack"
	"strings"
)

type Pair struct {
	Sign     string
	Priority int
}

func main() {
	var stack stack.Stack
	var output string

	input := "(a+b)*c"
	output = reversePolishNotation(input, stack)
	fmt.Println(output)
}

func reversePolishNotation(input string, stack stack.Stack) string {
	priorities := defOperationsPriority()
	output := ""
	j := -1
	splittedInput := strings.Split(input, "")
	for i, v := range splittedInput {
		if i < j {
			continue
		}
		if v == ")" {
			element := ""
			for element != "(" {
				topEl, _ := stack.Pop()
				if topEl == "(" {
					element = topEl
					continue
				}
				element = topEl
				output += topEl
				i++
			}
			j = i
			continue
		}
		if v == "(" || v == "+" || v == "-" || v == "/" || v == "*" {
			topElement, full := stack.Pop()
			if !full {
				stack.Push(v)
				continue
			}
			if getPriority(priorities, v) > getPriority(priorities, topElement) {
				stack.Push(topElement)
				stack.Push(v)
			} else {
				for getPriority(priorities, v) < getPriority(priorities, topElement) {
					output += topElement
					topElement, _ = stack.Pop()
				}
			}
		} else {
			output += v
		}
	}
	for !stack.IsEmpty() {
		el, _ := stack.Pop()
		output += el
	}
	return output
}

func defOperationsPriority() [6]Pair {
	return [6]Pair{
		{
			Sign:     "(",
			Priority: 0,
		},
		{
			Sign:     ")",
			Priority: 1,
		},
		{
			Sign:     "+",
			Priority: 2,
		},
		{
			Sign:     "-",
			Priority: 2,
		},
		{
			Sign:     "*",
			Priority: 3,
		},
		{
			Sign:     "/",
			Priority: 3,
		},
	}
}

func getPriority(pairs [6]Pair, operation string) int {
	for _, v := range pairs {
		if v.Sign == operation {
			return v.Priority
		}
	}
	return -1
}
