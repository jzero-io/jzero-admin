package serverless

import (
	"path/filepath"

	coresvc "github.com/jzero-io/jzero-admin/core-engine/svc"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"helloworld/internal/config"
	"helloworld/internal/custom"
	"helloworld/internal/global"
	"helloworld/internal/handler"
	"helloworld/internal/svc"
)

type Serverless struct {
	SvcCtx        *svc.ServiceContext                                   // 服务上下文
	HandlerFunc   func(server *rest.Server, svcCtx *svc.ServiceContext) // 服务路由
	RouteCodesMap map[string]string                                     // 路由与 code 码映射
}

// New serverless function
func New(coreSvcCtx *coresvc.ServiceContext) *Serverless {
	cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
		Type: "yaml",
	}, subscriber.MustNewFsnotifySubscriber(filepath.Join("plugins", "helloworld", "etc", "etc.yaml"), subscriber.WithUseEnv(true)))

	c, err := cc.GetConfig()
	logx.Must(err)

	customServer := custom.New(c)
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
