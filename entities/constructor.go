package entities

// NewDevice : init
func NewDevice() Device {
	dev := Device{}
	dev.FriendlyName = "lapis"
	dev.Manufacturer = "Silicondust"
	// Built in Transcoder
	dev.ModelNumber = "HDTC-2US"
	dev.FirmwareName = "hdhomeruntc_atsc"
	dev.DeviceID = "12345678"
	dev.DeviceAuth = ""
	dev.BaseURL = ""
	dev.LineupURL = ""
	return dev
}

// NewDeviceXML : init /device.xml
func NewDeviceXML() *DeviceXML {
	d := NewDevice()
	v := &DeviceXML{
		Xmlns:   "urn:schemas-upnp-org:device-1-0",
		URLBase: d.BaseURL,
	}
	v.SpecVersion = append(v.SpecVersion, SpecChild{
		Major: 1,
		Minor: 0,
	})
	v.Device = append(v.Device, DeviceChild{
		DeviceType:   "urn:schemas-upnp-org:device:MediaServer:1",
		FriendlyName: d.FriendlyName,
		Manufacturer: d.Manufacturer,
		ModelName:    d.ModelNumber,
		ModelNumber:  d.ModelNumber,
		SerialNumber: "001",
		UDN:          d.DeviceID,
	})

	return v
}
