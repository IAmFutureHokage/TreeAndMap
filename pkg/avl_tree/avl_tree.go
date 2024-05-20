package avltree

import (
	"encoding/json"
	"log"
)

type AVLNode struct {
	Key    int32    `json:"key"`
	Value  any      `json:"value"`
	Left   *AVLNode `json:"left,omitempty"`
	Right  *AVLNode `json:"right,omitempty"`
	Height byte     `json:"height"`
}

func NewAVLNode(key int32, value any) *AVLNode {
	return &AVLNode{
		Key:    key,
		Value:  value,
		Left:   nil,
		Right:  nil,
		Height: 1,
	}
}

func (n *AVLNode) fixHeight() {
	hl, hr := n.Left.height(), n.Right.height()
	if hl > hr {
		n.Height = hl + 1
	} else {
		n.Height = hr + 1
	}
}

func (n *AVLNode) getBalance() int {
	hr, hl := n.Right.height(), n.Left.height()
	return int(hr) - int(hl)
}

func (n *AVLNode) height() byte {
	if n == nil {
		return 0
	}
	return n.Height
}

func (n *AVLNode) rotateRight() *AVLNode {
	if n == nil || n.Left == nil {
		return n
	}
	newRoot := n.Left
	n.Left = newRoot.Right
	newRoot.Right = n
	n.fixHeight()
	newRoot.fixHeight()
	return newRoot
}

func (n *AVLNode) rotateLeft() *AVLNode {
	if n == nil || n.Right == nil {
		return n
	}
	newRoot := n.Right
	n.Right = newRoot.Left
	newRoot.Left = n
	n.fixHeight()
	newRoot.fixHeight()
	return newRoot
}

func (n *AVLNode) Balance() *AVLNode {
	if n == nil {
		return nil
	}
	n.fixHeight()

	balance := n.getBalance()
	if balance > 1 {
		if n.Right != nil && n.Right.getBalance() < 0 {
			n.Right = n.Right.rotateRight()
		}
		return n.rotateLeft()
	}
	if balance < -1 {
		if n.Left != nil && n.Left.getBalance() > 0 {
			n.Left = n.Left.rotateLeft()
		}
		return n.rotateRight()
	}

	return n
}

func (n *AVLNode) insert(key int32, value any) *AVLNode {
	if n == nil {
		return NewAVLNode(key, value)
	}
	if key < n.Key {
		n.Left = n.Left.insert(key, value)
	} else if key > n.Key {
		n.Right = n.Right.insert(key, value)
	} else {
		n.Value = value
	}
	return n.Balance()
}

func (n *AVLNode) findMin() *AVLNode {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n
	}
	return n.Left.findMin()
}

func (n *AVLNode) removeMin() *AVLNode {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n.Right
	}
	n.Left = n.Left.removeMin()
	return n.Balance()
}

func (n *AVLNode) remove(key int32) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.Key {
		n.Left = n.Left.remove(key)
	} else if key > n.Key {
		n.Right = n.Right.remove(key)
	} else {
		if n.Left == nil && n.Right == nil {
			return nil
		}
		if n.Left == nil {
			return n.Right
		}
		if n.Right == nil {
			return n.Left
		}

		minNode := n.Right.findMin()
		n.Key = minNode.Key
		n.Value = minNode.Value
		n.Right = n.Right.removeMin()
	}
	return n.Balance()
}

func (n *AVLNode) find(key int32) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.Key {
		return n.Left.find(key)
	}
	if key > n.Key {
		return n.Right.find(key)
	}
	return n
}

type AVLTree struct {
	Root *AVLNode `json:"root"`
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		Root: nil,
	}
}

func (t *AVLTree) Insert(key int32, value any) {
	t.Root = t.Root.insert(key, value)
}

func (t *AVLTree) Remove(key int32) {
	if t.Root != nil {
		t.Root = t.Root.remove(key)
	}
}

func (t *AVLTree) Find(key int32) any {
	if t.Root == nil {
		return nil
	}
	node := t.Root.find(key)
	if node == nil {
		return nil
	}
	return node.Value
}

func (t *AVLTree) ToJson() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Printf("Error serializing AVL Tree to JSON: %s", err)
		return ""
	}
	return string(jsonData)
}
