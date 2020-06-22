package chapter2

import (
	"fmt"
	"testing"
)

func TestBinarySqrt_MySqrt(t *testing.T) {
	bs := &BinarySqrt{}
	result := bs.MySqrt(5, 10)
	fmt.Printf("%f\n", result)
}

func TestBinarySqrt_MySqrt2(t *testing.T) {
	bs := &BinarySqrt{}
	result := bs.accuracy(5)
	fmt.Printf("%.5f\n", result)
}
