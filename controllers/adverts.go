package controllers

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdvertGet handles GET /adverts/:id
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
	c.HTML(http.StatusOK, "advert/show", h)
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
