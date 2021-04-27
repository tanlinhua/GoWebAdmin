package validator

import (
	"reflect"
	"regexp"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/tanlinhua/go-web-admin/pkg/trace"
)

// 验证器
func Validate(data interface{}) (string, int) {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		trace.Error("zhTrans err:" + err.Error())
	}
	//自定义校验字段及翻译方法
	validate.RegisterValidation("phone", checkPhone)
	validate.RegisterTranslation("phone", trans, registerTranslator("phone", "{0}格式不符合要求"), customTranslate)
	validate.RegisterValidation("state", checkState)
	validate.RegisterTranslation("state", trans, registerTranslator("state", "{0}不符合要求"), customTranslate)

	err = validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err.Translate(trans), 0
		}
	}
	return "SUCCSE", 1
}

// 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans unTrans.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// 自定义字段的翻译方法
func customTranslate(trans unTrans.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}

func checkPhone(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())
	return ok
}

func checkState(fl validator.FieldLevel) bool {
	state := fl.Field().Int()
	if state != 0 && state != 1 {
		return false
	}
	return true
}
