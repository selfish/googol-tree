package util

import (
	"math/rand"
	"sort"
	"time"
)

func MinIntSlice(intSlice []int) int {
	sort.Ints(intSlice)
	return intSlice[0]
}

func MaxIntSlice(intSlice []int) int {
	sort.Ints(intSlice)
	return intSlice[len(intSlice)-1]
}

func RandomIntSlice(sliceSize int, maxValue int) []int {
	slice := make([]int, sliceSize, sliceSize)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < sliceSize; i++ {
		slice[i] = rand.Intn(maxValue)
	}
	return slice
}

func CountWins(boards []bool) int {
	wins := 0
	for _, DidWin := range boards {
		if DidWin == true {
			wins++
		}
	}
	return wins
}