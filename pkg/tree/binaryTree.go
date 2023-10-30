package tree

import (
	"go-tools/pkg/interface/orderedtype"
	"go-tools/pkg/interface/tree"
)

/*
写二叉树是为了写二叉搜索树，写二叉搜索树是为了写红黑树，写红黑树是为了写map，哈哈哈套娃
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

func NewBinaryTree[T orderedtype.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{nil}
}

func (bt *BinaryTree[T]) Insert(data T) {
	if bt.root == nil {
		bt.root = &tree.Node[T]{Content: data}
		return
	} else {
		bt.insertWithoutRecursion(data, bt.root)
	}
}

//递归调用的插入，量大的时候性能太差，写着玩
/*func (bt *BinaryTree[T]) insert(data T, node *tree.Node[T]) {
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
}*/

func (bt *BinaryTree[T]) insertWithoutRecursion(data T, node *tree.Node[T]) {
	newNode := &tree.Node[T]{Content: data}

	curr := node
	for {
		if data < curr.Content {
			if curr.Left == nil {
				curr.Left = newNode
				return
			}
			curr = curr.Left

		} else if data > curr.Content {
			if curr.Right == nil {
				curr.Right = newNode
				return
			}
			curr = curr.Right
		} else { //二叉树遇到相同的值就没有必要添加了
			return
		}
	}
}

func (bt *BinaryTree[T]) PreorderTraverse() { //先序遍历需要用到栈这种结构
	if bt.root == nil {
		return
	}

	node := bt.root

	stack := []*tree.Node[T]{node} //让栈中先有根节点
	//由于先序遍历是根节点-》左节点-》右节点

	for len(stack) > 0 {
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		print(currNode.Content, " ")

		if currNode.Right != nil {
			stack = append(stack, currNode.Right)
		}

		if currNode.Left != nil {
			stack = append(stack, currNode.Left)
		}
	}
}

func (bt *BinaryTree[T]) InorderTraversal() {
	if bt.root == nil {
		return
	}
	var stack []*tree.Node[T]
	currentNode := bt.root

	for len(stack) > 0 || currentNode != nil {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
		}

		if currentNode == nil {
			currentNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			print(currentNode.Content, " ")
			currentNode = currentNode.Right
		}

	}
}

func (bt *BinaryTree[T]) PostOrderTraversal() {
	if bt.root == nil {
		return
	}

	var stack []*tree.Node[T]
	var lastRead *tree.Node[T]
	currNode := bt.root
	for len(stack) > 0 || currNode != nil {
		for currNode != nil {
			stack = append(stack, currNode)
			currNode = currNode.Left
		}

		currNode = stack[len(stack)-1]

		if currNode.Right == nil || currNode.Right == lastRead {
			stack = stack[:len(stack)-1]
			print(currNode.Content, " ")
			lastRead = currNode
			currNode = nil
		} else {
			currNode = currNode.Right
		}
	}
}

func (bt *BinaryTree[T]) Clear() { //清空树节点
	bt.root = nil
}

func (bt *BinaryTree[T]) exist(data T) {

}

func (bt *BinaryTree[T]) Delete(data T) {
	deleteNode[T](bt.root, data)
}

func deleteNode[T orderedtype.Ordered](node *tree.Node[T], data T) *tree.Node[T] {
	if node == nil {
		return nil
	}
}
