# [VIM](https://github.com/wsdjeg/vim-galore-zh_cn)

## Vim 从入门到精通
https://github.com/wsdjeg/vim-galore-zh_cn

## 安装
centos
```shell
yum -y install vim

# 卸载老的vim
yum remove vim-* -y

# 下载第三方yum源
wget -P /etc/yum.repos.d/  https://copr.fedorainfracloud.org/coprs/lbiaggi/vim80-ligatures/repo/epel-7/lbiaggi-vim80-ligatures-epel-7.repo

# install vim
yum  install vim-enhanced sudo -y

# 验证vim版本
rpm -qa |grep vim
```

ubuntu
```shell
sudo apt-get remove vim-common # 如果报错卸载不兼容
sudo apt-get install vim
```

## 安装插件
```
1. 安装插件管理器,下载plug.vim到~/.vim/autoload目录. https://github.com/junegunn/vim-plug,

curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

2. vim ~/.vimrc

call plug#begin('~/.vim/plugged')
Plug 'scrooloose/nerdtree'
Plug 'fatih/vim-go'
call plug#end()

3. :source ~/.vimrc
4. :PlugInstall
```

## 插件

https://zhuanlan.zhihu.com/p/139847548

https://github.com/rafi/vim-config

https://github.com/PegasusWang/linux_config

https://github.com/tao12345666333/vim

## 搜索并优化

> golang vim

> php vim

> vue vim

## 快捷键

https://www.jianshu.com/p/868e63940e11

https://www.cnblogs.com/sinsenliu/p/9353939.html