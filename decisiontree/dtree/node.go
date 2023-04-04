package dtree

import (
	"fmt"
	"gobyexample/decisiontree/utils"
	"gobyexample/decisiontree/votemachine"
	"gobyexample/ppds/tree"
)

type Node struct {
	Id          string
	Name        string
	Parent      *Node
	AllChildren []*Node
	Voted       map[int]int
	IsOuput     bool
	NodeType    string // "SingleChoice" or "MultipleChoice"
	AllData     interface{}
}

// return something that is printable
func (n *Node) Data() interface{} {
	if n.IsOuput == true {
		return n.Id + "*"
	}
	return n.Id
}

// cannot return n.children directly.
// https://github.com/golang/go/wiki/InterfaceSlice
func (n *Node) Children() (c []tree.Node) {
	for _, child := range n.AllChildren {
		c = append(c, tree.Node(child))
	}
	return
}

// nodeType == "SingleChoice" then option is int; nodeType == "MultiChoice" then option is []string
func (this *Node) Vote(who string, options []int) (selectedOption int, votedResult map[int]int) {
	voted := this.Voted
	if this.NodeType == "SingleChoice" {
		singleChoice := options[0]
		voted[singleChoice] += 1
		optionName := this.AllChildren[singleChoice].Name
		fmt.Printf("%s vote %d [%s]. There are %d person(s) choose %s\n", who, options[0], optionName, voted[singleChoice], optionName)
		if voted[singleChoice] == 3 { // let's move on
			return singleChoice, voted
		}
	} else if this.NodeType == "MultipleChoice" {
		choices := make([]int, 0)
		for _, opt := range options {
			choices = append(choices, utils.InterfaceToInt(opt))
		}
		str := ""
		max, options := utils.ConverToMultipleChoiceData(this.Data)
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
	if this.NodeType == "MultipleChoice" {
		// take max
		max := utils.InterfaceToInt(this.AllData.(map[string]interface{})["max"])
		options := make([]interface{}, 0)
		i := 0
		for opt, _ := range input {
			if i <= max {
				options = append(
					options, originalOptions[opt])
			}
			i++
		}
		this.AllData.(map[string]interface{})["options"] = options
	}
}

func CreateEmptyNode(name string, isOutput bool, nodeType string, data interface{}) *Node {
	node := Node{
		Name:        name,
		AllChildren: []*Node{},
		Voted:       make(map[int]int),
		IsOuput:     isOutput,
		NodeType:    nodeType,
		AllData:     data,
	}
	return &node
}

//	func CreateNodeWithChildren(name string, children []*Node, isOutput bool, nodeType string, data interface{}) *Node {
//		node := Node{
//			Name:     name,
//			AllChildren: children,
//			Voted:    make(map[int]int),
//			IsOuput:  isOutput,
//			NodeType: nodeType,
//			AllData:     data,
//		}
//		for _, child := range children {
//			child.parent = &node
//		}
//		return &node
//	}
func (this *Node) Attach(child *Node) *Node {
	child.Parent = this
	this.AllChildren = append(this.AllChildren, child)
	return child
}
func (this *Node) Print() {
	fmt.Printf("%s has following children:\n", this.Name)
	for i := range this.AllChildren {
		fmt.Printf("- opt %d: %s\n", i, this.AllChildren[i].Name)
	}
	fmt.Printf("\n")
}
func (this *Node) Choose(idx int) *Node {
	if idx < len(this.AllChildren) {
		return this.AllChildren[idx]
	}
	return nil
}
func (this *Node) IsValidChoice(choices []int) bool {
	if this.NodeType == "SingleChoice" {
		idx := choices[0]
		if idx < len(this.AllChildren) {
			return true
		}
		return false
	} else if this.NodeType == "MultipleChoice" {
		// should check if all choices are valid
		return true
	}
	return false
}

func (this *Node) ConvertOptionIdxToString(choices []int) []string {
	nodeType := this.NodeType
	result := make([]string, 0)
	if nodeType == "SingleChoice" {
		data := this.AllData.(votemachine.SingleChoiceData)
		result = append(result, data.Options[choices[0]])
	} else if nodeType == "MultipleChoice" {
		data := this.AllData.(votemachine.SingleChoiceData)
		for _, choice := range choices {
			result = append(result, data.Options[choice])
		}
	}
	return result
}
