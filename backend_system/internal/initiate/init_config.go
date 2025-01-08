package initiate

import (
	"github.com/redis/go-redis/v9"
	"github.com/wqh/smart/school/system/internal/errorx"
	"log"

	"github.com/spf13/viper"
	"github.com/wqh/smart/school/system/internal/configs"
	"github.com/wqh/smart/school/system/internal/initiate/db"
	"gorm.io/gorm"
)

var (
	Config           configs.Config
	DB               *gorm.DB
	RDB              *redis.Client
	LOGIN_UNKNOWN    = errorx.NewError(202, "用户不存在")
	INVALID_PASSWORD = errorx.NewError(203, "密码错误")
	INVALID_EMAIL    = errorx.NewError(203, "账户邮箱错误")
	INVALID_ARGUMENT = errorx.NewError(400, "参数错误")
	NOT_FOUND        = errorx.NewError(404, "资源不存在")
	ERROR            = errorx.NewError(400, "操作失败")
	TO_MANY          = errorx.NewError(429, "操作过于频繁")
	UNAUTHENTICATED  = errorx.NewError(401, "您还未登录")
	INNER_ERROR      = errorx.NewError(500, "内部错误")
	SecretKey        = "asdfghjklqwertyuiopzxcvbnm"
	Expiry           = 60 * 60 * 24
)

// 初始化配置文件
func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/resources")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Error unmarshalling config, %s", err)
	}
	db.InitMysql(DB, Config)
	db.InitRedis(RDB, Config)
}
