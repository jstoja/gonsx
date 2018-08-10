package edge

import (
	"github.com/tadaweb/gonsx/api"
	"net/http"
)

// CreateEdgeAPI struct
type CreateEdgeAPI struct {
	*api.BaseAPI
}

// NewCreate returns a new object of CreateEdgeAPI
func NewCreate(edgeList *Edges, edgeID string) *CreateEdgeAPI {

	this := new(CreateEdgeAPI)

	this.BaseAPI = api.NewBaseAPI(http.MethodPost, "/api/4.0/edges", edgeList, new(Edges))
	return this
}

// GetResponse returns the ResponseObject.
func (createAPI CreateEdgeAPI) GetResponse() *Edges {
	return createAPI.ResponseObject().(*Edges)
}
