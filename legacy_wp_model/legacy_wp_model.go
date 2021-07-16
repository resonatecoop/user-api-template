package legacy_wp_model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

// WpUser ...
type WpUser struct {
	bun.BaseModel `bun:"rsntr_users,alias:u"`
	ID            uint64         `bun:"ID,pk,type:serial,notnull"`
	Email         string         `bun:"user_email,type:varchar(100);unique;notnull"`
	Login         string         `bun:"user_login,type:varchar(60)"`
	Registered    time.Time      `bun:"user_registered,type:datetime"`
	Nicename      string         `bun:"user_nicename,type:varchar(50)"`
	DisplayName   string         `bun:"display_name,type:varchar(250)"`
	Password      sql.NullString `bun:"user_pass,type:varchar(255)"`
}

// UserMeta ...
type WpUserMeta struct {
	bun.BaseModel `bun:"rsntr_usermeta,alias:um"`
	ID            uint64 `bun:"umeta_id,pk,auto_increment,notnull"`
	UserId        uint64 `bun:"user_id,notnull"`
	MetaKey       string `bun:"meta_key,type:varchar(255)"`
	MetaValue     string `bun:"meta_value,type:longtext"`
}
