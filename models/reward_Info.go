package models

type Reward_Info struct {
	Reward_id    int32        `json:"reward_id" gorm:"primaryKey;type:int"`
	UserID       int32        `json:"user_id" gorm:"index;type:int"`
	User         *User        `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	House_no     int32        `json:"house_no" gorm:"not null;type:int"`
	ItemID       int32        `json:"item_id" gorm:"index;type:int"`
	Item         *Reward_Item `json:"item" gorm:"not null;foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Lane         string       `json:"lane" gorm:"not null;type:varchar(10)"`
	Village_no   string       `json:"village_no" gorm:"not null;type:varchar(20)"`
	Village      string       `json:"village" gorm:"not null;type:varchar(50)"`
	Road         string       `json:"road" gorm:"not null;type:varchar(50)"`
	Sub_district string       `json:"sub_district" gorm:"not null;type:varchar(60)"`
	District     string       `json:"district" gorm:"not null;type:varchar(60)"`
	Province     string       `json:"province" gorm:"not null;type:varchar(60)"`
	Postal       int32        `json:"postal" gorm:"not null;type:int"`
}
