package serverless

import (
	"path/filepath"

    "github.com/jzero-io/jzero/core/configcenter"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	coresvc "github.com/jzero-io/jzero-admin/core-engine/svc"

	"{{ .Module }}/internal/config"
	"{{ .Module }}/internal/custom"
	"{{ .Module }}/internal/global"
    "{{ .Module }}/internal/handler"
    "{{ .Module }}/internal/svc"
)

type Serverless struct {
	SvcCtx        *svc.ServiceContext                                   // 服务上下文
	HandlerFunc   func(server *rest.Server, svcCtx *svc.ServiceContext) // 服务路由
	RouteCodesMap map[string]string                                     // 路由与 code 码映射
}

// New serverless function
func New(coreSvcCtx *coresvc.ServiceContext) *Serverless {
	cc := configcenter.MustNewConfigCenter[config.Config](configcenter.Config{
		Type: "yaml",
	}, subscriber.MustNewFsnotifySubscriber(filepath.Join("plugins", "{{ .DirName }}", "etc", "etc.yaml"), subscriber.WithUseEnv(true)))
	global.ServiceContext.ConfigCenter = cc

    customServer := custom.New()
    logx.Must(customServer.Init())

	svcCtx := svc.NewServiceContext(cc, handler.Route2Code, svc.WithCoreServiceContext(coreSvcCtx))
	global.ServiceContext = *svcCtx

	group := service.NewServiceGroup()
    group.Add(customServer)
    group.Start()

	return &Serverless{
		SvcCtx:      svcCtx,
		HandlerFunc: handler.RegisterHandlers,
	}
}