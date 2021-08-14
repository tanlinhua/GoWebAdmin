# VIM

## 安装
centos: yum -y install vim
```
# 卸载老的vim
yum remove vim-* -y

# 下载第三方yum源
wget -P /etc/yum.repos.d/  https://copr.fedorainfracloud.org/coprs/lbiaggi/vim80-ligatures/repo/epel-7/lbiaggi-vim80-ligatures-epel-7.repo

# install vim
yum  install vim-enhanced sudo -y

# 验证vim版本
rpm -qa |grep vim
```

ubuntu: 

## 安装插件
```
1. 安装插件管理器,https://github.com/junegunn/vim-plug
下载plug.vim到~/.vim/autoload目录
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim

2. vim ~/.vimrc
set number
set encoding=utf-8
filetype plugin indent on
syntax on

call plug#begin('~/.vim/plugged')
Plug 'fatih/vim-go'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
call plug#end()

3. :source ~/.vimrc
4. :PlugInstall

https://zhuanlan.zhihu.com/p/139847548
vim-startify
nerdtree
vim-interestingwords
vim-fugitive

```

## 快捷键

https://www.jianshu.com/p/868e63940e11

### 一、移动光标

| 按键    | 功能           |
| ------- | -------------- |
| h,j,k,l | 上，下，左，右 |
| ctrl-e  | 移动页面       |
| ctrl-f  | 上翻一页       |
| ctrl-b  | 下翻一页       |
| ctrl-u  | 上翻半页       |
| ctrl-d  | 下翻半页       |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |
|         |                |

