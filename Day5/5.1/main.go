package main

import (
	"fmt"

)

var pi float64= 3.141

type shape interface {
	area() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	l, b float64
}

func (c *circle) area() float64{

	return pi*(float64(c.radius))*(float64(c.radius))
}

func (r *rectangle)area() float64{

	return (r.l*r.b)
}

func (c * circle) String() string {

	return fmt.Sprintf("Circle has radius of : %v cm",c.radius)
}

func (r* rectangle) String() string {

	return fmt.Sprintf("Rectangle has sides : L = %v and B = %v cm",r.l,r.b)
}

func cal(s shape) float64{

	switch s.(type){
    case *circle:
		fmt.Println("Area of Cicle is : ")
		return s.area()
	case *rectangle:
		fmt.Println("Area of Rectangle is : ")
		return s.area()	

	}

	return 0.0
}

func main(){

	var C shape
	var R shape

	C=&circle{10.2}

	R=&rectangle{l:20.2,b:20.2}

	fmt.Println("---------- Interfaces ----------")
	fmt.Println("Details of Shapes :\n")
	fmt.Println(C,"\n",R,"\n")

	fmt.Println(cal(C))
	fmt.Println(cal(R))

}
