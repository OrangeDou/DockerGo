package main

import (
	"gocker/cgroups"
	"gocker/cgroups/subsystems"
	"gocker/container"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	// docker-go as a cgroup name
	// 创建cgroupmanager，调用set和apply设置资源限制并使限制在容器上生效
	cgroupManager := cgroups.NewCgroupManager("docker-go")
	defer cgroupManager.Destory()

	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)
	// 设置完限制后，初始化容器
	sendInitCommand(comArray, writePipe)
	// 等待command 结束
	parent.Wait()
	os.Exit(-1)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
