package role

import (
	"context"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/role"
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
	roles, total, err := l.svcCtx.Model.ManageRole.PageByCondition(l.ctx, nil, condition.NewChain().
		Page(req.Current, req.Size).
		OrderByDesc(manage_role.CreateTime).
		Like(manage_role.Name, "%"+req.RoleName+"%", condition.WithSkip(req.RoleName == "")).
		Like(manage_role.Code, "%"+req.RoleCode+"%", condition.WithSkip(req.RoleCode == "")).
		Equal(manage_role.Status, req.Status, condition.WithSkip(req.Status == "")).
		Build()...)

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
