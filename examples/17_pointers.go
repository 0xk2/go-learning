package examples

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func swap(pa *int, pb *int) {
	fmt.Println("value of pa: ", pa, "; value of pb: ", pb)
	fmt.Println("address of pa: ", &pa, "; address of pb: ", &pb)
	tmp := *pb
	*pb = *pa
	*pa = tmp
}

func GoPointers() {
	i := 1
	fmt.Println("initial: ", i)
	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)

	a := 1
	b := 3
	fmt.Printf("before swap a: %d, b: %d \n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap a: %d, b: %d \n", a, b)
}
