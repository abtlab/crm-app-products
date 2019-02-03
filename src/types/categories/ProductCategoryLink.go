package products

import "time"

// ProductCategoryLink :
type ProductCategoryLink struct {
	ID         string
	ProductID  string
	CategoryID string
	CreateDate time.Time
}
