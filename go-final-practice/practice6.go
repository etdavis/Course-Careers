/* Question 6
In Go, write a function named flattenBST that accepts a pointer to the root node of a binary
search tree (BST) and two integer values min and max. The function should return a slice of
integers representing a flattened, inorder traversal of the binary search tree, but only including
the nodes whose values fall within the inclusive range min to max.
A tree is "flattened" by performing an inorder traversal, which visits the left subtree, the root
node, and then the right subtree. The function should return a slice that contains the values of
the nodes in the order they were visited.
The function signature in Go is: func flattenBST(root *TreeNode, min int, max int) []int
The TreeNode struct is defined as follows:

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
} */

//go:build practice6
//run with $ go run -tags practice6 practice6.go
package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

// Got this implementation from Tutorialspoint and I'm commenting it up to process and understand it
func flattenBST(root *TreeNode, min int, max int) []int {
	result := make([]int, 0) // the slice we're going to return
   	stack := make([]*TreeNode, 0) // holds our nodes as we wait to process them
   	curr := root

   	for curr != nil || len(stack) > 0 {
      	for curr != nil {
        stack = append(stack, curr) // adds current node to stack
        curr = curr.Left // moves to the left node
      	} //exits once the current node is nil meaning there's no left node
      	curr = stack[len(stack)-1] // moves the current node back to the parent
      	stack = stack[:len(stack)-1] // removes the current node from the stack
		if curr.Value >= min && curr.Value <= max { // check if current value is within specified range
			result = append(result, curr.Value) // appends current node's value to returning array
		}
      	curr = curr.Right // since current is processed, move to right node
   	}
   	return result
}

func main() {
	root1 := &TreeNode{
		Value: 20,
		Left: &TreeNode{
			Value: 10,
			Left: &TreeNode{
				Value: 5,
			},
			Right: &TreeNode{
				Value: 15,
			},
		},
		Right: &TreeNode{
			Value: 30,
			Right: &TreeNode{
				Value: 35,
			},
		},
	}
	min1 := 10
	max1 := 30

	fmt.Println(flattenBST(root1, min1, max1))
}