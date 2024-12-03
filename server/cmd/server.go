package cmd

import (
	"fmt"
	"net/http"
	"os"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"server/internal/config"
	"server/internal/handler"
	"server/internal/middleware"
	"server/internal/svc"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server server",
	Long:  "server server",
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(cfgFile, &c, conf.UseEnv())
		config.C = c

		// set up logger
		if err := logx.SetUp(c.Log.LogConf); err != nil {
			logx.Must(err)
		}
		if c.Log.LogConf.Mode != "console" {
			logx.AddWriter(logx.NewWriter(os.Stdout))
		}

		ctx := svc.NewServiceContext(c, handler.Route2Code)
		run(ctx)
	},
}

func run(svcCtx *svc.ServiceContext) {
	server := rest.MustNewServer(svcCtx.Config.Rest.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Headers", "X-Request-Id")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	}, nil, "*"))
	middleware.Register(server)

	// server add api handlers
	handler.RegisterHandlers(server, svcCtx)

	// server add custom routes
	svcCtx.Custom.AddRoutes(server)

	group := service.NewServiceGroup()
	group.Add(server)
	group.Add(svcCtx.Custom)

	printBanner(svcCtx.Config)
	fmt.Printf("\nUsing Database: %s\n", svcCtx.Config.DatabaseType)
	fmt.Printf("%s conf: %s\n", svcCtx.Config.DatabaseType, svc.BuildDataSource(svcCtx.Config))
	fmt.Printf("Using Cache: %s\n", svcCtx.Config.CacheType)
	logx.Infof("Starting rest server at %s:%d...", svcCtx.Config.Rest.Host, svcCtx.Config.Rest.Port)
	group.Start()
}

func printBanner(c config.Config) {
	figure.NewColorFigure(c.Banner.Text, c.Banner.FontName, c.Banner.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
