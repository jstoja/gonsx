package edge

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
	"strconv"
)

// GetEdgeAPI base object.
type GetEdgeAPI struct {
	*api.BaseAPI
}

// NewGet returns the api object of GetEdgeAPI
func NewGet(edgeID string, index int) *GetEdgeAPI {
	this := new(GetEdgeAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/4.0/edges/"+edgeID+"/s/"+strconv.Itoa(index), nil, new(Edge))
	return this
}

// GetResponse returns ResponseObject of GetEdgeAPI
func (getAPI GetEdgeAPI) GetResponse() Edge {
	return *getAPI.ResponseObject().(*Edge)
}
