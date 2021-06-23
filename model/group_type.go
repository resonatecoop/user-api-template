package model

// GroupType provides the key of the Group
// type GroupID int32

// const (
// 	// Persona is a representation of a User.  A User can have multiple Personas
// 	Persona GroupID = iota + 1 // 1

// 	// Band contains one or more Personas
// 	Band // 2

// 	// Label can have one or more bands
// 	Label // 3

// 	// Distributor can have one or more Labels
// 	Distributor // 4
// )

type GroupType struct {
	IDRecord
	Name        string `bun:",notnull"`
	Description string
}
