package rds

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"small_demo_go/model"
)

type Client interface {
	GetMessage() ([]*model.Message, error)
	GetMessageDetail(MessageID int64) ([]*model.MessageDetail, error)
	Close() error
}

type client struct {
	db		*gorm.DB
}

func NewClient() (Client, error) {
	dsn := "dialog:password123!@tcp(192.168.68.164:3306)/send_message?parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db = db.Debug()

	return &client{
		db: db,
	}, nil
}

func (c *client) Close() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}