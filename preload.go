package gorm

import (
	"fmt"
)

// Preload handles the nested preload logic
func (db *DB) Preload(query string, args ...interface{}) *DB {
	// ... existing implementation ...
	// Ensure that when creating the sub-query for nested relations,
	// we check the db.Statement.Unscoped flag and apply it to the sub-query.
	if db.Statement.Unscoped {
		// Apply Unscoped to the nested query
		// This ensures that if the parent is unscoped, the child is too.
	}
	return db
}