package user

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/internal/svc"
	types "github.com/jzero-io/jzero-admin/server/internal/types/manage/user"
)

type Delete struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Delete {
	return &Delete{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, r: r,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}

	err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err = l.svcCtx.Model.ManageUser.DeleteByCondition(l.ctx, session, condition.Condition{
			Field:    "id",
			Operator: condition.In,
			Value:    req.Ids,
		}); err != nil {
			return err
		}
		if err = l.svcCtx.Model.ManageUserRole.DeleteByCondition(l.ctx, session, condition.NewChain().In("user_id", req.Ids).Build()...); err != nil {
			return err
		}
		return nil
	})

	return
}
