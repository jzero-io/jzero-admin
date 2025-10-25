package role

import (
	"context"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	roles, total, err := l.svcCtx.Model.ManageRole.PageByCondition(l.ctx, nil, condition.Condition{
		Operator: condition.Limit,
		Value:    req.Size,
	}, condition.Condition{
		Operator: condition.Offset,
		Value:    (req.Current - 1) * req.Size,
	}, condition.Condition{
		Skip:     req.RoleName == "",
		Field:    manage_role.Name,
		Operator: condition.Like,
		Value:    "%" + req.RoleName + "%",
	}, condition.Condition{
		Skip:     req.RoleCode == "",
		Field:    manage_role.Code,
		Operator: condition.Like,
		Value:    "%" + req.RoleCode + "%",
	}, condition.Condition{
		Skip:     req.Status == "",
		Field:    manage_role.Status,
		Operator: condition.Equal,
		Value:    req.Status,
	})

	var records []types.ManageRole
	for _, role := range roles {
		records = append(records, types.ManageRole{
			Uuid:       role.Uuid,
			RoleCode:   role.Code,
			RoleName:   role.Name,
			RoleDesc:   role.Desc,
			Status:     role.Status,
			CreateTime: role.CreateTime.Format(time.DateTime),
			UpdateTime: role.UpdateTime.Format(time.DateTime),
		})
	}

	resp = &types.ListResponse{
		Records: records,
		PageResponse: types.PageResponse{
			Current: req.Current,
			Size:    req.Size,
			Total:   total,
		},
	}
	return
}
