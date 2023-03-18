package models

type Reward_Info struct {
	Reward_id int32        `json:"reward_id" gorm:"primaryKey;type:int"`                                                       // Reward_id is the id of the reward
	UserID    int32        `json:"user_id" gorm:"index;type:int;not null" validate:"required,number"`                          // UserID is the id of the user
	User      *User        `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`       // User is the user that made the reward
	ItemID    int32        `json:"item_id" gorm:"index;type:int;not null" validate:"required,number"`                          // ItemID is the id of the item
	Item      *Reward_Item `json:"item" gorm:"not null;foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`       // Item is the item that made the reward
}
