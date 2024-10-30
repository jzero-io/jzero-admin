package user

import (
	"context"
	"time"

	"github.com/guregu/null/v5"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/internal/model/system_user_role"
	"server/internal/svc"
	types "server/internal/types/system/user"
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
	user, err := l.svcCtx.Model.SystemUser.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	user.Username = req.Username
	user.Nickname = req.NickName
	user.Email = null.StringFrom(req.UserEmail).NullString
	user.Phone = null.StringFrom(req.UserPhone).NullString
	user.Gender = req.UserGender
	user.Status = req.Status
	user.UpdateTime = time.Now()

	if err = l.svcCtx.Model.SystemUser.Update(l.ctx, user); err != nil {
		return nil, err
	}

	// 更新 system_user_role 表
	if err = l.svcCtx.Model.SystemUserRole.DeleteByCondition(l.ctx, condition.Condition{
		Field:    "user_id",
		Operator: condition.Equal,
		Value:    req.Id,
	}); err != nil {
		return nil, err
	}
	var bulk []*system_user_role.SystemUserRole
	var roleIds []uint64
	roles, err := l.svcCtx.Model.SystemRole.FindByCondition(l.ctx, condition.Condition{
		Field:    "code",
		Operator: condition.In,
		Value:    req.UserRoles,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range roles {
		roleIds = append(roleIds, v.Id)
	}
	for _, v := range roleIds {
		bulk = append(bulk, &system_user_role.SystemUserRole{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			UserId:     int64(user.Id),
			RoleId:     int64(v),
		})
	}

	if err = l.svcCtx.Model.SystemUserRole.BulkInsert(l.ctx, bulk); err != nil {
		return nil, err
	}

	return
}
