# DockerGo

使用 go 手写 docker

## Introduce

    主要介绍从 0 实现 docker 需要掌握的必备知识，帮助阅读代码中的注释。代码参考书籍《自己动手写 docker》，开发基于 Ubuntu。

### Linux proc 文件系统

## Steps

### 1.实现 Run 命令

```shell
docker run -ti [cmmand] (需在 linux 系统运行，并已执行 mount -t proc proc /proc)
```

### 2.增加容器资源限制

    包括对容器的内存限制、CPU 时间片权重、CPU 核心数进行限制，主要实现了以下命令:

```shell
docker run -ti -m 100m -cpuset 1 -cpushare 512 /bin/sh
```

    将每个容器作为一个 subsystem，包含四个接口，实现 subsystem 的 Cgroup 资源限制，

在/sys/fs/cgroup/memory 中创建文件夹对应创建的 cgroup，就可以做内存限制。（实际上是在创建一个新的 cgroup，这个 cgroup 可以用来限制和监控一组进程的内存使用），创建 cgroup 后，就可以将进程添加到这个 cgroup 中，通常通过写入 cgroup.procs 文件来实现。一旦进程被添加到 cgroup，它将受到该 cgroup 设置的内存限制。该部分是通过将容器进程移动到 subsystem 创建的 cgroup 中进行资源限制。
