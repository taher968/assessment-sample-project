package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {

	assert.Equal(t, 10, addTwoNumbers(5, 10))

}
