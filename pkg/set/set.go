package set

import (
	"fmt"
	"go-tools/pkg/Interface/iterator/iteratorble"
)

type Iterator struct { //适用于set集合的基础迭代器，因为没办法调用包外的结构体作为接收者，只能再封装一层
	*iteratorble.BaseIterator
	set *Set
}

// 使用go语言自带的map实现set,由于golang的map是用hash实现的，故这个set是hashSet
// 此set非线程安全
// 已实现迭代器接口

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

func (firstSet *Set) Union(secondSet Set) { //求Set并集
	for key := range secondSet.set {
		if !firstSet.Contains(key) {
			firstSet.Add(key)
		}
	}
}

func (firstSet *Set) Intersection(secondSet Set) { //求Set交集
	newSet := &Set{}
	newSet.set = make(map[any]nothing)
	for key := range secondSet.set {
		if firstSet.Contains(key) {
			newSet.Add(key)
		}
	}
	*firstSet = *newSet
}

func (set *Set) getBaseIterator() *iteratorble.BaseIterator { //获得基础迭代器内容
	retBaseIterator := &iteratorble.BaseIterator{}

	elements := make([]any, 0, len(set.set))
	for element := range set.set {
		elements = append(elements, element)
	}

	retBaseIterator.HasNextAfterRemove = false
	retBaseIterator.Elements = elements
	retBaseIterator.Index = 0

	return retBaseIterator
}

func (set *Set) Iterator() *Iterator { //获得set的迭代器
	baseIterator := set.getBaseIterator()
	retI := Iterator{baseIterator, set}
	return &retI
}

func (its *Iterator) Next() (any, bool) {
	if its.BaseIterator.Index >= len(its.BaseIterator.Elements) {
		return nil, false
	}

	currentElement := its.BaseIterator.Elements[its.BaseIterator.Index]
	its.BaseIterator.Index++
	its.BaseIterator.HasNextAfterRemove = true
	return currentElement, true
}

func (its *Iterator) Remove() bool {
	if its.BaseIterator.HasNextAfterRemove == true {
		indexOfRemoveElement := its.BaseIterator.Index - 1
		if indexOfRemoveElement >= 0 && its.BaseIterator.Index < len(its.BaseIterator.Elements) && its.BaseIterator.HasNextAfterRemove {
			if its.set.Delete(its.Elements[indexOfRemoveElement]) == false {
				return false
			}
			its.BaseIterator.Elements = append(its.BaseIterator.Elements[:indexOfRemoveElement], its.BaseIterator.Elements[indexOfRemoveElement+1:]...)
			return true
		}
	}
	return false
}

func (its *Iterator) HasNext() bool {
	return its.BaseIterator.Index < len(its.BaseIterator.Elements)
}
