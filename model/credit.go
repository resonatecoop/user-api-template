package model

import (
	uuid "github.com/google/uuid"
)

// Credit
type Credit struct {
	IDRecord
	UserID uuid.UUID `bun:"type:uuid,notnull"`
	User   *User     `bun:"rel:has-one"`
	Total  int64     `bun:",notnull,default:128"`
}

const euro = float64(1.25)
const credit = int64(1022)

// ConvertCreditToEuro converts credit value to euro
func ConvertCreditToEuro(total int64) float64 {
	return float64(total) / float64(credit) * euro
}

// CalculateCost
func CalculateCost(count int64) int64 {
	cost := int64(0)
	n := int64(2)

	if count > 8 {
		return cost
	}

	for n < count {
		cost *= 2
		n++
	}

	return cost
}

// CalculateRemainingCost
func CalculateRemainingCost(count int64) int64 {
	cost := int64(0)
	n := int64(2)

	if count > 8 {
		return cost
	}

	for n < count {
		cost += CalculateCost(n)
		n++
	}

	return 1022 - cost
}
