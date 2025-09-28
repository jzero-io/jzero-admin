package cmd

import (
	"net/http"
	"os"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jzero-io/jzero-admin/server/internal/config"
	"github.com/jzero-io/jzero-admin/server/internal/custom"
	"github.com/jzero-io/jzero-admin/server/internal/global"
	"github.com/jzero-io/jzero-admin/server/internal/handler"
	"github.com/jzero-io/jzero-admin/server/internal/middleware"
	"github.com/jzero-io/jzero-admin/server/internal/svc"
	"github.com/jzero-io/jzero-admin/server/plugins"
)

type Server struct {
	Rest   *rest.Server
	Custom *custom.Custom
}

func NewServer(restConf rest.RestConf) *Server {
	restServer := rest.MustNewServer(restConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.ErrorCtx(r.Context(), w, err)
	}), rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Headers", "X-Request-Id")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	}, nil, "*"))
	return &Server{
		Rest:   restServer,
		Custom: custom.New(restServer),
	}
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server server",
	Long:  "server server",
	Run: func(cmd *cobra.Command, args []string) {
		cc := configurator.MustNewConfigCenter[config.Config](configurator.Config{
			Type: "yaml",
		}, subscriber.MustNewFsnotifySubscriber(cfgFile, subscriber.WithUseEnv(true)))

		c, err := cc.GetConfig()
		logx.Must(err)

		// set up logger
		logx.Must(logx.SetUp(c.Log.LogConf))
		if c.Log.LogConf.Mode != "console" {
			logx.AddWriter(logx.NewWriter(os.Stdout))
		}

		printBanner(c.Banner)
		printVersion()
		logx.Infof("Starting rest server at %s:%d...", c.Rest.Host, c.Rest.Port)

		NewServer(c.Rest.RestConf).run(cc)
	},
}

func (s *Server) run(cc configurator.Configurator[config.Config]) {
	logx.Must(s.Custom.Init(cc))

	svcCtx := svc.NewServiceContext(cc, handler.Route2Code)
	svcCtx.Middleware = middleware.NewMiddleware(svcCtx, handler.Route2Code)
	global.ServiceContext = *svcCtx
	middleware.Register(s.Rest)

	// server add api handlers
	handler.RegisterHandlers(s.Rest, svcCtx)

	plugins.LoadPlugins(s.Rest, *svcCtx)

	group := service.NewServiceGroup()
	group.Add(s.Rest)
	group.Add(s.Custom)
	group.Start()
}

func printBanner(c config.BannerConf) {
	figure.NewColorFigure(c.Text, c.FontName, c.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
