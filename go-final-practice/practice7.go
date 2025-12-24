/* Question 7 (No Test Cases)
In Go, create a generic Set[T comparable] struct that represents a mathematical set. The struct
should contain a map to store the elements of the set.
Implement the following methods for the Set[T comparable] struct:
	- NewSet() *Set[T]: This function should return a pointer to an empty Set[T].
	- Add(element T): This method should add an element to the set. If the element already
	exists in the set, the method should have no effect (recall that sets do not have duplicate
	elements).
	- Remove(element T): This method should remove an element from the set. If the element
	does not exist, the method should have no effect.
	- Contains(element T) bool: This method should return true if the element exists in the set
	and false otherwise.
	- Size() int: This method should return the number of elements in the set.
The Set[T comparable] struct is generic and can be used with any comparable type T. This
means it should work with types like int, string, etc. However, keep in mind that slices and maps
are not comparable in Go, so they cannot be used as T.
Remember to make sure all the methods are safe to use even if they are called in the "wrong"
order. For example, calling Remove on an element that has not been added to the set should
not cause a runtime error. */

//go:build practice7
//run with $ go run -tags practice7 practice7.go
package main

import (
	"fmt"
)

// included peeks at solution
type Set[T comparable] struct {
	elements map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: map[T]bool{}}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = true
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func main() {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(3)
	s.Remove(2)
	fmt.Println(s.Contains(1)) // true
	fmt.Println(s.Contains(2)) // false
	fmt.Println(s.Size()) // 2
}