package db

import "context"

type InventoryRepository struct{ db *DB }

func NewInventoryRepository(db *DB) *InventoryRepository { return &InventoryRepository{db: db} }

func (r *InventoryRepository) Exists(ctx context.Context, productCode string) (bool, error) {
    var count int64
    if err := r.db.DB.WithContext(ctx).Model(&InventoryItem{}).Where("product_code = ?", productCode).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}
