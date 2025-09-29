package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/core/logx"
)

type ValidatorMiddleware struct {
	instance     *validator.Validate
	zhTranslator unTrans.Translator
}

func NewValidatorMiddleware() *ValidatorMiddleware {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return getLabelValue(field)
	})

	zhTranslator, _ := uni.GetTranslator("zh_Hans_CN")
	err := zhTrans.RegisterDefaultTranslations(validate, zhTranslator)
	logx.Must(err)

	return &ValidatorMiddleware{
		instance:     validate,
		zhTranslator: zhTranslator,
	}
}

func (v *ValidatorMiddleware) Validate(r *http.Request, data any) error {
	var errJoin error

	err := v.instance.Struct(data)
	if err != nil {
		for _, ve := range err.(validator.ValidationErrors) {
			if v.zhTranslator != nil && r.Context().Value("lang").(string) == "zh-CN" {
				errJoin = errors.Join(errors.New(ve.Translate(v.zhTranslator)))
			} else {
				errJoin = errors.Join(ve)
			}
		}
	}
	return errJoin
}

func getLabelValue(field reflect.StructField) string {
	tags := []string{"label", "json", "form", "path"}
	label := ""

	for _, tag := range tags {
		label = field.Tag.Get(tag)
		if label != "" {
			return label
		}
	}
	return ""
}

func RegisterTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans unTrans.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

func Translate(trans unTrans.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	if len(strings.Split(fe.Namespace(), ".")) >= 2 {
		return fmt.Sprintf("%s%s", strings.Split(fe.Namespace(), ".")[1], msg)
	}
	return msg
}
