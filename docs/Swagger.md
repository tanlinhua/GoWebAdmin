# [Swaggo](https://github.com/swaggo)

### [帮助文档](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)

### 安装swag cli tool:
    go get -u github.com/swaggo/swag/cmd/swag
    
### 测试是否安装成功
    swag -version

### 使用
```
import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "swagger.demo/docs" // 执行docs.go的init函数
)

// 初始化swagger
func initSwagger(e *gin.Engine) {
	disablingKey := "GO_API_SWAGGER_DISABLE"
	if config.AppMode != "debug" {
		os.Setenv(disablingKey, "true") // 禁用swagger
	}
	e.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, disablingKey))
}
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

### [注解](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)
```
// @Tags 标签列表
// @Summary 简短摘要
// @Description 描述信息
// @accept API可以使用的MIME类型的列表 #mime类型
// @Param 参数名 #参数类型 #数据类型(struct忽略关键字swaggerignore:"true") 是否必填 "描述"
// @Produce API可以生成的MIME类型的列表 #mime类型
// @Response 以空格分隔的成功响应。return code,{param type},data type,comment
// @Router /api/user/login [post]
```

### 注意事项:
1. 假如func方法头标注的swagger注释不正确，在执行swag init会报错，自行根据报错信息去修改；
2. 访问swagger控制台报错404 page not found，是因为没有添加swagger的路由：
    router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler, url))；
3. 访问swagger控制台报错Failed to load spec，是因为没有import引入执行swag init生成的swagger的docs文件夹；
