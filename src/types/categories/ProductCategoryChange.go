package products

import "time"

// ProductCategoryChange :
type ProductCategoryChange struct {
	ID            string
	ChangerUserID string
	CategoryID    string
	Items         string
	Date          time.Time
}
