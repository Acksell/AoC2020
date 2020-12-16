package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/Acksell/aoc2020/util"
)

type Range struct {
	min, max int
}

func (rng Range) Contains(i int) bool {
	return rng.min <= i && i <= rng.max
}

type Rule struct {
	name   string
	ranges []Range
}

func (r Rule) Satisfies(i int) bool {
	for _, rng := range r.ranges {
		if rng.Contains(i) {
			return true
		}
	}
	return false
}

// NewRule takes a string of format "fieldname: 123-321 or 666-999" and returns a Rule.
func NewRule(s string) (Rule, error) {
	line := strings.Split(s, ": ")
	fieldname := line[0]
	rulestring := line[1]
	rangestrings := strings.Split(rulestring, " or ")
	ranges := make([]Range, 0)
	for _, rng := range rangestrings {
		minmax := strings.Split(rng, "-")
		min, err := strconv.Atoi(minmax[0])
		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			return Rule{}, err
		}
		r := Range{min, max}
		ranges = append(ranges, r)
	}
	return Rule{fieldname, ranges}, nil
}

type Ticket []int

func (t Ticket) InvalidFields(r Rule) {

}

func NewTicket(s string) (Ticket, error) {
	ints := strings.Split(s, ",")
	ticket := make(Ticket, 0)
	for _, n := range ints {
		i, err := strconv.Atoi(n)
		if err != nil {
			return Ticket{}, err
		}
		ticket = append(ticket, i)
	}
	return ticket, nil
}

type TicketValidator struct {
	rules   []Rule
	ticket  Ticket
	tickets []Ticket
}

func (tv TicketValidator) ValidTickets() ([]Ticket, int) {
	errorRate := 0 // as defined by the task (sum of invalid values)
	valid := make([]Ticket, 0)
	for _, t := range tv.tickets {
		isValid := true
		for _, v := range t {
			satisfied := false
			for _, r := range tv.rules {
				if r.Satisfies(v) {
					satisfied = true
					break
				}
			}
			if !satisfied {
				errorRate += v
				isValid = false
			}
		}
		if isValid {
			valid = append(valid, t)
		}
	}
	return valid, errorRate
}

func (tv TicketValidator) MatchRules() []Rule {
	// for each ticket, keep track of which rules are possible.
	possible := make([]util.IntSet, len(tv.ticket))
	for f := range tv.ticket {
		ruleset := make(util.IntSet)
		for i := range tv.rules {
			ruleset[i] = true // initialize value to true so that we can use AND operator to carry answers.
		}
		possible[f] = ruleset // create set of ints for each field.
	}
	valid, _ := tv.ValidTickets()
	for _, ticket := range valid {
		for field, v := range ticket {
			for i, r := range tv.rules {
				if candidate := r.Satisfies(v); possible[field][i] && !candidate {
					delete(possible[field], i)
				}
			}
		}
	}
	result := make([]Rule, len(possible))
	found := 0
	// finally remove those that we know for certain until all found.
	// assumes there is always at least one that we know for certain at each step in the reduction.
	for found < len(possible)-1 {
		for field, ruleset := range possible {
			if len(ruleset) == 1 {
				for rule := range ruleset {
					result[field] = tv.rules[rule]
					found++
					for j := range possible {
						delete(possible[j], rule)
					}
				}
			}
		}
	}
	return result
}

func NewTicketValidator(_ticket string, _rules, _tickets []string) (TicketValidator, error) {
	rules := make([]Rule, 0)
	for _, s := range _rules {
		rule, err := NewRule(s)
		if err != nil {
			return TicketValidator{}, err
		}
		rules = append(rules, rule)
	}
	ticket, err := NewTicket(_ticket)
	if err != nil {
		return TicketValidator{}, err
	}
	tickets := make([]Ticket, 0)
	for _, s := range _tickets {
		if len(s) == 0 { // Not a valid ticket
			continue
		}
		t, err := NewTicket(s)
		if err != nil {
			return TicketValidator{}, err
		}
		tickets = append(tickets, t)
	}
	return TicketValidator{rules, ticket, tickets}, nil
}

const inputFilePath = "../../inputs/tickets.txt"

func main() {
	buf, _ := ioutil.ReadFile(inputFilePath)
	input := string(buf)
	content := strings.Split(input, "\n\n")
	tv, _ := NewTicketValidator(strings.Split(content[1], "\n")[1], strings.Split(content[0], "\n"), strings.Split(content[2], "\n")[1:])

	_, errRate := tv.ValidTickets()
	fmt.Println(errRate) // part 1

	rules := tv.MatchRules()
	prod := 1
	for f, rule := range rules {
		if len(rule.name) >= 9 && rule.name[:9] == "departure" {
			prod *= tv.ticket[f]
		}
	}
	fmt.Println(prod) // part 2
}
