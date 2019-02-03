package products

import "time"

// ProductAttributeLink :
type ProductAttributeLink struct {
	ID          string
	ProductID   string
	AttributeID string
	Value       string
	CreateDate  time.Time
}
