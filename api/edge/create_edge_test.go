package edge

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var createEdgeAPI *CreateEdgeAPI

func createSetup() {
	edgesList := createObject()
	createEdgeAPI = NewCreate(edgesList, "edge-1")
}

func createObject() *Edges {
	edgesList := new(Edges)
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

func TestCreateMethod(t *testing.T) {
	createSetup()
	assert.Equal(t, http.MethodPost, createEdgeAPI.Method())
}

func TestCreateEndpoint(t *testing.T) {
	createSetup()
	assert.Equal(t, "/api/4.0/edges/edge-1/s/?action=patch", createEdgeAPI.Endpoint())
}

func TestCreateMarshalling(t *testing.T) {
	createSetup()
	expectedXML := "<s><><name>first</name><mtu>1500</mtu><type>internal</type><isConnected>true</isConnected><connectedToId>virtualwire-1</connectedToId><addressGroups></addressGroups><index>1</index></><><name>second</name><mtu>1500</mtu><type>internal</type><isConnected>true</isConnected><connectedToId>virtualwire-1</connectedToId><addressGroups></addressGroups><index>2</index></></s>"

	xmlBytes, err := xml.Marshal(createEdgeAPI.RequestObject())

	assert.Nil(t, err)
	assert.Equal(t, expectedXML, string(xmlBytes))
}
