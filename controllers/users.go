package controllers

import (
	"fmt"
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/system"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// handle GET /user/myinfo
func GetSelfInfo(c *gin.Context) {
	h := DefaultH(c)
	s := sessions.Default(c)
	db := models.GetBD()
	id := s.Get(USER_ID)
	u := models.User{}
	info := models.ExtraUserInfo{}
	db.First(&u, id).Related(&info, "InfoRefer")
	if info.ID == 0 {
		system.LogErr(db.Save(&info).Error)
		system.LogErr(db.Model(&u).Update("InfoRefer", uint(info.ID)).Error)
	}
	h["name"] = u.UserName
	h["info"] = info
	fmt.Println(info)
	c.HTML(http.StatusOK, "myinfo.html", h)
}

// handle POST /user/myinfo, expect ExtraUserInfo{} form
func PostSelfInfo(c *gin.Context) {
	db := models.GetBD()
	s := sessions.Default(c)
	id := s.Get(USER_ID)
	u := models.User{}
	data := models.ExtraUserInfo{}
	info := models.ExtraUserInfo{}
	if err := c.ShouldBind(&data); err != nil {
		s.AddFlash("Ошибка заполнения полей")
		c.Redirect(http.StatusFound, "/user/myinfo")
		return
	}
	db.First(&u, id).Related(&info, "InfoRefer")
	if info.ID == 0 {
		db.Save(&data)
		db.Model(&u).Update("InfoRefer", uint(data.ID))
	} else {
		db.Model(&info).Update(models.ExtraUserInfo{
			ContactEmail: data.ContactEmail,
			AboutMe:      data.AboutMe,
			PhoneNumber:  data.PhoneNumber,
			Photo:        data.Photo,
		})
	}
	c.Redirect(http.StatusFound, "/user/myinfo")
	return
}
