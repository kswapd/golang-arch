package base

import "fmt"

const (
	red   = true
	black = false
)

type node struct {
	key, value  int
	color       bool
	left, right *node
}

type rbTree struct {
	root *node
}

func newNode(key, value int) *node {
	return &node{key: key, value: value, color: red}
}

func (n *node) flipColor() {
	n.color = !n.color
	n.left.color, n.right.color = !n.left.color, !n.right.color
}

func (n *node) rotateLeft() *node {
	x := n.right
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = red

	return x
}

func (n *node) rotateRight() *node {
	x := n.left
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = red

	return x
}

func isRed(n *node) bool {
	if n == nil {
		return false
	}
	return n.color == red
}

func put(n *node, key, value int) *node {
	if n == nil {
		return newNode(key, value)
	}

	if key < n.key {
		n.left = put(n.left, key, value)
	} else if key > n.key {
		n.right = put(n.right, key, value)
	} else {
		n.value = value
	}

	if isRed(n.right) && !isRed(n.left) {
		n = n.rotateLeft()
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = n.rotateRight()
	}
	if isRed(n.left) && isRed(n.right) {
		n.flipColor()
	}

	return n
}

// Put inserts a key-value pair into the tree and returns the new root node.
func (t *rbTree) Put(key, value int) {
	t.root = put(t.root, key, value)
	t.root.color = black
}

func get(n *node, key int) int {
	for n != nil {
		if key < n.key {
			n = n.left
		} else if key > n.key {
			n = n.right
		} else {
			return n.value
		}
	}
	return -1
}

// Get searches for the value of a given key in the tree.
func (t *rbTree) Get(key int) int {
	return get(t.root, key)
}

func RbTest() {
	tree := &rbTree{}
	tree.Put(3, 30)
	tree.Put(1, 10)
	tree.Put(2, 20)

	fmt.Println(tree.Get(1)) // Output: 10
	fmt.Println(tree.Get(2)) // Output: 20
	fmt.Println(tree.Get(3)) // Output: 30
}
