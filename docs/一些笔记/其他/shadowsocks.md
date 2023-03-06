# Windwos Server

1. 下载一个已经编译好的最新版本，[下载 shadowsocks-libqss-v2.0.2-win64.7z](https://github.com/shadowsocks/libQtShadowsocks/releases)

2. 解压后得到 shadowsocks-libqss.exe,在同目录下创建两个文件

   shadowsocks-server.bat

   ```
   ::This batch will run shadowsocks-libqss in server mode using the config.json file in current folder as the configuration

   @echo off
   ::this script is updated for version 1.7.0
   shadowsocks-libqss.exe -c config.json -S
   ```

   config.json

   ```
   {
       "server":"0.0.0.0",
       "server_port":8388,
       "local_address":"127.0.0.1",
       "local_port":1080,
       "password":"barfoo!",
       "timeout":600,
       "method":"rc4-md5",
       "http_proxy": false,
       "auth": false
   }
   ```

3. Tips: config.json 中,server_port 端口,password 密码,method 加密方式,客户端对应填入即可.

# Linux

## Docker

```
安装docker
$ curl -fsSL https://get.docker.com | bash -s docker --mirror aliyun
启动docker
$ sudo systemctl start docker
开机自启
$ sudo systemctl enable docker
拉取镜像
$ docker pull shadowsocks/shadowsocks-libev
运行
$ docker run -e PASSWORD=ZheShiMima888 -e METHOD=chacha20-ietf-poly1305 -p 8865:8388 -p 8865:8388/udp -d --restart always shadowsocks/shadowsocks-libev

停止所有容器
$ docker stop $(docker ps -a -q)
删除所有容器
$ docker rm $(docker ps -a -q)
```

## Outline-server

1. 安装服务端管理器
   [下载地址 1.github](https://github.com/Jigsaw-Code/outline-server/releases)
   [下载地址 2.官网](https://s3.amazonaws.com/outline-vpn/index.html#/zh-CN/home)

2. 管理器获取安装命令执行
   ```
   yum -y install wget
   sudo bash -c "$(wget -qO- https://raw.githubusercontent.com/Jigsaw-Code/outline-server/master/src/server_manager/install_scripts/install_server.sh)"
   ```
3. 安装成功后,将安装脚本的输出结果粘贴到管理器界面。

4. 管理器获取 ss://协议链接,给客户端使用

5. 可能需要用到的 sh 命令:
   ```
   service docker restart
   firewall-cmd --zone=public --add-port=1024-65535/tcp --permanent
   firewall-cmd --reload
   ```

# [客户端](https://shadowsocks.org/en/download/clients.html)

1. [Windows 客户端 (Shadowsocks-4.4.0.185.zip)](https://github.com/shadowsocks/shadowsocks-windows/releases)

2. [Mac 客户端 (ShadowsocksX-NG.1.9.4.zip)](https://github.com/shadowsocks/ShadowsocksX-NG/releases)

3. [Android 客户端 (shadowsocks--universal-5.2.3.apk | shadowsocks-tv--universal-5.2.3.apk)](https://github.com/shadowsocks/shadowsocks-android/releases)

4. iOS 客户端

   > 国内 ID: AppStore 搜索下载 Outline

   > 外服 ID: Potatso 或 ShadowSocks

# Github

```
http://tool.chinaz.com/dns/

然后在输入框输入github.com
然后你就会看到检测出来的好多IP
找到一个TTL值最小的IP地址

添加到C:\Windows\System32\drivers\etc\hosts文件里面

52.74.223.119 github.com

一般是立刻生效。
没有的话，手动在 CMD 敲入：ipconfig /flushdns
```

# 问题

> OpenSSL SSL_connect: Connection was aborted in connection to github.com:443

```
设置代理
git config  --global  http.proxy 'http://localhost:1080'
取消代理
git config  --global  --unset http.proxy
```

# Telegram.MTProto

[仓库](https://github.com/TelegramMessenger/MTProxy)

[教程 1](https://www.it610.com/article/1305300310650032128.htm)

[一键安装脚本](https://github.com/ToyoDAdoubi/doubi)

```
wget -N --no-check-certificate https://raw.githubusercontent.com/ToyoDAdoubi/doubi/master/mtproxy.sh && chmod +x mtproxy.sh && bash mtproxy.sh
```

更多自行搜索： Telegram MTProto
