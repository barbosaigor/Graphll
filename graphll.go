// Graphll is a package that contains a graph data structure.
// 
// Each node has a weight and name attached,
// Graphll is a map of nodes with their respective deps.
package graphll

import (
	"fmt"
	"strings"
	"errors"
)

var (
	ErrElementNotFound = errors.New("Element not found")
)

type GraphLLNode struct {
	weight uint32
	deps set
}

type GraphLL map[string]GraphLLNode

func New() GraphLL {
	return make(GraphLL)
}

func (g GraphLL) Add(name string, weight uint32, deps []string) {
	if _, ok := g[name]; !ok {
		g[name] = GraphLLNode{weight, newSet(deps)}
	} else {
		g[name].deps.union(newSet(deps))
	}
}

func (g GraphLL) GetDeps(name string) ([]string, error) {
	if _, ok := g[name]; !ok {
		return nil, ErrElementNotFound
	}
	return g[name].deps.toSlice(), nil
}

func (g GraphLL) GetWeight(name string) (uint32, error) {
	if _, ok := g[name]; !ok {
		return 0, ErrElementNotFound
	}
	return g[name].weight, nil
}

func (g GraphLL) String() string {
	if len(g) == 0 {
		return ""
	}
	var lstr []string
	for name, node := range g {
		lstr = append(lstr, fmt.Sprintf("[%v|%v]: %v", name, node.weight, node.deps))
	}
	return strings.Join(lstr, "\n")
}