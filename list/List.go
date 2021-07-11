package list

import (
	"fmt"
	"math/rand"
	"ranking/node"
)

// zSkipList max layer
const ZSLML = 32
// zKipList
const ZSLP = 0.4

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
	var layer = make([]*node.LayerNode, ZSLML)

	for i := 0; i < ZSLML; i++ {
		layer[i] = &node.LayerNode{Span: uint64(i)}
	}

	header.Layer = layer
	zsl.Header = &header

	return &zsl
}

func (this *ZSkipList) Add(key, score int64)node.Node {

	op := this.Header
	needUpdateLayer := make(map[int]*node.Node)
	rank := make([]uint64,ZSLML)

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

	layers := getRandLayer()

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

func (this *ZSkipList) Del(key, score int64)bool {

	needUpdateLayer := make(map[int]*node.Node)

	op := this.Header
	for i := this.Layers - 1; i >= 0 ; i-- {
		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Score < score {
			op = op.Layer[i].ForwardNode
		}
		needUpdateLayer[i] = op
	}

	op = op.Layer[0].ForwardNode
	if op != nil && op.Score == score && op.Key == key {
		for i := 0; i < this.Layers; i++ {
			if needUpdateLayer[i].Layer[i].ForwardNode == op {
				needUpdateLayer[i].Layer[i].Span += op.Layer[i].Span - 1
				needUpdateLayer[i].Layer[i].ForwardNode = op.Layer[i].ForwardNode
			}else{
				needUpdateLayer[i].Layer[i].Span -= 1
			}
		}

		if op.Layer[0].ForwardNode != nil {
			op.Layer[0].ForwardNode.BackwardNode = op.BackwardNode
		}else{
			this.Tail = op.BackwardNode
		}

		for this.Layers > 1 && this.Header.Layer[this.Layers - 1].ForwardNode == nil {
			this.Layers --
		}
		this.Len --
		return true
	}

	return false
}


func getRandLayer()int  {
	layers := 1
	for rand.Float32() < ZSLP {
		layers ++
	}
	if  layers < ZSLML{
		return layers
	}
	return ZSLML
}

func (this *ZSkipList) GraphPrint()interface{}  {

	op := this.Header
	data := make([]interface{},0)

	res := make(map[string]interface{})
	res["key"] =  "head"
	res["score"] = "head"
	res["point"] = fmt.Sprintf("%p",op)
	res["span"] = op.Layer[0].Span

	layer := make([]interface{},len(op.Layer))
	for i := 0; i < len(op.Layer); i++ {
		layer[i] = fmt.Sprintf("%p",op.Layer[i].ForwardNode)
	}
	res["layer"] = layer
	data = append(data, res)

	for op.Layer[0].ForwardNode != nil {
		curNode := op.Layer[0].ForwardNode

		res := make(map[string]interface{})
		res["key"] =  curNode.Key
		res["score"] = curNode.Score
		res["point"] = fmt.Sprintf("%p",curNode)
		res["span"] = op.Layer[0].Span

		layer := make([]interface{},len(curNode.Layer))
		for i := 0; i < len(curNode.Layer); i++ {
			layer[i] = fmt.Sprintf("%p",curNode.Layer[i].ForwardNode)
		}
		res["layer"] = layer
		data = append(data, res)
		op = curNode
	}
	return data
}
