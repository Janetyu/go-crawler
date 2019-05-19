package monitor

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/disk"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func AllCheck(c *gin.Context)  {
	var msg []byte
	healthMsg := "Bad"
	diskMsg := "Bad"
	cpuMsg := "Bad"
	ramMsg := "Bad"
	resp, err := http.Get( "http://127.0.0.1:8989/sd/health")
	if err == nil && resp.StatusCode == 200 {
		msg, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			healthMsg = string(msg)
		}
	}

	resp, err = http.Get( "http://127.0.0.1:8989/sd/disk")
	if err == nil && resp.StatusCode == 200 {
		msg, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			diskMsg = string(msg)
		}
	}

	resp, err = http.Get( "http://127.0.0.1:8989/sd/cpu")
	if err == nil && resp.StatusCode == 200 {
		msg, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			cpuMsg = string(msg)
		}
	}

	resp, err = http.Get( "http://127.0.0.1:8989/sd/ram")
	if err == nil && resp.StatusCode == 200 {
		msg, err = ioutil.ReadAll(resp.Body)
		if err == nil {
			ramMsg = string(msg)
		}
	}

	data := map[string]interface{}{
		"health": healthMsg,
		"disk": diskMsg,
		"cpu": cpuMsg,
		"ram": ramMsg,
	}

	// JSONP解决跨域问题
	c.JSONP(http.StatusOK, data)
}

// HealthCheck Ping-Pong 后的结果显示'OK' 服务器健康情况
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, message)
}

// DiskCheck 服务器硬盘情况
func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}

// CPUCheck 服务器CPU情况
func CPUCheck(c *gin.Context) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | Cores: %d", text, l1, l5, l15, cores)
	c.String(status, message)
}

// RAMCheck 服务器内存使用量
func RAMCheck(c *gin.Context) {
	u, _ := mem.VirtualMemory()

	usedMB := int(u.Used) / MB
	usedGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	c.String(status, message)
}

