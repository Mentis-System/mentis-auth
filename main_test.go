package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	assert.Equal(t, "Hello", GetMessage(), "Error Test")
}
