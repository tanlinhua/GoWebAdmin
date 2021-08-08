# Ubuntu一些笔记

## 安装VSCode

1. 以 sudo 用户身份运行下面的命令，更新软件包索引，并且安装依赖软件：
```
sudo apt update
sudo apt install software-properties-common apt-transport-https wget
```
2. 使用 wget 命令插入 Microsoft GPG key ：
```
wget -q https://packages.microsoft.com/keys/microsoft.asc -O- | sudo apt-key add -
启用 Visual Studio Code 源仓库，输入：
sudo add-apt-repository "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main"
```
3. 一旦 apt 软件源被启用，安装 Visual Studio Code 软件包：
```
sudo apt install code
```
4. 当一个新版本被发布时，你可以通过你的桌面标准软件工具，或者在你的终端运行命令，来升级 Visual Studio Code 软件包：
```
sudo apt update
sudo apt upgrade
```

## 安装Golang
```
ubuntu
先安装最新的golang源，否则安装到的是老版本
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go

dpkk -L golang-go(查看安装目录,/usr/lib/go)

centos
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo yum -y install golang
```

## 前端
```
```

## Composer
```
wget -O composer-setup.php https://getcomposer.org/installer
sudo php composer-setup.php --install-dir=/usr/local/bin --filename=composer
composer -v
composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
更新
sudo composer self-update
```

## 一些配置
```
桌面配置工具
sudo apt  install gnome-tweaks
终端输入 gnome-tweaks 命令打开
```
```
点击任务栏最小化
gsettings set org.gnome.shell.extensions.dash-to-dock click-action 'minimize'
```