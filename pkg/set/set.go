package set

import "fmt"

//使用go语言自带的map实现set,由于golang的map是用hash实现的，故这个set是hashSet
//此set非线程安全
//没有实现迭代器接口，仅测试

type nothing struct{}

type Set struct {
	set map[any]nothing //set不对外开放,map的value是空结构体不占空间
}

func NewSet() *Set {
	retSet := &Set{}
	retSet.set = make(map[any]nothing)

	return retSet
}

func (set *Set) Print() {
	fmt.Printf("set contont:")
	for k := range set.set {
		fmt.Printf("(%T)%v, ", k, k)
	}
	fmt.Printf("\n")
}

func (set *Set) Add(element any) {
	set.set[element] = struct{}{}
}

func (set *Set) Contains(element any) bool {
	_, ok := set.set[element]
	return ok
}

func (set *Set) Delete(element any) bool {
	if !set.Contains(element) {
		return false
	}
	delete(set.set, element)
	return true
}

func (set *Set) Len() int {
	return len(set.set)
}
