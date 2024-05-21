package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFractionOfZero(t *testing.T) {
	data := []struct {
		name     string
		data     []int
		expected float64
	}{
		{name: "test 1", data: []int{-4, 3, -9, 0, 4, 1}, expected: 1.0 / 6.0},
		{name: "test 2", data: []int{1, 1, 0, -1, -1}, expected: 1.0 / 5.0},
	}
	for _, datum := range data {
		t.Run(datum.name, func(t *testing.T) {
			assert.Equal(t, fractionOfZero(datum.data), datum.expected)
		})
	}
}

func TestFractionOfNegative(t *testing.T) {
	data := []struct {
		name     string
		data     []int
		expected float64
	}{
		{name: "test 1", data: []int{-4, 3, -9, 0, 4, 1}, expected: 2.0 / 6.0},
		{name: "test 2", data: []int{1, 1, 0, -1, -1}, expected: 2.0 / 5.0},
	}
	for _, datum := range data {
		t.Run(datum.name, func(t *testing.T) {
			assert.Equal(t, fractionOfNegative(datum.data), datum.expected)
		})
	}
}

func TestFractionOfPositive(t *testing.T) {
	data := []struct {
		name     string
		data     []int
		expected float64
	}{
		{name: "test 1", data: []int{-4, 3, -9, 0, 4, 1}, expected: 3.0 / 6.0},
		{name: "test 2", data: []int{1, 1, 0, -1, -1}, expected: 2.0 / 5.0},
	}
	for _, datum := range data {
		t.Run(datum.name, func(t *testing.T) {
			assert.Equal(t, fractionOfPositive(datum.data), datum.expected)
		})
	}
}
