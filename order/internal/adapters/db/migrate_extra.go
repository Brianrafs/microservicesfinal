package db

func AutoMigrateExtra(db *DB) error {
    return db.DB.AutoMigrate(&InventoryItem{})
}
