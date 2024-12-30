package menu

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	null "github.com/guregu/null/v5"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/menu"
)

type Add struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Add {
	return &Add{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Add) Add(req *types.AddRequest) (resp *types.AddResponse, err error) {
	err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err = l.svcCtx.Model.ManageMenu.Insert(l.ctx, session, &manage_menu.ManageMenu{
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
			MultiTab:        null.IntFrom(cast.ToInt64(req.MultiTab)).NullInt64,
			FixedIndexInTab: null.IntFromPtr(req.FixedIndexInTab).NullInt64,
			Query:           null.StringFrom(marshal(req.Query)).NullString,
			ButtonCode:      null.StringFrom(req.ButtonCode).NullString,
			Permissions:     null.StringFrom(marshal(req.Permissions)).NullString,
			Constant:        cast.ToInt64(req.Constant),
		}); err != nil {
			return err
		}
		return nil
	})
	return
}

func marshal(typeDefine any) string {
	marshalStr, _ := json.Marshal(typeDefine)
	return string(marshalStr)
}
