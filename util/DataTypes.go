package util

import "fmt"

// Int64Set is a set of integers.
type Int64Set map[uint64]bool

// IntSet is a set of integers.
type IntSet map[int]bool

// StringSet is a set of strings.
type StringSet map[string]bool

// NewStringSet returns a set of the provided strings.
func NewStringSet(set ...string) StringSet {
	ss := make(StringSet)
	for _, key := range set {
		ss[key] = true
	}
	return ss
}

// Union returns the union between two StringSets.
func (ss1 StringSet) Union(ss2 StringSet) StringSet {
	union := make(StringSet)
	for s := range ss1 {
		union[s] = true
	}
	for s := range ss2 {
		union[s] = true
	}
	return union
}

/****
 **** Graph stuff.
 ****/
// type Edgy interface {
//
// }

// Edge is a weighted directed edge from node `From` to `To`
type Edge struct {
	From   *Node
	To     *Node
	Weight int
}

// WEdge is a weighted directed edge from node `From` to `To`
// type WEdge struct {
// 	Edge
// 	Weight int
// }

// Node is a node in a graph.
type Node struct {
	ID  string
	In  map[string]Edge
	Out map[string]Edge
}

// Graph is a set of nodes which have edges.
type Graph map[string]Node

// NewNode returns a new node given an ID.
func NewNode(id string) Node {
	ein := make(map[string]Edge, 1)
	eout := make(map[string]Edge, 1)
	return Node{id, ein, eout}
}

// NewEdge returns a new Edge struct.
func NewEdge(n1 *Node, n2 *Node, weight int) Edge {
	return Edge{n1, n2, weight}
}

// NewWEdge returns a new WEdge struct.
// func NewWEdge(n1 Node, n2 Node, weight int) WEdge {
// 	e := NewEdge(n1, n2)
// 	return WEdge{e, weight}
// }

// AddOutgoing adds an outgoing node and returns the modified inputs.
func (n Node) AddOutgoing(n2 Node, weight int) (Node, Node) {
	new := Node{n2.ID, n2.In, n2.Out} // n2 -> new.
	new.In[n.ID] = NewEdge(&new, &n, weight)
	n.Out[new.ID] = NewEdge(&n, &new, weight)
	// return this node and the new added node. (with properties updated)
	return n, Node{new.ID, new.In, new.Out}
}

// SumWeightOut Returns the outgoing sum of the weights from a node.
func (n Node) SumWeightOut() int {
	sum := 0
	for _, edge := range n.Out {
		// Eventually no remaining node satisfies this since graph is acyclic (base condition)
		if len(edge.To.Out) != 0 {
			sum += edge.Weight * (edge.To.SumWeightOut() + 1)
		} else {
			sum += edge.Weight
		}
	}
	return sum
}

// AllPredecessors returns IDs of all the predecessors to the node. Assumes acyclic.
func (n Node) AllPredecessors() StringSet {
	unique := make(StringSet)
	for id, edge := range n.In {
		if unique[edge.To.ID] { // already visited.
			continue
		} else {
			unique[id] = true
			// Update the unique ones found. Eventually no remaining node satisfies this (base condition)
			if len(edge.To.In) != 0 {
				unique = unique.Union(edge.To.AllPredecessors())
			}
		}
	}
	return unique
}

func (n Node) String() string {
	var sin []string
	var sout []string
	s := n.ID
	for _, in := range n.In {
		sin = append(sin, in.To.ID)
	}
	for _, out := range n.Out {
		sout = append(sout, out.To.ID)
	}
	return fmt.Sprintf("{%v %v %v}", s, sin, sout)
}
