package access

// New instantiates new Access config service
func New(NoTokenMethods string, PublicMethods string, ReadWriteMethods string) *AccessConfig {
	return &AccessConfig{
		NoTokenMethods: NoTokenMethods,
		PublicMethods:  PublicMethods,
		WriteMethods:   ReadWriteMethods,
	}
}

// Access contains the configuration that governs service access
type AccessConfig struct {
	// Secret key used for signing.
	NoTokenMethods string
	PublicMethods  string
	WriteMethods   string
}
