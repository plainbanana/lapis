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
type Service []struct {
	ID                 int64  `json:"id"`
	ServiceID          int    `json:"serviceId"`
	NetworkID          int    `json:"networkId"`
	Name               string `json:"name"`
	Type               int    `json:"type"`
	LogoID             int    `json:"logoId"`
	RemoteControlKeyID int    `json:"remoteControlKeyId"`
	Channel            struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
	} `json:"channel"`
	HasLogoData bool `json:"hasLogoData"`
}

// Channel : json
type Channel struct {
	Type     string `json:"type"`
	Channel  string `json:"channel"`
	Name     string `json:"name"`
	Services []struct {
		ID        int    `json:"id"`
		ServiceID int    `json:"serviceId"`
		NetworkID int    `json:"networkId"`
		Name      string `json:"name"`
	} `json:"services"`
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
	ID          string `xml:"id,attr"`
	DisplayName struct {
		Lang        string `xml:"lang,attr"`
		DisplayName string `xml:",chardata"`
	} `xml:"display-name"`
	Icon struct {
		Src string `xml:"src,attr"`
	} `xml:"icon"`
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

// MirakurunPrograms : json
type MirakurunPrograms []struct {
	ID          int64  `json:"id"`
	EventID     int    `json:"eventId"`
	ServiceID   int    `json:"serviceId"`
	NetworkID   int    `json:"networkId"`
	StartAt     int64  `json:"startAt"`
	Duration    int    `json:"duration"`
	IsFree      bool   `json:"isFree"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Video       struct {
		Type          string `json:"type"`
		Resolution    string `json:"resolution"`
		StreamContent int    `json:"streamContent"`
		ComponentType int    `json:"componentType"`
	} `json:"video"`
	Audio struct {
		SamplingRate  int `json:"samplingRate"`
		ComponentType int `json:"componentType"`
	} `json:"audio"`
	Genres []struct {
		Lv1 int `json:"lv1"`
		Lv2 int `json:"lv2"`
		Un1 int `json:"un1"`
		Un2 int `json:"un2"`
	} `json:"genres,omitempty"`
	RelatedItems []struct {
		ServiceID int `json:"serviceId"`
		EventID   int `json:"eventId"`
	} `json:"relatedItems,omitempty"`
	Extended struct {
		ProgrammeContent   string `json:"番組内容"`
		Performer          string `json:"出演者"`
		OriginalScreenplay string `json:"原作・脚本"`
		Director           string `json:"監督・演出"`
		Production         string `json:"制作"`
		Music              string `json:"音楽"`
	} `json:"extended,omitempty"`
}
