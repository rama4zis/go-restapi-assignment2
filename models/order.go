package models

type Order struct {
	Id           uint      `json:"order_id" gorm:"primary_key"`
	OrderedAt    string    `json:"ordered_at"`
	CustomerName string    `json:"customer_name"`
	Items        []Product `json:"items" gorm:"many2many:order_items;constraint:OnDelete:CASCADE;"`
}
