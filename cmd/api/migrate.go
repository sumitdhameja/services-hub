package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/sumitdhameja/services-hub/internal/logger"
	"github.com/sumitdhameja/services-hub/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "creates schema needed and populates with fake data",
	Long:  `Creates DB schema and populates with fake data`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCMD)
}

func migrate() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/service_hub?charset=UTF8&parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	})
	if err != nil {
		logger.Fatal("Failed to connect database: ", err)
	}
	seed(db)
}
func seed(db *gorm.DB) {
	models.AutoMigrate(db)
	maxRecords := 20
	minRecords := 5

	for i := 0; i < rand.Intn(maxRecords-minRecords)+minRecords; i++ {
		db.Create(&models.User{
			BaseModel: models.BaseModel{
				ID:        uuid.NewV4().String(),
				CreatedOn: time.Now(),
				UpdatedOn: time.Now(),
			},
			Email: faker.Email(),
			Name:  faker.Name(),
		})
	}
	var users []models.User
	db.Find(&users)
	for _, u := range users {
		// random services for users
		for i := 0; i < rand.Intn(maxRecords-minRecords)+minRecords; i++ {
			db.Create(&models.Service{
				BaseModel: models.BaseModel{
					ID:        uuid.NewV4().String(),
					CreatedOn: time.Now(),
					UpdatedOn: time.Now(),
				},
				Title:       faker.Sentence(),
				Description: faker.Paragraph(),
				UserID:      u.ID,
			})
		}

	}
	var services []models.Service
	db.Find(&services)
	for _, s := range services {
		// random service versions for services
		for i := 0; i < rand.Intn(maxRecords-minRecords)+minRecords; i++ {
			db.Create(&models.ServiceVersion{
				BaseModel: models.BaseModel{
					ID:        uuid.NewV4().String(),
					CreatedOn: time.Now(),
					UpdatedOn: time.Now(),
				},
				Version:        fmt.Sprintf("v%v", rand.Intn(10)),
				ServiceID:      s.ID,
				URL:            faker.URL(),
				OtherCoolStuff: faker.Sentence(),
			})
		}

	}
}
