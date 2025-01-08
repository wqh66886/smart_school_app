package initiate

import (
	"log"

	"github.com/spf13/viper"
	"github.com/wqh/smart/school/system/internal/configs"
	"github.com/wqh/smart/school/system/internal/initiate/db"
	"gorm.io/gorm"
)

var (
	Config configs.Config
	DB     *gorm.DB
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
}
