package main

import (
	"atom-go/config"
	"atom-go/internal/router"
	"atom-go/pkg/dal"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {

	// 初始化配置
	config.InitConfig()

	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := config.Data.Mysql.Username + ":" + config.Data.Mysql.Password + "@tcp(" + config.Data.Mysql.Host + ":" + strconv.Itoa(config.Data.Mysql.Port) + ")/" + config.Data.Mysql.Database + "?charset=" + config.Data.Mysql.Charset + "&parseTime=True&loc=Local"

	// 初始化数据访问层
	dal.InitDal(&dal.Config{
		GomrConfig: &dal.GomrConfig{
			Dialector: mysql.Open(dsn),
			Opts: &gorm.Config{
				SkipDefaultTransaction: true, // 跳过默认事务
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
				Logger: logger.New(log.Default(), logger.Config{
					// LogLevel: logger.Silent, // 不打印日志
					LogLevel:                  logger.Error, // 打印错误日志
					IgnoreRecordNotFoundError: true,
				}),
			},
			MaxOpenConns: config.Data.Mysql.MaxOpenConns,
			MaxIdleConns: config.Data.Mysql.MaxIdleConns,
		},
		// RedisConfig: &dal.RedisConfig{
		// 	Host:     config.Data.Redis.Host,
		// 	Port:     config.Data.Redis.Port,
		// 	Database: config.Data.Redis.Database,
		// 	Password: config.Data.Redis.Password,
		// },
	})

	// 设置模式
	gin.SetMode(config.Data.App.Server.Mode)

	// 初始化gin
	server := gin.New()

	// 使用恢复中间件
	server.Use(gin.Recovery())

	// 设置静态资源目录
	// 如果前端使用的是 history 路由模式，需要使用 nginx 代理
	// 注释 server.Static("/admin", "web/admin")
	// 如果前后端不分离方式部署需要配置前端为 hash 路由模式
	// 解除 server.Static("/admin", "web/admin") 注释
	// 并在项目根目录下创建 web/admin 目录，将前端打包后的 dist 内的文件复制到该目录下
	// server.Static("/admin", "web/admin")
	// 设置文件存储静态资源
	server.Static(config.Data.App.StoragePath, config.Data.App.StoragePath)

	// 注册后台路由
	router.AdminRegister(server)

	server.Run(":" + strconv.Itoa(config.Data.App.Server.Port))
}
