package util

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomStr(t *testing.T) {
	strLength := rand.Intn(20)
	randStr := RandomStr(int64(strLength))
	require.NotEmpty(t, randStr)
	require.True(t, len(randStr) == strLength)

	// does not end with an \s or empty char
	require.NotNil(t, randStr[strLength-1])
	require.True(t, rune(randStr[strLength-1]) != rune(' '))
}

func TestRandomInt(t *testing.T) {
	max := 10
	for i := 0; i < 50; i++ {
		result := RandomInt(max)
		assert.True(t, result >= 0 && result < max)
	}
}
