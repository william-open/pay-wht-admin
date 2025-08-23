package main

import (
	"log"
	"os"
	"strconv"
	"time"
	"wht-admin/app/router"
	"wht-admin/config"
	"wht-admin/framework/dal"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	// 数据库连接字符串
	dsn := config.Data.Mysql.Username + ":" + config.Data.Mysql.Password + "@tcp(" + config.Data.Mysql.Host + ":" + strconv.Itoa(config.Data.Mysql.Port) + ")/" + config.Data.Mysql.Database + "?charset=" + config.Data.Mysql.Charset + "&parseTime=True&loc=Local"
	orderDsn := config.Data.MysqlOrder.Username + ":" + config.Data.MysqlOrder.Password + "@tcp(" + config.Data.MysqlOrder.Host + ":" + strconv.Itoa(config.Data.MysqlOrder.Port) + ")/" + config.Data.MysqlOrder.Database + "?charset=" + config.Data.MysqlOrder.Charset + "&parseTime=True&loc=Local"

	// 自定义GORM日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢SQL阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	// 初始化数据访问层
	dal.InitDal(&dal.Config{
		GomrConfig: &dal.GomrConfig{
			Dialector: mysql.Open(dsn),
			Opts: &gorm.Config{
				SkipDefaultTransaction: true, // 跳过默认事务
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true, // 使用单数表名
				},
				Logger: newLogger, // 使用自定义日志配置
			},
			MaxOpenConns: config.Data.Mysql.MaxOpenConns,
			MaxIdleConns: config.Data.Mysql.MaxIdleConns,
		},
		GomrConfigOrder: &dal.GomrConfigOrder{
			Dialector: mysql.Open(orderDsn),
			Opts: &gorm.Config{
				SkipDefaultTransaction: true, // 跳过默认事务
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true, // 使用单数表名
				},
				Logger: newLogger, // 使用自定义日志配置
			},
			MaxOpenConns: config.Data.MysqlOrder.MaxOpenConns,
			MaxIdleConns: config.Data.MysqlOrder.MaxIdleConns,
		},
		RedisConfig: &dal.RedisConfig{
			Host:     config.Data.Redis.Host,
			Port:     config.Data.Redis.Port,
			Database: config.Data.Redis.Database,
			Password: config.Data.Redis.Password,
		},
	})

	// 设置gin模式
	gin.SetMode(config.Data.Server.Mode)

	// 初始化gin
	server := gin.New()

	// 设置受信任的代理IP
	server.SetTrustedProxies([]string{"127.0.0.1"})

	// 使用恢复中间件
	server.Use(gin.Recovery())

	// 设置文件资源目录
	server.Static(config.Data.Ruoyi.UploadPath, config.Data.Ruoyi.UploadPath)

	// 注册路由
	router.Register(server)

	// 启动服务器
	server.Run(":" + strconv.Itoa(config.Data.Server.Port))
}
