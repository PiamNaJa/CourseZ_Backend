package models

type Reward_Info struct {
	Reward_id    int32        `json:"reward_id" gorm:"primaryKey;type:int"`                                                 // Reward_id is the id of the reward
	UserID       int32        `json:"user_id" gorm:"index;type:int;not null" validate:"required,number"`                    // UserID is the id of the user
	User         *User        `json:"user" gorm:"not null;foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // User is the user that made the reward
	House_no     string       `json:"house_no" gorm:"not null;type:varchar(10)" validate:"required,max=10"`                 // House_no is the house number of the address
	ItemID       int32        `json:"item_id" gorm:"index;type:int;not null" validate:"required,number"`                    // ItemID is the id of the item
	Item         *Reward_Item `json:"item" gorm:"not null;foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Lane         string       `json:"lane" gorm:"not null;type:varchar(10)" validate:"required,max=10"`         // Lane is the lane of the address
	Village_no   string       `json:"village_no" gorm:"not null;type:varchar(20)" validate:"required,max=20"`   // Village_no is the village number of the address
	Village      string       `json:"village" gorm:"not null;type:varchar(50)" validate:"required,max=50"`      // Village is the village of the address
	Road         string       `json:"road" gorm:"not null;type:varchar(50)" validate:"required,max=50"`         // Road is the road of the address
	Sub_district string       `json:"sub_district" gorm:"not null;type:varchar(60)" validate:"required,max=60"` // Sub_district is the sub_district of the address
	District     string       `json:"district" gorm:"not null;type:varchar(60)" validate:"required,max=60"`     // District is the district of the address
	Province     string       `json:"province" gorm:"not null;type:varchar(60)" validate:"required,max=60"`     // Province is the province of the address
	Postal       int32        `json:"postal" gorm:"not null;type:int" validate:"required,number"`               // Postal is the postal code of the address
}
