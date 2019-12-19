# Graph LL

Graphll is an efficient weighted dependecy graph data structure. 
Easy to use, concise and simple, optimized for graph readings and writes.

## Documentation

**Add** a node to graph


| Node | Weight | Dependency |
| ------------- |:-----------------:| -------------:|
| A | 80 | B, C, D |
| B | 90 | D, E, X, Y |
| C | 10 | F, G, H |


```golang
g := graphll.New()
g.Add("a", 80, []string{"b", "c"})
g.Add("a", 80, []string{"d"})
g.Add("b", 90, []string{"d", "e"})
g.Add("c", 10, []string{"f", "g", "h"})
g.Add("b", 90, []string{"x", "y"})
```

**Get Weight** of a node

```golang
g := graphll.New()
g.Add("a", 80, []string{"b", "c"})
weight, err := g.GetWeight("a") // weight == 80
```

**Get Dependencies** of a node

```golang
g := graphll.New()
g.Add("a", 80, []string{"b", "c"})
deps, err := g.GetDeps("a") // deps == []string{"b", "c"}
```