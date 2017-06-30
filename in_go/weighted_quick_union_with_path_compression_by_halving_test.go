package main

import (
	"path/filepath"
	"testing"
)

func TestWeightedQuickUnionWithPathCompressionByHalving(t *testing.T) {
	testDataDir := "../test_data/"
	testFiles := []string{
		"graphs/input.json",
		"graphs/input2.json",
	}
	for _, file := range testFiles {
		t.Run(file, func(t *testing.T) {
			gr := loadGraph(filepath.Join(testDataDir, file))
			input := gr.Input
			result := WeightedQuickUnionWithPathCompressionByHalving(len(input), input)
			if !edgesEqual(result, gr.Result) {
				t.Fatalf("For %v\nexpected %v\ngot %v\n", input, gr.Result, result)
			}
		})
	}
}
