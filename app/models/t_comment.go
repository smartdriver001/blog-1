package models

import (
	"blog/app/support"
	"errors"
	"time"
)

//Comment model.
type Comment struct {
	Id         int64     `xorm:"not null pk autoincr INT(11)"`
	Content    string    `xorm:"TEXT"`
	Auther     int64     `xorm:"INT(11)"`
	BlogId     int64     `xorm:"INT(11)"`
	CommentId  int64     `xorm:"INT(11)"`
	CreateTime time.Time `xorm:"DATETIME"`
	Status     int       `xorm:"INT(1)"`
}

//New comment.
func (c *Comment) NewComment() error {

	if c.Content == "" || c.Auther == 0 || c.BlogId == 0 {
		return errors.New("Content, auther or blog id can't be null.")
	}

	model := new(Comment)
	model.Content = c.Content
	model.Auther = c.Auther
	model.BlogId = c.BlogId

	if c.CommentId != 0 {
		model.CommentId = c.CommentId
	}

	model.CreateTime = time.Now()
	model.Status = c.Status
	has, err := support.Xorm.InsertOne(&model)

	if has <= 0 {
		return err
	}

	return nil
}
