package collections

import "fmt"

type iterator func(int, interface{})

type treeNode struct {
	left  *treeNode
	right *treeNode
	key   int
	value interface{}
}

type Tree struct {
	root *treeNode
}

func (t *Tree) Insert(k int, v interface{}) {
	node := &treeNode{left: nil, right: nil, key: k, value: v}
	if t.root == nil {
		t.root = node
		return
	}
	t.root = t.insert_recursive(k, v, t.root)
}

func (t *Tree) insert_recursive(k int, v interface{}, node *treeNode) *treeNode {
	if node == nil {
		node = &treeNode{left: nil, right: nil, key: k, value: v}
		return node
	}
	if node.key == k {
		return nil
	}
	if node.key < k {
		node.left = t.insert_recursive(k, v, node.left)
		return node
	}
	node.right = t.insert_recursive(k, v, node.right)
	return node
}

func (t *Tree) Get(k int) (interface{}, error) {
	if t.root == nil {
		return nil, fmt.Errorf("Empty list")
	}
	return t.get_recursive(k, t.root)
}

func (t *Tree) get_recursive(k int, node *treeNode) (interface{}, error) {
	if node == nil {
		return nil, fmt.Errorf("No such element")
	}

	if node.key == k {
		return node.value, nil
	}
	if node.key < k {
		return t.get_recursive(k, node.left)
	}
	return t.get_recursive(k, node.right)
}

func (t *Tree) Iterate(f iterator) {
	iterate(f, t.root)
}

func iterate(f iterator, node *treeNode) {
	if node == nil {
		return
	}
	f(node.key, node.value)
	defer iterate(f, node.left)
	defer iterate(f, node.right)
}
