package response

type DeviceInfo struct {
	SN             string
	UUID           string
	Name           string
	ModelName      string
	ModelID        string
	ProcessorName  string
	ProcessorSpeed string
	Manufacturer   string
	OSName         string
	OSVersion      string
	HardwareUUID   string
	Status         int
	Uptime         int64
}

// type ListDevices struct {
// 	PageResponse
// 	Devices []DeviceInfo
// }
