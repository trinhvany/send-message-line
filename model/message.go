package model

type Message struct {
	ID						int64		`gorm:"primary_key;column:message_id"`
	MessageTitle			string		`gorm:"column:message_title"`
	Status					string		`gorm:"column:status"`
	DeleteFlg				string		`gorm:"column:delete_flg"`
}

func (Message) TableName() string {
	return "send_message.messages"
}