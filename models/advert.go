package models

type Advert struct {
	Model
	UserID  int    `gorm:"index"`
	Title   string `gorm:"type:varchar(100);unique_index"`
	Content string
	Price   int
	Img     string
}
