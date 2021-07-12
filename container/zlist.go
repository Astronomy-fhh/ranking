package container

import (
	"math/rand"
)

// zSkipList max layer
const ZSLML = 32
// zKipList
const ZSLP = 0.4

type zList interface {
	Add(key string, score int64) Obj
	Del(key string, score int64) bool
	Update(key string, curScore, newScore int64) Obj
	GetRangeByRank(start, end int64) []Obj
	GetRevRangeByRank(start, end int64) []Obj
	GetRangeByScore(start, end int64) []Obj
	GetRevRangeByScore(start, end int64) []Obj
}


type ZSkipList struct {
	Header *Node
	Tail   *Node
	Len    int64
	Layers int
}

func NewZSkipList() *ZSkipList {
	var zsl ZSkipList
	zsl.Layers = 1
	zsl.Len = 0
	header := Node{}
	var layer = make([]*LayerNode, ZSLML)

	for i := 0; i < ZSLML; i++ {
		layer[i] = &LayerNode{Span: int64(i)}
	}

	header.Layer = layer
	zsl.Header = &header
	return &zsl
}

func (zkl *ZSkipList) Add(key string, score int64) Obj {

	op := zkl.Header
	needUpdateLayer := make(map[int]*Node)
	rank := make([]int64, ZSLML)

	// to find a op node

	// skip the layer
	for i := zkl.Layers - 1; i >= 0; i-- {
		if i == zkl.Layers {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		// skip the nodes of each layer
		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Obj.Score < score {
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
	layerNode := make([]*LayerNode,layers)
	for i := 0; i < layers; i++ {
		layerNode[i] = &LayerNode{}
	}
	
	obj := &Obj{Key: key, Score: score}
	op = &Node{
		Obj: obj,
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
	return *op.Obj
}

func (zkl *ZSkipList) Del(key string, score int64)bool {

	needUpdateLayer := make(map[int]*Node)

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Obj.Score < score {
			op = op.Layer[i].ForwardNode
		}
		needUpdateLayer[i] = op
	}

	op = op.Layer[0].ForwardNode
	if op != nil && op.Obj.Score == score && op.Obj.Key == key {
		zkl.delNode(op,needUpdateLayer)
		return true
	}

	return false
}

func (zkl *ZSkipList) Update(key string, curScore,newScore int64)Obj {

	needUpdateLayer := make(map[int]*Node)
	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0; i-- {
		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Obj.Score < curScore {
			op = op.Layer[i].ForwardNode
		}
		needUpdateLayer[i] = op
	}

	op = op.Layer[0].ForwardNode
	if op != nil && op.Obj.Score == curScore && op.Obj.Key == key {
		if (op.BackwardNode == nil || op.BackwardNode.Obj.Score < newScore) && (op.Layer[0].ForwardNode == nil || op.Layer[0].ForwardNode.Obj.Score > newScore){
			op.Obj.Score = newScore
			return *op.Obj
		}

		zkl.delNode(op,needUpdateLayer)
	}
	return 	zkl.Add(key, newScore)
}

func (zkl *ZSkipList) GetRangeByRank(start,end int64) []Obj {

	res := make([]Obj,0)
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
		res = append(res, *op.Layer[0].ForwardNode.Obj)
		op = op.Layer[0].ForwardNode
		idx ++
	}
	return res
}

func (zkl *ZSkipList) GetRevRangeByRank(start,end int64)[]Obj {
	res := make([]Obj,0)
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
		res = append(res, *op.Obj)
		op = op.BackwardNode
		idx++
	}

	return res
}

func (zkl *ZSkipList) GetRangeByScore(start,end int64)[]Obj {
	res := make([]Obj,0)
	rank := make([]int64,zkl.Layers)

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		if i == zkl.Layers - 1 {
			rank[i] = 0
		}else {
			rank[i] = rank[i+1]
		}

		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Obj.Score < start {
			rank[i] += op.Layer[i].Span
			op = op.Layer[i].ForwardNode
		}
	}

	for op.Layer[0].ForwardNode != nil && op.Layer[0].ForwardNode.Obj.Score <= end {
		res = append(res, *op.Layer[0].ForwardNode.Obj)
		op = op.Layer[0].ForwardNode
	}
	return res
}

func (zkl *ZSkipList) GetRevRangeByScore(start,end int64)[]Obj {
	res := make([]Obj,0)
	rank := make([]int64,zkl.Layers)

	op := zkl.Header
	for i := zkl.Layers - 1; i >= 0 ; i-- {
		if i == zkl.Layers - 1 {
			rank[i] = 0
		}else {
			rank[i] = rank[i+1]
		}

		for op.Layer[i].ForwardNode != nil && op.Layer[i].ForwardNode.Obj.Score <= end {
			rank[i] += op.Layer[i].Span
			op = op.Layer[i].ForwardNode
		}
	}


	for op.BackwardNode != nil && op.Obj.Score >= start {
		res = append(res, *op.Obj)
		op = op.BackwardNode
	}

	return res
}


func (zkl *ZSkipList) delNode(op *Node,  needUpdateLayer map[int]*Node)  {
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

func (zkl *ZSkipList) isRangeValid(min,max int64)bool {
	if min < 0 || (max != -1 && max < min) {
		return false
	}
	return true
}

func getRandLayer()int  {
	layers := 1
	for rand.Float32() < ZSLP {
		layers ++
	}
	if  layers < ZSLML {
		return layers
	}
	return ZSLML
}