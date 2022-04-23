package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sumitdhameja/services-hub/config"
	"github.com/sumitdhameja/services-hub/internal/logger"
)

var (
	cfg *config.Config
	env string
)

var rootCmd = &cobra.Command{
	Use:   "server-api",
	Short: "Welcome to Service Hub API",
	Long: `Commands available: server/migrate. Server is used to run API(default port is 8000). Use migrate command 
	to create schema & populate database`,
}

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&env, "env", "dev", "environment to be used (default is dev)")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	var err error
	if cfg, err = config.Load(strings.ToLower(env)); err != nil {
		panic(err)
	}
	logger.InitializeWithWriter(logger.LevelFromString(cfg.LogLevel), os.Stdout)
}
