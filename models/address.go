package models

type Address struct {
	Address_id   int32  `json:"address_id" gorm:"primaryKey;type:int"`                                    // Address_id is the id of the address
	UserID       int32  `json:"user_id" gorm:"not null;type:int" validate:"required,number"`              // UserID is the id of the user
	House_no     string `json:"house_no" gorm:"not null;type:varchar(10)" validate:"required,max=10"`     // House_no is the house number of the address
	Lane         string `json:"lane" gorm:"not null;type:varchar(10)" validate:"max=10"`                  //ซอย
	Village_no   string `json:"village_no" gorm:"not null;type:varchar(20)" validate:"max=20"`            // หมู่
	Village      string `json:"village" gorm:"not null;type:varchar(50)" validate:"max=50"`               // หมู่บ้าน
	Road         string `json:"road" gorm:"not null;type:varchar(50)" validate:"max=50"`                  // Road is the road of the address
	Sub_district string `json:"sub_district" gorm:"not null;type:varchar(60)" validate:"required,max=60"` // ตำยล
	District     string `json:"district" gorm:"not null;type:varchar(60)" validate:"required,max=60"`     // อำเภอ
	Province     string `json:"province" gorm:"not null;type:varchar(60)" validate:"required,max=60"`     // จังหวัด
	Postal       int32  `json:"postal" gorm:"not null;type:int" validate:"required,number"`               // Postal is the postal code of the address
}
