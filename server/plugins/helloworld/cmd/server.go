package cmd

import (
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/jzero-io/jzero/core/configcenter/subscriber"
	"github.com/spf13/cobra"
	configurator "github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"

	"helloworld/internal/config"
	"helloworld/internal/custom"
	"helloworld/internal/global"
	"helloworld/internal/handler"
	"helloworld/internal/middleware"
	"helloworld/internal/svc"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "helloworld server",
	Long:  "helloworld server",
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

		restServer := rest.MustNewServer(c.Rest.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
			httpx.ErrorCtx(r.Context(), w, err)
		}), rest.WithCustomCors(func(header http.Header) {
			header.Set("Access-Control-Allow-Origin", "*")
			header.Add("Access-Control-Allow-Headers", "X-Request-Id")
			header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		}, nil, "*"))

		customServer := custom.New()

		logx.Must(customServer.Init())
		svcCtx := svc.NewServiceContext(cc, handler.Route2Code)
		global.ServiceContext = *svcCtx
		middleware.Register(restServer)
		handler.RegisterHandlers(restServer, svcCtx)

		group := service.NewServiceGroup()
		group.Add(restServer)
		group.Add(customServer)
		group.Start()
	},
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
