package main

import (
	//"./datastructures/basic"
	"fmt"
	"./fbrandes/algorithms/searching"
	"./fbrandes/algorithms/sorting"
)

func main() {
	a := []int{3, 6, 5, 79, 8, 8, 498, 47, 5}
	fmt.Println(sorting.ShellSort(a))

	a = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16}
	fmt.Println(searching.BinarySearch(a, 6))
	fmt.Println(searching.LinearSearch(a, 6))

	//var stack basic.Stack
	//stack.New()
	//stack.Push(basic.Item(3))
	//stack.Push(basic.Item(2))
	//stack.Push(basic.Item(1))
	//fmt.Println(stack)
	//stack.Pop()
	//fmt.Println(stack)
	//stack.Push(basic.Item(4))
	//fmt.Println(stack)
	//fmt.Println(stack.Peek())
	//fmt.Println(stack.IsEmpyt())
	//stack.Pop()
	//stack.Pop()
	//stack.Pop()
	//fmt.Println(stack.IsEmpyt())
}
