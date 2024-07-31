# DockerGo

使用 go 手写 docker

## Introduce

主要介绍从 0 实现 docker 需要掌握的必备知识，帮助阅读代码中的注释。代码参考书籍《自己动手写 docker》，开发基于 Ubuntu。

### Linux proc 文件系统

## Steps

### 实现 Run 命令

docker run -ti [cmmand]
(需在linux系统运行，并已执行mount -t proc proc /proc)

### 增加容器资源限制
