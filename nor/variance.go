package nor

import "math"

func Variance(data []int, mean float64) float64 {
	var sumdef float64
	datalen := float64(len(data))
	// get the sum of (n-mean)2
	for _, n := range data {
		sumdef += math.Pow(float64(n)-mean, 2.0)
	}
	// sum of (n-mean)2 / len
	variance := sumdef / datalen
	return variance
}