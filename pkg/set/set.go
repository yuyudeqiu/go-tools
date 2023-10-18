package set

import (
	"fmt"
)

// 使用go语言自带的map实现set,由于golang的map是用hash实现的，故这个set是hashSet
// 此set非线程安全
// 没有实现迭代器接口，仅测试
type IteratorSet struct { //set集合的迭代器
	index int
	set   *Set
}

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

func (set *Set) Len() int { //获得set的长度
	return len(set.set)
}

func (set *Set) Size() int {
	return len(set.set)
}

func (set *Set) Clear() {
	set.set = make(map[any]nothing)
}

func mapToArray(set *Set) {

}

func (firstSet *Set) Union(secondSet Set) { //求Set并集
	for key := range secondSet.set {
		if !firstSet.Contains(key) {
			firstSet.Add(key)
		}
	}
}

func (firstSet *Set) Intersection(secondSet Set) { //求Set交集
	newSet := &Set{}
	for key := range secondSet.set {
		if firstSet.Contains(key) {
			newSet.Add(key)
		}
	}
	firstSet = newSet
}

func (set *Set) GetIterator() *IteratorSet { //获得set接口的迭代器
	retIterator := &IteratorSet{}
	retIterator.set = set
	retIterator.index = 0

	return retIterator
}
