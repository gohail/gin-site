package models

type Advert struct {
	Model
	UserID    uint64
	Title     string `gorm:"type:varchar(100)" form:"title" binding:"required"`
	Content   string `form:"content" binding:"required"`
	Price     string `form:"price" binding:"required"`
	Img       string
	IsPublish bool
}
