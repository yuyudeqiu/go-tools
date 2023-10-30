package test

import (
	tree2 "go-tools/pkg/tree"
	"testing"
)

func Test(t *testing.T) {
	tree := tree2.NewBinaryTree[int]()

	tree.Insert(6)
	tree.Insert(12)
	tree.Insert(1)
	tree.Insert(78)
	tree.Insert(16)
	tree.Insert(7)
	tree.Insert(5)

	tree.PreorderTraverse()
	print("\n")
	tree.InorderTraversal()
	print("\n")
	tree.PostOrderTraversal()

	tree2 := tree2.NewBinaryTree[int]()

	tree2.Insert(4)
	tree2.Insert(2)
	tree2.Insert(7)
	tree2.Insert(1)
	tree2.Insert(3)
	tree2.Insert(7)
	tree2.Insert(5)
	tree2.Insert(8)
	tree2.Insert(6)
}
