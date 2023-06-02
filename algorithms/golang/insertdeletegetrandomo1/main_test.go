package insertdeletegetrandomo1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedSet(t *testing.T) {
	randSet := Constructor()

	assert.Equal(t, true, randSet.Insert(1))
	assert.Equal(t, false, randSet.Insert(1))
	assert.Equal(t, true, randSet.Insert(2))
	assert.Equal(t, true, randSet.Remove(1))
	assert.Equal(t, false, randSet.Remove(1))
	assert.Equal(t, true, randSet.Remove(2))
}
