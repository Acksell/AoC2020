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
func ToBags(bags *Bags) func(string) error {
	toBags := func(s string) error {
		rule := strings.Split(s, " bags contain ")
		id := rule[0]
		if _, found := (*bags)[id]; !found {
			(*bags)[id] = util.NewNode(id)
		}
		if rule[1] != "no other bags." {
			rule1 := rule[1][:len(rule[1])-1] // strip the trailing dot "."
			contents := strings.Split(rule1, ", ")
			for _, smallbag := range contents {
				smallID := strings.Split(smallbag, " bag")[0] // remove the suffix "bag(s)"
				// "1 vibrant yellow" => "vibrant yellow". Assumes only single digit qty.
				qty, err := strconv.Atoi(string(smallID[0]))
				smallID = smallID[2:]
				if err != nil {
					return err
				}
				if _, found := (*bags)[smallID]; !found { // Add bag with this id.
					(*bags)[smallID] = util.NewNode(smallID)
				}
				temp := (*bags)[id] // we know this exists
				(*bags)[smallID] = temp.AddOutgoing((*bags)[smallID], qty)
				(*bags)[id] = temp
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
