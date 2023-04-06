package store

import (
	"fmt"

	"github.com/ShinyShield/AbyShopServ/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() {

	databasesConfig := config.GetDatabaseConfig()

	var err error
	DB, err = gorm.Open("mysql", mysqlSource(&databasesConfig))
	if err != nil {
		panic("init mysql failed: " + err.Error())
	}

	DB.LogMode(databasesConfig.Logmode)
}

func mysqlSource(config *config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=UTC&multiStatements=true",
		config.Username,
		config.Password,
		config.Name,
		config.Host,
		config.Port,
		config.Encoding,
	)
	//	return fmt.Sprintf(":"+config.Username+"@tcp("+config.Password+":"+config.Name+")/"+config.Host+"?charset="++"&parseTime=true&loc=UTC&multiStatements=true",
	//"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=UTC&multiStatements=true"
}
