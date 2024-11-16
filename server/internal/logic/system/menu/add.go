package menu

import (
	"context"
	"encoding/json"
	"time"

	"github.com/guregu/null/v5"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/model/system_menu"
	"server/internal/svc"
	types "server/internal/types/system/menu"
)

type Add struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext) *Add {
	return &Add{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add) Add(req *types.AddRequest) (resp *types.AddResponse, err error) {
	_, err = l.svcCtx.Model.SystemMenu.Insert(l.ctx, nil, &system_menu.SystemMenu{
		CreateTime:      time.Now(),
		UpdateTime:      time.Now(),
		Status:          req.Status,
		ParentId:        int64(req.ParentId),
		MenuType:        req.MenuType,
		MenuName:        req.MenuName,
		HideInMenu:      cast.ToInt64(req.HideInMenu),
		ActiveMenu:      null.StringFrom(req.ActiveMenu).NullString,
		Order:           int64(req.Order),
		RouteName:       req.RouteName,
		RoutePath:       req.RoutePath,
		Component:       req.Component,
		Icon:            req.Icon,
		IconType:        req.IconType,
		I18nKey:         req.I18nKey,
		KeepAlive:       cast.ToInt64(req.KeepAlive),
		Href:            null.StringFrom(req.Href).NullString,
		MultiTab:        null.IntFrom(cast.ToInt64(req.MutiTab)).NullInt64,
		FixedIndexInTab: cast.ToInt64(req.FixedIndexInTab),
		Query:           null.StringFrom(marshal(req.Query)).NullString,
		Buttons:         null.StringFrom(marshal(req.Buttons)).NullString,
		Constant:        cast.ToInt64(req.Constant),
	})
	return
}

func marshal(typeDefine any) string {
	marshalStr, _ := json.Marshal(typeDefine)
	return string(marshalStr)
}
