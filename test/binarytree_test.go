package test

import (
	tree2 "go-tools/pkg/tree"
	"testing"
)

func Test(t *testing.T) {
	//tree := tree2.NewBinaryTree[int]()
	//
	//tree.Insert(6)
	//tree.Insert(12)
	//tree.Insert(1)
	//tree.Insert(78)
	//tree.Insert(16)
	//tree.Insert(7)
	//tree.Insert(5)
	//
	//tree.PreorderTraverse()
	//print("\n")
	//tree.InorderTraversal()
	//print("\n")
	//tree.PostOrderTraversal()

	treeTwo := tree2.NewBinaryTree[int]()

	treeTwo.Insert(4)
	treeTwo.Insert(2)
	treeTwo.Insert(7)
	treeTwo.Insert(1)
	treeTwo.Insert(3)
	treeTwo.Insert(7)
	treeTwo.Insert(5)
	treeTwo.Insert(8)
	treeTwo.Insert(6)

	treeTwo.PreorderTraverse()
	treeTwo.Delete(7)
	print("\n")
	treeTwo.PreorderTraverse()

}
