package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

var dbConfig *DatabaseConfig

func LoadDbConfig() {
	data, err := ioutil.ReadFile("config/dbconfig.json")
	if err != nil {
		panic(err)
	}
	config := &DatabaseConfig{}
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	fmt.Println("LOGG:")
	fmt.Println(config)
	dbConfig = config
}

//GetConnectionString returns a database connection string
func GetConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name)
}
