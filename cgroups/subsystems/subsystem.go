package subsystems

type ResourceConfig struct {
	MemoryLimit string // 内存
	CpuShare    string // CPU时间片
	CpuSet      string //CPU核心数
}

// 每个subsystem可以实现下面四个接口
type Subsystem interface {
	Name() string                               // 返回subsystem的名字
	Set(path string, res *ResourceConfig) error // 设置某个cgroup在subsystem的资源限制
	Apply(path string, pid int) error           // 将进程添加到某个cgroup中（将进程ID写入cgroup的cgroup.procs文件）
	Remove(path string) error                   //移除某个cgroup
}

// 通过不同的subsystem初始化实例 创建资源限制处理链数组
var (
	SubsystemsIns = []Subsystem{
		&CpusetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSystem{},
	}
)
