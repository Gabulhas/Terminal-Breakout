package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func floatRange(min, max float64) float64 {

	return (min + rand.Float64()*(max-min))

}

func intRange(min, max int) int {

	return (rand.Intn(max-min+1) + min)

}
