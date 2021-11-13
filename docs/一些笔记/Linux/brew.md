# brew
```
本地软件库列表：brew ls
查找软件：brew search google（其中google替换为要查找的关键字）
查看brew版本：brew -v  更新brew版本：brew update
安装cask软件：brew install --cask firefox 把firefox换成你要安装的
启动服务： brew services start redis@4.0
查看服务： brew services list

查询可更新的包： brew outdated
更新所有软件： brew upgrade
清理缓存： brew cleanup
```

## ngnix
```shell
安装： brew install nginx
启动： brew services start nginx
重启： brew services restart nginx
停止： brew services stop nginx
查看： cat /usr/local/etc/nginx/nginx.conf
编辑： vi /usr/local/etc/nginx/nginx.conf
查看安装目录： brew list nginx

localhost:8080

nginx配置tp5.0
/usr/local/etc/nginx/servers/php-web-admin.conf
server {
    listen       1992;
    server_name  localhost;

    access_log  /Users/tanlinhua/Documents/Code/logs/access.log;
    error_log   /Users/tanlinhua/Documents/Code/logs/error.log;

    index index.php index.html index.htm;
    root /Users/tanlinhua/Documents/Code/PhpWebAdmin/public/;
    location / {
    if (!-e $request_filename ) {
        rewrite ^(.*)$ /index.php?s=/$1 last;
        break;
        }
        try_files $uri $uri/ /index.php?$args;
    }
    location ~ \.php$ {
        fastcgi_pass 127.0.0.1:9000;
        fastcgi_index index.php;
        fastcgi_split_path_info       ^(.+\.php)(/?.+)$;
        fastcgi_param PATH_INFO       $fastcgi_path_info;
        fastcgi_param PATH_TRANSLATED $document_root$fastcgi_path_info;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include                       fastcgi_params;
    }
    location ~ /\.ht {
        deny  all;
    }
}

nginx 打开nginx
nginx  重新加载配置|重启|停止|退出
nginx -s reload|reopen|stop|quit
nginx -t 测试配置是否有语法错误
```

## php
```
brew安装php

brew search php  使用此命令搜索可用的PHP版本
brew install php@7.4 使用此命令安装指定版本的php
brew install brew-php-switcher 安装php多版本切换工具
brew-php-switcher 7.4 切换PHP版本到7.4（需要brew安装多个版本）

php -v & php -m

启动php-fpm
brew services start php@7.4
or
/usr/local/opt/php@7.4/sbin/php-fpm --nodaemonize

brew安装PHP扩展

通过brew安装的PHP版本中自带了pecl,可以直接使用
pecl version 查看版本信息
pecl help 可以查看命令帮助
pecl search xdebug  搜索可以安装的扩展信息
pecl install xdebug 安装扩展
pecl install http://pecl.php.net/get/redis-4.2.0.tgz 安装指定版本扩展
默认扩展.so文件会被编译到/usr/local/Cellar/php@7.2/7.2.15/pecl/目录中，此目录实际上是软链接到了/usr/local/lib/php/pecl目录。

查看扩展是否开启的方法：
1. php -m
2. phpinfo()
3. extension_loaded() //直接判断扩展是否加载
4. function_exists() //判断扩展库下的方法是否存在
5. php --ri 扩展名 //查看扩展版本信息
```

## mysql
如果之前没有安装过MySQL 5.7
```
brew install mysql@5.7  // 安装
brew link --force mysql@5.7 // 链接
brew services start mysql@5.7 // 启动服务
echo 'export PATH="/usr/local/opt/mysql@5.7/bin:$PATH"' >> ~/.zshrc // 输出到环境变量
```
如果之前安装了 MySQL 5.7
```
brew uninstall mysql@5.7
rm -rf /usr/local/var/mysql
rm /usr/local/etc/my.cnf

--- 👇和第一种情况一样

brew install mysql@5.7  // 安装
brew link --force mysql@5.7 // 链接
brew services start mysql@5.7 // 启动服务
echo 'export PATH="/usr/local/opt/mysql@5.7/bin:$PATH"' >> ~/.zshrc // 输出到环境变量
```
❗️重要的一步，设置安全的访问
```
mysql_secure_installation
```

mysql 密码： root+root

## redis
```
brew install redis@4.0
```

## nodejs
```shell
brew install nvm

mkdir ~/.nvm

vim ~/.zshrc  在 ~/.zshrc 配置文件后添加如下内容
export NVM_DIR="$HOME/.nvm"
  [ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
  [ -s "/usr/local/opt/nvm/etc/bash_completion.d/nvm" ] && . "/usr/local/opt/nvm/etc/bash_completion.d/nvm"

source ~/.zshrc
echo $NVM_DIR
nvm --help
nvm ls-remote

nvm install 14
nvm uninstall 14

nvm list 是查找本电脑上所有的node版本
nvm install <version> 安装指定版本node
nvm use <version> 切换使用指定的版本node
nvm ls 列出所有版本
nvm current 显示当前版本
nvm alias <name> <version> ## 给不同的版本号添加别名
nvm unalias <name> ## 删除已定义的别名
nvm reinstall-packages <version> ## 在当前版本node环境下，重新全局安装指定版本号的npm包
```