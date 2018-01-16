package funcPro

import (
	"fmt"
	"testing"
)

//from blog -> http://blog.sina.com.cn/s/blog_6e1bd8350101k4r3.html
func TestFilter(t *testing.T) {
	Students := []*Student{
		{"Danny", 15, 165},
		{"Jacky", 16, 180},
		{"Alan", 17, 172},
		{"Sandy", 18, 168},
	}
	result1 := []*Student{}
	for _, s := range Students {
		if s.Age > 16 {
			result1 = append(result1, s)
		}
	}
	fmt.Println("result1")
	for _, s := range result1 {
		fmt.Println(s.Name)
	}
	result2 := []*Student{}
	for _, s := range Students {
		if (s.Age > 15) && (s.Height > 170) {
			result2 = append(result2, s)
		}
	}
	fmt.Println("result2")
	for _, s := range result2 {
		fmt.Println(s.Name)
	}
}

//测试 函数式
func TestFuncProFilter(t *testing.T) {
	Students := []*Student{
		{"Danny", 15, 165},
		{"Jacky", 16, 180},
		{"Alan", 17, 172},
		{"Sandy", 18, 168},
	}
	result1 := Filter(Students, AgeGreatThanFunc(16))
	fmt.Println("result1")
	printArr(result1)
	result2 := Filter(Students, ComplexFunc(15, 170))
	fmt.Println("result2")
	printArr(result2)
}
func printArr(result1 []*Student) {
	for _, s := range result1 {
		fmt.Println(s.Name)
	}
}
