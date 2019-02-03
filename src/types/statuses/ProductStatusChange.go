package products

import "time"

// ProductStatusChange :
type ProductStatusChange struct {
	ID            string
	ChangerUserID string
	StatusID      string
	Items         string
	Date          time.Time
}
