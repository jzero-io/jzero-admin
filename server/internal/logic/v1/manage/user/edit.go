package user

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_role"
	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/v1/manage/user"
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
	data, err := l.svcCtx.Model.ManageUser.FindOneByUuid(l.ctx, nil, req.Uuid)
	if err != nil {
		return nil, err
	}

	newData := lo.FromPtr(data)
	newData.Username = req.Username
	newData.Nickname = req.NickName
	newData.Email = req.UserEmail
	newData.Phone = req.UserPhone
	newData.Gender = req.UserGender
	newData.Status = req.Status

	if err = l.svcCtx.Model.ManageUser.Update(l.ctx, nil, lo.ToPtr(newData)); err != nil {
		return nil, err
	}

	// 更新 system_user_role 表
	if err = l.svcCtx.Model.ManageUserRole.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    manage_user_role.UserUuid,
		Operator: condition.Equal,
		Value:    req.Uuid,
	}); err != nil {
		return nil, err
	}
	var bulk []*manage_user_role.ManageUserRole
	roles, err := l.svcCtx.Model.ManageRole.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    manage_role.Code,
		Operator: condition.In,
		Value:    req.UserRoles,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range roles {
		bulk = append(bulk, &manage_user_role.ManageUserRole{
			Uuid:     uuid.New().String(),
			UserUuid: data.Uuid,
			RoleUuid: v.Uuid,
		})
	}

	if err = l.svcCtx.Model.ManageUserRole.BulkInsert(l.ctx, nil, bulk); err != nil {
		return nil, err
	}

	return
}
