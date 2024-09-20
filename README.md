# 通过Golang动手实现一个迷你Docker
---

Docker 是基于 Linux Kernel 的 Namespace 和 Cgroup 等技术实现的。
> 所以 Docker 本身是 Linux 系统的，所以 Windows 系统上需要安装一个虚拟机，然后在这个虚拟机上安装 Docker，这样就可以在 Windows 系统上使用 Docker 了。

## 容器技术原理
### chroot
chroot 是 Linux 系统中的一个命令，用于改变当前进程的根目录，即改变进程的视图，使其只能访问当前目录及其子目录。

chroot 的作用是使进程只能访问当前目录及其子目录，从而限制进程的访问能力。
chroot 一般用于 root 用户，用于限制普通用户的访问能力，防止普通用户越权访问系统文件。

### 什么是 Namespace
Namespace 是 Linux 内核提供的一种机制，用于隔离进程的运行环境，包括进程可见的视图、进程的系统资源限制等。

### 什么是 Cgroup
Cgroup 是 Linux 内核提供的一种机制，用于限制、记录、隔离进程组所使用的物理资源（如 CPU、内存、磁盘 IO、网络等）。

### 什么是 UnionFS
UnionFS 是一种分层、轻量级并且高性能的文件系统，它支持对文件系统的修改作为一次提交来一层层的叠加，同时可以将不同目录挂载到同一个虚拟文件系统下（unite several directories into a single virtual filesystem）。