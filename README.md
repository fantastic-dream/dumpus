# Dumpus

## 概述
Dumpus（Dumbo Octopus）是一个系统信息采集器，可以将采集到CPU、内存、磁盘、操作系统等信息
通过TCP发送到配置文件指定的服务器上。支持Windows、Linux操作系统平台，其他操作系统理论上支持，
但不保证可用性。

## 配置
配置文件`dumpus.conf`位于安装路径的`conf`目录下，配置项如下：

| 配置项         | 值类型        | 默认值 | 说明           |
|   ----        | ----         | ----  | ----           |
| server.ipAddr |   string     |       |   服务器IP地址   |
| server.port   |   int        | 14869 |   服务器端口    |

## 用法
Dumpus可以以守护进程的方式运行，也可以直接运行。命令`dumpus`提供了以下几个可指定的参数：
```shell
Usage: dumpus [-irh]
Options:
  -h    This help.
  -i    Install service.
  -r    Remove service.
```

### 直接运行

在命令行环境下，直接运行`dump`，而不指定参数即可。
```shell
> dumpus
time="2019-02-28T13:04:33+08:00" level=info msg="Service starting"
time="2019-02-28T13:04:33+08:00" level=info msg="Service started."
time="2019-02-28T13:04:33+08:00" level=info msg="Loading configuration."
time="2019-02-28T13:04:33+08:00" level=info msg="Loaded configuration: {\"server\":{\"ipAddr\":\"192.168.0.2\",\"port\":14869}}"
time="2019-02-28T13:04:33+08:00" level=info msg="Starting to send info."
time="2019-02-28T13:04:33+08:00" level=info msg="Started sending info."
```

### 守护进程

以守护进程的方式运行，需要先进行服务安装（需要有足够的权限）：
```shell
> dumpus -i
time="2019-02-28T13:02:36+08:00" level=info msg="Installing service."
time="2019-02-28T13:02:37+08:00" level=info msg="Install service finished."
time="2019-02-28T13:02:37+08:00" level=info msg="Staring service."
time="2019-02-28T13:02:37+08:00" level=info msg="Start service finished."
```

安装后会自动启动服务，运行期间的服务管理可通过对应平台的服务管理机制进行：
Windows平台为服务器管理器，Linux平台为Systemd。

如果要卸载服务，则指定`-r`参数（需要有足够的权限）：
```shell
> dumpus -r
time="2019-02-28T13:03:48+08:00" level=info msg="Stopping service."
time="2019-02-28T13:03:48+08:00" level=info msg="Stop service finished."
time="2019-02-28T13:03:48+08:00" level=info msg="Removing service."
time="2019-02-28T13:03:48+08:00" level=info msg="Remove service finished."
```

## 日志

Dumpus的运行日志记录文件`dumpus.log`保存在安装路径下的`log`目录中。