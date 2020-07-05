package slice_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Words   string
	Number  int
	Decimal float32
}

func TestSubtractString(t *testing.T) {
	original := []interface{}{"one", "two", "three", "four"}
	subtract := []interface{}{"one", "two", "three"}
	result := Subtract(original, subtract)
	expected := []interface{}{"four"}
	assert.Equal(t, expected, result)
}

func TestSubtractInteger(t *testing.T) {
	original := []interface{}{1, 2, 3, 4}
	subtract := []interface{}{1, 2, 3}
	result := Subtract(original, subtract)
	expected := []interface{}{4}
	assert.Equal(t, expected, result)
}

func TestSubtractFloat(t *testing.T) {
	original := []interface{}{1.1, 2.2, 3.3, 4.4}
	subtract := []interface{}{1.1, 2.2, 3.3}
	result := Subtract(original, subtract)
	expected := []interface{}{4.4}
	assert.Equal(t, expected, result)
}

func TestSubtractObject(t *testing.T) {
	original := []interface{}{testStruct{Words: "one", Number: 1, Decimal: 1.1}, testStruct{Words: "two", Number: 2, Decimal: 2.2}}
	subtract := []interface{}{testStruct{Words: "one", Number: 1, Decimal: 1.1}}
	result := Subtract(original, subtract)
	expected := []interface{}{testStruct{Words: "two", Number: 2, Decimal: 2.2}}
	assert.Equal(t, expected, result)
}

func TestElementsInBoth(t *testing.T) {
	original := []interface{}{"one", "two", "three", "four"}
	subtract := []interface{}{"one", "two", "three"}
	result := ElementsInBoth(original, subtract)
	expected := []interface{}{"one", "two", "three"}
	assert.Equal(t, expected, result)
}

func TestElementsInBothInteger(t *testing.T) {
	original := []interface{}{1, 2, 3, 4}
	subtract := []interface{}{1, 2, 3}
	result := ElementsInBoth(original, subtract)
	expected := []interface{}{1, 2, 3}
	assert.Equal(t, expected, result)
}

func TestElementsInBothFloat(t *testing.T) {
	original := []interface{}{1.1, 2.2, 3.3, 4.4}
	subtract := []interface{}{1.1, 2.2, 3.3}
	result := ElementsInBoth(original, subtract)
	expected := []interface{}{1.1, 2.2, 3.3}
	assert.Equal(t, expected, result)
}

func TestElementsInBothObject(t *testing.T) {
	original := []interface{}{testStruct{Words: "one", Number: 1, Decimal: 1.1}, testStruct{Words: "two", Number: 2, Decimal: 2.2}}
	subtract := []interface{}{testStruct{Words: "one", Number: 1, Decimal: 1.1}}
	result := ElementsInBoth(original, subtract)
	expected := []interface{}{testStruct{Words: "one", Number: 1, Decimal: 1.1}}
	assert.Equal(t, expected, result)
}

func TestPop(t *testing.T) {
	original := []interface{}{1.1, 2.2, 3.3, 4.4}
	result := Pop(original, 1)
	expect := []interface{}{1.1, 3.3, 4.4}
	assert.Equal(t, expect, result)
}
