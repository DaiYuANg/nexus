package runtime

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"time"
)

func BuildRuntimeInfo() *SystemInfo {
	// 示例：获取 CPU 信息
	cpuUsage, _ := cpu.Percent(0, false)
	//cpuFreq, _ := cpu.Info()
	cpuCount, _ := cpu.Counts(false)
	// 示例：获取内存信息
	memStats, _ := mem.VirtualMemory()

	// 示例：获取磁盘信息
	diskStats, _ := disk.Usage("/")

	// 示例：获取网络信息
	netStats, _ := net.IOCounters(true)

	// 示例：获取进程信息
	//processes, _ := process.Processes()

	// 示例：获取启动时间
	uptime := time.Now().Unix()

	// 组装成结构体
	systemInfo := SystemInfo{
		CPU: CPUInfo{
			CPUCount: cpuCount,
			CPUUsage: cpuUsage[0],
			//CPUFrequency: float64(cpuFreq),
			CPUTimes: cpuUsage,
		},
		Memory: MemoryInfo{
			Total:       memStats.Total,
			Available:   memStats.Available,
			Used:        memStats.Used,
			UsedPercent: memStats.UsedPercent,
			SwapTotal:   memStats.SwapTotal,
			SwapUsed:    memStats.SwapCached,
			SwapFree:    memStats.SwapFree,
		},
		Disk: DiskInfo{
			DiskUsage: DiskUsageInfo{
				Total:       diskStats.Total,
				Used:        diskStats.Used,
				Free:        diskStats.Free,
				UsedPercent: diskStats.UsedPercent,
			},
		},
		Network: NetworkInfo{
			Interfaces: []NetworkInterfaceInfo{
				{
					Name:      netStats[0].Name,
					IP:        "192.168.1.1", // 示例 IP 地址
					Mac:       netStats[0].Name,
					TxBytes:   netStats[0].BytesSent,
					RxBytes:   netStats[0].BytesRecv,
					TxPackets: netStats[0].PacketsSent,
					RxPackets: netStats[0].PacketsRecv,
				},
			},
		},
		Processes: []ProcessInfo{
			{
				PID:        1,
				Name:       "init",
				CPUPercent: 0.1,
				MemPercent: 0.05,
			},
		},
		Sensors: SensorsInfo{
			Temperatures: []SensorTemperature{
				{
					SensorName:  "CPU",
					Temperature: 45.5,
				},
			},
		},
		BootTime: time.Unix(int64(uptime), 0),
	}
	return &systemInfo
}
