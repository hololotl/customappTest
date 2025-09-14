package internal

import (
	"math/rand"
)

type ParetoAlgorithm struct {
	rtp float64
}

func NewParetoAlgorithm(rtp float64) *ParetoAlgorithm {
	return &ParetoAlgorithm{rtp: rtp}
}

func (k *ParetoAlgorithm) GetMultiplier() float64 {
	U := rand.Float64()
	if U < 1-k.rtp {
		return 1.0
	}
	V := rand.Float64()
	m := 1 / (1.0 - V)
	return m
}
