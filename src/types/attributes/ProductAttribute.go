package products

import "github.com/abtlab/crm-app-utils/types"

// ProductAttribute :
type ProductAttribute struct {
	ID       string
	AcountID string
	Type     common.AttributeType
	Key      string
}
