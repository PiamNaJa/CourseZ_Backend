package models

type Reward_Item struct {
	Item_id      int32  `json:"item_id" gorm:"primaryKey;type:int"`
	Item_name    string `json:"item_name" gorm:"not null;type:varchar(50)"`
	Item_title   string `json:"item_title" gorm:"not null;type:varchar(50)"`
	Item_cost    int32  `json:"item_cost" gorm:"not null;type:int"`
	Item_picture string `json:"item_picture" gorm:"not null;type:varchar(255)"`
}
