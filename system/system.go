package system

import (
	"encoding/gob"
	"github.com/Sirupsen/logrus"
	"github.com/andreyberezin/gin-site/models"
)

func InitModelsForSession() {
	gob.Register(models.Login{})
	gob.Register(models.Register{})
}

func LogErr(e error) {
	if e != nil {
		logrus.Error(e)
	}
}

func LogInfo(e error) {
	if e != nil {
		logrus.Error(e)
	}
}
