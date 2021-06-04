package csv

/*

//导出csv

import "github.com/gocarina/gocsv"

	//创建文件流
    f, err := os.Create("dictList.csv")
    //关闭流
    defer f.Close()
    //写入UTF-8 格式
    f.WriteString("\xEF\xBB\xBF")
    var newContent [][]string
    //添加数据
    newContent = append(newContent, []string{"1", "2", "3", "4", "5", "6"})
    //保存文件流
    err = gocsv.MarshalFile(newContent, f)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    //传输文件流（使用gin或者http的时候向前端发送的流文件）
    c.File("dictList.csv")
    //删除文件
    _ = os.Remove("dictList.csv")

*/

/*

//导入csv

	//获取文件头
    file, err := c.FormFile("file")
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        return
    }
    //获取文件名
    fileName := file.Filename
    //获取文件后缀名
    fileSuffix := path.Ext(fileName)
    //后缀名判断
    if fileSuffix != ".csv" {
        response.FailWithMessage("文件后缀格式有误", c)
        return
    }
    //SaveUploadedFile(文件头，保存路径)
    if err := c.SaveUploadedFile(file, fileName); err != nil {
        response.FailWithMessage("保存失败", c)
        os.Remove(fileName)
        return
    }
    //打开流
    clientsFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
    if err != nil {
        response.FailWithMessage(err.Error(), c)
        os.Remove(fileName)
        return
    }
    //关闭流
    defer clientsFile.Close()
    //解析csv文件到结构体，clients为自己定义的结构体
    if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil {
        response.FailWithMessage("文件内容格式有误", c)
        os.Remove(fileName)
        return
    }
	for _, client := range clients {
		.....
		//遍历clients，每个结构体参数用client来获取，并按照需求进行处理
	}
	//删除上传文件
    os.Remove(fileName)
    insert.Commit()
    response.OkWithMessage("导入成功", c)
*/

// 转自链接：https://learnku.com/articles/49971
