package models

type Inbox struct {
	Inbox_id          int32  `json:"inbox_id" gorm:"primaryKey;type:int"`
	User1ID           int32  `json:"user1_id" gorm:"not null;type:int;index"`                                                // owner inbox or receiver
	User1             *User  `json:"user1" gorm:"not null;foreignKey:User1ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // user_id
	User2ID           int32  `json:"user2_id" gorm:"not null;type:int;index"`                                                // user_id
	User2             *User  `json:"user2" gorm:"not null;foreignKey:User2ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // user_id
	Last_message      string `json:"last_message" gorm:"not null;type:text"`
	LastMessageUserID int32  `json:"last_message_user_id" gorm:"not null;type:int"`
	LastTimeMessage   int64  `json:"last_time_message" gorm:"not null;type:bigint;autoCreateTime"`
}

type ChatRoom struct {
	Inbox_id      int32           `json:"inbox_id" gorm:"primaryKey;autoIncrement:false;type:int"`                                   // owner inbox or receiver                                                    // user_id
	Conversations []*Conversation `json:"conversations" gorm:"foreignKey:ChatRoom_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // message
}

type Conversation struct {
	Conversation_id int32  `json:"conversation_id" gorm:"primaryKey;type:int"`
	ChatRoom_id     int32  `json:"chatroom_id" gorm:"type:int"`
	Sender_id       int32  `json:"sender_id" gorm:"not null;type:int"` // user_id
	Message         string `json:"message" gorm:"type:text"`           // message
	CreatedAt       int64  `json:"created_at" gorm:"autoCreateTime"`
}
