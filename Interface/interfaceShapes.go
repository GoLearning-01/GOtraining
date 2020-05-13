package main

import "fmt"

type shapes interface {
	area() float64
	perimeter() float64
}

type triangle struct {
	length float64
	height float64
}

type square struct {
	length float64
}

func (t triangle) area() float64 {
	return 0.5 * t.length * t.height
}

func (s square) area() float64 {
	return s.length * s.length
}

func (t triangle) perimeter() float64 {
	return t.length * 3
}

func (s square) perimeter() float64 {
	return s.length * 4
}

func measure(g shapes) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {

	s := square{5}
	t := triangle{2, 6}

	fmt.Println("")
	measure(s)
	measure(t)

	fmt.Println("")
	fmt.Println(s.area())
	fmt.Println(s.perimeter())

	fmt.Println("")
	fmt.Println(t.area())
	fmt.Println(t.perimeter())

}
