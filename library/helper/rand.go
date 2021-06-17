package helper

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandIntn(n int) int {
	return rand.Intn(n)
}
