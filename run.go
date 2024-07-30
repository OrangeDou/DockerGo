package main

import (
	"gocker/container"
	"os"

	log "github.com/sirupsen/logrus"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	// 等待command 结束
	parent.Wait()
	os.Exit(-1)
}
