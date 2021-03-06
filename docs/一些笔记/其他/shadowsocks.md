# Windwos Server

1. 下载一个已经编译好的最新版本，[下载 shadowsocks-libqss-v2.0.2-win64.7z](https://github.com/shadowsocks/libQtShadowsocks/releases)

2. 解压后得到shadowsocks-libqss.exe,在同目录下创建两个文件

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
3. Tips: config.json中,server_port端口,password密码,method加密方式,客户端对应填入即可.

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
    [下载地址1.github](https://github.com/Jigsaw-Code/outline-server/releases) 
    [下载地址2.官网](https://s3.amazonaws.com/outline-vpn/index.html#/zh-CN/home)

2. 管理器获取安装命令执行
    ```
    yum -y install wget
    sudo bash -c "$(wget -qO- https://raw.githubusercontent.com/Jigsaw-Code/outline-server/master/src/server_manager/install_scripts/install_server.sh)"
    ```
3. 安装成功后,将安装脚本的输出结果粘贴到管理器界面。

4. 管理器获取ss://协议链接,给客户端使用

4. 可能需要用到的sh命令:
    ```
    service docker restart
    firewall-cmd --zone=public --add-port=1024-65535/tcp --permanent
    firewall-cmd --reload
    ```

# [客户端](https://shadowsocks.org/en/download/clients.html)

1. [Windows客户端 (Shadowsocks-4.4.0.185.zip)](https://github.com/shadowsocks/shadowsocks-windows/releases)

2. [Mac客户端 (ShadowsocksX-NG.1.9.4.zip)](https://github.com/shadowsocks/ShadowsocksX-NG/releases)

3. [Android客户端 (shadowsocks--universal-5.2.3.apk | shadowsocks-tv--universal-5.2.3.apk)](https://github.com/shadowsocks/shadowsocks-android/releases)

4. iOS客户端
    > 国内ID: AppStore搜索下载Outline

    > 外服ID: Potatso或ShadowSocks

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