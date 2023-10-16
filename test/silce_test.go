package test

import (
	"testing"

	"go-tools/pkg/slice"
)

func TestList_Add(t *testing.T) {
	list := slice.NewList[int]()
	list.Add(1)
	list.Print()
	list.Add(2)
	list.Print()
	list.Add(3)
	list.Print()

	list.Delete(1)
	list.Print()

}
