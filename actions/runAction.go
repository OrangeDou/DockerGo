package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

// 解析docker run命令
func RunAction(context *cli.Context) error {
	if len(context.Args()) < 1 {
		return fmt.Errorf("Missing container command")
	}
	// 获取参数列表
	var cmdArray []string
	for _, arg := range context.Args() {
		cmdArray = append(cmdArray, arg)
	}

	// // image name
	// imageName := cmdArray[0]
	// cmdArray = cmdArray[1:]

	// 从cli.Context中获取名为ti的布尔标志的值。ti标志通常用于分配一个伪终端（pseudo-TTY）
	createTty := context.Bool("ti")
	// 从cli.Context中获取名为d的布尔标志的值，用于控制容器是否在后台运行。
	detach := context.Bool("d")

	if createTty && detach {
		return fmt.Errorf("ti and d paramter can not both provided")
	}

	cmd := context.Args().Get(0)
	Run(createTty, cmd)
	return nil
}
