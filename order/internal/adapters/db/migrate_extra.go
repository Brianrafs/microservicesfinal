package db

func AutoMigrateExtra(adapter *Adapter) error {
	return adapter.db.AutoMigrate(&InventoryItem{})
}
