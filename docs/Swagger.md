# [Swaggo](https://github.com/swaggo)

### [帮助文档](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

### 安装swag cli tool:
    go get -u github.com/swaggo/swag/cmd/swag
### 测试是否安装成功
    swag -version
### 导入
```
import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "swagger.demo/docs" // 执行docs.go的init函数
)
```
### 初始化
```
swag init (注解有做修改就需要执行一次)

生成

./docs
├── docs.go
├── swagger.json
└── swagger.yaml
```

### Swagger的注解:
```
// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags 用户信息   //swagger API分类标签, 同一个tag为一组
// @accept json  //浏览器可处理数据类型，浏览器默认发 Accept: */*
// @Produce  json  //设置返回数据的类型和编码
// @Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
// @Param name query string false "name"
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 400 {object} Res {"code":200,"data":null,"msg":""}
// @Router /test/{id} [get]    //路由信息，一定要写上
```

### 注意事项:
1. 假如func方法头标注的swagger注释不正确，在执行swag init会报错，自行根据报错信息去修改；
2. 访问swagger控制台报错404 page not found，是因为没有添加swagger的路由：
    router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler, url))；
3. 访问swagger控制台报错Failed to load spec，是因为没有import引入执行swag init生成的swagger的docs文件夹；
