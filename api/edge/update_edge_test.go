package edge

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var updateEdgeAPI *UpdateEdgeAPI

func setupUpdate() {
	edge := Edge{Name: "foo", Mtu: 1500}
	updateEdgeAPI = NewUpdate("edge-1", 1, edge)
}

func TestUpdateMethod(t *testing.T) {
	setupUpdate()
	assert.Equal(t, http.MethodPut, updateEdgeAPI.Method())
}

func TestUpdateEndpoint(t *testing.T) {
	setupUpdate()
	assert.Equal(t, "/api/4.0/edges/edge-1/s/1", updateEdgeAPI.Endpoint())
}
