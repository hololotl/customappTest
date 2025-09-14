package algorithms

import (
	"fmt"
	"math"
	"math/rand"
)

func GetPar(rtp float64) float64 {
	U := rand.Float64()
	if U < 1-rtp {
		return 1.0
	}
	V := rand.Float64()
	m := 1 / (1.0 - V)
	return m
}

func TestPar() {
	var MAE float64  // absolute error
	var MAPE float64 // Mean Absolute Percentage Error
	var PWT float64  // Percent within tolerance ∣y−x∣≤ t

	tryes := 100
	for i := 0; i < tryes; i++ {
		rtp := rand.Float64()
		size := 100000
		arr := make([]float64, size)
		for i := 0; i < size; i++ {
			arr[i] = 9000 + rand.Float64()*(1000)
		}
		res := comparepar(arr, rtp)
		MAE += math.Abs(rtp - res)
		MAPE += math.Abs(rtp-res) / rtp
		if math.Abs(rtp-res) < rtp*0.15 {
			PWT += 1
		}
	}
	fmt.Println("Constant. Mean absolute error:", MAE/float64(tryes))
	fmt.Println("Constant. Percent within tolerance", PWT/float64(tryes)*100)
	fmt.Println("Constant. Mean Absolute Percentage Error", MAPE/float64(tryes)*100)
}

func comparepar(arr []float64, rtp float64) float64 {
	var sum float64
	c := 0
	for i := 0; i < len(arr); i++ {
		m := GetPar(rtp)
		if m > arr[i] {
			sum += arr[i]
			c += 1
		}
	}
	return sum / float64(len(arr))
}
