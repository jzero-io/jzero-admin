package role

import (
	"context"
	"net/http"

	"github.com/google/uuid"
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
	var homeMenuUuid string
	if home, err := l.svcCtx.Model.ManageMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().Equal("route_path", "/home").Build()...); err != nil {
		return nil, err
	} else {
		homeMenuUuid = home.Uuid
	}

	err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		roleUuid := uuid.New().String()
		if err = l.svcCtx.Model.ManageRole.InsertV2(l.ctx, session, &manage_role.ManageRole{
			Uuid:     roleUuid,
			Code:     req.RoleCode,
			Name:     req.RoleName,
			Desc:     req.RoleDesc,
			CreateBy: authInfo.Uuid,
			Status:   req.Status,
		}); err != nil {
			return err
		}

		// 添加首页路由
		if err = l.svcCtx.Model.ManageRoleMenu.InsertV2(l.ctx, session, &manage_role_menu.ManageRoleMenu{
			Uuid:     uuid.New().String(),
			CreateBy: authInfo.Uuid,
			RoleUuid: roleUuid,
			MenuUuid: homeMenuUuid,
			IsHome:   cast.ToInt64(true),
		}); err != nil {
			return err
		}

		return nil
	})

	return
}
