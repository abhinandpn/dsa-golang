package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}
type LinkedList struct {
	Head *Node
	Tail *Node
	len  int
}

func (l *LinkedList) AddTail(val int) {
	Newnode := &Node{value: val}
	if l.Head == nil {
		l.Head = Newnode
		l.len++
	} else {
		l.Tail.next = Newnode
		Newnode.prev = l.Tail
		l.len++
	}
	l.Tail = Newnode

	// return

}
func (l *LinkedList) AddHead(val int) {
	newNode := &Node{value: val}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	temp := newNode
	l.Head.prev = temp
	temp.next = l.Head
	l.Head = temp
	l.len++
}

func (l *LinkedList) AddSpec(val, idx int) {

	Newnode := *&Node{value: val}

	if idx == 1 {

		l.Head.prev = &Newnode
		Newnode.next = l.Head
		l.Head = &Newnode
		return
	}
	// fmt.Println(l.len)
	if idx > l.len {
		fmt.Println("Index not found")
		// return
	}
	if idx <= l.len {

		temp := l.Head
		pvs := l.Head
		for i := 1; i < idx; i++ {
			pvs = temp
			temp = temp.next

		}
		pvs.next = &Newnode
		Newnode.prev = pvs
		Newnode.next = temp
		temp.prev = &Newnode
		l.len++
		// fmt.Println(l.len)
		return
	}

}

func (l *LinkedList) DelHead() {
	if l.len == 1 {
		fmt.Println("Only one value is found")
		l.Head = nil
		return
	}
	l.Head = l.Head.next
	l.Head.prev = nil
	l.len--
	// fmt.Println(l.len)
}
func (l *LinkedList) DelTail() {
	if l.len == 1 {
		fmt.Println("Only one value is found")
		l.Tail = nil
		return
	}
	l.Tail = l.Tail.prev
	l.Tail.next = nil
	l.len--
}
func (l *LinkedList) DelSpc(idx int) {

	if idx == 1 {
		l.DelHead()
		return
	}
	if idx == l.len {
		l.DelTail()
		return
	}
	if idx <= l.len {
		temp := l.Head
		prv := l.Head
		for i := 1; i < idx; i++ {
			temp = temp.next
			prv = temp.prev
		}
		prv.next = temp.next
		temp.next.prev = prv
		l.len--
		return
	} else if idx > l.len {
		fmt.Println("Index not found")
	}

}
func (l LinkedList) Display() {

	if l.Head == nil {
		fmt.Println("No Linked List Found")
	}
	temp := l.Head
	for temp != nil {
		fmt.Print(" --> ", temp.value)
		temp = temp.next
	}

}
func (l LinkedList) DelVal(val int) {
	// fmt.Println("hh")
	temp := l.Head
	inx := 0
	for i := 1; i < l.len; i++ {
		if temp.value != val {
			temp = temp.next
			inx++
		} else if temp.value == val {
			fmt.Println("Index value is : ", inx+1)
			return
		}

	}
	if temp.value != val {
		fmt.Println("Value is not found")
		return
	}

}
func (l *LinkedList) Delposi(indx int) {
	temp := l.Head
	prv := l.Head
	for i := 1; i < indx; i++ {
		temp = temp.next
		prv = temp.prev
	}
	prv.next = temp.next
	temp.next.prev = prv
	l.len--

}
func (l LinkedList) RevDis() {
	if l.Tail == nil {
		fmt.Println("No value found")
		return
	}
	temp := l.Tail
	for temp != nil {
		fmt.Print(" <-- ", temp.value)
		temp = temp.prev
	}
}

// func main() {
// 	New := &LinkedList{}
// 	New.AddTail(6)
// 	New.AddTail(3)
// 	New.AddTail(44)
// 	New.AddTail(69)
// 	New.AddTail(8)

// 	New.Display()
// 	fmt.Print("\n")
// 	// New.AddHead(55)
// 	// New.Display()
// 	fmt.Println("\n............................")
// 	// New.DelVal(40)
// 	New.Delposi(3)
// 	// New.DelSpc(8)
// 	// New.DelHead()
// 	// New.DelTail()
// 	// New.AddSpec(10, 6)
// 	// New.RevDis()
// 	New.Display()
// }
