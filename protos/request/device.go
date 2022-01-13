package request

type DeviceInfo struct {
	SN             string `json:"sn" xorm:"not null unique VARCHAR(255) comment('设备SN')"`
	Name           string `json:"name" xorm:"not null  VARCHAR(255) comment('设备名')"`
	ModelName      string `json:"model_name" xorm:" null  VARCHAR(255) comment('设备名,比如树莓派/MacBook')"`
	ModelID        string `json:"model_id" xorm:" null  VARCHAR(255) comment('设备名,比如树莓派4B/MacBook16,1')"`
	ProcessorName  string `json:"processor_name" xorm:" null  VARCHAR(255) comment('处理器名字')"`
	ProcessorSpeed string `json:"processor_speed" xorm:" null  VARCHAR(255) comment('处理器速度')"`
	Manufacturer   string `json:"manufacturer" xorm:" null  VARCHAR(255) comment('设备制造商,比如,亚博')"`
	OSName         string `json:"os_name" xorm:" null  VARCHAR(255) comment('OS名字')"`
	OSVersion      string `json:"os_version" xorm:" null  VARCHAR(255) comment('OS Version')"`
	HardwareUUID   string `json:"hardware_uuid" xorm:" null  VARCHAR(255) comment('硬件UUID')"`
}
