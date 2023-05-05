package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomStr(length int64) string {
	var sb strings.Builder

	if length <= 0 {
		return ""
	}

	for i := 0; i < int(length); i++ {
		sb.WriteRune(rune('a' + rand.Intn(26)))
	}

	return sb.String()
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
