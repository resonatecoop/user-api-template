package legacy_model

import (
	"github.com/uptrace/bun"
)

// Credit ...
type Credit struct {
	bun.BaseModel `bun:"credits,alias:c"`
	ID            int `bun:"cid,pk,auto_increment,notnull"`
	UserId        int `bun:"uid"`
	Total         int `bun:"total"`
}

// Play ...
type Play struct {
	bun.BaseModel `bun:"plays,alias:p"`
	ID            int `bun:"pid,pk,auto_increment,notnull"`
	UserId        int `bun:"uid"`
	TrackId       int `bun:"tid"`
	Status        int `bun:"event"`
	Date          int `bun:"date"`
}

// Track
type Track struct {
	bun.BaseModel `bun:"tracks,alias:t"`
	ID            int    `bun:"tid,pk,auto_increment,notnull"`
	UserId        int    `bun:"uid"`
	Status        int    `bun:"status"`
	Date          int    `bun:"date"`
	Artist        string `bun:"track_artist"`
	Name          string `bun:"track_name"`
	Album         string `bun:"track_album"`
	AlbumArtist   string `bun:"track_album_artist"`
	Duration      string `bun:"track_duration"`
	Composer      string `bun:"track_composer"`
	Year          string `bun:"track_year"`
	AudioFile     string `bun:"track_url"`
	Cover         string `bun:"track_cover_art"`
	Number        string `bun:"track_number"`
}
