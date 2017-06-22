package tree

type Tree struct {
    Root *Node
    size int
    m    int
}

type Node struct {
    Parent *Node
    Weight float64
    Number int
    Children []*Node
}
