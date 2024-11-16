package menu

import (
	"context"

	"github.com/guregu/null/v5"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system/menu"
)

type Edit struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEdit(ctx context.Context, svcCtx *svc.ServiceContext) *Edit {
	return &Edit{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Edit) Edit(req *types.EditRequest) (resp *types.EditResponse, err error) {
	one, err := l.svcCtx.Model.SystemMenu.FindOne(l.ctx, nil, req.Id)
	if err != nil {
		return nil, err
	}

	one.Status = req.Status
	one.ParentId = int64(req.ParentId)
	one.MenuType = req.MenuType
	one.MenuName = req.MenuName
	one.HideInMenu = cast.ToInt64(req.HideInMenu)
	one.ActiveMenu = null.StringFrom(req.ActiveMenu).NullString
	one.Order = int64(req.Order)
	one.RouteName = req.RouteName
	one.RoutePath = req.RoutePath
	one.Component = req.Component
	one.Icon = req.Icon
	one.IconType = req.IconType
	one.I18nKey = req.I18nKey
	one.KeepAlive = cast.ToInt64(req.KeepAlive)
	one.Href = null.StringFrom(req.Href).NullString
	one.MultiTab = null.IntFrom(cast.ToInt64(req.MutiTab)).NullInt64
	one.FixedIndexInTab = cast.ToInt64(req.FixedIndexInTab)
	one.Query = null.StringFrom(marshal(req.Query)).NullString
	one.Buttons = null.StringFrom(marshal(req.Buttons)).NullString
	one.Constant = cast.ToInt64(req.Constant)

	err = l.svcCtx.Model.SystemMenu.Update(l.ctx, nil, one)
	return
}
