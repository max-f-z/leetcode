package main

import (
	"reflect"
	"testing"
)

func Test3607(t *testing.T) {
	tests := []struct {
		name        string
		c           int
		connections [][]int
		queries     [][]int
		expected    []int
	}{
		{
			name: "all connected stations with offline queries",
			c:    5,
			connections: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			queries: [][]int{
				{1, 3}, // Query station 3 (online) -> return 3
				{2, 1}, // Mark station 1 offline
				{1, 1}, // Query station 1 (offline) -> return smallest available (2)
				{2, 2}, // Mark station 2 offline
				{1, 2}, // Query station 2 (offline) -> return smallest available (3)
			},
			expected: []int{3, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processQueries(tt.c, tt.connections, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("processQueries() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
