package shell

import (
	"errors"
	"os/exec"
)

type NginxShell struct{}

// Restart 重启nginx
func (NginxShell) Restart() error {
	err := exec.Command("systemctl", "restart", "nginx").Run()
	if err != nil {
		return errors.New("重启nginx失败:" + err.Error())
	}
	return nil
}
