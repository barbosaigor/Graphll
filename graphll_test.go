package graphll

import "testing"

func populateGraphLL() GraphLL {
	g := New()
	g.Add("a", 80, []string{"b", "c", "d"})
	g.Add("b", 90, []string{"d", "e", "f"})
	g.Add("c", 10, []string{"d", "e"})
	g.Add("d", 5, nil)
	g.Add("e", 5, nil)
	g.Add("f", 15, []string{"a"})
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
	if len(g) != 6 {
		t.Errorf("Add is not attaching the correct number of nodes, expected %d but got %d", 6, len(g))
	}
	if g.String() == "" {
		t.Error("Add String is not returning a valid string")
	}
}

func TestRemove(t *testing.T) {
	g := populateGraphLL()
	g.Remove("a")
	g.Remove("f")
	if len(g) != 4 {
		t.Errorf("Add is not attaching the correct number of nodes, expected %d but got %d", 4, len(g))
	}
}

func TestWeight(t *testing.T) {
	g := populateGraphLL()
	weight, err := g.Weight("b")
	if err != nil {
		t.Errorf("weight should return %v, but got %v", 90, weight)
	}
	weight, err = g.Weight("k")
	if err == nil {
		t.Errorf("weight should return an error, but got %v", weight)
	}
}

func TestDeps(t *testing.T) {
	g := populateGraphLL()
	deps, err := g.Deps("a")
	if err != nil {
		t.Error("deps should return deps, but got an error")
		return
	}
	expected := []string{"b", "c", "d"}
	if !compareStringSlices(expected, deps) {
		t.Errorf("deps should return %v, but got %v", expected, deps)
	}
}
