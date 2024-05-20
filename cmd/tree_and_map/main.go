package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	avltree "tree_and_map/pkg/avl_tree"
	hashmap "tree_and_map/pkg/hash_map"
)

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	keys := make([]int32, 10000)
	for i := range keys {
		keys[i] = rng.Int31n(20001) - 10000
	}

	tree := avltree.NewAVLTree()
	hashMap := hashmap.NewHashMap(10000)

	start := time.Now()
	for _, key := range keys {
		tree.Insert(key, key)
	}
	elapsedTreeInsert := time.Since(start)
	start = time.Now()
	for _, key := range keys {
		hashMap.Insert(key, key)
	}
	elapsedMapInsert := time.Since(start)

	start = time.Now()
	for _, key := range keys {
		tree.Find(key)
	}
	elapsedTreeFind := time.Since(start)
	start = time.Now()
	for _, key := range keys {
		hashMap.Find(key)
	}
	elapsedMapFind := time.Since(start)

	saveData(tree, "./output/tree_output.json")
	saveData(hashMap, "./output/map_output.json")

	start = time.Now()
	for _, key := range keys {
		tree.Remove(key)
	}
	elapsedTreeRemove := time.Since(start)
	start = time.Now()
	for _, key := range keys {
		hashMap.Remove(key)
	}
	elapsedMapRemove := time.Since(start)

	saveData(tree, "./output/tree_output_empty.json")
	saveData(hashMap, "./output/map_output_empty.json")

	fmt.Printf("AVL Tree - Insert: %s, Find: %s, Remove: %s\n", elapsedTreeInsert, elapsedTreeFind, elapsedTreeRemove)
	fmt.Printf("HashMap - Insert: %s, Find: %s, Remove: %s\n", elapsedMapInsert, elapsedMapFind, elapsedMapRemove)
}

func saveData(data interface{ ToJson() string }, filePath string) {
	jsonOutput := data.ToJson()
	if jsonOutput == "" {
		fmt.Println("Error serializing data to JSON")
		return
	}
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, err = file.WriteString(jsonOutput)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
