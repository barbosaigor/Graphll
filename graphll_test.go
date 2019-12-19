package graphll

import (
	// "fmt"
	"testing"
)

func populateGraphLL() GraphLL {
	g := New()
	g.Add("a", 80, []string{"b", "c"})
	g.Add("a", 80, []string{"d"})
	g.Add("b", 90, []string{"d", "e"})
	g.Add("c", 10, []string{"f", "g", "h"})
	g.Add("b", 90, []string{"x", "y"})
	return g
}

func compareStringSlices(a, b []string) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestAdd(t *testing.T) {
	g := populateGraphLL()
	if g.String() == "" {
		t.Error("Add elements to graph not working")
	}
}

func TestGetWeight(t *testing.T) {
	g := populateGraphLL()
	weight, err := g.GetWeight("b")
	if err != nil {
		t.Errorf("Get weight should return %v, but got %v", 90, weight)
	}
	weight2, err2 := g.GetWeight("k")
	if err2 == nil {
		t.Errorf("Get weight should return an error, but got %v", weight2)
	}
}

func TestGetDeps(t *testing.T) {
	g := populateGraphLL()
	deps, err := g.GetDeps("a")
	if err != nil {
		t.Error("Get deps should return deps, but have an error")
		return
	}
	expected := []string{"b", "c", "d"}
	if !compareStringSlices(expected, deps) {
		t.Errorf("Get deps should return %v, but got %v", expected, deps)
	}
}