# go web项目通用模板


## 一些笔记 

+ 模板引擎循环判断等方式输出数据
    + https://blog.csdn.net/sryan/article/details/52353937
    + https://www.jianshu.com/p/17873107c0fd
    + https://studygolang.com/articles/28508?fr=sidebar

+ go web项目如何优雅的更新
    + https://blog.csdn.net/qihoo_tech/article/details/104386331

+ go cache
    + https://studygolang.com/articles/25283
    + https://github.com/patrickmn/go-cache

+ go交叉编译
    + SET GOOS=linux
    + SET GOARCH=amd64
    + go build main.go

+ 程序目录执行后台运行命令
    + nohup ./main >> /www/wwwroot/main.go.nohup.output.`date +%Y-%m-%d`.log 2>&1 &
    + ps -ef|grep main
    + ->kill -9 pid

+ 并发相关
    + sync.WaitGroup channel
    + https://blog.csdn.net/m0_37422289/article/details/105328796
    + https://my.oschina.net/xlplbo/blog/682884
    + https://learnku.com/go/t/23456/using-the-go-language-to-handle-1-million-requests-per-minute

+ 解决1W并发http.Get报错(socket: too many open files)
    + 永久修改open files 方法
    + vim /etc/security/limits.conf  
    + 在最后加入  
    + * soft nofile 65535
    + * hard nofile 65535
    + soft/hard前面的 * 表示所有用户


## TODO
+ 1.admin/role/per -> crud(https://www.17sucai.com/preview/1528155/2019-03-31/ladmin/index.html)
+ 2.admin/main的菜单根据权限列表来展示
+ xss/sql注入/csrf (https://github.com/xinliangnote/PHP/blob/master/04-PHP%20WEB%20%E5%AE%89%E5%85%A8%E9%98%B2%E5%BE%A1.md)
+ bug -> admin/main 刷新丢失已经打开的tab
+ golang googleauthenticator
    + https://blog.csdn.net/weixin_33958585/article/details/92494417
    + 参考:composer require "phpgangsta/googleauthenticator:dev-master"
+ api模块增加(go gin swagger)