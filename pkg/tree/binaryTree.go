package tree

type Node[T any] struct { //二叉树，节点存储的数据的类型最好是统一的，因此需要一个统一类型
	content T
	left    *Node[T]
	right   *Node[T]
}
