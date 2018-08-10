package edge

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setup() (edgesList *Edges) {
	edgesList = &Edges{}
	first := Edge{
		Name:          "first",
		ConnectedToID: "virtualwire-1",
		Type:          "internal",
		Index:         1,
		Mtu:           1500,
		IsConnected:   true,
	}
	second := Edge{
		Name:          "second",
		ConnectedToID: "virtualwire-1",
		Type:          "internal",
		Index:         2,
		Mtu:           1500,
		IsConnected:   true,
	}
	edgesList.s = []Edge{first, second}
	return edgesList
}

func TestFilterByName(t *testing.T) {
	edges := setup()

	firstFiltered := edges.FilterByName("first")
	assert.Equal(t, 1, firstFiltered.Index)

	secondFiltered := edges.FilterByName("second")
	assert.Equal(t, 2, secondFiltered.Index)
}

func TestStringImplementation(t *testing.T) {
	edge := setup()
	expectedOutput := "[{{ } first  1500 internal true virtualwire-1 {[]} 1} {{ } second  1500 internal true virtualwire-1 {[]} 2}]"
	assert.Equal(t, expectedOutput, fmt.Sprint(edge))
}
