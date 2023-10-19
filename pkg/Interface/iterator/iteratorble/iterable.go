package iteratorble

type BaseIterator struct { //基础迭代器
	Index              int   //迭代器目前处在的位置
	Elements           []any //迭代器内部遍历用的数组
	HasNextAfterRemove bool  //在上次Remove后是否next
}

type Iterable interface { //继承这个接口表示可迭代
	GetBaseIterator() *BaseIterator //获得实现该接口的对象的基础迭代器
}
