package random

import (
	"math/rand"

	"go.x2ox.com/sorbifolia/coarsetime"
)

type MathRand struct {
	randBytes    []byte
	randBytesLen int
}

func init() {
	rand.Seed(coarsetime.FloorTime().UnixNano())
}

func NewMathRand() RandString {
	return &MathRand{
		randBytes:    []byte(randBytes),
		randBytesLen: randBytesLen,
	}
}

func (r MathRand) RandString(length int) string {
	arr := make([]byte, length)
	for i := range arr {
		arr[i] = r.randBytes[rand.Intn(r.randBytesLen)]
	}
	return string(arr)
}

func (r MathRand) SetRandBytes(data []byte) RandString {
	if len(data) > 256 {
		panic("data too long")
	}
	if hasRepeat(data) {
		panic("not repeatable")
	}
	r.randBytes = data
	r.randBytesLen = len(data)
	return r
}
