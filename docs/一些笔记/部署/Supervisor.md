# Supervisor的使用

## 简介
Supervisord 是用 Python 实现的一款的进程管理工具，supervisord 要求管理的程序是非 daemon 程序，supervisord 会帮你把它转成 daemon 程序，因此如果用 supervisord 来管理进程，进程需要以非daemon的方式启动。

## 安装
```
# yum install 的方式
yum install -y supervisor

# easy_install的方式
yum install -y python-setuptools
easy_install supervisor
echo_supervisord_conf >/etc/supervisord.conf

# 宝塔安装
```
查看是否安装成功

$ echo_supervisord_conf

## Supervisor配置
如果使用yum install -y supervisor的命令安装，会生成默认配置/etc/supervisord.conf和目录/etc/supervisord.d

如果没有则自行创建。(echo_supervisord_conf > /etc/supervisord.conf)

在/etc/supervisord.d的目录下创建conf和log两个目录，conf用于存放管理进程的配置，log用于存放管理进程的日志。
```
cd /etc/supervisord.d
mkdir conf log
```

修改/etc/supervisord.conf的[include]部分，即载入/etc/supervisord.d/conf目录下的所有配置。
```
vi /etc/supervisord.conf
...
[include]
files = supervisord.d/conf/*.conf
...
```

## 应用配置
说明
```
#项目名
[program:blog]
#脚本目录
directory=/opt/bin
#脚本执行命令
command=/usr/bin/python /opt/bin/test.py
#supervisor启动的时候是否随着同时启动，默认True
autostart=true
#当程序exit的时候，这个program不会自动重启,默认unexpected，设置子进程挂掉后自动重启的情况，有三个选项，false,unexpected和true。如果为false的时候，无论什么情况下，都不会被重新启动，如果为unexpected，只有当进程的退出码不在下面的exitcodes里面定义的
autorestart=false
#这个选项是子进程启动多少秒之后，此时状态如果是running，则我们认为启动成功了。默认值为1
startsecs=1
#脚本运行的用户身份 
user = test
#日志输出 
stderr_logfile=/tmp/blog_stderr.log 
stdout_logfile=/tmp/blog_stdout.log 
#把stderr重定向到stdout，默认 false
redirect_stderr = true
#stdout日志文件大小，默认 50MB
stdout_logfile_maxbytes = 20MB
#stdout日志文件备份数
stdout_logfile_backups = 20
```
创建管理应用的配置, 例如:创建/etc/supervisord.d/conf/GoWebAdmin.conf配置。(可以创建多个应用配置)
```
[program:GoWebAdmin]
directory=/www/wwwroot/GoWebAdmin
command=/www/wwwroot/GoWebAdmin/main
autostart=true
autorestart=false
stderr_logfile=/etc/supervisord.d/log/GoWebAdmin_stderr.log
stdout_logfile=/etc/supervisord.d/log/GoWebAdmin.log
```

## Surpervisor的启动
```
# supervisord二进制启动
supervisord -c /etc/supervisord.conf
# 检查进程
ps aux | grep supervisord
```

## Surpervisor开机自启
在目录/usr/lib/systemd/system/ 新建文件supervisord.service,并添加配置内容
```
[Unit]
Description=Process Monitoring and Control Daemon
After=rc-local.service nss-user-lookup.target

[Service]
Type=forking
ExecStart=/usr/bin/supervisord -c /etc/supervisord.conf
ExecStop=/usr/bin/supervisord shutdown
ExecReload=/usr/bin/supervisord reload
killMode=process
Restart=on-failure
RestartSec=42s

[Install]
WantedBy=multi-user.target
```
```
systemctl start supervisord.service     //启动supervisor并加载默认配置文件
systemctl enable supervisord.service    //将supervisor加入开机启动项
systemctl is-enabled supervisord        //验证一下是否为开机启动
```
## systemctl常用命令
以firewalld.service为例
```
启动一个服务：systemctl start firewalld.service
关闭一个服务：systemctl stop firewalld.service
重启一个服务：systemctl restart firewalld.service
显示一个服务的状态：systemctl status firewalld.service

在开机时启用一个服务：systemctl enable firewalld.service
在开机时禁用一个服务：systemctl disable firewalld.service
查看服务是否开机启动：systemctl is-enabled firewalld.service

查看开机启动的服务列表：systemctl list-unit-files|grep enabled
查看启动失败的服务列表：systemctl --failed
```

## supervisorctl常用命令
- supervisorctl stop GoWebAdmin         停止某一个进程
- supervisorctl start GoWebAdmin        启动某个进程。
- supervisorctl restart GoWebAdmin      重启某个进程。
- supervisorctl status                  查看进程状态。
- supervisorctl stop groupworker        重启所有属于名为 groupworker 这个分组的进程(start,restart 同理)。
- supervisorctl stop all                停止全部进程，注：start、restart、stop 都不会载入最新的配置文件。
- supervisorctl reload                  载入最新的配置文件，停止原有进程并按新的配置启动、管理所有进程。
- supervisorctl update                  根据最新的配置文件，启动新配置或有改动的进程，配置没有改动的进程不会受影响而重启。

例如
```
$ supervisorctl status
GoWebAdmin                       RUNNING   pid 2459, uptime 0:03:50
```

## 常见问题
```
unix:///var/run/supervisor.sock no such file
问题描述：安装好supervisor没有开启服务直接使用supervisorctl报的错
解决办法：supervisord -c /etc/supervisord.conf
```
```
command中指定的进程已经起来，但supervisor还不断重启
问题描述：command中启动方式为后台启动，导致识别不到pid，然后不断重启，这里使用的是elasticsearch，command指定的是$path/bin/elasticsearch -d
解决办法：supervisor无法检测后台启动进程的pid，而supervisor本身就是后台启动守护进程，因此不用担心这个
```
```
启动了多个supervisord服务，导致无法正常关闭服务
问题描述：在运行supervisord -c /etc/supervisord.conf之前，直接运行过supervisord -c /etc/supervisord.d/xx.conf导致有些进程被多个superviord管理，无法正常关闭进程。
解决办法：使用ps -fe | grep supervisord查看所有启动过的supervisord服务，kill相关的进程。
```

## 遇到的坑

### socket: too many open files

```
vim /etc/security/limits.conf
在最后加入
* soft nofile 65535
* hard nofile 65535

* soft nproc 65535
* hard nproc 65535

tips↓
* 表示所有用户
soft/hard 软硬限制
nproc 最大线程数 / nofile 最大文件数
```

```
supervisor 控制的程序
CentOS 上使用系统自带的 supervisor，使用 systemd 启动 supervisord 的服务。被 supervisor 管理的程序，
继承的是 systemd 对应的限制，如果需要修改的话，就需要在启动.service 文件里面修改对应的限制

方法1:
vi /usr/lib/systemd/system/supervisord.service

[Service]
Type=forking
LimitNOFILE=102400
LimitNPROC=102400

方法2:
vi /etc/supervisord.conf
minfds=102400
```

```
一些命令
查看当前程序nofile限制: cat /proc/39977/limits
临时修改: prlimit --nofile=65536:65536 --pid 39977
查看当前系统打开的文件数量,代码如下:
lsof | wc -l
watch "lsof | wc -l"

查看某一进程的打开文件数量,代码如下:
lsof -p 3490 | wc -l
```
[或者通过此方案限制并发数](../../pkg/gpool/docs/demo.md)