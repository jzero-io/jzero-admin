package role

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/role"
)

type GetAll struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewGetAll(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetAll {
	return &GetAll{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *GetAll) GetAll(req *types.GetAllRequest) (resp []types.GetAllResponse, err error) {
	roles, err := l.svcCtx.Model.ManageRole.FindByCondition(l.ctx, nil)
	if err != nil {
		return nil, err
	}
	var list []types.GetAllResponse

	for _, role := range roles {
		if role.Status == "1" {
			list = append(list, types.GetAllResponse{
				Id:       role.Id,
				RoleCode: role.Code,
				RoleName: role.Name,
			})
		}
	}

	return list, nil
}
