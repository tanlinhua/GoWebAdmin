package captcha

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 设置自带的store
var store = base64Captcha.DefaultMemStore

// 生成验证码
func CaptchaMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	var captDriver base64Captcha.DriverMath
	// dight 数字验证码 / audio 语音验证码 / string 字符验证码 / math 数学验证码(加减乘除) / chinese中文验证码-有bug

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverMath{
		Height:          47,
		Width:           145,
		NoiseCount:      0,
		ShowLineOptions: 1 | 3,
		// Length:          4,
		// Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	captDriver = captchaConfig
	driver = captDriver.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

// 验证captcha是否正确
func CaptchaVerify(id string, capt string) bool {
	return store.Verify(id, capt, true) //无论对错清理该验证码
}
