package examples

import "fmt"

type mrect struct {
	width, height int
}

func (r *mrect) area() int {
	return r.width * r.height
}

/*
Methods can be defined for either pointer or value receiver types.
Here’s an example of a value receiver
*/
func (r mrect) perim() int {
	return 2 * r.width * r.height
}

func (r *mrect) doubleWidth() {
	r.width *= 2
}

func (r mrect) doubleWidthWithValueReceiver() {
	r.width *= 2
}

func GoMethods() {
	r := mrect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perum: ", r.perim())

	rp := &r
	/**
	 Go automatically handles conversion between values and pointers for method calls.
	 You may want to use a pointer receiver type to avoid:
	 - copying on method calls or
	 - to allow the methods to mutate (change) the receiving struct
	**/
	fmt.Println("area: ", rp.area())
	fmt.Println("area: ", rp.perim())

	fmt.Printf("original width: %d, then double with pointer receiver, and ", r.width)
	r.doubleWidth()
	fmt.Printf("we have: %d \n", r.width)

	fmt.Printf("Now, double again with value receiver, and ")
	r.doubleWidthWithValueReceiver()
	fmt.Printf("we have: %d - mean nothing change \n", r.width)
}
