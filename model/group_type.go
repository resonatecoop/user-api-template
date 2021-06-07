package model

// GroupType provides the key of the Group
type GroupType int32

const (
	// Persona is a representation of a User.  A User can have multiple Personas
	Persona GroupType = iota + 1 // 1

	// Band contains one or more Personas
	Band // 2

	// Label can have one or more bands
	Label // 3

	// Distributor can have one or more Labels
	Distributor // 4
)
