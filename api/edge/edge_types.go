package edge

import "encoding/xml"

// Edges top level list object
type Edges struct {
	XMLName xml.Name `xml:"s"`
	s       []Edge   `xml:""`
}

// Edge object within Edges list.
type Edge struct {
	XMLName       xml.Name      `xml:""`
	Name          string        `xml:"name"`
	Label         string        `xml:"label,omitempty"`
	Mtu           int           `xml:"mtu"`
	Type          string        `xml:"type"`
	IsConnected   bool          `xml:"isConnected"`
	ConnectedToID string        `xml:"connectedToId"`
	AddressGroups AddressGroups `xml:"addressGroups"`
	Index         int           `xml:"index,omitempty"`
}

// AddressGroups within Edge.
type AddressGroups struct {
	AddressGroups []AddressGroup `xml:"addressGroup"`
}

// AddressGroup object within AddressGroup list.
type AddressGroup struct {
	XMLName        xml.Name `xml:"addressGroup"`
	PrimaryAddress string   `xml:"primaryAddress"`
	SubnetMask     string   `xml:"subnetMask"`
}
