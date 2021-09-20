# Docker

## 前戏
```sh
安装
curl -fsSL https://get.docker.com | bash -s docker --mirror aliyun
查看
docker -v
docker info
启动
systemctl start docker
停止docker
systemctl stop docker
重启
systemctl restart docker
重载
systemctl reload
状态
systemctl status
开机自启
systemctl enable docker

换源
vi /etc/docker/daemon.json
#添加如下网易镜像源
{
"registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"]
}

为了避免每次命令都输入sudo，可以设置用户权限，注意执行后须注销重新登录
sudo usermod -a -G docker $USER

查看镜像
docker images
删除镜像
docker rmi IMAGEID

拉取镜像
$ docker pull shadowsocks/shadowsocks-libev
运行
$ docker run -e PASSWORD=ZheShiMima888 -p 38388:8388 -p 38388:8388/udp -d --restart always shadowsocks/shadowsocks-libev

查看运行中的容器
docker ps
查看所有运行中的容器
docker ps -a


```