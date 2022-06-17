package model

import (
	"database/sql"
	"time"

	uuid "github.com/google/uuid"
)

// AuthUser represents data stored in session/context for a user
type AuthUser struct {
	ID       uuid.UUID
	TenantID int32
	Username string
	Email    string
	Role     AccessRole
}

// User basic definition of a User and its meta
type User struct {
	IDRecord
	Username               string `bun:",notnull,unique"`
	FullName               string
	FirstName              string
	LastName               string
	EmailConfirmed         bool   `bun:"default:false"`
	Country                string `bun:"type:varchar(2)"`
	Member                 bool   `bun:"default:false,notnull"`
	NewsletterNotification bool
	FollowedGroups         []uuid.UUID  `bun:",type:uuid[],array"`
	OwnerOfGroups          []*UserGroup `bun:"rel:has-many"`
	TenantID               int32
	RoleID                 int32
	LastLogin              time.Time
	LastPasswordChange     time.Time
	Password               sql.NullString `bun:"type:varchar(60)"`
	Token                  string
	//	Email                  string `bun:",unique,notnull"`
	// FavoriteTracks []uuid.UUID `bun:",type:uuid[]" pg:",array"`
	// Playlists      []uuid.UUID `bun:",type:uuid[]" pg:",array"`
	// Plays []Track `pg:"many2many:plays"` Payment API
}

// UpdateLoginDetails updates login related fields
func (u *User) UpdateLoginDetails(token string) {
	u.Token = token
	t := time.Now()
	u.LastLogin = t
}

//DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Empty, error)

// Delete deletes a provided User
// func (u *User) Delete(db *bun.DB) (error, string) {

// 	// ctx context.Context
// 	pgerr := tx.Model(u).
// 		Column("user.followed_groups", "OwnerOfGroups").
// 		WherePK().
// 		Select()
// 	if pgerr != nil {
// 		return pgerr, "user"
// 	}

// 	//		Column("user.favorite_tracks", "user.followed_groups", "user.playlists", "OwnerOfGroups").
// 	// if len(u.FavoriteTracks) > 0 {
// 	// 	_, pgerr = tx.Exec(`
// 	//     UPDATE tracks
// 	//     SET favorite_of_users = array_remove(favorite_of_users, ?)
// 	//     WHERE id IN (?)
// 	//   `, u.Id, pg.In(u.FavoriteTracks))
// 	// 	if pgerr != nil {
// 	// 		return pgerr, "track"
// 	// 	}
// 	// }

// 	if len(u.FollowedGroups) > 0 {
// 		_, pgerr = tx.Exec(`
//       UPDATE user_groups
//       SET followers = array_remove(followers, ?)
//       WHERE id IN (?)
//     `, u.ID, pg.In(u.FollowedGroups))
// 		if pgerr != nil {
// 			return pgerr, "user_group"
// 		}
// 	}

// 	if len(u.OwnerOfGroups) > 0 {
// 		for _, group := range u.OwnerOfGroups {
// 			if pgerr, table := group.Delete(tx); pgerr != nil {
// 				return pgerr, table
// 			}
// 		}
// 	}

// 	// if len(u.Playlists) > 0 {
// 	// 	for _, playlistId := range u.Playlists {
// 	// 		playlist := &TrackGroup{Id: playlistId}
// 	// 		if pgerr, table := playlist.Delete(tx); pgerr != nil {
// 	// 			return pgerr, table
// 	// 		}
// 	// 	}
// 	// }

// 	pgerr = tx.Delete(u)
// 	if pgerr != nil {
// 		return pgerr, "user"
// 	}

// 	return nil, ""
// }

// FollowGroup causes a User to follow a UserGroup
// func (u *User) FollowGroup(db *pg.DB, userGroupID uuid.UUID) (error, string) {
// 	var table string
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err, table
// 	}
// 	defer tx.Rollback()

// 	// Add userGroupID to user FollowedGroups
// 	userGroupIDArr := []uuid.UUID{userGroupID}
// 	_, pgerr := tx.ExecOne(`
//     UPDATE users
//     SET followed_groups = (select array_agg(distinct e) from unnest(followed_groups || ?) e)
//     WHERE id = ?
//   `, pg.Array(userGroupIDArr), u.ID)
// 	if pgerr != nil {
// 		table = "user"
// 		return pgerr, table
// 	}

