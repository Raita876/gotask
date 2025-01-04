package main

import (
	"log"
	"os"

	"github.com/avast/retry-go/v4"
	"github.com/raita876/gotask/internal/adapter/rest"
	"github.com/raita876/gotask/internal/infra/db/mysql"
	"github.com/raita876/gotask/internal/usecase"
	"github.com/urfave/cli/v2"

	mysqldriver "gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	version string
	name    string
)

const (
	USAGE      = "This is a REST API for task management"
	USAGE_TEXT = "gotask [OPTION]..."
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Version:   version,
		Name:      name,
		Usage:     USAGE,
		UsageText: USAGE_TEXT,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dsn",
				Aliases: []string{"d"},
				Value:   "mysql:mysql@tcp(mysql:3306)/task_db?charset=utf8mb4&parseTime=True&loc=Local",
				Usage:   "specify data source name",
			},
			&cli.StringFlag{
				Name:    "addr",
				Aliases: []string{"a"},
				Value:   "0.0.0.0:8080",
				Usage:   "specify addr",
			},
		},
		Action: func(ctx *cli.Context) error {
			dsn := ctx.String("dsn")
			addr := ctx.String("addr")

			gormDB, err := retry.DoWithData(
				func() (*gorm.DB, error) {
					gormDB, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{})
					log.Println(err)
					return gormDB, err
				},
			)
			if err != nil {
				log.Fatalf("Failed to connect to database: %v", err)
			}

			userRepository := mysql.NewDBUserRepository(gormDB)
			taskRepository := mysql.NewDBTaskRepository(gormDB)

			userUseCase := usecase.NewUserInteractor(userRepository)
			taskUseCase := usecase.NewTaskInteractor(taskRepository)

			e := rest.Setup(
				userUseCase,
				taskUseCase,
				name, version, USAGE,
			)

			return e.Start(addr)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Failed to run server: %v\n", err)
	}

}
