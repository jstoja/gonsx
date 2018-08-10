package edge

import "fmt"

func (s Edges) String() string {
	return fmt.Sprintf("%v", s.s)
}

//func (s Edge) String() string {
//	return fmt.Sprintf("index: %s, name: %s", s.Index, s.Name)
//}

// FilterByName filters Edges list by given name of  and returns
// the found  object.
func (s Edges) FilterByName(name string) *Edge {
	var edgeFound Edge
	for _, edge := range s.s {
		if edge.Name == name {
			edgeFound = edge
			break
		}
	}
	return &edgeFound
}
