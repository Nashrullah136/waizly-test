package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinOfFour(t *testing.T) {
	data := []struct {
		name     string
		input    []int64
		expected int64
	}{
		{name: "normal test", input: []int64{1, 2, 3, 4, 5}, expected: 10},
		{name: "int 32 overflow test", input: []int64{4e9, 5e9, 6e9, 7e9, 8e9}, expected: 22e9},
	}
	for _, datum := range data {
		t.Run(datum.name, func(t *testing.T) {
			assert.Equal(t, minOfFour(datum.input), datum.expected)
		})
	}
}

func TestMaxOfFour(t *testing.T) {
	data := []struct {
		name     string
		input    []int64
		expected int64
	}{
		{name: "normal test", input: []int64{1, 2, 3, 4, 5}, expected: 14},
		{name: "int 32 overflow test", input: []int64{4e9, 5e9, 6e9, 7e9, 8e9}, expected: 26e9},
	}
	for _, datum := range data {
		t.Run(datum.name, func(t *testing.T) {
			assert.Equal(t, maxOfFour(datum.input), datum.expected)
		})
	}
}

func TestSum(t *testing.T) {
	data := []struct {
		name     string
		input    []int64
		expected int64
	}{
		{name: "normal test", input: []int64{1, 2, 3, 4, 5}, expected: 15},
		{name: "int 32 overflow test", input: []int64{4e9, 5e9, 6e9, 7e9, 8e9}, expected: 30e9},
	}
	for _, datum := range data {
		t.Run(datum.name, func(t *testing.T) {
			assert.Equal(t, sum(datum.input...), datum.expected)
		})
	}
}
