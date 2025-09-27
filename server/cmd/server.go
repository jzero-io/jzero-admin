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

		printBanner(c)
		printVersion()
		logx.Infof("Starting rest server at %s:%d...", c.Rest.Host, c.Rest.Port)

		svcCtx := svc.NewServiceContext(cc, handler.Route2Code)
		global.ServiceContext = *svcCtx
		run(svcCtx)
	},
}

func run(svcCtx *svc.ServiceContext) {
	server := rest.MustNewServer(svcCtx.MustGetConfig().Rest.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.ErrorCtx(r.Context(), w, err)
	}), rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Headers", "X-Request-Id")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	}, nil, "*"))

	ctm := custom.New(server)
	logx.Must(ctm.Init())

	middleware.Register(server)

	// server add api handlers
	handler.RegisterHandlers(server, svcCtx)

	plugins.LoadPlugins(server, *svcCtx)

	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(ctm)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
