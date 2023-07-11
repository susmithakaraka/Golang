package main

import "fmt"

type emp struct {
	id     int
	name   string
	salary int
}

func main() {

	//Arithmetic Operators
	var a int = 2
	var b int = 6
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Addition:", a+b)
	fmt.Println("Substraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", b/a)
	fmt.Println("Remainder:", b%a)
	fmt.Println()

	//Logical Operators
	var x bool = true
	var y bool = false
	fmt.Println("Logical Operators:")
	if !(x && y) {
		fmt.Println("AND")
	}
	if x || y {
		fmt.Println("OR")
	}
	fmt.Println()
	//Relational Operators
	var c int = 6
	fmt.Println("Relational Operators:")
	if b == c {
		fmt.Println("equals to")
	}
	if a != b {
		fmt.Println("Not equal to")
	}
	if a < b && a < c {
		fmt.Println("Less than")
	}
	if a > b || a > c {
		fmt.Println("greater than")
	}
	fmt.Println()

	//for loop with slice
	var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	fmt.Println("for loop with slice")
	for i, j := range months {
		fmt.Println("the index is", i, "and the value is", j)
	}
	fmt.Println()

	//structure
	fmt.Println("Structure of an employee")
	var e1 emp
	e1.id = 20153
	e1.name = "John"
	e1.salary = 5000
	fmt.Println(e1)

	var e2 = emp{20154, "Alan", 4600}
	fmt.Println(e2)

}
