package products

import (
	"github.com/shopspring/decimal"
	"time"
)

// Product :
type Product struct {
	ID              string
	AcountID        string
	ParentProductID string
	StatusID        string
	Name            string
	VendorCode      string
	Price           decimal.Decimal
	Image           []byte
	IsHidden        bool
	CreateDate      time.Time
	Version         int64
}
