package runtime

import "time"

type SystemInfo struct {
	CPU       CPUInfo
	Memory    MemoryInfo
	Disk      DiskInfo
	Network   NetworkInfo
	Processes []ProcessInfo
	Sensors   SensorsInfo
	BootTime  time.Time
}

type CPUInfo struct {
	CPUCount     int       `json:"cpu_count"`     // CPU 核心数
	CPUTimes     []float64 `json:"cpu_times"`     // 各个 CPU 核心的使用时间（用户、系统、空闲等）
	CPUUsage     float64   `json:"cpu_usage"`     // CPU 总体使用率
	CPUFrequency float64   `json:"cpu_frequency"` // CPU 当前频率
}

type MemoryInfo struct {
	Total       uint64  `json:"total"`        // 总内存
	Available   uint64  `json:"available"`    // 可用内存
	Used        uint64  `json:"used"`         // 已使用内存
	UsedPercent float64 `json:"used_percent"` // 已使用的内存百分比
	SwapTotal   uint64  `json:"swap_total"`   // 总交换空间
	SwapUsed    uint64  `json:"swap_used"`    // 已使用交换空间
	SwapFree    uint64  `json:"swap_free"`    // 可用交换空间
}

type DiskInfo struct {
	Partitions []DiskPartitionInfo `json:"partitions"` // 磁盘分区信息
	DiskUsage  DiskUsageInfo       `json:"disk_usage"` // 磁盘使用情况
}

type DiskPartitionInfo struct {
	Device     string `json:"device"`      // 分区设备名
	MountPoint string `json:"mount_point"` // 挂载点
	Fstype     string `json:"fstype"`      // 文件系统类型
}

type DiskUsageInfo struct {
	Total       uint64  `json:"total"`        // 总磁盘空间
	Used        uint64  `json:"used"`         // 已使用磁盘空间
	Free        uint64  `json:"free"`         // 剩余磁盘空间
	UsedPercent float64 `json:"used_percent"` // 已使用磁盘空间百分比
}

type NetworkInfo struct {
	Interfaces []NetworkInterfaceInfo `json:"interfaces"` // 网络接口信息
}

type NetworkInterfaceInfo struct {
	Name      string `json:"name"`       // 网络接口名称
	IP        string `json:"ip"`         // 接口 IP 地址
	Mac       string `json:"mac"`        // 接口 MAC 地址
	TxBytes   uint64 `json:"tx_bytes"`   // 发送的字节数
	RxBytes   uint64 `json:"rx_bytes"`   // 接收的字节数
	TxPackets uint64 `json:"tx_packets"` // 发送的数据包数
	RxPackets uint64 `json:"rx_packets"` // 接收的数据包数
	TxErrors  uint64 `json:"tx_errors"`  // 发送错误数
	RxErrors  uint64 `json:"rx_errors"`  // 接收错误数
	TxDropped uint64 `json:"tx_dropped"` // 发送丢包数
	RxDropped uint64 `json:"rx_dropped"` // 接收丢包数
}

type ProcessInfo struct {
	PID        int32   `json:"pid"`         // 进程 ID
	Name       string  `json:"name"`        // 进程名称
	CPUPercent float64 `json:"cpu_percent"` // 进程 CPU 使用率
	MemPercent float64 `json:"mem_percent"` // 进程内存使用率
}

type SensorsInfo struct {
	Temperatures []SensorTemperature `json:"temperatures"` // 温度传感器信息
}

type SensorTemperature struct {
	SensorName  string  `json:"sensor_name"` // 温度传感器名称
	Temperature float64 `json:"temperature"` // 温度值（摄氏度）
}
