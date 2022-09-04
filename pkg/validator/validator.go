package validator

import (
	"reflect"
	"regexp"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

// 验证器
func Validate(data interface{}) error {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}

	//自定义校验字段及翻译方法
	validate.RegisterValidation("phone", checkPhone)
	validate.RegisterTranslation("phone", trans, registerTranslator("phone", "{0}格式不符合要求"), customTranslate)

	validate.RegisterValidation("status", checkStatus)
	validate.RegisterTranslation("status", trans, registerTranslator("status", "{0}不符合要求"), customTranslate)

	err = validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
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

// 自定义校验手机号码
func checkPhone(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())
	return ok
}

// 自定义校验状态码
func checkStatus(fl validator.FieldLevel) bool {
	status := fl.Field().Int()
	if status != 0 && status != 1 {
		return false
	}
	return true
}
