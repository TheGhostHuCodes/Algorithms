package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type graph struct {
	Input  []edge
	Result []edge
}

func loadGraph(path string) graph {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open JSON file.\nError: %s\n", err)
	}
	dec := json.NewDecoder(f)
	var g graph
	err = dec.Decode(&g)
	f.Close()
	if err != nil {
		log.Fatalf("Could not decode JSON file.\nError: %s\n", err)
	}
	return g
}

func main() {
	testDataPath := "../test_data/"
	gr := loadGraph(filepath.Join(testDataPath, "graphs/input.json"))
	input := gr.Input
	QuickFind(len(input), input)
	QuickUnion(len(input), input)

	gr2 := loadGraph(filepath.Join(testDataPath, "graphs/input2.json"))
	input2 := gr2.Input
	QuickFind(len(input2), input2)
	QuickUnion(len(input2), input2)
}
