package middleware

import (
	"net/http"
	"reflect"

	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/pkg/errors"
)

type Validator struct{}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(r *http.Request, data any) (err error) {
	validate := validator.New()
	var trans unTrans.Translator

	lang := r.Context().Value("lang")
	switch lang {
	case "zh-CN":
		uni := unTrans.New(zh_Hans_CN.New())
		trans, _ = uni.GetTranslator("zh_Hans_CN")
		err = zhTrans.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			return err
		}
	case "en-US":
		uni := unTrans.New(en_US.New())
		trans, _ = uni.GetTranslator("en_US")
		err = enTrans.RegisterDefaultTranslations(validate, trans)
	}

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return getLabelValue(field)
	})

	err = validate.Struct(data)
	if err != nil {
		for _, ve := range err.(validator.ValidationErrors) {
			if trans != nil {
				return errors.Errorf(ve.Translate(trans))
			}
			return ve
		}
	}
	return nil
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
