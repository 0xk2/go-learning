package examples

import "fmt"

// type State struct {
// 	started bool
// }

// func (s *State) start() {
// 	s.started = true
// }

// func (s *State) stop() {
// 	s.started = false
// }

// func (s State) isStarted() bool {
// 	return s.started
// }

// type Engine struct { // what data an engine can hold
// 	manufacturer string
// 	horsePower   int
// 	state        State
// }

type Engine struct { // what data an engine can hold
	manufacturer string
	horsePower   int
	started      bool
}

func (e *Engine) start() {
	e.started = true
}

func (e *Engine) stop() {
	e.started = false
}

func (e Engine) isStarted() bool {
	return e.started
}

type Vehicle struct { // what data a vehicle can hold
	noOfSeat  int
	noOfWheel int
	engine    Engine
}

type Project struct {
	name    string
	started bool
}

func (p *Project) start() {
	p.started = true
}

func (p *Project) stop() {
	p.started = false
}

func (p Project) isStarted() bool {
	return p.started
}

func newVehicle(engineManufacture string, horsePower int, noOfSeat int, noOfWheel int) *Vehicle {
	v := Vehicle{
		noOfSeat:  noOfSeat,
		noOfWheel: noOfWheel,
		engine: Engine{
			manufacturer: engineManufacture,
			horsePower:   horsePower,
			started:      false,
		},
	}
	return &v
}

func GoStructEmbedding() {
	myCar := newVehicle("Huyndai", 1500, 4, 4)
	myCar.engine.start()
	fmt.Printf("is %s engine started? %t \n", myCar.engine.manufacturer, myCar.engine.isStarted())

	p := Project{"Hectagon Chain", false}

	type startable interface {
		isStarted() bool
		start()
		stop()
	}

	var s1 startable = &myCar.engine
	var s2 startable = &p
	fmt.Printf("Access myCar engine status through s1: %t\n", s1.isStarted())
	s2.start()
	s2.stop()
	fmt.Printf("Access project status through s2: %t\n", s2.isStarted())
}
