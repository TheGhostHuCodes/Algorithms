package main

import (
	"encoding/json"
	"log"
	"os"
)

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
