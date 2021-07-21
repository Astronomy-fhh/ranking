package db

import pb "ranking/proto"

type Node struct {
	Obj *pb.Obj
	Layer        [] *LayerNode
	BackwardNode *Node
}

type LayerNode struct {
	ForwardNode *Node
	Span int64
}


