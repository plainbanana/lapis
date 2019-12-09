package entities

import "encoding/xml"

// LapisVersion is version
var LapisVersion = "1.1.0"

// LapisPort is port of lapis
var LapisPort string

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

// Guide : xml
type Guide struct {
	XMLName           xml.Name         `xml:"tv"`
	GeneratorInfoName string           `xml:"generator-info-name,attr"`
	Channel           []ChannelGuide   `xml:"channel"`
	Programme         []ProgrammeGuide `xml:"programme"`
}

// ChannelGuide : xml
type ChannelGuide struct {
	ID          string           `xml:"id,attr"`
	DisplayName DisplayNameGuide `xml:"display-name"`
}

// DisplayNameGuide : xml
type DisplayNameGuide struct {
	Lang        string `xml:"lang,attr"`
	DisplayName string `xml:",chardata"`
}

// ProgrammeGuide : xml
type ProgrammeGuide struct {
	Start      string          `xml:"start,attr"`
	Stop       string          `xml:"stop,attr"`
	Channel    string          `xml:"channel,attr"`
	Title      TitleGuide      `xml:"title"`
	SubTitle   TitleGuide      `xml:"sub-title"`
	Desc       DescGuide       `xml:"desc"`
	Category   CategoryGuide   `xml:"category"`
	EpisodeNum EpisodeNumGuide `xml:"episode-num"`
}

// TitleGuide : xml
type TitleGuide struct {
	Lang  string `xml:"lang,attr"`
	Title string `xml:",chardata"`
}

// DescGuide : xml
type DescGuide struct {
	Lang string `xml:"lang,attr"`
	Desc string `xml:",chardata"`
}

// CategoryGuide : xml
type CategoryGuide struct {
	Lang     string `xml:"lang,attr"`
	Category string `xml:",chardata"`
}

// EpisodeNumGuide : xml
type EpisodeNumGuide struct {
	System     string `xml:"system,attr"`
	EpisodeNum string `xml:",chardata"`
}

// Schedule : json
type Schedule struct {
	ID        string          `json:"id"`
	Category  string          `json:"category"`
	Title     string          `json:"title"`
	FullTitle string          `json:"fulltitle"`
	Detail    string          `json:"detail"`
	Start     int64           `json:"start"`
	End       int64           `json:"end"`
	Seconds   int             `json:"seconds"`
	Desc      string          `json:"description"`
	Extra     ExtraSchedule   `json:"extra"`
	Channel   ChannelSchedule `json:"channel"`
	SubTitle  string          `json:"subTitle"`
	Episode   int             `json:"episode"`
}

// Schedules : json
type Schedules []Schedule

// ExtraSchedule : json
type ExtraSchedule struct {
	Desc  string `json:"番組内容"`
	Act   string `json:"出演者"`
	Music string `json:"原作・脚本"`
}

// ChannelSchedule : json
type ChannelSchedule struct {
	SID  int    `json:"sid"`
	Name string `json:"name"`
}
