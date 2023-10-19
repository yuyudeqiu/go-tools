package test

import (
	"fmt"
	"testing"
)
import "go-tools/pkg/set"

func TestSet(t *testing.T) {
	mySet := set.NewSet()
	mySet.Add(1)
	mySet.Add("new")
	mySet.Add("new")
	mySet.Add("new2")
	mySet.Add(true)

	mySet.Print()
	fmt.Printf("mySet.Len():%v \n", mySet.Len())
	fmt.Printf("mySet.Contains(\"new\"):%v \n", mySet.Contains("new"))
	fmt.Printf("mySet.Contains(\"good\"):%v \n", mySet.Contains("good"))

	mySet.Delete("new")
	fmt.Printf("after mySet.Delete(\"new\"),mySet.Contains(\"new\"):%v \n", mySet.Contains("new"))
	fmt.Printf("mySet.Len():%v \n", mySet.Len())
	mySet.Print()

	fmt.Print("==============================================================================\n\n")

	newSet := set.NewSet()
	newSet.Add(false)
	mySet.Print()
	newSet.Print()

	fmt.Printf("after union: ")
	mySet.Union(*newSet)
	mySet.Print()
}

func TestIteratorSet(t *testing.T) { //测试set集合的迭代器
	newSet := set.NewSet()
	newSet.Add(1)
	newSet.Add(2)
	newSet.Add(3)
	newSet.Add(4)
	newSet.Add(5)
	newSet.Add(6)
	newSet.Add(7)
	newSet.Add(8)
	newSet.Add(9)

	it := newSet.Iterator()
	newSet.Print()

	for it.HasNext() {
		element, _ := it.Next()
		if element == 3 || element == 7 {
			it.Remove()
		}
	}

	newSet.Print()

}
