package learning

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type that implements the Shape interface
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	//TODO implement me
	panic("implement me")
}

func (c Circle) Perimeter() float64 {
	//TODO implement me
	panic("implement me")
}

// Rectangle type that implements the Shape interface
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	//TODO implement me
	panic("implement me")
}

func (r Rectangle) Perimeter() float64 {
	//TODO implement me
	panic("implement me")
}

func calculateArea(shape interface{}) {
	switch s := shape.(type) {
	case Circle:
		fmt.Printf("Circle area: %.2f\n", s.Area())
	case Rectangle:
		fmt.Printf("Rectangle area: %.2f\n", s.Area())
	default:
		fmt.Println("Unknown shape")
	}
}

func interfaces() {
	var s Shape

	s = Circle{radius: 5}
	fmt.Println("C Area:", s.Area())
	fmt.Println("C Perimeter:", s.Perimeter())

	s = Rectangle{length: 4, width: 3}
	fmt.Println("R Area:", s.Area())
	fmt.Println("R Perimeter:", s.Perimeter())

	c := Circle{radius: 5}
	r := Rectangle{length: 4, width: 3}
	calculateArea(c)
	calculateArea(r)

}
