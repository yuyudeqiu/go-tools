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

	mySet.Print()
	fmt.Printf("mySet.Len():%v \n", mySet.Len())
	fmt.Printf("mySet.Contains(\"new\"):%v \n", mySet.Contains("new"))
	fmt.Printf("mySet.Contains(\"good\"):%v \n", mySet.Contains("good"))

	mySet.Delete("new")
	fmt.Printf("after mySet.Delete(\"new\"),mySet.Contains(\"new\"):%v \n", mySet.Contains("new"))
	fmt.Printf("mySet.Len():%v \n", mySet.Len())

}
