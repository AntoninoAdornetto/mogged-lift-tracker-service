package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomStr(t *testing.T) {
	strLength := 10
	randStr := RandomStr(int64(strLength))
	require.NotEmpty(t, randStr)
	require.True(t, len(randStr) == strLength)

	lastChar := randStr[strLength-1]
	require.NotEqual(t, lastChar, ' ')
	require.NotEqual(t, lastChar, '\n')
	require.NotEqual(t, lastChar, '\r')
}

func TestRandomInt(t *testing.T) {
	min := 1
	max := 10
	for i := 0; i < 100; i++ {
		result := RandomInt(min, max)
		require.True(t, result >= min && result <= max)
	}
}
