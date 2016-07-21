package collections

import "fmt"

type LinkedNode struct {
	next *LinkedNode
	item interface{}
}

type LinkedList struct {
	first *LinkedNode
	last  *LinkedNode
	size  int
}

func (this *LinkedList) Insert(elem interface{}) {
	var node = &LinkedNode{nil, elem}
	if this.size == 0 {
		this.first = node
	} else {
		copy := *this.first
		node.next = &copy
		this.first = node
	}
	this.size++
}

func (this *LinkedList) Pop() (interface{}, error) {
	if this.size == 0 {
		return 0, fmt.Errorf("Cannot pop from empty list")
	}
	ret := this.first.item
	this.first = this.first.next
	this.size--
	return ret, nil
}

func (this *LinkedList) Iterate(f func(interface{})) {

	for elem := this.first; elem != nil; elem = elem.next {
		f(elem.item)
	}
}

func (this *LinkedList) Contains(e interface{}) bool {
	for elem := this.first; elem != nil; elem = elem.next {

		if elem.item == e {
			return true
		}
	}
	return false
}

func (this *LinkedList) Filter(p func(interface{}) bool) *LinkedList {
	nextLinked := LinkedList{}
	for elem := this.first; elem != nil; elem = elem.next {

		if p(elem.item) {
			nextLinked.Insert(elem.item)
		}
	}
	return &nextLinked
}
