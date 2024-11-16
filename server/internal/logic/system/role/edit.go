package role

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/svc"
	types "server/internal/types/system/role"
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
	role, err := l.svcCtx.Model.SystemRole.FindOne(l.ctx, nil, req.Id)
	if err != nil {
		return nil, err
	}
	role.Code = req.RoleCode
	role.Name = req.RoleName
	role.Desc = req.RoleDesc
	role.Status = req.Status
	err = l.svcCtx.Model.SystemRole.Update(l.ctx, nil, role)
	return
}
