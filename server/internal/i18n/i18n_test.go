package i18n

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	c := Conf{
		Dir: "",
	}

	trans := NewTranslator(c, LocaleFS)

	res := trans.Trans(context.WithValue(context.Background(), "lang", "zh-CN"), "manage.menu.existSubMenu")
	assert.Equal(t, "存在子菜单, 请先删除子菜单", res)

	res = trans.Trans(context.WithValue(context.Background(), "lang", "en-US"), "manage.menu.existSubMenu")
	assert.Equal(t, "Exist sub menu, please delete first", res)
}
