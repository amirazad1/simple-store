// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Category struct {
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	ID        int32     `json:"id"`
}

type Factor struct {
	CustomerName   sql.NullString `json:"customer_name"`
	CustomerMobile sql.NullString `json:"customer_mobile"`
	Seller         string         `json:"seller"`
	CreatedAt      time.Time      `json:"created_at"`
	ID             int64          `json:"id"`
}

type FactorDetail struct {
	FactorID  int64 `json:"factor_id"`
	ProductID int64 `json:"product_id"`
	SaleCount int32 `json:"sale_count"`
	SalePrice int64 `json:"sale_price"`
	ID        int64 `json:"id"`
}

type Product struct {
	Name          string         `json:"name"`
	Brand         sql.NullString `json:"brand"`
	Model         sql.NullString `json:"model"`
	InitNumber    int32          `json:"init_number"`
	PresentNumber int32          `json:"present_number"`
	BuyPrice      int64          `json:"buy_price"`
	BuyDate       time.Time      `json:"buy_date"`
	SalePrice     sql.NullInt64  `json:"sale_price"`
	CreatedAt     time.Time      `json:"created_at"`
	ID            int64          `json:"id"`
}

type ProductCategory struct {
	ProductID  int64     `json:"product_id"`
	CategoryID int32     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	ID         int64     `json:"id"`
}

type User struct {
	Username          string       `json:"username"`
	Password          string       `json:"password"`
	FullName          string       `json:"full_name"`
	Mobile            string       `json:"mobile"`
	PasswordChangedAt sql.NullTime `json:"password_changed_at"`
	CreatedAt         time.Time    `json:"created_at"`
}
