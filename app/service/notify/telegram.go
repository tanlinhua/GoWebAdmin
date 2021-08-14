package notify

import (
	"github.com/tanlinhua/go-web-admin/app/config"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
	"github.com/tanlinhua/go-web-admin/pkg/utils"
	"github.com/tidwall/gjson"
)

// 发送telegram消息
func SendTelegarmMessage(msg string) {
	url := config.TgHost + "bot" + config.TgToken + "/sendMessage"

	var params = make(map[string]string)
	params["chat_id"] = config.TgChatId
	params["text"] = msg

	ok, result := utils.HttpGet(url, params)

	tg_ok := gjson.Get(result, "ok")

	// 如果发送失败暂时只记录日志
	if !ok || tg_ok.Value() != true {
		trace.Error(result)
	}
}

// 1.创建机器人
// https://telegram.me/botfather 手机端搜索 BotFather

// 2.获取chat_id
// 将电报BOT添加到组中,发送消息
// 获取您的BOT的更新列表：
// https://api.telegram.org/bot<YourBOTToken>/getUpdates
// 例如：
// https://api.telegram.org/bot123456789:jbd78sadvbdy63d37gda37bd8/getUpdates
// 个人: result.message.chat.id
// 群聊: result.my_chat_member.chat.id

// 3.发送消息
// https://api.telegram.org/bot1913301606:AAFUIyYRDYgLEOvjHtn_H3nxjQGMB-idtu8/sendMessage?chat_id=-582110667&text=helloworld

// 其他功能↓
// TG.SDK https://github.com/go-telegram-bot-api/telegram-bot-api
