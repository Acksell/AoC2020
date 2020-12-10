package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

// Bags is a map of all the bags. Acyclicity is assumed by the input.
type Bags util.Graph

const inputFilePath = "../../inputs/bags.txt"

// Load parses the input to Nodes and edges.
func (b *Bags) Load(s string) error {
	rule := strings.Split(s, " bags contain ")
	id := rule[0]
	if _, found := (*b)[id]; !found {
		(*b)[id] = util.NewNode(id)
	}
	if rule[1] != "no other bags." {
		rule1 := rule[1][:len(rule[1])-1] // strip the trailing dot "."
		contents := strings.Split(rule1, ", ")
		for _, smallbag := range contents {
			smallID := strings.Split(smallbag, " bag")[0] // remove the suffix "bag(s)"
			qty, err := strconv.Atoi(string(smallID[0]))  // Assumes only single digit quantity for now.
			// "1 vibrant yellow" => "vibrant yellow". Assumes only single digit quantity for now.
			smallID = smallID[2:]
			if err != nil {
				return err
			}
			if _, found := (*b)[smallID]; !found { // Add bag with this id.
				(*b)[smallID] = util.NewNode(smallID)
			}
			(*b)[id], (*b)[smallID] = (*b)[id].AddOutgoing((*b)[smallID], qty)
		}
	}
	return nil
}

func main() {
	bags := make(Bags)
	err := util.ReadLines(inputFilePath, &bags)
	if err != nil {
		fmt.Println(err)
	}
	predecessors := bags["shiny gold"].AllPredecessors()
	fmt.Println(len(predecessors))
	sum := bags["shiny gold"].SumWeightOut()
	fmt.Println(sum)
}
