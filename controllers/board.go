package controllers

import (
	"fmt"
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainBoard(c *gin.Context) {
	h := DefaultH(c)
	db := models.GetBD()
	var adverts []models.Advert
	system.LogInfo(db.Order("created_at").Where("is_publish = ?", true).Find(&adverts).Error)
	h["advertsList"] = adverts
	h["count"] = countAdverts(len(adverts))
	h["Title"] = "Доска объявлений"
	c.HTML(http.StatusOK, "board.html", h)
}

func countAdverts(num int) string {
	switch num {
	case 0:
		return "Обявления не найдены"
	default:
		return fmt.Sprintf("Найдено %d обявлений", num)
	}
}
