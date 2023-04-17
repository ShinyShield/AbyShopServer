package store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init() error {
	var err error
	//databasesConfig := config.GetDatabaseConfig()

	DB, err = gorm.Open("mysql", "root:admin@tcp(localhost:3306)/abyshopdb?charset=&parseTime=true&loc=UTC&multiStatements=true")
	/* mysqlSource(&databasesConfig))*/
	if err != nil {
		return err
	}

	//DB.LogMode(databasesConfig.Logmode)
	return nil
}
func Close() {
	DB.Close()
}
