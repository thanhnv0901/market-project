package marketdb

import (
	"fmt"
	"log"
	"market_apis/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	configuation           = configs.GetConfig()
	marketDB     *MarketDB = nil
)

// MarketDB ..
type MarketDB struct {
	host     string
	port     int
	username string
	password string
	database string
	Conn     *gorm.DB
}

// GetConnection ..
func (m *MarketDB) GetConnection() *gorm.DB {
	return m.Conn
}

// newMarketConnection ..
func newMarketConnection(host string, port int, user string, password string, database string) (*MarketDB, error) {

	if host == "" || port == 0 || user == "" || password == "" || database == "" {
		return nil, fmt.Errorf("Error in Open Postgre Connection, wrong in configuation")
	}

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		host, port, user, password, database)

	db, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("Error in Open Postgre Connection to Market DB, ")
	}

	tmp := &MarketDB{
		host:     host,
		port:     port,
		username: user,
		password: password,
		database: database,
		Conn:     db,
	}

	return tmp, nil
}

func init() {

	var err error
	marketDB, err = newMarketConnection(configuation.MarketPostgreDBHost, configuation.MarketPostgreDBPort, configuation.MarketPostgreDBUsername, configuation.MarketPostgreDBPassword, configuation.MarketPostgreDatabase)
	if err != nil {
		log.Fatalln(err)
	}

}

// GetMarketDB ..
func GetMarketDB() *MarketDB {
	return marketDB
}
