package hashmap

import (
	"encoding/json"
	"log"
)

type hashKey int32

func (key hashKey) Hash() int {
	hash := int(key)
	if hash < 0 {
		return -hash
	}
	return hash
}

type Node struct {
	Key   int32 `json:"key"`
	Value any   `json:"value"`
	Next  *Node `json:"next,omitempty"`
}

func NewNode(key int32, value any, next *Node) *Node {
	return &Node{key, value, next}
}

type HashMap struct {
	Buckets []*Node `json:"buckets,omitempty"`
}

func NewHashMap(size int32) *HashMap {
	return &HashMap{
		Buckets: make([]*Node, size),
	}
}

func (hm *HashMap) Insert(key int32, value any) {
	bucketIndex := hashKey(key).Hash() % len(hm.Buckets)
	node := hm.Buckets[bucketIndex]
	for node != nil {
		if node.Key == key {
			node.Value = value
			return
		}
		node = node.Next
	}
	hm.Buckets[bucketIndex] = NewNode(key, value, hm.Buckets[bucketIndex])
}

func (hm *HashMap) Find(key int32) any {
	bucketIndex := hashKey(key).Hash() % len(hm.Buckets)
	node := hm.Buckets[bucketIndex]
	for node != nil {
		if node.Key == key {
			return node.Value
		}
		node = node.Next
	}
	return nil
}

func (hm *HashMap) Remove(key int32) {
	bucketIndex := hashKey(key).Hash() % len(hm.Buckets)
	nodePrev := hm.Buckets[bucketIndex]
	if nodePrev != nil {
		if nodePrev.Key == key {
			hm.Buckets[bucketIndex] = nodePrev.Next
			return
		}
		node := nodePrev.Next
		for node != nil {
			if node.Key == key {
				nodePrev.Next = node.Next
				return
			}
			nodePrev = node
			node = node.Next
		}
	}
}

func (hm *HashMap) ToJson() string {
	jsonData, err := json.MarshalIndent(hm, "", "  ")
	if err != nil {
		log.Printf("Error serializing Hash Map to JSON: %s", err)
		return ""
	}
	return string(jsonData)
}
