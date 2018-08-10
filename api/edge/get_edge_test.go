package edge

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getEdgeAPI *GetEdgeAPI

func setupGet() {
	getEdgeAPI = NewGet("edge-1", 10)
}

func TestGetMethod(t *testing.T) {
	setupGet()
	assert.Equal(t, http.MethodGet, getEdgeAPI.Method())
}

func TestGetEndpoint(t *testing.T) {
	setupGet()
	assert.Equal(t, "/api/4.0/edges/edge-1/s/10", getEdgeAPI.Endpoint())
}

func TestGetUnMarshalling(t *testing.T) {
	setupGet()
	xmlContent := []byte("<><name>Name</name></>")

	xmlerr := xml.Unmarshal(xmlContent, getEdgeAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Equal(t, "Name", getEdgeAPI.GetResponse().Name)
}
