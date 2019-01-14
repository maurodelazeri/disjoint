# disjoint

disjoint algorithm to union / find sets based on map and list in golang.

## Install

`$ go get github.com/maurodelazeri/disjoint`

## Usage

```go
uf := union_find.New()

uf.Union(1, 2, 3)
fmt.Println(uf.Find(1))    // Prints [1, 2, 3]
fmt.Println(uf.Find(2))    // Prints [1, 2, 3]
fmt.Println(uf.Find(2, 3)) // Prints [1, 2, 3]
fmt.Println(uf.InSameSet(1, 2)) // Prints true
fmt.Println(uf.InSameSet(1, 3)) // Prints true
fmt.Println(uf.InSameSet(2, 3)) // Prints true

uf.Union(4, 5, 6)
fmt.Println(uf.Find(4))    // Prints [4, 5, 6]
fmt.Println(uf.Find(5))    // Prints [4, 5, 6]
fmt.Println(uf.Find(4, 6)) // Prints [4, 5, 6]
fmt.Println(uf.InSameSet(4, 5)) // Prints true
fmt.Println(uf.InSameSet(4, 6)) // Prints true
fmt.Println(uf.InSameSet(5, 6)) // Prints true

fmt.Println(uf.Find(1, 6))      // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.InSameSet(1, 4)) // Prints false
fmt.Println(uf.InSameSet(1, 5)) // Prints false
fmt.Println(uf.InSameSet(2, 6)) // Prints false

uf.Union(2, 6)
fmt.Println(uf.Find(1)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(2)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(3)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(4)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(5)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(1, 6)) // Prints [1, 2, 3, 4, 5, 6]

fmt.Println(uf.InSameSet(1, 4)) // Prints true
fmt.Println(uf.InSameSet(1, 5)) // Prints true
fmt.Println(uf.InSameSet(2, 6)) // Prints true

fmt.Println(uf.RemoveSet(1)) // Prints [1, 2, 3, 4, 5, 6]
fmt.Println(uf.Find(3))      // Prints [ ]
```
