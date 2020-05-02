package graphll

import "strings"

// set implements a set data structure
type set map[string]struct{}

func newSet(values []string) set {
	s := make(set)
	for _, value := range values {
		s.add(value)
	}
	return s
}

func newEmptySet() set {
	return make(set)
}

func (s set) add(key string) {
	s[key] = struct{}{}
}

func (s set) union(b set) {
	for key := range b {
		s.add(key)
	}
}

func (s set) toSlice() []string {
	var elements []string
	for key := range s {
		elements = append(elements, key)
	}
	return elements
}

func (s set) String() string {
	var lstr []string
	for key := range s {
		lstr = append(lstr, key)
	}
	return strings.Join(lstr, ", ")
}
