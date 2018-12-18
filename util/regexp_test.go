package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexp(t *testing.T) {
	value := "create namespace arbol on service"
	exp := "^create namespace [a-zA-Z_.-]+ on [a-zA-Z]+$"
	_, err := Match(exp, value)
	assert.Nil(t, err)
	//assert.Equal(t, 2, len(res))
}
