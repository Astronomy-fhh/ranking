package node

type Node struct {
	Key int64
	Score int64
	Layer        [] *LayerNode
	BackwardNode *Node
}

type LayerNode struct {
	ForwardNode *Node
	Span int64
}

func newNode(layer,score int64)Node  {
	var node Node
	node.Score = score
	return node
}



