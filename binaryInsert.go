package main

import "fmt"

type Node struct {
	value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) insert(value int) {
	bst.Root = insertNode(bst.Root, value)
}

func insertNode(node *Node, value int) *Node {

	if node == nil {
		return &Node{value: value}

	}

	if value < node.value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.value {
		node.Right = insertNode(node.Right, value)
	}
	return node

}

func (bst BST) inorderTraversal(node *Node) {

	if node != nil {
		bst.inorderTraversal(node.Left)
		fmt.Print(node.value, " ")
		bst.inorderTraversal(node.Right)
	}
}

func main() {

	bst := &BST{}
	bst.insert(5)
	bst.insert(3)
	bst.insert(7)

	fmt.Println("Inorder Traversal of BST:")
	bst.inorderTraversal(bst.Root)

}
