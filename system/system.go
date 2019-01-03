package system

import (
	"encoding/gob"
	"github.com/andreyberezin/gin-site/models"
)

func InitModelsForSession() {
	gob.Register(models.Login{})
	gob.Register(models.Register{})
}
