package grpc

import (
	"context"
	golog "log"
	"net"
	"os"
	"time"

	v1 "github.com/coding-standard/golang-project-layout/api/golang-project-layout/v1"
	"github.com/coding-standard/golang-project-layout/internal/mid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"

	"github.com/coding-standard/golang-project-layout/internal/dao"
	"github.com/coding-standard/golang-project-layout/internal/dao/mysql"
	"github.com/coding-standard/golang-project-layout/internal/service"
	"github.com/coding-standard/golang-project-layout/pkg/db"
	"github.com/coding-standard/golang-project-layout/pkg/log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm/logger"
)

func Run(ctx context.Context, network, address string) error {
	//init grpc server and run
	l, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	go func() {
		defer func() error {
			if err := l.Close(); err != nil {
				return err
			}
			return nil
		}()
		<-ctx.Done()
	}()

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(mid.Auth), selector.MatchFunc(mid.AllButHealthZ)),
		),
		grpc.ChainStreamInterceptor(
			selector.StreamServerInterceptor(auth.StreamServerInterceptor(mid.Auth), selector.MatchFunc(mid.AllButHealthZ)),
		),
	)

	var daoInterface dao.Interface
	if daoInterface, err = initDao(); err != nil {
		return err
	}
	demoService := service.NewDemoService()
	v1.RegisterDemoServer(s, demoService)
	demoDbService := service.NewDemoDbService(daoInterface)
	// This is generated by protoc
	v1.RegisterDemoDbServer(s, demoDbService)

	go func() {
		defer s.GracefulStop()
		<-ctx.Done()
	}()

	go func() error {
		log.L(ctx).Infof("grpc listen on:%s\n", address)
		if err := s.Serve(l); err != nil {
			return err
		}
		return nil
	}()

	return nil
}

func initDao() (dao.Interface, error) {
	newLogger := logger.New(
		golog.New(os.Stdout, "", golog.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.LogLevel(viper.GetInt("data.database.log-level")),
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	options := db.Options{
		Host:                  viper.GetString("data.database.host"),
		Port:                  viper.GetString("data.database.port"),
		Username:              viper.GetString("data.database.user"),
		Password:              viper.GetString("data.database.password"),
		Database:              viper.GetString("data.database.database"),
		MaxIdleConnections:    viper.GetInt("data.database.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("data.database.max-open-connections"),
		MaxConnectionLifeTime: time.Duration(viper.GetInt("data.database.max-connection-lifetime")) * time.Second,
		Logger:                newLogger,
	}
	databaseDao, err := mysql.GetDao(&options)
	if err != nil {
		return nil, err
	}
	return databaseDao, nil
}
