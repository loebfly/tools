package ipT

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// IsCorrect 判断是否是IP地址
func (Enter) IsCorrect(ip string) bool {
	reg, _ := regexp.Compile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`)
	return reg.MatchString(ip)
}

// GetLocal return local ip address
func (Enter) GetLocal() string {
	address, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("GetLocal err:", err)
		return ""
	}
	for _, value := range address {
		if ipNet, ok := value.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.String()[:7] != "169.254" {
			if ipNet.IP.To4() != nil {
				ip := ipNet.IP.String()
				return ip
			}
		}
	}

	ip := "127.0.0.1"
	return ip
}

// GetV4s return all non-loop back IPv4 addresses
func (Enter) GetV4s() ([]string, error) {
	var ips []string
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}

func (Enter) GetV4sByInterface(name string) ([]string, error) {
	var ips []string

	iFace, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}

	address, err := iFace.Addrs()
	if err != nil {
		return nil, err
	}

	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}

// IsIntranet return true if ip is intranet
func (Enter) IsIntranet(ip string) bool {
	if strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "192.168.") {
		return true
	}

	if strings.HasPrefix(ip, "172.") {
		// 172.16.0.0-172.31.255.255
		arr := strings.Split(ip, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

// IsPortUsed return true if port is used
func (Enter) IsPortUsed(port int) bool {
	sysType := runtime.GOOS
	var (
		output         []byte
		checkStatement string
	)
	if sysType == "linux" {
		checkStatement = fmt.Sprintf("netstat -anp | grep %d ", port)
		output, _ = exec.Command("sh", "-c", checkStatement).CombinedOutput()
	} else if sysType == "windows" {
		checkStatement = fmt.Sprintf("netstat -ano -p tcp | findstr %d", port)
		output, _ = exec.Command("cmd", "/c", checkStatement).CombinedOutput()
	} else {
		return false
	}

	if len(output) > 0 {
		return false
	}
	return true
}
