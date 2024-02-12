package myuuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	u1 := New()
	u2 := New()

	assert.NotEqual(t, u1, u2)
}
