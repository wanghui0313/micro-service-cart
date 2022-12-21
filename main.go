package main

import (
	"fmt"
	"github.com/asim/go-micro/plugins/wrapper/ratelimiter/ratelimit/v3"
	"github.com/asim/go-micro/v3/registry"
	ratelimit2 "github.com/juju/ratelimit"
	"github.com/opentracing/opentracing-go"
	"github.com/wanghui0313/micro-service-cart/domain/repository"
	"github.com/wanghui0313/micro-service-cart/handler"
	pb "github.com/wanghui0313/micro-service-cart/proto"
	"github.com/wanghui0313/micro-service-common/common"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"

	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/go-micro/plugins/v3/registry/consul"
	opentracing2 "github.com/go-micro/plugins/v3/wrapper/trace/opentracing"
	cartservice "github.com/wanghui0313/micro-service-cart/domain/service"
	"gorm.io/driver/mysql"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "micro/config")
	if err != nil {
		logger.Error(err)
	}

	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "localhost:6831")
	if err != nil {
		logger.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//限流,每10ms填充一个令牌
	b := ratelimit2.NewBucket(10*time.Millisecond, 1)
	// Create service
	srv := service.NewService(
		service.Name("go.micro.service.cart"),
		service.Version("latest"),
		//这里设置地址和要暴露的端口
		service.Address("0.0.0.0:8087"),
		//添加consul作为注册中心
		service.Registry(consulRegistry),
		//绑定链路追踪
		service.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		service.WrapHandler(ratelimit.NewHandlerWrapper(b, false)),
	)

	//获取mysql配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := getMysqlDb(mysqlInfo.User, mysqlInfo.Pwd, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database)
	if err != nil {
		logger.Fatal(err)
	}
	cartService := cartservice.NewCartService(repository.NewCartRepository(db))
	if err = cartService.InitTable(); err != nil {
		logger.Fatal(err)
	}

	srv.Init()
	// Register handler
	pb.RegisterCartHandler(srv.Server(), handler.New(cartService))

	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

func getMysqlDb(user, pwd, host, port, database string) (*gorm.DB, error) {
	//创建数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //禁止复表
			},
		})
	return db, err
}
