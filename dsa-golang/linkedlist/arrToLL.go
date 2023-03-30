package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func convertToLinkedList(arr []int) *Node {

	head := &Node{arr[0], nil}

	curr := head

	for i := 1; i < len(arr); i++ {
		curr.next = &Node{arr[i], nil}
		curr = curr.next
	}

	return head
}

func printLinkedList(head *Node) {
	for curr := head; curr != nil; curr = curr.next {
		fmt.Print(curr.data, " -> ")
	}

	fmt.Println("nil")
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	head := convertToLinkedList(arr)

	printLinkedList(head)
}
