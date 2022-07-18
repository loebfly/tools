package shell

type Enter struct {
	Firewall FirewallShell // 防火墙命令
	Nginx    NginxShell    // nginx命令
}
