package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/olegovin/todo-app"
	"github.com/olegovin/todo-app/pkg/handler"
	"github.com/olegovin/todo-app/repository"
	"github.com/olegovin/todo-app/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("ошибка инициализации конфигурации: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBname:   "postgres",
		SSLmode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("ошибка инициализации базы данных: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("ошибка при запуске: %s", err.Error())

		}

	}()

	logrus.Print("todoapp let'go!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("todoapp exit")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down %s ", err.Error())

	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close≠ %s ", err.Error())

	}

}

func initConfig() error {
	viper.AddConfigPath("/Users/olegvinogradov/todo-app/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
