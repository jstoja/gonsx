package edge

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var getAllEdgesAPI *GetAllEdgesAPI

func setupGetAll() {
	getAllEdgesAPI = NewGetAll("edge-1")
}

func TestGetAllMethod(t *testing.T) {
	setupGetAll()
	assert.Equal(t, http.MethodGet, getAllEdgesAPI.Method())
}

func TestGetAllEndpoint(t *testing.T) {
	setupGetAll()
	assert.Equal(t, "/api/4.0/edges/edge-1/s", getAllEdgesAPI.Endpoint())
}

func TestGetAllUnMarshalling(t *testing.T) {
	setupGetAll()
	xmlContent := []byte("<s><><name>Name</name></><><name>Name2</name></></s>")

	xmlerr := xml.Unmarshal(xmlContent, getAllEdgesAPI.ResponseObject())

	assert.Nil(t, xmlerr)
	assert.Len(t, getAllEdgesAPI.GetResponse().s, 2)
	assert.Equal(t, "Name", getAllEdgesAPI.GetResponse().s[0].Name)
}
