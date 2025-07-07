package metrics

import (
	"fmt"
	"log/slog"
	"os"
	"hotel-service/logger"

	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
)

var (
	numCPU int
)

func init() {
	numCPU = runtime.NumCPU()
}

type cpuMetrics struct {
	SystemPerc float64 `json:"sysCpuPerc"`
	AppPerc    float64 `json:"appCpuPerc"`
	Cores      int     `json:"cores"`
}

type memoryMetrics struct {
	SystemPerc  float64 `json:"sysMemPerc"`
	SystemUsed  uint64  `json:"sysMemUsedMb"`
	SystemTotal uint64  `json:"sysMemTotalMb"`
	AppPerc     float64 `json:"appMemPerc"`
	AppUsed     uint64  `json:"appMemUsedMb"`
}

type applicationStats struct {
	Goroutines int `json:"goroutines"`
	Threads    int `json:"threads"`
}

type systemMetrics struct {
	CPUPerc   float64
	CoreCount int
	MemPerc   float64
	MemTotal  uint64
	MemUsed   uint64
}

type applicationMetrics struct {
	CPUPerc     float64
	MemPerc     float64
	MemUsed     uint64
	Goroutines  int
	ThreadCount int
}

func getSystemMetrics() (*systemMetrics, error) {
	cpuPerc, err := cpu.Percent(0, false)
	if err != nil {
		return nil, fmt.Errorf("error getting CPU usage: %v", err)
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error getting memory stats: %v", err)
	}

	return &systemMetrics{
		CPUPerc:   roundToTwoDecimals(cpuPerc[0]),
		CoreCount: numCPU,
		MemPerc:   roundToTwoDecimals(vmStat.UsedPercent),
		MemTotal:  vmStat.Total,
		MemUsed:   vmStat.Used,
	}, nil
}

func getApplicationMetrics() (*applicationMetrics, error) {
	proc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return nil, fmt.Errorf("error getting process: %v", err)
	}

	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		return nil, fmt.Errorf("error getting process CPU percent: %v", err)
	}

	memInfo, err := proc.MemoryInfo()
	if err != nil {
		return nil, fmt.Errorf("error getting process memory info: %v", err)
	}

	memPercent, err := proc.MemoryPercent()
	if err != nil {
		return nil, fmt.Errorf("error getting memory percent: %v", err)
	}

	numThreads, err := proc.NumThreads()
	if err != nil {
		return nil, fmt.Errorf("error getting thread count: %v", err)
	}

	return &applicationMetrics{
		CPUPerc:     roundToTwoDecimals(cpuPercent),
		MemPerc:     float64(memPercent),
		MemUsed:     memInfo.RSS,
		Goroutines:  runtime.NumGoroutine(),
		ThreadCount: int(numThreads),
	}, nil
}
func roundToTwoDecimals(num float64) float64 {
	return float64(int64(num*100)) / 100
}
func bytesToMegabytes(bytes uint64) uint64 {
	return bytes / 1024 / 1024
}

type Metrics struct {
	Version  string           `json:"version"`
	CPU      cpuMetrics       `json:"cpu"`
	Memory   memoryMetrics    `json:"memory"`
	AppStats applicationStats `json:"appStats"`
}

func CollectMetrics(version string) (*Metrics, error) {
	sysMetrics, err := getSystemMetrics()
	if err != nil {
		slog.Error("Error collecting system metrics", "error", err)
		return nil, err
	}

	appMetrics, err := getApplicationMetrics()
	if err != nil {
		slog.Error("Error collecting application metrics", "error", err)
		return nil, err
	}

	metrics := &Metrics{
		Version: version,
		CPU: cpuMetrics{
			SystemPerc: sysMetrics.CPUPerc,
			AppPerc:    appMetrics.CPUPerc,
			Cores:      sysMetrics.CoreCount,
		},
		Memory: memoryMetrics{
			SystemPerc:  sysMetrics.MemPerc,
			SystemUsed:  bytesToMegabytes(sysMetrics.MemUsed),
			SystemTotal: bytesToMegabytes(sysMetrics.MemTotal),
			AppPerc:     roundToTwoDecimals(appMetrics.MemPerc),
			AppUsed:     bytesToMegabytes(appMetrics.MemUsed),
		},
		AppStats: applicationStats{
			Goroutines: appMetrics.Goroutines,
			Threads:    appMetrics.ThreadCount,
		},
	}

	slog.Info("Metrics", logger.Extra(metrics))

	return metrics, nil
}
