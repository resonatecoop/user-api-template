package legacy_wp_model

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

// WpUser ...
type WpUser struct {
	bun.BaseModel `bun:"rsnrt_users,alias:u"`
	ID            uint64         `gorm:"column:ID;primary_key;auto_increment;not_null"`
	Email         string         `gorm:"column:user_email" sql:"type:varchar(100);unique;not null"`
	Login         string         `gorm:"column:user_login" sql:"type:varchar(60)"`
	Registered    time.Time      `gorm:"column:user_registered" sql:"type:datetime"`
	Nicename      string         `gorm:"column:user_nicename" sql:"type:varchar(50)"`
	DisplayName   string         `gorm:"column:display_name" sql:"type:varchar(250)"`
	Password      sql.NullString `gorm:"column:user_pass" sql:"type:varchar(255)"`
}

// UserMeta ...
type WpUserMeta struct {
	bun.BaseModel `bun:"rsntr_usermeta,alias:u"`
	ID            uint64 `gorm:"column:umeta_id;primary_key;auto_increment;not_null"`
	UserId        uint64 `gorm:"column:user_id" sql:"not null"`
	MetaKey       string `gorm:"column:meta_key" sql:"type:varchar(255)"`
	MetaValue     string `gorm:"column:meta_value" sql:"type:longtext"`
}
