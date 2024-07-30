package container

import (
	"os"
	"syscall"

	"log"
)

/*
*

	这里的init是在容器内部执行的，也就是说代码执行到这里，容器已经创建出来了，这是本容器执行的第一个进程
	使用mount先去挂载proc文件系统，以便使用ps查看系统当前进程情况
	MS_NOEXEC：本文件系统不允许运行其他程序
	MS_NOSUID：不允许set_user_id group_id
	MS_NODEV：default

*
*/
func RunContainerInitProcess(command string, args []string) error {
	log.Printf("command %s", command)
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		log.Print(err.Error())
	}
	return nil
}
