package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title string
	Price uint
}

type ProductData struct {
	gorm.Model
	Name  string
	Value string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	migrate(db)
	add_records(db)
}

func migrate(db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&Product{}, &ProductData{})
}

func add_records(db *gorm.DB) {
	// Create
	db.Create(&Product{Title: "Laptop", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                     // find product with integer primary key
	db.First(&product, "title = ?", "Laptop") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Title: "Tablet"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Title": "Ultrabook"})

	// Delete - delete product
	db.Delete(&product, 1)
}
