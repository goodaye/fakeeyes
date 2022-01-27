package request

// DeviceInfo 设备信息
type DeviceInfo struct {
	SN             string `json:"sn"  validate:"required"`
	Name           string `json:"name" validate:"required"`
	ModelName      string `json:"model_name" `
	ModelID        string `json:"model_id" `
	ProcessorName  string `json:"processor_name" `
	ProcessorSpeed string `json:"processor_speed" `
	Manufacturer   string `json:"manufacturer" `
	OSName         string `json:"os_name" `
	OSVersion      string `json:"os_version" `
	HardwareUUID   string `json:"hardware_uuid"`
}

type DescribeDevice struct {
	UUID string `json:"uuid"`
	SN   string `json:"sn"`
}
