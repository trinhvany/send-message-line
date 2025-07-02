package system

const (
	DeleteFlagNotDeleted = "0"
	DeleteFlagDeleted    = "1"
)

const (
	StatusProcessing 		= "1"
)

const (
	LinePushURL = "https://api.line.me/v2/bot/message/push"
	LinePushTemplate = `{
		"to": "<<TO>>",
		"messages": <<MESSAGES>>
	}`
	UserIDFake = "" // "UserID1, UserID2, ..."
	Secret_key = "" // Channel secret key
)

const (
	Account = ""
	Password = ""
	IP = ""
	Port = ""
	Schema = "send_message"
)