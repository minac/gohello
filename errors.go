package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z := 1.0
		for i := 0; i < 10; i++ {
			y := (z*z - x) / (2 * z)
			if math.Abs(z-math.Sqrt(x)) < 0.01 {
				fmt.Println("break", "iter", i, ": math.Abs(z-math.Sqrt(x))", math.Abs(z-math.Sqrt(x)))
				break
			} else {
				z = z - y
			}
			fmt.Println("iter", i, ": y", y, "result", z)
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
