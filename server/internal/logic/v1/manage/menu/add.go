package menu

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/menu"
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
		if err = l.svcCtx.Model.ManageMenu.InsertV2(l.ctx, session, &manage_menu.ManageMenu{
			Uuid:            uuid.New().String(),
			Status:          req.Status,
			ParentUuid:      req.ParentUuid,
			MenuType:        req.MenuType,
			MenuName:        req.MenuName,
			HideInMenu:      cast.ToInt64(req.HideInMenu),
			ActiveMenu:      req.ActiveMenu,
			Order:           req.Order,
			RouteName:       req.RouteName,
			RoutePath:       req.RoutePath,
			Component:       req.Component,
			Icon:            req.Icon,
			IconType:        req.IconType,
			I18nKey:         req.I18nKey,
			KeepAlive:       cast.ToInt64(req.KeepAlive),
			Href:            req.Href,
			MultiTab:        cast.ToInt64(req.MultiTab),
			FixedIndexInTab: cast.ToInt64(req.FixedIndexInTab),
			Query:           marshal(req.Query),
			ButtonCode:      req.ButtonCode,
			Permissions:     marshal(req.Permissions),
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
