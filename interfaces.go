// Copyright © 2019 Developer Network, LLC
//
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package graph

// Edge is the interface that defines an edge in the graph
type Edge interface {
	Parent() *Node
	Child() *Node
	Value() interface{}
}

// DirectedEdge is the interface that defines a directional edge in the graph
type DirectedEdge interface {
	Edge
	Directional() bool
}

// DirectedWeightedEdge is the interface that defines a directional weighted edge in the graph
type DirectedWeightedEdge interface {
	DirectedEdge
	Weight() float64
}

// WeightedEdge is the interface that defines an undirected weighted edge in the graph
type WeightedEdge interface {
	Edge
	Weight() float64
}
