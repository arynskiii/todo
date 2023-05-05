package main

import (
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db:%s", err.Error())
	}
	repos := repository.NewRepository(*db)
	service := service.NewService(*repos)
	handler := handler.NewHandler(*service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
