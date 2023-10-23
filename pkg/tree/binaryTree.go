package tree

import (
	"go-tools/pkg/interface/orderedtype"
	"go-tools/pkg/interface/tree"
)

/*
关于二叉树递归和非递归哪个好，问了一下gpt
在某些特定情况下，使用递归进行二叉树插入可能会比非递归插入更好。以下是一些情况：
小规模数据：对于小规模的数据集，递归插入通常更加简洁和易于实现。递归的代码结构更易于理解和维护，在这种情况下，递归的性能可能更好。
平衡二叉树：在平衡二叉树（如AVL树、红黑树等）的插入操作中，采用递归实现通常更为简单和清晰。递归可以自然地处理树的旋转和平衡操作，避免了繁琐的代码和条件逻辑。
动态树结构：对于动态修改频繁的树结构，如伸展树、Treap树等，递归插入可以更好地展现其动态性。通过递归，在插入节点时可以自然地调整树的结构和平衡，使其保持高效的操作。

总结：其实不嫌麻烦就是非递归的好。因为非递归性能好，递归仅仅是易于理解。
*/

type BinaryTree[T orderedtype.Ordered] struct {
	root *tree.Node[T] //为了确保二叉树内类型一致，需要对二叉树进行封装
}

func (bt *BinaryTree[T]) Insert(data T) {
	if bt.root == nil {
		bt.root = &tree.Node[T]{Content: data}
		return
	} else {
		bt.insert(data, bt.root)
	}
}

func (bt *BinaryTree[T]) insert(data T, node *tree.Node[T]) { //递归调用的插入，性能太差，没用
	if data < node.Content {
		if node.Left == nil {
			node.Left = &tree.Node[T]{Content: data}
		} else {
			bt.insert(data, node.Left)
		}
	} else {
		if node.Right == nil {
			node.Right = &tree.Node[T]{Content: data}
		} else {
			bt.insert(data, node.Right)
		}
	}
}

func (bt *BinaryTree[T]) Traverse() {

}
