package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertTo(t *testing.T) {

	palabra := "AC03AAA"
	resultado := NewResult("AC", 3, "AAA")
	r := ConverterTo(palabra)
	assert.Equal(t, r, resultado, "la conversion fallo")

}

func TestContainsCharValid(t *testing.T) {

	palabra := "$$"
	assert.Equal(t, false, containsCharValid(palabra), "Fallo la comparacion de caracteres palabra1 ")

	palabra1 := "AA"
	assert.Equal(t, true, containsCharValid(palabra1), "la conversion fallo p2")

	palabra2 := "AB02AA"
	assert.Equal(t, false, containsCharValid(palabra2), "la conversion fallo p3")

}
