package model

// MembershipClass
type MembershipClass struct {
	IDRecord
	Name      string `bun:",notnull"`
	ProductID string `bun:",notnull"`
	PriceID   string `bun:",notnull"`
}
