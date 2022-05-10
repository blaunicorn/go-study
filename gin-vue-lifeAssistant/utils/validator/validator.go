package validator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

func Validate(data interface{}) (string, int) {
	// Validate instantiation
	validate := validator.New()
	// Translation instantiation
	uni := unTrans.New(zh_Hans_CN.New())
	// Introducing translation methods
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	// Registered translation method
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	// 验证是否是结构体，并结合模型的validat 进行验证
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), 400
		}
	}
	return "", 200
}
