package listIterator

import "go-tools/pkg/interface/iterator/baseIterator"

type ListInterface interface { //有序集合的迭代器接口,在无序迭代器的基础上增加了若干方法
	baseIterator.BasicInterface
	set()               //替换迭代器前面的元素
	add()               //添加一个元素在迭代器前面
	nextIndex() int     //返回下一个元素的索引
	previousIndex() int //返回上一个元素的索引
}
