package models

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID			uint		`json:"orderId"`
	ProductID		uint		`json:"productId"`
	Product			Product		`json:"product"`
	Quantity		int			`gorm:"not null" json:"quantity"`
	Price			float64		`gorm:"not null" json:"price"`
}

type Order struct {
	gorm.Model
	UserID			uint		`json:"userId"`
	User			User		`json:"user"`
	Items			[]OrderItem	`gorm:"foreignKey:OrderID" json:"items"`
	TotalPrice		float64		`json:"totalPrice"`
	Status			string 		`json:"status"`
}