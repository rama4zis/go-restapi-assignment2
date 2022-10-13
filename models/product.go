package models

type Product struct {
	Code        string `json:"item_code" gorm:"primary_key"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
}
