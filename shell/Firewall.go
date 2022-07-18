package shell

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
)

type FirewallShell struct{}

// AddPort /* 添加端口号 */
func (fw FirewallShell) AddPort(port string) error {
	err := exec.Command("firewall-cmd", "--zone=public", "--add-port="+port+"/tcp", "--permanent").Run()
	if err != nil {
		return errors.New("add port failed: " + err.Error())
	}
	err = exec.Command("firewall-cmd", "--reload").Run()
	if err != nil {
		return errors.New("firewall reload failed:" + err.Error())
	}
	return nil
}

// RemovePort /* 删除端口号 */
func (fw FirewallShell) RemovePort(port string) error {
	err := exec.Command("firewall-cmd", "--zone=public", "--remove-port="+port+"/tcp", "--permanent").Run()
	if err != nil {
		return errors.New("remove port failed: " + err.Error())
	}
	err = exec.Command("firewall-cmd", "--reload").Run()
	if err != nil {
		return errors.New("firewall reload failed:" + err.Error())
	}
	return nil
}

// AddPortForwarding /* 添加端口转发 */
func (fw FirewallShell) AddPortForwarding(srcPort, dstIp, dstPort string) error {
	_ = exec.Command("firewall-cmd", "--add-masquerade").Run()
	err := exec.Command("firewall-cmd", fmt.Sprintf("--add-forward-port=port=%s:proto=tcp:toport=%s:toaddr=%s", srcPort, dstPort, dstIp)).Run()
	if err != nil {
		return errors.New("add port forwarding failed: " + err.Error())
	}
	err = exec.Command("firewall-cmd", "--runtime-to-permanent").Run()
	if err != nil {
		return errors.New("firewall reload failed:" + err.Error())
	}
	return nil
}

// RemovePortForwarding /* 删除端口转发 */
func (fw FirewallShell) RemovePortForwarding(srcPort, dstIp, dstPort string) error {
	err := exec.Command("firewall-cmd", fmt.Sprintf("--remove-forward-port=port=%s:proto=tcp:toport=%s:toaddr=%s", srcPort, dstPort, dstIp)).Run()
	if err != nil {
		return errors.New("remove port forwarding failed: " + err.Error())
	}
	err = exec.Command("firewall-cmd", "--runtime-to-permanent").Run()
	if err != nil {
		return errors.New("firewall reload failed:" + err.Error())
	}
	return nil
}

// IsPortUsed /* 检查端口是否被占用 */
func (fw FirewallShell) IsPortUsed(proto string, port int) bool {
	if proto != "tcp" && proto != "udp" {
		return true
	}
	if proto == "tcp" {
		_, err := net.Dial(proto, fmt.Sprintf(":%d", port))
		if err != nil {
			return false
		}
	} else {
		udpAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
		if err != nil {
			return true
		}
		listener, err := net.ListenUDP("udp", udpAddress)
		if err == nil {
			listener.Close()
			return false
		}
	}
	return true
}
