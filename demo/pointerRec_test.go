package demo

import (
	"fmt"
	"math"
	"testing"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Belarger(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Belarger1(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Distance() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 指针 receiver
func TestPointerMethod(t *testing.T) {
	v := Vertex{3.0, 4.0}
	fmt.Println(" Original: ", v, v.Distance())
	fmt.Println("----------------------------")

	v.Belarger1(5.0)
	fmt.Println(" no pointer method: ", v, v.Distance())
	fmt.Println("----------------------------")

	v.Belarger(5.0)
	fmt.Println(" pointer method: ", v, v.Distance())
	fmt.Println("----------------------------")
}
