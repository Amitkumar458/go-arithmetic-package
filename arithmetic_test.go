package arithmetic

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3, 4)
	expected := 9.0
	if result != expected {
		t.Errorf("Add(2, 3, 4) = %v; want %v", result, expected)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(6, 4)
	expected := 2.0
	if expected != result {
		t.Errorf("Substract(6, 4) = %v; want %v", result, expected)
	}
}
