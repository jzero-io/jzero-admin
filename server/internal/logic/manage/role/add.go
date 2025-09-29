package role

import (
	"context"
	"net/http"
	"time"

	"github.com/guregu/null/v5"
	"github.com/jzero-io/jzero-admin/core-engine/helper/auth"
	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role_menu"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
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
	authInfo, err := auth.Info(l.ctx)
	if err != nil {
		return nil, err
	}

	// role code 唯一
	if _, err := l.svcCtx.Model.ManageRole.FindOneByCondition(l.ctx, nil, condition.NewChain().Equal("code", req.RoleCode).Build()...); err == nil {
		return nil, errors.New("角色编码已存在")
	}

	// find home menu
	var homeMenuId uint64
	if home, err := l.svcCtx.Model.ManageMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().Equal("route_path", "/home").Build()...); err != nil {
		return nil, err
	} else {
		homeMenuId = home.Id
	}

	err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if _, err = l.svcCtx.Model.ManageRole.Insert(l.ctx, session, &manage_role.ManageRole{
			Code:       req.RoleCode,
			Name:       req.RoleName,
			Desc:       req.RoleDesc,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			CreateBy:   null.IntFrom(int64(authInfo.Id)).NullInt64,
			Status:     req.Status,
		}); err != nil {
			return err
		}

		// get role id
		role, err := l.svcCtx.Model.ManageRole.FindOneByCondition(l.ctx, session, condition.NewChain().Equal("code", req.RoleCode).Build()...)
		if err != nil {
			return err
		}

		// 添加首页路由
		if _, err = l.svcCtx.Model.ManageRoleMenu.Insert(l.ctx, session, &manage_role_menu.ManageRoleMenu{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			CreateBy:   null.IntFrom(int64(authInfo.Id)).NullInt64,
			RoleId:     int64(role.Id),
			MenuId:     int64(homeMenuId),
			IsHome:     cast.ToInt64(true),
		}); err != nil {
			return err
		}

		return nil
	})

	return
}
