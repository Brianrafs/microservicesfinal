package db

import (
	"context"
)

type InventoryRepository struct {
	adapter *Adapter
}

func NewInventoryRepository(adapter *Adapter) *InventoryRepository {
	return &InventoryRepository{adapter: adapter}
}

func (r *InventoryRepository) Exists(ctx context.Context, productCode string) (bool, error) {
	var count int64
	if err := r.adapter.db.WithContext(ctx).Model(&InventoryItem{}).
		Where("product_code = ?", productCode).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
