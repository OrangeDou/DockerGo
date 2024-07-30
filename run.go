package main

import (
	"dockergo/container"
	"os"

	log "github.com/sirupsen/logrus"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	// 等待command结束
	parent.Wait()
	os.Exit(-1)
}
