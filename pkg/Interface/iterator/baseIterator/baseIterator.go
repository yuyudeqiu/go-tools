package baseIterator

type BasicInterface interface { //无序集合的迭代器接口
	HasNext() bool     //是否有下一个元素
	Next() (any, bool) //前往下一个元素，返回该元素,如果到达了集合的末尾，bool改为false
	Remove() bool      //删除上一个遍历的元素
}
