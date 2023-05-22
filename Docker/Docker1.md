### 容器技术原理
#### chroot
`chroot 实在 Unix 和 Linux 系统的一个操作，针对正在运作的软件进程和它的子进程，改变它外显的根目录。
一个运行在这个环境下，经由 chroot 设置根目录的程序，它不能够对这个指定根目录之外的文件进行访问动作，不能读取，也不能修改它的内容。`

#### Namespace
`Linux内核的一项功能，该功能对内核资源进行隔离，使得容器中的进程都可以在单独的命名空间中运行，并且只可以访问当前容器命名空间的资源。`
> pid namespace: 用于隔离进程ID </br>
`int pid = clone(main_function, stack_size, CLONE_NEWPID | SIGCHLD, NULL);` <br>
> net namespace: 隔离网络接口，在虚拟的net namespace内用户可以拥有自己独立的IP、路由、端口等。</br>
> mnt namespace: 文件系统挂载点隔离。 </br>
> ipc namespace: 信号量，消息队列和共享内存的隔离。<br>
> uts namespace: 主机名和域名的隔离。<br>
#### Cgroups
`Cgroups 是一种 Linux 内核功能，可以限制和隔离进程的资源使用情况（CPU、内存、磁盘 I/O、网络等）。在容器的实现中，Cgroups 通常用来限制容器的 CPU 和内存等资源的使用`

#### 联合文件系统 UnionFS
`联合文件系统，又叫 UnionFS，是一种通过创建文件层进程操作的文件系统，因此，联合文件系统非常轻快。Docker 使用联合文件系统为容器提供构建层，使得容器可以实现写时复制以及镜像的分层构建和存储。常用的联合文件系统有 AUFS、Overlay 和 Devicemapper 等`