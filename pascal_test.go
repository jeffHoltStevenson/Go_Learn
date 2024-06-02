
package main

import (
	"reflect"
	"testing"
)

// generates correct first row of Pascal's triangle
func TestPascalTriangleFirstRow(t *testing.T) {
	expected := []int{1}
	result := pascal_triangle(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestPascalTriangleSecondRow(t *testing.T) {
	expected := []int{1,1}
	result := pascal_triangle(2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestPascalTriangleThirdRow(t *testing.T) {
	expected := []int{1,2,1}
	result := pascal_triangle(3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestPascalTriangleFourthRow(t *testing.T) {
	expected := []int{1,3,3,1}
	result := pascal_triangle(4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestPascalTriangleFifthRow(t *testing.T) {
	expected := []int{1,4,6,4,1}
	result := pascal_triangle(5)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestPascalTriangleSixthRow(t *testing.T) {
	expected := []int{1,5,10,10,5,1}
	result := pascal_triangle(6)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}