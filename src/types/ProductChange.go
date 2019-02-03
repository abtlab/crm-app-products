package products

import "time"

// ProductChange :
type ProductChange struct {
	ID            string
	ChangerUserID string
	ProductID     string
	Items         string
	Date          time.Time
}
