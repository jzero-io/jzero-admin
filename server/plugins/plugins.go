// Code generated by jzero. DO NOT EDIT.
package plugins

import (
	hello "github.com/jzero-io/jzero-admin-plugins/hello/serverless"
	"github.com/jzero-io/jzero-admin/server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

type CoreSvcCtx = svc.ServiceContext

func LoadPlugins(server *rest.Server, svcCtx CoreSvcCtx) {

	{
		serverless := hello.New(svcCtx)
		serverless.HandlerFunc(server, serverless.SvcCtx)
	}

}
