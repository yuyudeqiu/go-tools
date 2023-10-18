package iterator

type iterable interface {
	getIterator() //获得实现该接口的对象的迭代器
}
