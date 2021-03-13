package access

// New instantiates new Access config service
func New(NoTokenMethods string, PublicUserMethods string) *AccessConfig {
	return &AccessConfig{
		NoTokenMethods:    NoTokenMethods,
		PublicUserMethods: PublicUserMethods,
	}
}

// Access contains the configuration that governs service access
type AccessConfig struct {
	// Secret key used for signing.
	NoTokenMethods    string
	PublicUserMethods string
}
