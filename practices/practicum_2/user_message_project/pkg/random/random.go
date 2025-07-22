package random


import (
	"math/rand/v2"
)

func randRange(min, max int) int {
    return rand.IntN(max-min) + min
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.IntN(len(letterBytes))]
    }
    return string(b)
}
