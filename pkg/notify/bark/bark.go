package bark

import (
	"github.com/tanlinhua/go-web-admin/pkg/utils"
)

var token string = "yihfb28z2mT2pT7d57S35C"

func Send(title, content string) (bool, string) {
	url := "https://api.day.app/" + token + "/" + title + "/" + content
	ok, c := utils.HttpGet(url, nil)
	return ok, c
}

// https://github.com/Finb
// 1.iPhone安装Bark,https://apps.apple.com/cn/app/bark-customed-notifications/id1403753865
// 2.打开app获取token: https://api.day.app/token
// 3.Demo: https://api.day.app/yihfb28z2mT2pT7d57S35C/Title/Content
