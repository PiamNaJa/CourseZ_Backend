package models

type Reward_Item struct {
	Item_id      int32  `json:"item_id" gorm:"primaryKey;type:int"`                                         // Item_id is the id of the item
	Item_name    string `json:"item_name" gorm:"not null;type:varchar(50)" validate:"required,max=50"`      // Item_name is the name of the item
	Item_title   string `json:"item_title" gorm:"not null;type:varchar(50)" validate:"required,max=50"`     // Item_title is the title of the item
	Item_cost    int32  `json:"item_cost" gorm:"not null;type:int" validate:"required,number"`              // Item_cost is the cost of the item
	Item_picture string `json:"item_picture" gorm:"not null;type:varchar(255)" validate:"required,max=255"` // Item_picture is a link to the picture
}
