# VIM

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

" 示例
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

https://github.com/theniceboy/nvim

https://github.com/tao12345666333/vim

## 搜索并优化

> golang vim

> php vim

> vue vim

## 快捷键

https://www.jianshu.com/p/868e63940e11

https://www.cnblogs.com/sinsenliu/p/9353939.html

## .vimrc
```
let mapleader=','
inoremap jj <Esc>

set number
set encoding=utf-8
set foldenable

syntax on

call plug#begin('~/.vim/plugged')

Plug 'w0ng/vim-hybrid'
Plug 'scrooloose/nerdtree'
Plug 'fatih/vim-go'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'voldikss/vim-floaterm'
Plug 'mhinz/vim-startify'

call plug#end()


let g:go_version_warning = 0

nmap <leader>v :NERDTreeFind<cr>
nmap <leader>g :NERDTreeToggle<cr>
let NERDTreeShowHidden=1
let NERDTreeIgnore = [
            \ '\.git$', '\.hg$', '\.svn$', '\.stversions$', '\.pyc$', '\.pyo$', '\.svn$', '\.swp$',
            \ '\.DS_Store$', '\.sass-cache$', '__pycache__$', '\.egg-info$', '\.ropeproject$',
            \ ]


set background=dark
colorscheme hybrid
```

## PlugInstall失败
因为中国地区访问像 github 这种国外网站很不稳定, 所以在一般都是采用镜像网站的方式间接访问. 而 vim-plug 下载时都是用的实际网站, 我们可以修改 plug.vim 来将实际网站变为镜像网站, 提高下载成功率.

在 plug.vim 中搜索 github, 修改两条语句，即可成功下载GitHub资源。 
```
1 将该行
let fmt = get(g:, 'plug_url_format', 'https://git::@github.com/%s.git')
 改为
let fmt = get(g:, 'plug_url_format', 'https://git::@hub.fastgit.org/%s.git')
2 将改行
\ '^https://git::@github\.com', 'https://github.com', '')
改为
\ '^https://git::@hub.fastgit\.org', 'https://hub.fastgit.org', '')
```
上述两个步骤的目是让镜像网站代替实际网站, 这样能有效提高下载成功率。