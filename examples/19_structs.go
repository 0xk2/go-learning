package examples

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// return the point p; no new memory will be used after this
	return &p
}

// read this: https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values
func newPersonWithoutPointer(name string) person {
	p := person{name: name}
	p.age = 10
	// return a copy of p; new memory will be needed and &p will be clean (auto?)
	return p
}

/*
- Methods using receiver pointers are common; the rule of thumb for receivers is, "If in doubt, use a pointer."
- Slices, maps, channels, strings, function values, and interface values are implemented with pointers internally, and a pointer to them is often redundant.
- Elsewhere, use pointers for big structs or structs you'll have to change, and otherwise pass values, because getting things changed by surprise via a pointer is confusing.
*/

func GoStructs() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	j := newPerson("Jon")
	fmt.Println(j)
	fmt.Println(j.name)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	k := newPersonWithoutPointer("kiddie")
	k.name = "hommie"
	k.age = 7
	fmt.Println(k.name, k.age)
}
