package examples

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

/*** Code for lib ***/
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 52 letter, 52 = 110100b => keep 6 smallest bit, zer0 the rest
	letterIdxMask = 1<<letterIdxBits - 1 // 1b move right 6 pos -> 1000000, -1 and turn into 0111111
)
const (
	letterIdxMax = 63 / letterIdxBits
)

func RandString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 { // remain = 0 -> init new random number and assign to cache and restart remain
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		// shift cache 6 bit to the right
		cache = cache >> letterIdxBits
		remain--
	}
	return string(b)
}

/*** End code for lib ***/

/** Code for Blockchain Logic **/
// NOTE: do not use json in real life example
type Command string
type Status string

const (
	Create Command = "Create"
	Start  Command = "Start"
	Stop   Command = "Stop"
	Pause  Command = "Pause"
	Finish Command = "Finish"
)

const (
	JustCreated Status = "JustCreated"
	Doing       Status = "Doing"
	Paused      Status = "Paused"
	Stopped     Status = "Stopped"
	Finished    Status = "Finished"
)

type CreateInstruction struct {
	Title string
	Desc  string
}
type StartInstruction struct {
	Id             string
	EstDayToFinish int
}
type StopInstruction struct {
	Id     string
	Reason string
}
type PauseInstruction struct {
	Id         string
	EstWaitDay int
}
type FinishInstruction struct {
	Id             string
	CongratMessage string
}

type Instruction struct {
	Id   string
	C    Command
	Data []byte
}

type Block struct {
	BlockHash    string
	Instructions []Instruction
}

type Task struct {
	Id     string
	Title  string
	Desc   string
	Status Status
}

type Event struct {
	Id            string
	BlockHash     string
	InstructionId string
	Msg           string
}

var events = make([]Event, 0)

func logic(mem []byte, block Block) []byte {
	var tasks []Task
	newMem := mem
	err := json.Unmarshal(mem, &tasks)
	if err != nil {
		return mem
	}
	instructions := block.Instructions
	blockHash := block.BlockHash
	for _, ins := range instructions {
		switch ins.C {
		case Create:
			var createInstruction *CreateInstruction
			err = json.Unmarshal(ins.Data, &createInstruction)
			if err != nil {
				tasks = append(tasks, Task{
					Id:     RandString(8),
					Title:  createInstruction.Title,
					Desc:   createInstruction.Desc,
					Status: JustCreated,
				})
				emitEvent(blockHash, ins.Id, fmt.Sprintf("Create Task{%s, %s}", createInstruction.Title, createInstruction.Desc))
				newMem, _ = json.Marshal(tasks)
			}
			break
		case Start:
			var startInstrucion *StartInstruction
			err = json.Unmarshal(ins.Data, &startInstrucion)
			if err != nil {
				t := findTask(tasks, startInstrucion.Id)
				if t != nil && (t.Status == JustCreated || t.Status == Paused) {
					t.Status = Doing
					newMem, _ = json.Marshal(tasks)
					emitEvent(blockHash, ins.Id, fmt.Sprintf("Start Task #%s(%s), est to finish in %d", t.Id, t.Title, startInstrucion.EstDayToFinish))
				}
			}
			break
		case Stop:
			var stopInstrucion *StopInstruction
			err = json.Unmarshal(ins.Data, &stopInstrucion)
			if err != nil {
				t := findTask(tasks, stopInstrucion.Id)
				if t != nil && t.Status != Finished {
					t.Status = Stopped
					emitEvent(blockHash, ins.Id, fmt.Sprintf("Stop Task #%s(%s), because of %s", t.Id, t.Title, stopInstrucion.Reason))
					newMem, _ = json.Marshal(tasks)
				}
			}
			break
		case Pause:
			var pauseInstrucion *PauseInstruction
			err = json.Unmarshal(ins.Data, &pauseInstrucion)
			if err != nil {
				t := findTask(tasks, pauseInstrucion.Id)
				if t != nil && (t.Status == JustCreated || t.Status == Doing) {
					t.Status = Paused
					emitEvent(blockHash, ins.Id, fmt.Sprintf("Pause Task #%s(%s), est to wait %d day(s)", t.Id, t.Title, pauseInstrucion.EstWaitDay))
					newMem, _ = json.Marshal(tasks)
				}
			}
			break
		case Finish:
			var finishInstrucion *FinishInstruction
			err = json.Unmarshal(ins.Data, &finishInstrucion)
			if err != nil {
				t := findTask(tasks, finishInstrucion.Id)
				if t != nil && t.Status != Stopped {
					t.Status = Finished
					emitEvent(blockHash, ins.Id, fmt.Sprintf("Finish Task #%s(%s). %s", t.Id, t.Title, finishInstrucion.CongratMessage))
					newMem, _ = json.Marshal(tasks)
				}
			}
			break
		}
	}
	return newMem
}

func emitEvent(blockHash string, instructionId string, message string) {
	id := RandString(8)
	ev := Event{
		Id:            id,
		BlockHash:     blockHash,
		InstructionId: instructionId,
		Msg:           message,
	}
	fmt.Println("Emit: ", ev) // event will be sent to clients
	events = append(events, ev)
}

func findTask(tasks []Task, Id string) *Task {
	for _, t := range tasks {
		if t.Id == Id {
			return &t
		}
	}
	return nil
}

func GoBlockchanLogic() {
	// test your understanding of my code here hehe!
}
