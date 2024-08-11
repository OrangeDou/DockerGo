package main

import (
	"fmt"

	"gocker/cgroups/subsystems"
	"gocker/container"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `Create a container with namespace and cgroups limit ie: mydocker run -ti [image] [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			// 交互式终端
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		// 内存限制
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		// CPU权重
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		// 容器名
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
	},
	// run function
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}

		//get image name
		imageName := cmdArray[0]
		cmdArray = cmdArray[1:]

		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}
		tty := context.Bool("ti")
		containerName := context.String("name")
		volume := context.String("v")
		Run(tty, cmdArray, resConf, containerName, volume, imageName)
		return nil
	},
}

// initCommand 内部方法，禁止外部调用
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",

	// 获取command参数初始化
	Action: func(context *cli.Context) error {
		log.Info("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}

var listCommand = cli.Command{
	Name:  "ps",
	Usage: "list all containers",
	Action: func(context *cli.Context) error {
		container.ListContainers()
		return nil
	},
}
