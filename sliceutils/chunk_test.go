package sliceutils

import "testing"

func TestChunkIndexes(t *testing.T) {
	tests := []struct {
		array          []int
		size           int
		expectedChunks [][]int
	}{
		{
			array:          []int{},
			size:           2,
			expectedChunks: [][]int{},
		},
		{
			array:          []int{0},
			size:           2,
			expectedChunks: [][]int{{0}},
		},
		{
			array:          []int{0, 1},
			size:           2,
			expectedChunks: [][]int{{0, 1}},
		},
		{
			array:          []int{0, 1, 2},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2}},
		},
		{
			array:          []int{0, 1, 2, 3},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2, 3}},
		},
		{
			array:          []int{0, 1, 2, 3, 4},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2, 3}, {4}},
		},
		{
			array:          []int{0, 1, 2, 3, 4, 5},
			size:           2,
			expectedChunks: [][]int{{0, 1}, {2, 3}, {4, 5}},
		},
	}
	for _, test := range tests {
		indexes := ChunkIndexes(len(test.array), test.size)
		var chunks [][]int
		for i := 0; i < len(indexes)-1; i++ {
			chunks = append(chunks, test.array[indexes[i]:indexes[i+1]])
		}
		if len(chunks) != len(test.expectedChunks) {
			t.Fatalf("%v expected chunks is %v, but is %v", test.array, test.expectedChunks, chunks)
		}
		for i := range chunks {
			if len(chunks[i]) != len(test.expectedChunks[i]) {
				t.Fatalf("%v expected chunks is %v, but is %v", test.array, test.expectedChunks, chunks)
			}
			for j := range chunks[i] {
				if chunks[i][j] != test.expectedChunks[i][j] {
					t.Fatalf("%v expected chunks is %v, but is %v", test.array, test.expectedChunks, chunks)
				}
			}
		}
		t.Log(chunks)
	}

}
