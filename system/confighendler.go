package system

import (
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/andreyberezin/gin-site/models"
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
		LogErr(err)
	}
	config := &DatabaseConfig{}
	err = json.Unmarshal(data, config)
	if err != nil {
		LogErr(err)
	}
	logrus.Info(config)
	dbConfig = config
}

func InitAdminByConfig() {
	db := models.GetBD()
	data, err := ioutil.ReadFile("config/adminconfig.json")
	if err != nil {
		LogErr(err)
	}
	u := &models.User{}
	logrus.Info("Looking for an admin user in DB...")
	LogInfo(db.Where("role = ?", "admin").First(&u).Error)
	if u.ID == 0 {
		logrus.Info("Admin not found, creating a new admin record...")
		err := json.Unmarshal(data, u)
		if err != nil {
			LogErr(err)
		} else {
			LogErr(db.Create(u).Error)
		}
	}
}

//GetConnectionString returns a database connection string
func GetConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name)
}
