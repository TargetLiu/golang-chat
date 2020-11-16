# Golang简单聊天室示例

## 更新

2020-11-17 稍微更新了一下目录结构和启动参数，加入了一个简单的聊天协议

## 介绍

一个简单的基于GO语言的聊天室示例，参考了网上许多类似代码。代码有简单注释。

单执行文件，通过命令的方式选择启动服务端监听还是客户端连接。

功能：客户端昵称、服务端消息、服务端踢人命令。

## 使用方法

- 克隆代码

```
git clone https://github.com/TargetLiu/golang-chat
```

- 编译安装

```
make build
```

or 

```
make install
```

- 运行

```
golang-chat --help

Usage of golang-chat:
  -host string
        host (default "127.0.0.1")
  -port int
        port (default 6666)
  -type string
        start type [server|client] (default "server")
```

- 服务端命令

```
//发送消息直接输入即可
//踢人
kick|[客户端昵称]
```