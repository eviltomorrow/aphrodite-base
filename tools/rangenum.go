package tools

import (
	"math/rand"
	"time"
)

// GenerateRangeNum generate range num
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
