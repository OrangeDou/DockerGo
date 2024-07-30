package main

import (
	"fmt"

	"dockergo/container"

	"log"

	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `Create a container with namespace and cgroups limit ie: mydocker run -ti [image] [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	// run function
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

// initCommand 内部方法，禁止外部调用
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",

	// 获取command参数初始化
	Action: func(context *cli.Context) error {
		log.Print("init come on")
		cmd := context.Args().Get(0)
		log.Print("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
