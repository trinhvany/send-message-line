package model

type MessageDetail struct {
	ID						int64		`gorm:"primary_key;column:message_detail_id"`
	MessageID				int64		`gorm:"column:message_id"`
	ContentJson				string		`gorm:"column:content_json"`
	ContentType				int64		`gorm:"column:content_type"`
	Seq						int64		`gorm:"column:seq"`
	DeleteFlg				string		`gorm:"column:delete_flg"`
}

func (MessageDetail) TableName() string {
	return "send_message.message_detail"
}