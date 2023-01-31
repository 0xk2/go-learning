package examples

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		fmt.Println("tail is nil")
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		fmt.Println("tail not nil")
		lst.tail.next = &element[T]{val: v}
		// lst.tail = lst.tail.next
		lst.tail = &element[T]{val: v}
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func GoGenerics() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys: ", MapKeys(m))

	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(11)
	lst.Push(12)
	lst.Push(13)
	fmt.Println("list:", lst.GetAll())
}
