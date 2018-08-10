package edge

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
	"strconv"
)

// UpdateEdgeAPI struct
type UpdateEdgeAPI struct {
	*api.BaseAPI
}

// NewUpdate returns a new delete method object of UpdateEdgeAPI
func NewUpdate(edgeID string, index int, edge Edge) *UpdateEdgeAPI {
	this := new(UpdateEdgeAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodPut, "/api/4.0/edges/"+edgeID+"/s/"+strconv.Itoa(index), edge, nil)
	return this
}