// 	// Add userID to userGroup Followers
// 	userIDArr := []uuid.UUID{u.ID}
// 	_, pgerr = tx.ExecOne(`
//     UPDATE user_groups
//     SET followers = (select array_agg(distinct e) from unnest(followers || ?) e)
//     WHERE id = ?
//   `, pg.Array(userIDArr), userGroupID)
// 	if pgerr != nil {
// 		table = "user_group"
// 		return pgerr, table
// 	}
// 	return tx.Commit(), table
// }

// // UnfollowGroup removes the follow state of the supplied User from the supplied userGroup via the supplied userGroupID
// func (u *User) UnfollowGroup(db *pg.DB, userGroupID uuid.UUID) (error, string) {
// 	var table string
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err, table
// 	}
// 	// Rollback tx on error.
// 	defer tx.Rollback()

// 	// Remove userGroupId from user FollowedGroups
// 	_, pgerr := tx.ExecOne(`
//     UPDATE users
//     SET followed_groups = array_remove(followed_groups, ?)
//     WHERE id = ?
//   `, userGroupID, u.ID)
// 	if pgerr != nil {
// 		table = "user"
// 		return pgerr, table
// 	}

// 	// Remove userId from track FavoriteOfUsers
// 	_, pgerr = tx.ExecOne(`
//     UPDATE user_groups
//     SET followers = array_remove(followers, ?)
//     WHERE id = ?
//   `, u.ID, userGroupID)
// 	if pgerr != nil {
// 		table = "user_group"
// 		return pgerr, table
// 	}
// 	return tx.Commit(), table
// }

// func (u *User) RemoveFavoriteTrack(db *pg.DB, trackId uuid.UUID) (error, string) {
// 	var table string
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err, table
// 	}
// 	// Rollback tx on error.
// 	defer tx.Rollback()

// 	// Remove trackId from user FavoriteTracks
// 	_, pgerr := tx.ExecOne(`
//     UPDATE users
//     SET favorite_tracks = array_remove(favorite_tracks, ?)
//     WHERE id = ?
//   `, trackId, u.ID)
// 	if pgerr != nil {
// 		table = "user"
// 		return pgerr, table
// 	}

// 	// Remove userId from track FavoriteOfUsers
// 	_, pgerr = tx.ExecOne(`
//     UPDATE tracks
//     SET favorite_of_users = array_remove(favorite_of_users, ?)
//     WHERE id = ?
//   `, u.Id, trackId)
// 	if pgerr != nil {
// 		table = "track"
// 		return pgerr, table
// 	}
// 	return tx.Commit(), table
// }

// func (u *User) AddFavoriteTrack(db *pg.DB, trackId uuid.UUID) (error, string) {
// 	var table string
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err, table
// 	}
// 	// Rollback tx on error.
// 	defer tx.Rollback()

// 	// Add trackId to user FavoriteTracks
// 	trackIdArr := []uuid.UUID{trackId}
// 	_, pgerr := tx.ExecOne(`
//     UPDATE users
//     SET favorite_tracks = (select array_agg(distinct e) from unnest(favorite_tracks || ?) e)
//     WHERE id = ?
//   `, pg.Array(trackIdArr), u.Id)
// 	// WHERE NOT favorite_tracks @> ?
// 	if pgerr != nil {
// 		table = "user"
// 		return pgerr, table
// 	}

// 	// Add userId to track FavoriteOfUsers
// 	userIdArr := []uuid.UUID{u.Id}
// 	_, pgerr = tx.ExecOne(`
//     UPDATE tracks
//     SET favorite_of_users = (select array_agg(distinct e) from unnest(favorite_of_users || ?) e)
//     WHERE id = ?
//   `, pg.Array(userIdArr), trackId)
// 	if pgerr != nil {
// 		table = "track"
// 		return pgerr, table
// 	}
// 	return tx.Commit(), table
// }
