package main

import (
	"fmt"
)

type rect struct {
	width  int
	height int
	name   string
}

func (r rect) funcPerim() int {
	return 2*r.height + 2*r.width
}

func vals() int {
	return 3
}

func multiVals() (int, string) {
	return 7, "rah"
}

func main() {
	a := vals()
	fmt.Println("A IS :: ", a)
	a1, _ := multiVals()
	_, b1 := multiVals()
	fmt.Println("A IS :: ", a1, b1)

	r := rect{width: 10, height: 20, name: "rect1"}

	fmt.Println("rect : ", r.funcPerim())

}
