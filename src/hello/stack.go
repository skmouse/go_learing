package main

import "fmt"

type Stack struct {
	items []int
	n     int //栈大小
}

func (stack *Stack) len() int {
	return len(stack.items)
}

func (stack *Stack) cap() int {
	return cap(stack.items)
}

func (stack *Stack) empty() bool {
	return len(stack.items) == 0
}

func (stack *Stack) push(data int) bool {
	if stack.n == stack.len() {
		return false
	}

	stack.items = append(stack.items, data)
	return true
}

func (stack *Stack) pop() bool {
	if stack.len() == 0 {
		return false
	}
	stack.items = stack.items[:stack.len()-1]
	return true

}

func (stack Stack) r() {
	for k, v := range stack.items {
		fmt.Println(k, v)
	}

}

func main() {
	stack := Stack{[]int{}, 4}
	arr := [4]int{10, 20, 30, 40}

	for _, v := range arr {
		stack.push(v)
	}

	stack.pop()
	stack.pop()

	fmt.Println(stack)

}
