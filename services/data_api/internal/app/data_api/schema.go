package dataapi

import "time"

type supplier struct {
	tableName struct{} `sql:"suppliers,alias:s"`

	ID    string `sql:"id,pk"`
	Name  string `sql:"name"`
	Phone string `sql:"phone"`
}

type order struct {
	tableName struct{} `sql:"orders,alias:o"`

	ID          string    `sql:"id,pk"`
	Description string    `sql:"description"`
	PickupAt    time.Time `sql:"pickup_at"`
	Lockation   string    `sql:"lockation"`
	SupplierID  string    `sql:"supplier_id"`
	Status      string    `sql:"status"`
	ImageURL    string    `sql:"image_url"`
	UpdatedAt   time.Time `sql:"updated_at"`
	CreatedAt   time.Time `sql:"created_at"`

	Supplier *Supplier `pg:"fk:supplier_id,rel:has-one"`
}
