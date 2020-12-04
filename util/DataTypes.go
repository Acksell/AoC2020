package util

// IntSet is a set of integers.
type IntSet map[uint64]bool

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
