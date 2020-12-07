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

// ToBags parses the input to Nodes and WEdges (weighted edges).
func ToBags(b *Bags) func(string) error {
	bags := make(Bags)
	toBags := func(s string) error {
		if len(s) == 0 { // EOF write to b
			*b = bags
			return nil
		}
		rule := strings.Split(s, " bags contain ")
		id := rule[0]
		if _, found := bags[id]; !found {
			bags[id] = util.NewNode(id)
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
				if _, found := bags[smallID]; !found { // Add bag with this id.
					bags[smallID] = util.NewNode(smallID)
				}
				bags[id], bags[smallID] = bags[id].AddOutgoing(bags[smallID], qty)
			}
		}
		return nil
	}
	return toBags
}

func main() {
	bags := make(Bags)
	err := util.ReadLines(inputFilePath, ToBags(&bags))
	if err != nil {
		fmt.Println(err)
	}
	predecessors := bags["shiny gold"].AllPredecessors()
	fmt.Println(len(predecessors))
	sum := bags["shiny gold"].SumWeightOut()
	fmt.Println(sum)
}
