package dal

import "sync"

type Config struct {
	GomrConfig      *GomrConfig
	GomrConfigOrder *GomrConfigOrder
	RedisConfig     *RedisConfig
}

var once sync.Once

// 初始化数据访问层
func InitDal(config *Config) {

	once.Do(func() {
		// 初始化数据库
		if config.GomrConfig != nil {
			initGorm(config.GomrConfig)
		}
		// 初始化订单数据库
		if config.GomrConfigOrder != nil {
			initGormOrder(config.GomrConfigOrder)
		}
		// 初始化 Redis
		if config.RedisConfig != nil {
			initRedis(config.RedisConfig)
		}
	})
}
