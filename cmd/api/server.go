package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/sumitdhameja/services-hub/internal/errors"
	"github.com/sumitdhameja/services-hub/internal/logger"
	"github.com/sumitdhameja/services-hub/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var (
	VERSION = "0.0.0"
	COMMIT  = "unknown"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "API Server for Services Hub",
	Long: `This command runs API Server to serve requests. Default Port is 8000. 
	Configuration is found under config folder`,
	Run: func(cmd *cobra.Command, args []string) {
		serverInit()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func serverInit() {

	logger.Info("[-] SERVICE HUB")
	logger.Info("    - Version:", VERSION)
	logger.Info("    - Commit:", COMMIT)
	logger.Info("[-] SERVER")
	logger.Info("    - HOST: ", cfg.Server.Host)
	logger.Info("    - PORT: ", cfg.Server.Port)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/service_hub?charset=UTF8&parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		logger.Fatal("Failed to connect database: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal("Can't connect to database")
	}
	go models.AutoMigrate(db)

	sqlDB.SetMaxOpenConns(cfg.Database.MaxConnections)
	defer func() {
		sqlDB.Close()
		logger.Info("Closed db connection")
	}()

	router := gin.New()
	router.Use(errors.GinError())
	router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
