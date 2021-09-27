package csv

/*

//导出csv

func Export(c *gin.Context) {
	task_id := c.PostForm("tid")
	export_type := c.PostForm("e_type")
	fileName := "任务ID(" + task_id + ")_导出类型(" + export_type + ").csv"

	tid, _ := strconv.ParseInt(task_id, 10, 64)
	e_type, _ := strconv.Atoi(export_type)

	file, err := os.Create(fileName) //创建文件流
	if err != nil {
		response.New(c).Error(-1, err.Error())
	}
	file.WriteString("\xEF\xBB\xBF") //写入UTF-8 格式

	data := model.GetByTidAndStatus(tid, e_type)
	err = gocsv.MarshalFile(data, file)
	if err != nil {
		response.New(c).Error(-1, err.Error())
		return
	}

	c.Header("Content-type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename="+fileName)
	c.Header("Cache-Control", "must-revalidate,post-check=0,pre-check=0")
	c.Header("Expires", "0")
	c.Header("Pragma", "public")

	c.File(fileName)          //传输文件流
	file.Close()              //关闭流
	err = os.Remove(fileName) //删除文件
	if err != nil {
		trace.Error("SmsTaskExport,删除文件失败:" + err.Error())
	}
}

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

/*
// 通用csv数据
type CsvData struct {
	Data string `csv:"data"`
}

// 获取csv数据
func GetCsvData(c *gin.Context, name string) ([]*CsvData, error) {
	file, err := c.FormFile(name) //获取文件
	if err != nil {
		return nil, err
	}

	fileName := file.Filename        //获取文件名
	fileSuffix := path.Ext(fileName) //获取文件后缀名
	if fileSuffix != ".csv" {
		return nil, errors.New("文件格式错误")
	}

	if err := c.SaveUploadedFile(file, fileName); err != nil {
		os.Remove(fileName)
		return nil, err
	}
	//打开流
	clientsFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		os.Remove(fileName)
		return nil, err
	}

	data := []*CsvData{}
	if err := gocsv.UnmarshalFile(clientsFile, &data); err != nil {
		os.Remove(fileName)
		return nil, err
	}

	clientsFile.Close()
	if removeErr := os.Remove(fileName); removeErr != nil {
		return nil, removeErr
	}
	return data, nil
}
*/
