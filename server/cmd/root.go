package cmd

import (
	"log"
	"os"
	"strings"

	"github.com/a8m/envsubst"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	cfgFile    string
	cfgEnvFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "server root",
	Long:  "server root.",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "etc/etc.yaml", "config file (default is project root dir etc/etc.yaml")
	rootCmd.PersistentFlags().StringVar(&cfgEnvFile, "env", "etc/.env.yaml", "env file (default is project root dir etc/.env.yaml")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if len(os.Args) <= 1 || os.Args[1] != serverCmd.Name() {
		return
	}

	if _, err := os.Stat(cfgEnvFile); err == nil {
		data, err := envsubst.ReadFile(cfgEnvFile)
		if err != nil {
			log.Fatalf("envsubst error: %v", err)
		}
		var env map[string]any
		err = yaml.Unmarshal(data, &env)
		if err != nil {
			log.Fatalf("yaml unmarshal error: %v", err)
		}

		for k, v := range env {
			_ = os.Setenv(k, cast.ToString(v))
		}
	}

	externalEnvs := os.Environ()
	for _, v := range externalEnvs {
		splits := strings.Split(v, "=")
		if len(splits) == 2 {
			_ = os.Setenv(splits[0], splits[1])
		}
	}
}
