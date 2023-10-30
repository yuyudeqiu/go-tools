package tree

import "go-tools/pkg/interface/orderedtype"

type Node[T orderedtype.Ordered] struct { //二叉树，节点存储的数据的类型最好是统一的，因此加上[]修饰符
	Content T //T 是comparable类型，可以直接用比较符号比较
	Left    *Node[T]
	Right   *Node[T]
}

type BinaryTree[T orderedtype.Ordered] interface {
	Insert(T)
	Search(T) bool
	Delete(T) bool
}
