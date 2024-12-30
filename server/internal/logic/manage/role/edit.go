package role

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
)

type Edit struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewEdit(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Edit {
	return &Edit{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Edit) Edit(req *types.EditRequest) (resp *types.EditResponse, err error) {
	role, err := l.svcCtx.Model.ManageRole.FindOne(l.ctx, nil, req.Id)
	if err != nil {
		return nil, err
	}
	role.Code = req.RoleCode
	role.Name = req.RoleName
	role.Desc = req.RoleDesc
	role.Status = req.Status
	err = l.svcCtx.Model.ManageRole.Update(l.ctx, nil, role)
	return
}
