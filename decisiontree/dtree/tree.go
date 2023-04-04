package dtree

import (
	"fmt"
	"gobyexample/ppds/tree"
	"log"
)

type Tree struct {
	Name        string
	Description string
	Start       *Node
	Current     *Node
}

func CreateTree(root *Node, name string, description string) *Tree {
	tree := Tree{
		Start:       root,
		Current:     root,
		Name:        name,
		Description: description,
	}
	return &tree
}
func (this *Tree) Print() {
	tree.Print(this.Start)
}
func (this *Tree) PrintFromCurrent() {
	tree.Print(this.Current)
}
func (this *Tree) Choose(idx int) {
	nextNode := this.Current.Choose(idx)
	if nextNode == nil {
		fmt.Println(idx, " out of bound, no move")
	}
	if nextNode != nil {
		fmt.Printf("from %s choose: %d got %s\n", this.Current.Name, idx, nextNode.Name)
		this.Current = nextNode
		this.PrintFromCurrent()
	}
}
func (this *Tree) IsValidChoice(choices []int) bool {
	if this.Current == nil {
		log.Fatal("current node is nil")
	}
	return this.Current.IsValidChoice(choices)
}
