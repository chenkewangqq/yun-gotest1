package main

import (
	"flag"
	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"
)

func main() {
	var user string
	var passwd string
	var ip string
	var port string
	var ipandport string
	var action string
	klog.InitFlags(nil)
	flag.StringVar(&user, "user", "root", "输入ssh的用户名")
	flag.StringVar(&passwd, "passwd", "root", "输入ssh的密码")
	flag.StringVar(&ip, "ip", "", "输入ssh的ip地址")
	flag.StringVar(&port, "port", "22", "输入ssh的端口号")
	flag.StringVar(&action, "action", "", "输入执行动作 : restart 重启程序( 需要数参数 -user  xxx -passwd xxx  -ip xx -port xxx -action restart ) ")
	flag.Parse()
	defer klog.Flush()
	if action == "" {
		klog.Fatal("action is null , see help : ./pangolin-ssh-tools.exe -help ")
	}
	ipandport = ip + ":" + port
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", ipandport, config)
	if err != nil {
		klog.Fatal(err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		klog.Fatal(err)
	}
	defer session.Close()
	switch action {
	case "restart":
		res, err := session.CombinedOutput("cd /mnt/workspace/pangolin/bin/ ; /mnt/workspace/pangolin/bin/restart_pangolin.sh")
		if err != nil {
			klog.Fatal(err)
		}
		klog.Info("重启命令执行成功")
		klog.Info(string(res))
	case "ping":

	}

}
