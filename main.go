package main

import (
	"fmt"
	"math"
	"math/rand"
)

// GCDEuclidean calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	return a
}

// GCDRemainderRecursive calculates GCD recursively using remainder.
func GCDRemainderRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDRemainderRecursive(b, a%b)
}

// GCDRemainder calculates GCD iteratively using remainder.
func GCDRemainder(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func calculate_pi() chan float64 {
	bound := 1000
	times := 1000000000
	coprimes := 0
	count := 0
	c := make(chan float64)
	go func() {
		for i := 0; i < times; i++ {
			x1 := rand.Intn(bound)
			x2 := rand.Intn(bound)
			gcd := GCDRemainder(x1, x2)
			if gcd == 1 {
				coprimes++
			} else {
				count++
			}
			approx := math.Sqrt((float64(coprimes) / float64(count)) * 6)
			if i%10000 == 0 {
				c <- approx
			}
		}
		close(c)
	}()
	return c
}

func main() {
	c := calculate_pi()
	for a := range c {
		fmt.Println(a)
	}
	fmt.Println("done")
}
