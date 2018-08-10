package edge

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var deleteEdgeAPI *DeleteEdgeAPI

func setupDelete() {
	deleteEdgeAPI = NewDelete("edge-1", 1)
}

func TestDeleteMethod(t *testing.T) {
	setupDelete()
	assert.Equal(t, http.MethodDelete, deleteEdgeAPI.Method())
}

func TestDeleteEndpoint(t *testing.T) {
	setupDelete()
	assert.Equal(t, "/api/4.0/edges/edge-1/s/1", deleteEdgeAPI.Endpoint())
}
