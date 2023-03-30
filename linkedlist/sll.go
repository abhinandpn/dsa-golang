package main

import (
	"fmt"
	"os"
)

// Create Node
type Node struct {
	value int   // Node value
	next  *Node // Next node address
}

// Create LinkedList
type LinkedList struct {
	Head *Node // Head of LinkedList
	tail *Node
	len  int // Length of LinkedLits
}

// Add value into Head
func (l *LinkedList) AddHead(val int) {
	NewNode := &Node{value: val} //Creating TempNode as NewNode

	// Checking Head is nil
	if l.Head == nil {
		l.Head = NewNode
		l.len++ // Length Incresing
		return
	}

	// If we have Head value
	// Swapping
	temp := l.Head
	l.Head = NewNode
	l.Head.next = temp
	l.len++ // Length Incresing
}

// Add value in to tail
func (l *LinkedList) AddTail(val int) {
	NewNode := &Node{value: val} //Creating TempNode as NewNode

	// Checking Head is nil
	if l.Head == nil {
		l.Head = NewNode
		l.len++ // Length Incresing
		return
	}
	// Setting a temp Head
	temp := l.Head
	for temp.next != nil { // Check current tempHead.next is nil
		temp = temp.next // Changing the TempHead
	}
	temp.next = NewNode // Add New Tail Node
	l.len++             //length Incresing

}

// Add value in User defined
func (l *LinkedList) AddUserDef(indx, val int) {

	NewNode := &Node{value: val}

	if l.Head == nil {
		fmt.Println("No LinkedList Showing ...")
		var choice string
		fmt.Println("----------")
		fmt.Println("1. Add value in to head")
		fmt.Println("2. Exit")
		fmt.Print("Enter your choice : ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			New1 := &LinkedList{} // Creating new linkedList
			New1.AddHead(val)
			fmt.Print("New list is  : ")
			New1.Display()
		case "2":
			os.Exit(0)
		}

	}
	temp := l.Head
	if indx < l.len {
		if indx == 1 {
			NewNode.next = temp
			l.Head = NewNode
			return
		}

		for i := 1; i < indx-1; i++ {
			temp = temp.next
		}
		NewNode.next = temp.next
		temp.next = NewNode
		l.len++

	} else if indx == l.len+1 {
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = NewNode
		l.len++
		return
	}

}

// Display Linked List
func (l *LinkedList) Display() {
	if l.Head == nil { // Chech Head node nil
		fmt.Println("No value is there") // If head is nil
	}
	temp := l.Head    // Temp head for Looping
	for temp != nil { // Lopping head to last
		fmt.Print(" --> ", temp.value) // Print value
		temp = temp.next               // Temp head changing to Head.next
	}
}

func (l LinkedList) DisplayHead() {
	fmt.Println(l.Head)
}
func (l LinkedList) DisplayTail() {
	temp := l.Head
	for temp.next != nil {
		temp = temp.next
	}
	fmt.Println(temp)
}

func (l *LinkedList) AddNode(val int) {
	NewNode := &Node{value: val}
	if l.Head == nil {
		l.Head = NewNode
	} else {
		l.tail.next = NewNode
	}
	l.tail = NewNode
	l.len++

}

func (l *LinkedList) Dellndx(indx int) {
	temp := l.Head
	if indx > l.len+1 {
		fmt.Println("Not found")
	} else {
		if indx == 1 {
			l.Head = temp.next
			l.len--
			return
		}
		for i := 1; i < indx-1; i++ {
			temp = temp.next
		}
		if temp.next == nil {
			l.tail = temp
			l.len--
			return
		}
		temp.next = temp.next.next
		l.len--
		return
	}

}

func (l *LinkedList) DelVal(val int) {
	temp := l.Head
	for temp.next.value == val {
		temp = temp.next
		l.len--
		return

	}
	if temp.next.next == nil {
		l.tail = temp
		temp.next = nil
		l.len--
		return
	}
	temp.next = temp.next.next
	l.len--
	return
}
func (l LinkedList) IndxDsply(indx int) {
	temp := l.Head
	for i := 1; i < indx; i++ {
		temp = temp.next
	}
	fmt.Println("Value is : ", temp.value)
}

func (l LinkedList) DltDupl() {
	temp := l.Head
	for i := 1; i < l.len; i++ {

		for j := i + 1; j < l.len-1; j++ {
			tempnxt := temp.next
			if temp.value == tempnxt.value && tempnxt.next != nil {
				temp.next = tempnxt.next
				// fmt.Println("1")

			} else if tempnxt.next == nil {
				l.tail = temp
				temp.next = nil
				fmt.Println("3")
				return
			}
			tempnxt = tempnxt.next
			// fmt.Println("4")
		}
		temp = temp.next
		// fmt.Println("5")
	}
}
func (l *LinkedList) RevDis() {
	temp := l.Head
	var prev *Node
	for temp != nil {
		nxt := temp.next
		temp.next = prev
		prev = temp
		temp = nxt
	}
	l.Head = prev
}

// func main() {
// 	New := &LinkedList{} // Creating new linkedList

// 	New.AddNode(34)
// 	New.AddNode(56)
// 	New.AddNode(26)
// 	New.AddNode(66)
// 	New.AddNode(76)
// 	New.AddNode(26)
// 	New.Display()
// 	fmt.Println("\n..................")
// 	New.RevDis()
// 	New.Display()
// 	// New.AddNode(6)

// 	// New.DltDupl()
// 	// New.Display()

// 	// New.Display()
// 	// fmt.Println("\n----------------")
// 	// New.Dellndx(7)
// 	// New.Display()
// 	// fmt.Println("\n----------------")
// 	// fmt.Print("Head is : ")
// 	// New.DisplayHead()
// 	// fmt.Print("Tail is : ")
// 	// New.DisplayTail()
// 	// fmt.Println("\n--------------")
// 	// New.Display()
// 	// fmt.Println("\n--------------")
// 	// New.IndxDsply(3)
// 	// fmt.Println("\n--------------")

// 	// New.AddTail(10)
// 	// New.AddTail(20)
// 	// New.AddTail(30)
// 	// New.AddTail(40)
// 	// New.AddUserDef(5, 23)
// 	// New.Display()
// 	// New.DisplayTail()
// 	// fmt.Println("\n-------------------")
// 	// New.DisplayHead()
// 	// New.Display()

// 	// New.AddTail(10)
// 	// New.AddTail(20)
// 	// New.AddTail(30)
// 	// New.AddTail(40)
// 	// fmt.Print("Linked List is  : ")
// 	// New.Display()
// 	// fmt.Println("-------------------")
// 	// New.AddHead(34)
// 	// New.Display()

// }

// // Oprations

// /*
// TODO: Get
// TODO: Set
// TODO: Update
// */
