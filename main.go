package main

import (
	"fmt"
	"math"
)

func Bisection(f func(float64) float64, a, b, eps float64) float64 {
	var c float64
	var done bool
	for !done {
		c = (a + b) / 2
		if f(c) == 0 || (b-a)/2 < eps {
			return c
		}
		if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return a
}

func main() {
	fmt.Println(Bisection(func(x float64) float64 { return x*x*x - 2*x - 5 }, 0, 4, 0.0001))
	fmt.Println(Bisection(func(x float64) float64 { return 3*x*x*x*x - 4*x*x*x - 12*x*x - 5 }, 0, 4, 0.0001))
	fmt.Println(Bisection(func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 0, 4, 0.0001))
	fmt.Println(Bisection(func(x float64) float64 { return x - 5*math.Sin(x) - 3.5 }, 0, 4, 0.0001))

}
