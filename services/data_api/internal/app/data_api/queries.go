package dataapi

import "context"

func upsertOrder(ctx context.Context, o *order) error {
	_, err := db.Model(o).
		OnConflict("(id) DO UPDATE").
		Set("status = EXCLUDED.status = status").
		Set("updated_at = NOW()").
		Insert(ctx)
	return err
}
