package request

// DeviceInfo 设备信息
type DeviceInfo struct {
	SN           string `json:"sn"  validate:"required"`
	Name         string `json:"name" validate:"required"`
	ModelName    string `json:"model_name" `
	ModelID      string `json:"model_id" `
	CPUModel     string `json:"cpu_model" `
	CPUModelID   string `json:"cpu_model_id"`
	CPUSpeed     string `json:"cpu_speed" `
	CPUArch      string `json:"cpu_arch" `
	CPUVendor    string `json:"cpu_vendor" `
	CPUCores     int    `json:"cpu_cores" `
	CPUSocket    int    `json:"cpu_socket" `
	Manufacturer string `json:"manufacturer" `
	OSName       string `json:"os_name" `
	OSVersion    string `json:"os_version" `
	HardwareUUID string `json:"hardware_uuid"`
}

type DescribeDevice struct {
	UUID string `json:"uuid"`
	SN   string `json:"sn"`
}
