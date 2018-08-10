package edge

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// DeleteEdgeAPI struct
type DeleteEdgeAPI struct {
	*api.BaseAPI
}

// NewDelete returns a new delete method object of DeleteEdgeAPI
func NewDelete(edgeID string) *DeleteEdgeAPI {
	this := new(DeleteEdgeAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodDelete, "/api/4.0/edges/"+edgeID, nil, nil)
	return this
}
