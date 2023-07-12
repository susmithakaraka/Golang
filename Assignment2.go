package main

import "fmt"

type values struct {
	a, b int
}

type calculations interface {
	substraction() int
	addition() int
	multiplication() int
	divison() float64
}

func (v values) substraction() int {

	fmt.Println("substraction of a and b is:")
	x := &(v.a)
	return *x - v.b
}

func (v values) addition() int {

	fmt.Println("addition of a and b is:")
	return v.a + v.b
}

func (v values) multiplication() int {

	fmt.Println("multiplication of a and b is:")
	return v.a * v.b
}
func (v values) divison() float64 {
	fmt.Println("divison of a and b is:")
	return float64(v.a / v.b)
}

func cal(c calculations) {

	fmt.Println(c)
	fmt.Println(c.substraction())
	fmt.Println(c.addition())
	fmt.Println(c.multiplication())
	fmt.Println(c.divison())
}
func main() {
	v := values{a: 10, b: 5}
	cal(v)
}
