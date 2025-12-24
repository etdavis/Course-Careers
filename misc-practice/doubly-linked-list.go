package main

import "fmt"

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	value T
}

type LinkedList[T any] struct {
	head   *Node[T]
	length uint
}

func (ll *LinkedList[T]) index(idx uint) (T, bool) {
	// return value and "ok", indicating if index is valid
	if idx >= ll.length {
		var zero T
		return zero, false
	} else {
		currentNode := ll.head
		for i := 0; i < int(idx); i++ {
			currentNode = currentNode.next
		}
		return currentNode.value, true
	}
}

func (ll *LinkedList[T]) append(value T) {
	ll.length++
	newNode := &Node[T]{value: value}
	if ll.head == nil {
		ll.head = newNode
		newNode.next = newNode
		newNode.prev = newNode
	} else {
		tail := ll.head.prev
		tail.next = newNode
		newNode.prev = tail
		newNode.next = ll.head
		ll.head.prev = newNode
	}
}

func (ll *LinkedList[T]) pop() T {
	ll.length--
	tail := ll.head.prev
	if tail == ll.head {
		ll.head = nil
	} else {
		newTail := tail.prev
		newTail.next = ll.head
		ll.head.prev = newTail
	}
	return tail.value
}


func (ll *LinkedList[T]) printList() {
	for i := 0; i < int(ll.length); i++ {
		val, _ := ll.index(uint(i))
		fmt.Println(val)
	}

	//An alternative I found
	/*
	for currentNode, i := ll.head, 0; i < int(ll.length); currentNode = currentNode.next, i++ {
		fmt.Println(currentNode.value)
	}
	*/

	//I wrote this function before I realized that ll has a length property
	/*
	fmt.Println(ll.head.value)
	if ll.head.next != nil {	
		for currentNode := ll.head.next; true; currentNode = currentNode.next {
			if currentNode != ll.head {
				fmt.Println(currentNode.value)
			}
			else {
				return
			}
		}
	}
	*/
}

func main() {
	ll := LinkedList[int]{}
	ll.append(10)
	ll.append(20)
	ll.append(30)
	ll.printList()
	fmt.Println("Popped:", ll.pop())
	ll.printList()
}
