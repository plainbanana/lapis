package entities

import "encoding/xml"

// Device : device info
type Device struct {
	FriendlyName string `json:"FriendlyName"`
	Manufacturer string `json:"Manufacturer"`
	ModelNumber  string `json:"ModelNumber"`
	FirmwareName string `json:"FirmwareName"`
	DeviceID     string `json:"DeviceID"`
	DeviceAuth   string `json:"DeviceAuth"`
	BaseURL      string `json:"BaseURL"`
	LineupURL    string `json:"LineupURL"`
}

// Lineup : tv programs lineup
type Lineup struct {
	GuideNumber string `json:"GuideNumber"`
	GuideName   string `json:"GuideName"`
	URL         string `json:"URL"`
}

// DeviceXML :
type DeviceXML struct {
	XMLName     xml.Name      `xml:"root"`
	Xmlns       string        `xml:"xmlns,attr"`
	SpecVersion []SpecChild   `xml:"specVersion"`
	URLBase     string        `xml:"URLBace"`
	Device      []DeviceChild `xml:"device"`
}

// SpecChild :
type SpecChild struct {
	Major int `xml:"major"`
	Minor int `xml:"minor"`
}

// DeviceChild :
type DeviceChild struct {
	DeviceType   string `xml:"deviceType"`
	FriendlyName string `xml:"friendlyName"`
	Manufacturer string `xml:"manufacturer"`
	ModelName    string `xml:"modelName"`
	ModelNumber  string `xml:"modelNumber"`
	SerialNumber string `xml:"serialNumber"`
	UDN          string `xml:"UDN"`
}

// Service : json
type Service struct {
	ID        int    `json:"id"`
	ServiceID int    `json:"serviceId"`
	NetworkID int    `json:"networkId"`
	Name      string `json:"name"`
}

// Channel : json
type Channel struct {
	Type     string    `json:"type"`
	Channel  string    `json:"channel"`
	Name     string    `json:"name"`
	Services []Service `json:"services"`
}

// Channels : json
type Channels []Channel
