package examples

import "fmt"

type Command string

const (
	Queue   Command = "Queue"
	UnQueue Command = "UnQueue"
)

type Instruction struct {
	msg string
	c   Command
}

func logic(arr []string, ins Instruction) []string {
	if arr == nil {
		arr = make([]string, 0)
	}
	if ins.c == Queue {
		arr = append(arr, ins.msg)
		fmt.Println("Queue: ", ins.msg)
	} else if ins.c == UnQueue {
		if len(arr) > 0 {
			rs := arr[0]
			arr = arr[1:]
			fmt.Println("UnQueue: ", rs)
		}
	}
	return arr
}

func GoBlockchanLogic() {
	// init state 1
	arr := make([]string, 0)
	instructions := []Instruction{
		Instruction{msg: "1", c: Queue},
		Instruction{msg: "2", c: Queue},
		Instruction{msg: "x", c: UnQueue},
		Instruction{msg: "x", c: UnQueue},
		Instruction{"x", UnQueue},
		Instruction{"3", Queue},
		Instruction{"4", Queue},
	}
	for _, item := range instructions {
		arr = logic(arr, item)
	}
	fmt.Println(arr)
}
