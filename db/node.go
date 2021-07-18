package db

type Node struct {
	Obj *Obj
	Layer        [] *LayerNode
	BackwardNode *Node
}

type LayerNode struct {
	ForwardNode *Node
	Span uint64
}

type Obj struct {
	Key string
	Score uint64
}

