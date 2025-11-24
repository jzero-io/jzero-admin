package role

import (
	"context"
	"net/http"

	"github.com/samber/lo"
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
	data, err := l.svcCtx.Model.ManageRole.FindOneByUuid(l.ctx, nil, req.Uuid)
	if err != nil {
		return nil, err
	}

	newData := lo.FromPtr(data)
	newData.Code = req.RoleCode
	newData.Name = req.RoleName
	newData.Desc = req.RoleDesc
	newData.Status = req.Status

	err = l.svcCtx.Model.ManageRole.Update(l.ctx, nil, lo.ToPtr(newData))
	return
}
