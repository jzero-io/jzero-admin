package user

import (
	"context"
	"net/http"
	"time"

	"github.com/guregu/null/v5"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/internal/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/user"
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
	user, err := l.svcCtx.Model.ManageUser.FindOne(l.ctx, nil, req.Id)
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

	if err = l.svcCtx.Model.ManageUser.Update(l.ctx, nil, user); err != nil {
		return nil, err
	}

	// 更新 system_user_role 表
	if err = l.svcCtx.Model.ManageUserRole.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    "user_id",
		Operator: condition.Equal,
		Value:    req.Id,
	}); err != nil {
		return nil, err
	}
	var bulk []*manage_user_role.ManageUserRole
	var roleIds []uint64
	roles, err := l.svcCtx.Model.ManageRole.FindByCondition(l.ctx, nil, condition.Condition{
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
		bulk = append(bulk, &manage_user_role.ManageUserRole{
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			UserId:     int64(user.Id),
			RoleId:     int64(v),
		})
	}

	if err = l.svcCtx.Model.ManageUserRole.BulkInsert(l.ctx, nil, bulk); err != nil {
		return nil, err
	}

	return
}
