# Golang简单聊天室示例

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
go build
go install
```

- 运行

```
//服务端
golang-chat server [:端口号]
//客户端
golang-chat client [服务端IP地址:端口号] 
```

- 服务端命令

```
//发送消息直接输入即可
//踢人
kick|[客户端昵称]
```