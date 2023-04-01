package decisiontree

import (
	"fmt"
	"gobyexample/ppds/tree"
	"log"
)

type Node struct {
	id       string
	name     string
	parent   *Node
	children []*Node
	voted    map[int]int
	isOuput  bool
	nodeType string // "SingleChoice" or "MultipleChoice"
	data     interface{}
}

// return something that is printable
func (n *Node) Data() interface{} {
	if n.isOuput == true {
		return n.id + "*"
	}
	return n.id
}

// cannot return n.children directly.
// https://github.com/golang/go/wiki/InterfaceSlice
func (n *Node) Children() (c []tree.Node) {
	for _, child := range n.children {
		c = append(c, tree.Node(child))
	}
	return
}

// nodeType == "SingleChoice" then option is int; nodeType == "MultiChoice" then option is []string
func (this *Node) Vote(who string, options []int) (selectedOption int, votedResult map[int]int) {
	voted := this.voted
	if this.nodeType == "SingleChoice" {
		singleChoice := options[0]
		voted[singleChoice] += 1
		optionName := this.children[singleChoice].name
		fmt.Printf("%s vote %d [%s]. There are %d person(s) choose %s\n", who, options[0], optionName, voted[singleChoice], optionName)
		if voted[singleChoice] == 3 { // let's move on
			return singleChoice, voted
		}
	} else if this.nodeType == "MultipleChoice" {
		choices := make([]int, 0)
		for _, opt := range options {
			choices = append(choices, InterfaceToInt(opt))
		}
		str := ""
		max, options := ConverToMultipleChoiceData(this.data)
		for _, choice := range choices {
			str += options[choice] + ","
			voted[choice] += 1
		}
		choosen := make(map[int]int)

		fmt.Printf("%s vote [%s]; top %d choice win will\n", who, str, max)

		for opt, choice := range voted {
			if choice >= max {
				choosen[opt] = choice
			}
		}

		return 0, choosen
	}
	return -1, nil
}

func (this *Node) Start(input map[int]int, originalOptions []string) {
	if this.nodeType == "MultipleChoice" {
		// take max
		max := InterfaceToInt(this.data.(map[string]interface{})["max"])
		options := make([]interface{}, 0)
		i := 0
		for opt, _ := range input {
			if i <= max {
				options = append(
					options, originalOptions[opt])
			}
			i++
		}
		this.data.(map[string]interface{})["options"] = options
	}
}

func createEmptyNode(name string, isOutput bool, nodeType string, data interface{}) *Node {
	node := Node{
		name:     name,
		children: []*Node{},
		voted:    make(map[int]int),
		isOuput:  isOutput,
		nodeType: nodeType,
		data:     data,
	}
	return &node
}
func createNodeWithChildren(name string, children []*Node, isOutput bool, nodeType string, data interface{}) *Node {
	node := Node{
		name:     name,
		children: children,
		voted:    make(map[int]int),
		isOuput:  isOutput,
		nodeType: nodeType,
		data:     data,
	}
	for _, child := range children {
		child.parent = &node
	}
	return &node
}
func (this *Node) attach(child *Node) *Node {
	child.parent = this
	this.children = append(this.children, child)
	return child
}
func (this *Node) print() {
	fmt.Printf("%s has following children:\n", this.name)
	for i := range this.children {
		fmt.Printf("- opt %d: %s\n", i, this.children[i].name)
	}
	fmt.Printf("\n")
}
func (this *Node) choose(idx int) *Node {
	if idx < len(this.children) {
		return this.children[idx]
	}
	return nil
}
func (this *Node) isValidChoice(choices []int) bool {
	if this.nodeType == "SingleChoice" {
		idx := choices[0]
		if idx < len(this.children) {
			return true
		}
		return false
	} else if this.nodeType == "MultipleChoice" {
		// should check if all choices are valid
		return true
	}
	return false
}

type Tree struct {
	name        string
	description string
	start       *Node
	current     *Node
}

func createTree(root *Node, name string, description string) *Tree {
	tree := Tree{
		start:       root,
		current:     root,
		name:        name,
		description: description,
	}
	return &tree
}
func (this *Tree) print() {
	tree.Print(this.start)
}
func (this *Tree) printFromCurrent() {
	tree.Print(this.current)
}
func (this *Tree) choose(idx int) {
	nextNode := this.current.choose(idx)
	if nextNode == nil {
		fmt.Println(idx, " out of bound, no move")
	}
	if nextNode != nil {
		fmt.Printf("from %s choose: %d got %s\n", this.current.name, idx, nextNode.name)
		this.current = nextNode
		this.printFromCurrent()
	}
}
func (this *Tree) isValidChoice(choices []int) bool {
	if this.current == nil {
		log.Fatal("current node is nil")
	}
	return this.current.isValidChoice(choices)
}

func ConverToMultipleChoiceData(data interface{}) (int, []string) {
	tmp := data.(map[string]interface{})
	var max int
	options := make([]string, 0)
	for key, value := range tmp {
		if key == "options" {
			for _, opt := range value.([]interface{}) {
				options = append(options, opt.(string))
			}
		} else if key == "max" {
			max = InterfaceToInt(value)
		}
	}
	return max, options
}

func ConvertToSingleChoiceData(data interface{}) []string {
	tmp := data.(map[string]interface{})
	options := make([]string, 0)
	for key, value := range tmp {
		if key == "options" {
			for _, opt := range value.([]interface{}) {
				options = append(options, opt.(string))
			}
		}
	}
	return options
}
