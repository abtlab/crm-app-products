package products

import "time"

// ProductAttributeChange :
type ProductAttributeChange struct {
	ID            string
	ChangerUserID string
	AttributeID   string
	Items         string
	Date          time.Time
}
