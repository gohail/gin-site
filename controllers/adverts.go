package controllers

import (
	"github.com/Sirupsen/logrus"
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/system"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdvertGet handles GET /adverts/show/:id
func AdvertGet(c *gin.Context) {
	db := models.GetBD()
	advert := models.Advert{}
	id := c.Param("id")
	db.First(&advert, id)
	if advert.ID == 0 {
		c.HTML(http.StatusNotFound, "errors/404", nil)
		return
	}
	h := DefaultH(c)
	h["Title"] = advert.Title
	h["Advert"] = advert
	c.HTML(http.StatusOK, "advert.html", h)
}

// AdvertGet handles GET /adverts/all
func GetAllAdverts(c *gin.Context) {
	db := models.GetBD()
	var adverts []models.Advert
	db.Find(&adverts)
	h := DefaultH(c)
	h["Title"] = "List of adverts"
	h["Adverts"] = adverts
	c.HTML(http.StatusOK, "adverts/index", h)
}

// GetNewAdvertTemplate handles GET /advert/add_new
func GetNewAdvertTemplate(c *gin.Context) {
	s := sessions.Default(c)
	h := DefaultH(c)
	h["Flash"] = s.Flashes()
	h["Title"] = "Добавить объявление"
	c.HTML(http.StatusOK, "create_advert.html", h)
}

// PostAddNewAdvert handles Post /advert/add_new
func PostAddNewAdvert(c *gin.Context) {
	h := DefaultH(c)
	s := sessions.Default(c)
	db := models.GetBD()
	uId := s.Get(USER_ID)
	adv := models.Advert{}
	if err := c.ShouldBind(&adv); err != nil {
		logrus.Trace(err)
		s.AddFlash("Заполните пустые поля")
		system.LogWarn(s.Save())
		c.Redirect(http.StatusFound, "/advert/add_new")
		return
	}
	//fmt.Printf("%+v\n", adv)
	adv.UserID = uId.(uint64)
	err := db.Create(&adv).Error
	if err != nil {
		logrus.Trace(err)
		s.AddFlash("Ошибка сервера")
		system.LogWarn(s.Save())
		c.Redirect(http.StatusFound, "/advert/add_new")
		return
	}
	h["Advert"] = adv
	c.HTML(http.StatusOK, "confirm_advert.html", h)
	//c.Redirect(http.StatusFound, fmt.Sprintf("/advert/prove/%d", adv.ID))
}

//
func GetAdvertProve(c *gin.Context) {
	h := DefaultH(c)
	s := sessions.Default(c)
	db := models.GetBD()
	adv := models.Advert{}
	system.LogErr(db.First(&adv, c.Param("id")).Error)
	h["Title"] = "подтверждение"
	if adv.ID == 0 {
		s.AddFlash("Ошибка сервера")
		system.LogErr(s.Save())
		c.Redirect(http.StatusFound, "/advert/add_new")
		return
	}
	if adv.UserID != s.Get(USER_ID) {
		s.AddFlash("Ошибка доступа")
		system.LogErr(s.Save())
		c.Redirect(http.StatusFound, "/advert/add_new")
		return
	}
	err := db.Model(&adv).Update("is_publish", true).Error
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("Advert was updated successfully")
	}
	c.Redirect(http.StatusFound, "/board")
}
