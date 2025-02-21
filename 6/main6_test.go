package main

import "testing"

func TestRandomGenerator(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		max     int
		wantLen int
	}{
		{"generate 10 numbers", 10, 100, 10},
		{"generate 0 numbers", 0, 100, 0},
		{"generate 1 numbers", 1, 100, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RandomGenerator(tt.n, tt.max)

			count := 0
			for _ = range result {
				count++
			}

			if count != tt.wantLen {
				t.Errorf("Expected len %d, got %d", count, tt.wantLen)
			}

		})
	}
}
