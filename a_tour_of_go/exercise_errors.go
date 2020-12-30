package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	 return fmt.Sprintf("cannot Sqrt negative number:%v", float64(e))
}

func Sqrt(x float64) (float64, error){
	var err error
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(2.)
	s := float64(0)
	for i := 0; i < 10; i++ {
			z = z - (z*z - x)/(2*z)
			if math.Abs(z - s) < 1e-10 {
					break;
			}
			s = z
	}
	return z, err
}

func main() {
		v, err := Sqrt(-4)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
}
