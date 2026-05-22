package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{
			name: "Normal size 10",
			size: 10,
			want: 10,
		},
		{
			name: "Normal size 100",
			size: 100,
			want: 100,
		},
		{
			name: "Normal size 1000",
			size: 1000,
			want: 1000,
		},
		{
			name: "Zero size",
			size: 0,
			want: 0,
		},
		{
			name: "Negative size -1",
			size: -1,
			want: 0,
		},
		{
			name: "Negative size -100",
			size: -100,
			want: 0,
		},
		{
			name: "Size 1",
			size: 1,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			assert.Equal(t, tt.want, len(result), "Length mismatch")
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "Normal case with positive numbers",
			data: []int{1, 5, 3, 9, 2, 7},
			want: 9,
		},
		{
			name: "All positive numbers",
			data: []int{10, 20, 30, 40, 50},
			want: 50,
		},
		{
			name: "Single element positive",
			data: []int{42},
			want: 42,
		},
		{
			name: "All equal positive numbers",
			data: []int{7, 7, 7, 7, 7},
			want: 7,
		},
		{
			name: "Maximum at the beginning",
			data: []int{100, 1, 2, 3, 4, 5},
			want: 100,
		},
		{
			name: "Maximum at the end",
			data: []int{1, 2, 3, 4, 5, 100},
			want: 100,
		},
		{
			name: "Empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "Nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "Two elements",
			data: []int{10, 20},
			want: 20,
		},
		{
			name: "Two equal elements",
			data: []int{15, 15},
			want: 15,
		},
		{
			name: "Large positive numbers",
			data: []int{500000, 750000, 1000000, 250000},
			want: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Should not panic for any test case
			got := maximum(tt.data)
			assert.Equal(t, tt.want, got, "Maximum mismatch for data: %v", tt.data)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "Normal case with positive numbers",
			data: []int{1, 5, 3, 9, 2, 7, 4, 8, 6, 10, 11, 12, 13, 14, 15, 16},
			want: 16,
		},
		{
			name: "All positive numbers",
			data: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160},
			want: 160,
		},
		{
			name: "Single element",
			data: []int{42},
			want: 42,
		},
		{
			name: "All equal numbers",
			data: []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
			want: 7,
		},
		{
			name: "Maximum at the beginning",
			data: []int{100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			want: 100,
		},
		{
			name: "Maximum at the end",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 100},
			want: 100,
		},
		{
			name: "Size smaller than CHUNKS",
			data: []int{5, 3, 8, 1, 2},
			want: 8,
		},
		{
			name: "Exactly CHUNKS elements",
			data: []int{10, 20, 30, 40, 50, 60, 70, 80},
			want: 80,
		},
		{
			name: "Empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "Nil slice",
			data: nil,
			want: 0,
		},
		{
			name: "Size 1 with CHUNKS",
			data: []int{999},
			want: 999,
		},
		{
			name: "Large positive numbers",
			data: []int{500000, 750000, 1000000, 250000, 125000, 875000, 625000, 375000},
			want: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Should not panic for any test case
			got := maxChunks(tt.data)
			assert.Equal(t, tt.want, got, "MaxChunks mismatch for data: %v", tt.data)
		})
	}
}
