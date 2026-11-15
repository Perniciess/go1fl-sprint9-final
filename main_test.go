package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name       string
		size       int
		wantLength int
	}{
		{name: "отрицательный размер", size: -1, wantLength: 0},
		{name: "нулевой размер", size: 0, wantLength: 0},
		{name: "один элемент", size: 1, wantLength: 1},
		{name: "много элементов", size: 10000, wantLength: 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.wantLength {
				t.Fatalf("Длина слайса: получено %d, ожидалось %d", len(data), tt.wantLength)
			}
			for i, value := range data {
				if value < 1 || value > randomRange {
					t.Errorf("Элемент с индексом %d равен %d, ожидалось значение от 1 до %d", i, value, randomRange)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil-слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 3, 1}, want: 9},
		{name: "максимум в середине", data: []int{2, 9, 1}, want: 9},
		{name: "максимум в конце", data: []int{2, 3, 9}, want: 9},
		{name: "одинаковые элементы", data: []int{5, 5, 5}, want: 5},
		{name: "повторяющийся максимум", data: []int{9, 1, 9}, want: 9},
		{name: "нулевые элементы", data: []int{0, 0, 0}, want: 0},
		{name: "отрицательные элементы", data: []int{-5, -1, -9}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("Максимум слайса %v: получено %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil-слайс", data: nil, want: 0},
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "меньше восьми элементов", data: []int{1, 7, 3, 4, 5, 6, 2}, want: 7},
		{name: "ровно восемь элементов", data: []int{1, 2, 3, 8, 5, 6, 7, 4}, want: 8},
		{name: "максимум в остатке", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, want: 99},
		{name: "одинаковые элементы", data: []int{5, 5, 5, 5, 5, 5, 5, 5, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("Максимум по частям слайса %v: получено %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}
