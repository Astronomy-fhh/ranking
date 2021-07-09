package list

import (
	"GoLearn/GoZset/node"
	"math/rand"
)

// zKipList max layer
const ZML = 32

type ZSkipList struct {
	Header *node.Node
	Tail   *node.Node
	Len    uint64
	Layers int
}

func NewZSkipList() *ZSkipList {
	var zsl ZSkipList
	zsl.Layers = 1
	zsl.Len = 0
	header := node.Node{}
	var layer = make([]*node.LayerNode, ZML)

	for i := 0; i < ZML; i++ {
		layer[i] = &node.LayerNode{Span: uint64(i)}
	}

	header.Layer = layer
	zsl.Header = &header

	return &zsl
}

func (this *ZSkipList) Add(key, score int64)node.Node {

	op := this.Header
	needUpdateLayer := make(map[int]*node.Node)
	rank := make([]uint64,ZML)

	// to find a op node

	//skip the layer
	for i := this.Layers - 1; i >= 0; i-- {
		if i == this.Layers {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		// skip the nodes of each layer
		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Score < score {
			// no need to lower the layer
			// continue to maintain the info of the layer
			rank[i] = rank[i] + op.Layer[i].Span
			op = op.Layer[i].ForwardNode
		}
		needUpdateLayer[i] = op
	}

	layers := rand.Intn(20)

	if layers > this.Layers {
		for i := this.Layers; i < layers; i++ {
			needUpdateLayer[i] = this.Header
			needUpdateLayer[i].Layer[i].Span = this.Len
		}
		this.Layers = layers
	}

	layerNode := make([]*node.LayerNode,layers)
	for i := 0; i < layers; i++ {
		layerNode[i] = &node.LayerNode{}
	}
	op = &node.Node{
		Key: key,
		Score: score,
		Layer: layerNode,
	}

	for i := 0; i < layers; i++ {
		op.Layer[i].ForwardNode = needUpdateLayer[i].Layer[i].ForwardNode
		needUpdateLayer[i].Layer[i].ForwardNode = op
		op.Layer[i].Span = needUpdateLayer[i].Layer[i].Span - (rank[0] - rank[i])
		needUpdateLayer[i].Layer[i].Span = rank[0] - rank[i] + 1
	}

	for i := layers; i < this.Layers; i++ {
		needUpdateLayer[i].Layer[i].Span ++
	}

	if needUpdateLayer[0] != this.Header {
		op.BackwardNode = needUpdateLayer[0]
	}
	if op.Layer[0].ForwardNode != nil {
		op.Layer[0].ForwardNode.BackwardNode = op
	}else {
		this.Tail = op
	}

	this.Len++
	return *op
}
