package main

import (
	"testing"
)

func TestUnzip1(t *testing.T) {
	result := "aaaabccddddde"
	unzipped, err := UnzipString("a4bc2d5e")
	if err != nil {
		t.Errorf("Error wasn't nil")
	}
	if unzipped != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", unzipped, result)
	}
}

func TestUnzip2(t *testing.T) {
	result := "abcd"
	unzipped, err := UnzipString("abcd")
	if err != nil {
		t.Errorf("Error wasn't nil")
	}
	if unzipped != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", unzipped, result)
	}
}

func TestUnzip3(t *testing.T) {
	result := ""
	unzipped, err := UnzipString("45")
	if err == nil {
		t.Errorf("Error is nil")
	}
	if unzipped != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", unzipped, result)
	}
}

func TestUnzip4(t *testing.T) {
	result := ""
	unzipped, err := UnzipString("")
	if err != nil {
		t.Errorf("Error wasn't nil")
	}
	if unzipped != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", unzipped, result)
	}
}

func TestUnzipEscape1(t *testing.T) {
	result := "qwe45"
	unzipped, err := UnzipString("qwe\\4\\5")
	if err != nil {
		t.Errorf("Error wasn't nil")
	}
	if unzipped != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", unzipped, result)
	}
}

func TestUnzipEscape2(t *testing.T) {
	result := "qwe44444"
	unzipped, err := UnzipString("qwe\\45")
	if err != nil {
		t.Errorf("Error wasn't nil")
	}
	if unzipped != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", unzipped, result)
	}
}

func TestUnzipEscape3(t *testing.T) {
	result := "qwe\\\\\\\\\\"
	unzipped, err := UnzipString("qwe\\\\5")
	if err != nil {
		t.Errorf("Error wasn't nil")
	}
	if unzipped != result {
		t.Errorf("Result was incorrect, got: %s, want: %s.", unzipped, result)
	}
}
