package main

import (
	"testing"
)

func TestReadLines1(t *testing.T) {
	expected := []string{
		"I was a child",
		"Good 21st is nice",
		"Good 21st now",
		"1239 842 21",
		"10000",
	}
	result := readLines("1.txt")

	if len(expected) != len(result) {
		t.Errorf("Incorrect read lines count, got: %d, expected: %d.", len(result), len(expected))
	}
	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Incorrect line №%d, got: %s, expected: %s.", i, expected[i], result[i])
		}
	}
}

func TestReadLines2(t *testing.T) {
	expected := []string{
		"",
	}
	result := readLines("2.txt")

	if len(expected) != len(result) {
		t.Errorf("Incorrect read lines count, got: %d, expected: %d.", len(result), len(expected))
	}
	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Incorrect line №%d, got: %s, expected: %s.", i, expected[i], result[i])
		}
	}
}

func TestSortLines1(t *testing.T) {
	expected := []string{
		"10000",
		"1239 842 21",
		"Good 21st is nice",
		"Good 21st now",
		"I was a child",
	}

	flags := map[byte]int{}
	lines := []string{
		"I was a child",
		"Good 21st now",
		"1239 842 21",
		"10000",
		"Good 21st is nice",
	}
	sortLines(lines, flags)

	for i := range expected {
		if expected[i] != lines[i] {
			t.Errorf("Incorrect line №%d, got: %s, expected: %s.", i, expected[i], lines[i])
		}
	}
}

func TestSortLines2(t *testing.T) {
	expected := []string{
		"I was a child",
		"Good 21st now",
		"Good 21st is nice",
		"1239 842 21",
		"10000",
	}

	flags := map[byte]int{
		'r': 1,
	}
	lines := []string{
		"I was a child",
		"Good 21st now",
		"1239 842 21",
		"10000",
		"Good 21st is nice",
	}
	sortLines(lines, flags)

	for i := range expected {
		if expected[i] != lines[i] {
			t.Errorf("Incorrect line №%d, got: %s, expected: %s.", i, expected[i], lines[i])
		}
	}
}

func TestSortLines3(t *testing.T) {
	expected := []string{
		"10000",
		"1239 842 21",
		"I was a child",
		"Good 21st is nice",
		"Good 21st now",
	}

	flags := map[byte]int{
		'k': 2,
	}
	lines := []string{
		"I was a child",
		"Good 21st now",
		"1239 842 21",
		"10000",
		"Good 21st is nice",
	}
	sortLines(lines, flags)

	for i := range expected {
		if expected[i] != lines[i] {
			t.Errorf("Incorrect line №%d, got: %s, expected: %s.", i, expected[i], lines[i])
		}
	}
}

func TestSortLines4(t *testing.T) {
	expected := []string{
		"Good 21st is nice",
		"Good 21st now",
		"I was a child",
		"1239 842 21",
		"10000",
	}

	flags := map[byte]int{
		'n': 1,
	}
	lines := []string{
		"I was a child",
		"Good 21st now",
		"1239 842 21",
		"10000",
		"Good 21st is nice",
	}
	sortLines(lines, flags)

	for i := range expected {
		if expected[i] != lines[i] {
			t.Errorf("Incorrect line №%d, got: %s, expected: %s.", i, expected[i], lines[i])
		}
	}
}

func TestSortLines5(t *testing.T) {
	expected := []string{
		"10000",
		"1239 842 21",
		"Good 21st is nice",
		"Good 21st now",
		"I was a child",
	}

	flags := map[byte]int{
		'u': 1,
	}
	lines := []string{
		"I was a child",
		"Good 21st is nice",
		"Good 21st now",
		"1239 842 21",
		"10000",
		"Good 21st is nice",
	}
	sortLines(lines, flags)
	cutRepeatableLines(lines)

	for i := range expected {
		if expected[i] != lines[i] {
			t.Errorf("Incorrect line №%d, got: %s, expected: %s.", i, expected[i], lines[i])
		}
	}
}
