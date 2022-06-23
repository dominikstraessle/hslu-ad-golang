package graph

type GraphA struct {
	numberOfNodes   int
	nodeName        []string
	adjacencyMatrix [][]bool
}

func NewGraphA(nodeName []string, adjacencyMatrix [][]bool) *GraphA {
	return &GraphA{
		numberOfNodes:   len(nodeName),
		nodeName:        nodeName,
		adjacencyMatrix: adjacencyMatrix,
	}
}

func (g *GraphA) NodeName(k int) string {
	return g.nodeName[k]
}

func (g *GraphA) NumberOfNodes() int {
	return g.numberOfNodes
}

func (g *GraphA) isEdge(i, j int) bool {
	return g.adjacencyMatrix[i][j]
}

type RailwayNet1 struct {
	GraphA
}

func NewRailwayNet1(nodeName []string, adjacencyMatrix [][]bool) *RailwayNet1 {
	return &RailwayNet1{
		GraphA{
			numberOfNodes:   len(nodeName),
			nodeName:        nodeName,
			adjacencyMatrix: adjacencyMatrix,
		},
	}
}

func (r *RailwayNet1) AddStation(a string) {
	r.nodeName = append(r.nodeName, a)
	var paths []bool
	for i := 0; i < len(r.adjacencyMatrix); i++ {
		paths = append(paths, false)
	}
	r.adjacencyMatrix = append(r.adjacencyMatrix, paths)

	for i, m := range r.adjacencyMatrix {
		r.adjacencyMatrix[i] = append(m, false)
	}

	r.numberOfNodes = len(r.adjacencyMatrix)
}

func (r *RailwayNet1) AddPath(a, b string) {
	i, j := r.getIndexForStation(a, b)
	r.adjacencyMatrix[i][j] = true
}

func (r *RailwayNet1) RemovePath(a, b string) {
	i, j := r.getIndexForStation(a, b)
	r.adjacencyMatrix[i][j] = false
}

func (r *RailwayNet1) NumberOfStations() int {
	return r.NumberOfNodes()
}

func (r *RailwayNet1) NumberOfPaths() int {
	sum := 0

	for _, i := range r.adjacencyMatrix {
		for _, isPath := range i {
			if isPath {
				sum++
			}
		}
	}

	return sum
}

func (r *RailwayNet1) HasDirectPath(a, b string) bool {
	i, j := r.getIndexForStation(a, b)
	return r.adjacencyMatrix[i][j] || r.adjacencyMatrix[j][i]
}

func (r *RailwayNet1) getIndexForStation(a string, b string) (int, int) {
	i, j := 0, 0
	for k, n := range r.nodeName {
		if n == a {
			i = k
		}
		if n == b {
			j = k
		}
	}
	return i, j
}
