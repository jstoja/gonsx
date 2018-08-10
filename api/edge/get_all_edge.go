package edge

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// GetAllEdgesAPI base object.
type GetAllEdgesAPI struct {
	*api.BaseAPI
}

// NewGetAll returns the api object of GetAllEdgesAPI
func NewGetAll(edgeID string) *GetAllEdgesAPI {
	this := new(GetAllEdgesAPI)
	this.BaseAPI = api.NewBaseAPI(http.MethodGet, "/api/4.0/edges/"+edgeID+"/s", nil, new(Edges))
	return this
}

// GetResponse returns ResponseObject of GetAllEdgesAPI
func (getAllAPI GetAllEdgesAPI) GetResponse() *Edges {
	return getAllAPI.ResponseObject().(*Edges)
}
