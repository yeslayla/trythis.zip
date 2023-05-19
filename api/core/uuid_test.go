package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	assert := assert.New(t)

	one := GenerateID()
	assert.Greater(len(one), 0)

	two := GenerateID()
	assert.Greater(len(two), 0)

	assert.NotEqual(two, one, "Duplicate ID generated!")
}
