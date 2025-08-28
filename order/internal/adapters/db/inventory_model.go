package db

import "gorm.io/gorm"

type InventoryItem struct {
    gorm.Model
    ProductCode string `gorm:"uniqueIndex;size:128"`
    Name        string `gorm:"size:255"`
}
