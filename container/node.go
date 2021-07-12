package container

type Node struct {
	Obj *Obj
	Layer        [] *LayerNode
	BackwardNode *Node
}

type LayerNode struct {
	ForwardNode *Node
	Span int64
}

type Obj struct {
	Key string
	Score int64
}

