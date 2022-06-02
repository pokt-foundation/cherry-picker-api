package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandom_GetRandomHex(t *testing.T) {
	c := require.New(t)

	length := 20

	randomHexA, err := RandomHex(length)
	c.NoError(err)

	c.True(len(randomHexA) == length)

	randomHexB, err := RandomHex(length)
	c.NoError(err)
	c.Len(randomHexB, length)

	// This is non deterministic and could technically be equal due to entropy
	// but the possibility is so low is not worth the effort to prevent that.
	c.NotEqual(randomHexA, randomHexB)
}
