package lang

import (
	"fmt"

	"github.com/tanlinhua/go-web-admin/app/config"
)

// 根据字符获取多语言的具体内容
func Get(msg string) string {
	lang := config.LangDefault

	if _, ok := content[msg]; ok {
		arr := content[msg].(map[string]interface{})
		if _, ok := arr[lang]; ok {
			return fmt.Sprintf("%v", arr[lang])
		}
	}
	return msg
}

var (
	zh_cn = "zh-cn" //中文 https://www.cnblogs.com/striveLD/p/9953748.html
	en_ww = "en-ww" //英文
)

var content = map[string]interface{}{
	"成功": map[string]interface{}{
		zh_cn: "成功",
		en_ww: "success",
	},
	"失败": map[string]interface{}{
		zh_cn: "失败",
		en_ww: "fail",
	},
}
