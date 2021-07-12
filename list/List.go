package list

import (
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
	Len    int64
	Layers int
}

func NewZSkipList() *ZSkipList {
	var zsl ZSkipList
	zsl.Layers = 1
	zsl.Len = 0
	header := node.Node{}
	var layer = make([]*node.LayerNode, ZSLML)

	for i := 0; i < ZSLML; i++ {
		layer[i] = &node.LayerNode{Span: int64(i)}
	}

	header.Layer = layer
	zsl.Header = &header

	return &zsl
}

func (zkl *ZSkipList) Add(key, score int64)node.Node {

	op := zkl.Header
	needUpdateLayer := make(map[int]*node.Node)
	rank := make([]int64,ZSLML)

	// to find a op node

	// skip the layer
	for i := zkl.Layers - 1; i >= 0; i-- {
		if i == zkl.Layers {
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

	// if the level of new node is higher than header
	// init header's top layer span and append them to the needUpdateLayer
	if layers > zkl.Layers {
		for i := zkl.Layers; i < layers; i++ {
			needUpdateLayer[i] = zkl.Header
			needUpdateLayer[i].Layer[i].Span = zkl.Len
		}
		zkl.Layers = layers
	}

	// init a new node
	layerNode := make([]*node.LayerNode,layers)
	for i := 0; i < layers; i++ {
		layerNode[i] = &node.LayerNode{}
	}
	op = &node.Node{
		Key: key,
		Score: score,
		Layer: layerNode,
	}

	// update needUpdateLayer's layer pointer and span
	for i := 0; i < layers; i++ {
		op.Layer[i].ForwardNode = needUpdateLayer[i].Layer[i].ForwardNode
		needUpdateLayer[i].Layer[i].ForwardNode = op
		op.Layer[i].Span = needUpdateLayer[i].Layer[i].Span - (rank[0] - rank[i])
		needUpdateLayer[i].Layer[i].Span = rank[0] - rank[i] + 1
	}

	for i := layers; i < zkl.Layers; i++ {
		needUpdateLayer[i].Layer[i].Span ++
	}

	if needUpdateLayer[0] != zkl.Header {
		op.BackwardNode = needUpdateLayer[0]
	}

	// change tail pointer
	if op.Layer[0].ForwardNode != nil {
		op.Layer[0].ForwardNode.BackwardNode = op
	}else {
		zkl.Tail = op
	}

	zkl.Len++
	return *op
}

func (zkl *ZSkipList) Del(key, score int64)bool {

	needUpdateLayer := make(map[int]*node.Node)

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Score < score {
			op = op.Layer[i].ForwardNode
		}
		needUpdateLayer[i] = op
	}

	op = op.Layer[0].ForwardNode
	if op != nil && op.Score == score && op.Key == key {
		zkl.delNode(op,needUpdateLayer)
		return true
	}

	return false
}

func (zkl *ZSkipList) delNode(op *node.Node,  needUpdateLayer map[int]*node.Node)  {
	for i := 0; i < zkl.Layers; i++ {
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
		zkl.Tail = op.BackwardNode
	}

	for zkl.Layers > 1 && zkl.Header.Layer[zkl.Layers - 1].ForwardNode == nil {
		zkl.Layers --
	}
	zkl.Len --
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

func (zkl *ZSkipList) Update(Key, curScore,newScore int64)bool {

	needUpdateLayer := make(map[int]*node.Node)
	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0; i-- {
		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Score < curScore {
			op = op.Layer[i].ForwardNode
		}
		needUpdateLayer[i] = op
	}

	op = op.Layer[0].ForwardNode
	if op != nil && op.Score == curScore && op.Key == Key {
		if (op.BackwardNode == nil || op.BackwardNode.Score < newScore) && (op.Layer[0].ForwardNode == nil || op.Layer[0].ForwardNode.Score > newScore){
			op.Score = newScore
			return true
		}

		zkl.delNode(op,needUpdateLayer)
		zkl.Add(Key,newScore)
		return true
	}
	return false
}

func (zkl *ZSkipList) GetRangeByRank(start,end int64) []node.Node {

	res := make([]node.Node,0)
	rank := make([]int64,zkl.Layers)

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		if i == zkl.Layers - 1 {
			rank[i] = 0
		}else {
			rank[i] = rank[i+1]
		}

		for op.Layer[i].ForwardNode != nil && (rank[i] + op.Layer[i].Span) < start {
			rank[i] += op.Layer[i].Span
			op = op.Layer[i].ForwardNode
		}
	}

	idx := start
	for (idx <= end || end == -1) && op.Layer[0].ForwardNode != nil {
		res = append(res, *op.Layer[0].ForwardNode)
		op = op.Layer[0].ForwardNode
		idx ++
	}
	return res
}

func (zkl *ZSkipList) GetRevRangeByRank(start,end int64)[]node.Node  {
	res := make([]node.Node,0)
	rank := make([]int64,zkl.Layers)
	realRank := zkl.Len - start + 1

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		if i == zkl.Layers - 1 {
			rank[i] = 0
		}else {
			rank[i] = rank[i+1]
		}

		for op.Layer[i].ForwardNode != nil && (rank[i] + op.Layer[i].Span) <= realRank {
			rank[i] += op.Layer[i].Span
			op = op.Layer[i].ForwardNode
		}
	}

	idx := start
	for (idx <= end || end == -1) && op != nil && op.BackwardNode != nil {
		res = append(res, *op)
		op = op.BackwardNode
		idx++
	}

	return res
}

func (zkl *ZSkipList) GetRangeByScore(start,end int64)[]node.Node  {
	res := make([]node.Node,0)
	rank := make([]int64,zkl.Layers)

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		if i == zkl.Layers - 1 {
			rank[i] = 0
		}else {
			rank[i] = rank[i+1]
		}

		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Score < start {
			rank[i] += op.Layer[i].Span
			op = op.Layer[i].ForwardNode
		}
	}

	for op.Layer[0].ForwardNode != nil && op.Layer[0].ForwardNode.Score <= end {
		res = append(res, *op.Layer[0].ForwardNode)
		op = op.Layer[0].ForwardNode
	}
	return res
}

func (zkl *ZSkipList) GetRevRangeByScore(start,end int64)[]node.Node  {
	res := make([]node.Node,0)
	rank := make([]int64,zkl.Layers)

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		if i == zkl.Layers - 1 {
			rank[i] = 0
		}else {
			rank[i] = rank[i+1]
		}

		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Score <= end {
			rank[i] += op.Layer[i].Span
			op = op.Layer[i].ForwardNode
		}
	}


	for op.BackwardNode != nil && op.Score >= start {
		res = append(res, *op)
		op = op.BackwardNode
	}

	return res
}



func (zkl *ZSkipList) isRangeValid(min,max int64)bool {
	if min < 0 || (max != -1 && max < min) {
		return false
	}
	return true
}