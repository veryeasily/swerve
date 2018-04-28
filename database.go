package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "fmt"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

var loaded = false

func CreateDb(foo string) {
  db, err := gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()

  // Migrate the schema
  if !loaded {
    fmt.Printf("Migrating DB...\n")
    db.AutoMigrate(&Product{})
    loaded = true
  }

  // Create
  db.Create(&Product{Code: foo, Price: 1000})

  // Read
  var product Product
  db.First(&product, 1) // find product with id 1
  db.First(&product, "code = ?", foo) // find product with code foo

  // Update - update product's price to 2000
  db.Model(&product).Update("Price", 2000)

  // Delete - delete product
  // db.Delete(&product)
}
