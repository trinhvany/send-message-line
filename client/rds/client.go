package rds

import (
	"small_demo_go/model"
	"small_demo_go/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := system.Account + ":" + system.Password + "@tcp(" + system.IP + ":" + system.Port + ")/" + system.Schema +"?parseTime=True&loc=Local"
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