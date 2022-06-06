package quadratic

import (
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	res1, res2 := 0.0, 0.0
	temp := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	res1 = (temp - b) / (2 * a)
	res2 = -(temp + b) / (2 * a)
	return res1, res2
}
