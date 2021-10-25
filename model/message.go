package model

const (
	textMsg =iota
)

type Message struct {
	ID string `json:"id" gorm:"primaryKey"`
	//消息类型
    Type string
	//Payload interface{}
	//消息所属的会话 ID
	ConversationID string
	//消息所属会话的类型
	ConversationType string
	To string
	//发送方的 userID，在消息发送时，会默认设置为当前登录的用户
	From string
	//in 为收到的消息 out 为发出的消息
	Flow string
	//消息发送者的昵称
	Nick string
	//头像
	Avatar string
	//C2C 消息对端是否已读，true 标识对端已读
	IsPeerRead bool
	//是否被撤回的消息，true 标识被撤回的消息
	IsRevoked bool
}
