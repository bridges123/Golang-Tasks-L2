package main

import (
	"fmt"
	"testing"
)

func TestFindAnagramsByDict1(t *testing.T) {
	expected := map[string][]string{
		"пятак":  {"пятак", "пятка", "тяпка"},
		"листок": {"листок", "слиток", "столик"},
	}
	given := []string{"пятак", "листок", "пятка", "слиток", "столик", "тяпка"}
	result := FindAnagramsByDict(given)
	fmt.Println(result)

	for k, expValue := range expected {
		if resValue, ok := result[k]; !ok {
			t.Errorf("The expected key is missing %s", k)
		} else {
			if len(expValue) != len(resValue) {
				t.Errorf("Incorrect length of sets, got: %d, expected: %d.", len(resValue), len(expValue))
			}
			for i := range expValue {
				if expValue[i] != resValue[i] {
					t.Errorf("Incorrect element of set №%d, got: %s, expected: %s.", i, expValue[i], resValue[i])
				}
			}
		}
	}
}

func TestFindAnagramsByDict2(t *testing.T) {
	expected := map[string][]string{
		"маниша": {"маниша", "машина", "минаша"},
	}
	given := []string{"МАшина", "МИНАША", "влага", "манишА"}
	result := FindAnagramsByDict(given)

	for k, expValue := range expected {
		if resValue, ok := result[k]; !ok {
			t.Errorf("The expected key is missing %s", k)
		} else {
			if len(expValue) != len(resValue) {
				t.Errorf("Incorrect length of sets, got: %d, expected: %d.", len(resValue), len(expValue))
			}
			for i := range expValue {
				if expValue[i] != resValue[i] {
					t.Errorf("Incorrect element of set №%d, got: %s, expected: %s.", i, expValue[i], resValue[i])
				}
			}
		}
	}
}
