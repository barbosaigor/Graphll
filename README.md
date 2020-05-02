# Graph LL

Graphll is a simple but efficient _weighted dependency graph_ data structure, having O(1) complexity for read and write operations. 
Easy to use, concise and simple, optimized for graph readings and writes. 

## Documentation

**Add** attaches a node to the graph


| Node | Weight | Dependency |
| ------------- |:-----------------:| -------------:|
| A | 80 | B, C, D |
| B | 90 | D, E, F |
| C | 10 | D, E |
| D | 5 | - |
| E | 7 | - |
| F | 15 | A |


```golang
g := graphll.New()
g.Add("a", 80, []string{"b", "c", "d"})
g.Add("b", 90, []string{"d", "e", "f"})
g.Add("c", 10, []string{"d", "e"})
g.Add("d", 5, nil)
g.Add("e", 7, nil)
g.Add("f", 15, []string{"a"})
```

**Weight** gets the weight of a node  
```golang
g := graphll.New()
g.Add("a", 80, []string{"b", "c"})
weight, err := g.Weight("a") // weight == 80
```

**Deps** gets the dependencies of a node  
```golang
g := graphll.New()
g.Add("a", 80, []string{"b", "c"})
deps, err := g.Deps("a") // deps == []string{"b", "c"}
```
