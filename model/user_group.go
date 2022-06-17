package model

import (

	// "log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	// trackpb "user-api/rpc/track"
	//tagpb "github.com/resonatecoop/user-api/proto/api"

	pbUser "github.com/resonatecoop/user-api-template/proto/user"

	uuid "github.com/google/uuid"
)

// UserGroup represents a group of Users and maintains a set of metadata
type UserGroup struct {
	IDRecord
	DisplayName    string `bun:",unique,notnull"`
	Description    string
	ShortBio       string
	GroupEmail     string
	AddressID      uuid.UUID `bun:"type:uuid,notnull"` //for Country see User model
	Address        *StreetAddress
	TypeID         uuid.UUID `bun:"type:uuid,notnull"` //for e.g. Persona Type
	Type           *GroupType
	OwnerID        uuid.UUID   `bun:"type:uuid,notnull"`
	Owner          *User       `bun:"rel:has-one"`
	Links          []uuid.UUID `bun:",type:uuid[],array"`
	Members        []UserGroup `pg:"many2many:user_group_members,fk:user_group_id,joinFK:member_id"`
	MemberOfGroups []UserGroup `pg:"many2many:user_group_members,fk:member_id,joinFK:user_group_id"`
	Avatar         uuid.UUID   `bun:"type:uuid"`
	Banner         uuid.UUID   `bun:"type:uuid"`
	Tags           []uuid.UUID `bun:",type:uuid[],array"`
	// AdminUsers         []uuid.UUID `bun:",type:uuid[]" pg:",array"`
	// Followers          []uuid.UUID `bun:",type:uuid[]" pg:",array"`
	// RecommendedArtists []uuid.UUID `bun:",type:uuid[]" pg:",array"`
	// RecommendedBy      []uuid.UUID `bun:",type:uuid[]" pg:",array"`
	// PrivacyID          uuid.UUID   `bun:"type:uuid,notnull"`
	// Privacy            *UserGroupPrivacy
	// Kvstore            map[string]string `pg:",hstore"`
	// Publisher          map[string]string `pg:",hstore"`
	// Pro                map[string]string `pg:",hstore"`
}

// Address            *StreetAddress
// Type               GroupType

//TypeID             uuid.UUID `bun:"type:uuid,notnull"`
//AddressID          uuid.UUID `bun:"type:uuid,notnull"`
//HighlightedTracks    []uuid.UUID `bun:",type:uuid[]" pg:",array"`
//FeaturedTrackGroupID uuid.UUID   `bun:"type:uuid,default:uuid_nil()"`

// Members        []UserGroup `pg:"many2many:user_group_members,fk:user_group_id,joinFK:member_id"`
// MemberOfGroups []UserGroup `pg:"many2many:user_group_members,fk:member_id,joinFK:user_group_id"`

//OwnerOfTracks      []Track      `pg:"fk:user_group_id"`          // user group gets paid for these tracks
//ArtistOfTracks     []uuid.UUID  `bun:",type:uuid[]" pg:",array"` // user group displayed as artist for these tracks
//OwnerOfTrackGroups []TrackGroup `pg:"fk:user_group_id"`          // user group owner of these track groups
//LabelOfTrackGroups []TrackGroup `pg:"fk:label_id"`               // label of these track groups

// func (u *UserGroup) BeforeInsert(c context.Context, db orm.DB) error {
// 	newPrivacy := &UserGroupPrivacy{Private: false, OwnedTracks: true, SupportedArtists: true}
// 	_, pgerr := db.Model(newPrivacy).Returning("*").Insert()
// 	if pgerr != nil {
// 		return pgerr
// 	}
// 	u.PrivacyID = newPrivacy.ID

// 	return nil
// }

// Create creates a new UserGroup
// func (u *UserGroup) Create(db *pg.DB, userGroup *pbUser.UserGroup) (error, string) {
// 	var table string
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err, table
// 	}
// 	defer tx.Rollback()

// 	groupTaxonomy := new(GroupTaxonomy)
// 	pgerr := tx.Model(groupTaxonomy).Where("type = ?", userGroup.Type.Type).First()

// 	if pgerr != nil {
// 		return pgerr, "group_taxonomy"
// 	}
// 	u.TypeID = groupTaxonomy.ID

// 	var newAddress *StreetAddress
// 	if userGroup.Address != nil {
// 		newAddress = &StreetAddress{Data: userGroup.Address.Data}
// 		_, pgerr = tx.Model(newAddress).Returning("*").Insert()
// 		if pgerr != nil {
// 			return pgerr, "street_address"
// 		}
// 	}
// 	u.AddressID = newAddress.ID

// 	linkIDs, pgerr := getLinkIDs(userGroup.Links, tx)
// 	if pgerr != nil {
// 		return pgerr, "link"
// 	}
// 	u.Links = linkIDs

// 	// tagIDs, pgerr := GetTagIDs(userGroup.Tags, tx)
// 	// if pgerr != nil {
// 	// 	return pgerr, "tag"
// 	// }
// 	// u.Tags = tagIDs

// 	// recommendedArtistIDs, pgerr := GetRelatedUserGroupIDs(userGroup.RecommendedArtists, tx)
// 	// if pgerr != nil {
// 	// 	return pgerr, "user_group"
// 	// }
// 	// u.RecommendedArtists = recommendedArtistIDs

// 	_, pgerr = tx.Model(u).Returning("*").Insert()
// 	if pgerr != nil {
// 		fmt.Println("insert")
// 		return pgerr, "user_group"
// 	}

// 	// if len(recommendedArtistIDs) > 0 {
// 	// 	_, pgerr = tx.Exec(`
// 	//     UPDATE user_groups
// 	//     SET recommended_by = (select array_agg(distinct e) from unnest(recommended_by || ?) e)
// 	//     WHERE id IN (?)
// 	//   `, pg.Array([]uuid.UUID{u.ID}), pg.In(recommendedArtistIDs))
// 	// 	if pgerr != nil {
// 	// 		return pgerr, "user_group"
// 	// 	}
// 	// }

// 	pgerr = tx.Model(u).
// 		Column("Privacy").
// 		WherePK().
// 		Select()
// 	if pgerr != nil {
// 		return pgerr, "user_group"
// 	}

// 	// Building response
// 	userGroup.Address.Id = u.AddressID.String()
// 	userGroup.Type.ID = u.TypeID.String()
// 	// userGroup.Privacy = &pbUser.Privacy{
// 	// 	ID:               u.Privacy.ID.String(),
// 	// 	Private:          u.Privacy.Private,
// 	// 	OwnedTracks:      u.Privacy.OwnedTracks,
// 	// 	SupportedArtists: u.Privacy.SupportedArtists,
// 	// }

// 	return tx.Commit(), table
// }

// func (u *UserGroup) Update(db *pg.DB, userGroup *pbUser.UserGroup) (error, string) {
// 	var table string
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err, "user_group"
// 	}
// 	defer tx.Rollback()

// 	userGroupToUpdate := &UserGroup{IDRecord: IDRecord{ID: u.ID}}
// 	pgerr := tx.Model(userGroupToUpdate).
// 		Column("user_group.links", "Type").
// 		WherePK().
// 		Select()
// 	if pgerr != nil {
// 		return pgerr, "user_group"
// 	}

// 	columns := []string{
// 		"updated_at",
// 		// "pro",
// 		// "publisher",
// 		"links",
// 		// "tags",
// 		"display_name",
// 		// "avatar",
// 		"description",
// 		"short_bio",
// 		// "banner",
// 		"group_email_address",
// 	}

// 	// User group type - changes allowed: user => artist, user => label
// 	groupTaxonomy := new(GroupTaxonomy)
// 	pgerr = tx.Model(groupTaxonomy).Where("type = ?", userGroup.Type.Type).First()
// 	if pgerr != nil {
// 		return pgerr, "group_taxonomy"
// 	}
// 	if userGroupToUpdate.Type.ID != groupTaxonomy.ID {
// 		if userGroupToUpdate.Type.Type != "user" ||
// 			(userGroupToUpdate.Type.Type == "user" && !(groupTaxonomy.Type == "artist" || groupTaxonomy.Type == "label")) {
// 			twerr := twirp.InvalidArgumentError("type", "not allowed")
// 			return twerr.(error), "user_group"
// 		}
// 		u.TypeID = groupTaxonomy.ID
// 		columns = append(columns, "type_id")
// 	}

// 	// Update address
// 	addressID, twerr := uuid.Parse(userGroup.Address.Id)
// 	if twerr != nil {
// 		return twerr, "street_address"
// 	}
// 	address := &StreetAddress{ID: addressID, Data: userGroup.Address.Data}
// 	_, pgerr = tx.Model(address).Column("data").WherePK().Update()
// 	// _, pgerr := db.Model(address).Set("data = ?", pg.Hstore(userGroup.Address.Data)).Where("id = ?id").Update()
// 	if pgerr != nil {
// 		return pgerr, "street_address"
// 	}

// 	// Update privacy
// 	// privacyID, twerr := uuidpkg.GetUUIDFromString(userGroup.Privacy.ID)
// 	// if twerr != nil {
// 	// 	return twerr, "user_group_privacy"
// 	// }
// 	// privacy := &UserGroupPrivacy{
// 	// 	ID:               privacyID,
// 	// 	Private:          userGroup.Privacy.Private,
// 	// 	OwnedTracks:      userGroup.Privacy.OwnedTracks,
// 	// 	SupportedArtists: userGroup.Privacy.SupportedArtists,
// 	// }
// 	// _, pgerr = tx.Model(privacy).WherePK().Returning("*").UpdateNotNull()
// 	// if pgerr != nil {
// 	// 	return pgerr, "user_group_privacy"
// 	// }

// 	// Update tags
// 	// tagIDs, pgerr := GetTagIDs(userGroup.Tags, tx)
// 	// if pgerr != nil {
// 	// 	return pgerr, "tag"
// 	// }

// 	// Update links
// 	linkIDs, pgerr := getLinkIDs(userGroup.Links, tx)
// 	if pgerr != nil {
// 		return pgerr, "link"
// 	}
// 	// Delete links if needed
// 	linkIDsToDelete := uuidpkg.Difference(userGroupToUpdate.Links, linkIDs)
// 	if len(linkIDsToDelete) > 0 {
// 		_, pgerr = tx.Model((*Link)(nil)).
// 			Where("id in (?)", pg.In(linkIDsToDelete)).
// 			Delete()
// 		if pgerr != nil {
// 			return pgerr, "link"
// 		}
// 	}

// 	// Update user group
// 	// u.Tags = tagIDs
// 	u.Links = linkIDs
// 	// u.RecommendedArtists = recommendedArtistIDs
// 	u.UpdatedAt = time.Now()
// 	_, pgerr = tx.Model(u).
// 		Column(columns...).
// 		WherePK().
// 		Returning("*").
// 		Update()
// 	if pgerr != nil {
// 		return pgerr, "user_group"
// 	}

// 	return tx.Commit(), table
// }

// func SearchUserGroups(query string, db *pg.DB) (*pbUser.SearchResults, twirp.Error) {
// 	var userGroups []UserGroup

// 	pgerr := db.Model(&userGroups).
// 		Column("user_group.id", "user_group.display_name", "user_group.avatar", "Privacy", "Type").
// 		Where("to_tsvector('english'::regconfig, COALESCE(display_name, '') || ' ' || COALESCE(f_arr2str(tags), '')) @@ (plainto_tsquery('english'::regconfig, ?)) = true", query).
// 		Where("privacy.private = false").
// 		Select()
// 	if pgerr != nil {
// 		return nil, errorpkg.CheckError(pgerr, "user_group")
// 	}

// 	var people []*pbUser.RelatedUserGroup
// 	var artists []*pbUser.RelatedUserGroup
// 	var labels []*pbUser.RelatedUserGroup
// 	for _, userGroup := range userGroups {
// 		searchUserGroup := &pbUser.RelatedUserGroup{
// 			Id:          userGroup.ID.String(),
// 			DisplayName: userGroup.DisplayName,
// 			// Avatar:      userGroup.Avatar,
// 		}
// 		switch userGroup.TypeID {
// 		case "user":
// 			people = append(people, searchUserGroup)
// 		case "artist":
// 			artists = append(artists, searchUserGroup)
// 		case "label":
// 			labels = append(labels, searchUserGroup)
// 		}
// 	}
// 	return &pbUser.SearchResults{
// 		People:  people,
// 		Artists: artists,
// 		Labels:  labels,
// 	}, nil
// }

// func (u *UserGroup) Delete(tx *pg.Tx) (error, string) {
// 	pgerr := tx.Model(u).
// 		Column("user_group.links", "user_group.followers", "user_group.recommended_by", "user_group.recommended_artists", "Address", "Privacy",
// 			"OwnerOfTrackGroups", "LabelOfTrackGroups", "user_group.artist_of_tracks").
// 		WherePK().
// 		Select()
// 	if pgerr != nil {
// 		return pgerr, "user_group"
// 	}

// 	// These tracks contain the user group to delete as artist
// 	// so we have to remove it from the tracks' artists list
// 	// if len(u.ArtistOfTracks) > 0 {
// 	// 	_, pgerr = tx.Exec(`
// 	//     UPDATE tracks
// 	//     SET artists = array_remove(artists, ?)
// 	//     WHERE id IN (?)
// 	//   `, u.ID, pg.In(u.ArtistOfTracks))
// 	// 	if pgerr != nil {
// 	// 		return pgerr, "track"
// 	// 	}
// 	// }

// 	// These track groups contain the user group to delete as label
// 	// so we have to set their label_id as null
// 	// if len(u.LabelOfTrackGroups) > 0 {
// 	// 	_, pgerr = tx.Model(&u.LabelOfTrackGroups).
// 	// 		Set("label_id = uuid_nil()").
// 	// 		Update()
// 	// 	if pgerr != nil {
// 	// 		return pgerr, "track"
// 	// 	}
// 	// }

// 	// Delete track groups owned by user group to delete
// 	// if a track is a release (lp, ep, single), its tracks are owned by the same user group
// 	// and they'll be deleted as well
// 	// for _, trackGroup := range u.OwnerOfTrackGroups {
// 	// 	pgerr, table := trackGroup.Delete(tx)
// 	// 	if pgerr != nil {
// 	// 		return pgerr, table
// 	// 	}
// 	// }

// 	if len(u.Links) > 0 {
// 		_, pgerr = tx.Model((*Link)(nil)).
// 			Where("id in (?)", pg.In(u.Links)).
// 			Delete()
// 		if pgerr != nil {
// 			return pgerr, "link"
// 		}
// 	}

// if len(u.RecommendedBy) > 0 {
// 	_, pgerr = tx.Exec(`
//     UPDATE user_groups
//     SET recommended_artists = array_remove(recommended_artists, ?)
//     WHERE id IN (?)
//   `, u.ID, pg.In(u.RecommendedBy))
// 	if pgerr != nil {
// 		return pgerr, "user_group"
// 	}
// }

// if len(u.RecommendedArtists) > 0 {
// 	_, pgerr = tx.Exec(`
//     UPDATE user_groups
//     SET recommended_by = array_remove(recommended_by, ?)
//     WHERE id IN (?)
//   `, u.ID, pg.In(u.RecommendedArtists))
// 	if pgerr != nil {
// 		return pgerr, "user_group"
// 	}
// }

// if len(u.Followers) > 0 {
// 	_, pgerr = tx.Exec(`
//     UPDATE users
//     SET followed_groups = array_remove(followed_groups, ?)
//     WHERE id IN (?)
//   `, u.ID, pg.In(u.Followers))
// 	if pgerr != nil {
// 		return pgerr, "user"
// 	}
// }

// 	var userGroupMembers []UserGroupMember
// 	_, pgerr = tx.Model(&userGroupMembers).
// 		Where("user_group_id = ?", u.ID).
// 		WhereOr("member_id = ?", u.ID).
// 		Delete()
// 	if pgerr != nil {
// 		return pgerr, "user_group_member"
// 	}

// 	_, pgerr = tx.Model(u).WherePK().Delete()
// 	if pgerr != nil {
// 		return pgerr, "user_group"
// 	}

// 	_, pgerr = tx.Model(u.Address).WherePK().Delete()
// 	if pgerr != nil {
// 		return pgerr, "street_address"
// 	}

// 	// _, pgerr = tx.Model(u.Privacy).WherePK().Delete()
// 	// if pgerr != nil {
// 	// 	return pgerr, "user_group_privacy"
// 	// }

// 	return nil, ""
// }

func (u *UserGroup) AddRecommended(db *pg.DB, recommendedID uuid.UUID) (error, string) {
	var table string
	tx, err := db.Begin()
	if err != nil {
		return err, table
	}
	defer tx.Rollback()

	res, pgerr := tx.Exec(`
    UPDATE user_groups
    SET recommended_artists = (select array_agg(distinct e) from unnest(recommended_artists || ?) e)
    WHERE id = ?
  `, pg.Array([]uuid.UUID{recommendedID}), u.ID)
	if res.RowsAffected() == 0 {
		return pg.ErrNoRows, "user_group"
	}
	if pgerr != nil {
		return pgerr, "user_group"
	}

	res, pgerr = tx.Exec(`
    UPDATE user_groups
    SET recommended_by = (select array_agg(distinct e) from unnest(recommended_by || ?) e)
    WHERE id = ?
  `, pg.Array([]uuid.UUID{u.ID}), recommendedID)
	if res.RowsAffected() == 0 {
		return pg.ErrNoRows, "recommended"
	}
	if pgerr != nil {
		return pgerr, "user_group"
	}

	return tx.Commit(), table
}

func (u *UserGroup) RemoveRecommended(db *pg.DB, recommendedID uuid.UUID) (error, string) {
	var table string
	tx, err := db.Begin()
	if err != nil {
		return err, table
	}
	defer tx.Rollback()

	res, pgerr := tx.Exec(`
    UPDATE user_groups
    SET recommended_artists = array_remove(recommended_artists, ?)
    WHERE id = ?
  `, recommendedID, u.ID)
	if res.RowsAffected() == 0 {
		return pg.ErrNoRows, "user_group"
	}
	if pgerr != nil {
		return pgerr, "user_group"
	}

	res, pgerr = tx.Exec(`
    UPDATE user_groups
    SET recommended_by = array_remove(recommended_by, ?)
    WHERE id = ?
  `, u.ID, recommendedID)
	if res.RowsAffected() == 0 {
		return pg.ErrNoRows, "recommended"
	}
	if pgerr != nil {
		return pgerr, "user_group"
	}

	return tx.Commit(), table
}

// Select user groups in db with given 'ids'
// Return slice of UserGroup response
func GetRelatedUserGroups(ids []uuid.UUID, db orm.DB) ([]*pbUser.RelatedUserGroup, error) {
	groupsResponse := make([]*pbUser.RelatedUserGroup, len(ids))
	if len(ids) > 0 {
		var groups []UserGroup
		pgerr := db.Model(&groups).
			Where("id in (?)", pg.In(ids)).
			Select()
		if pgerr != nil {
			return nil, pgerr
		}
		for i, group := range groups {
			groupsResponse[i] = &pbUser.RelatedUserGroup{
				Id:          group.ID.String(),
				DisplayName: group.DisplayName,
				// Avatar:      group.Avatar,
			}
		}
	}

	return groupsResponse, nil
}

// Select user groups in db with given ids in 'userGroups'
// Return ids slice
// Used in CreateUserGroup/UpdateUserGroup to add/update ids slice to recommended Artists
func GetRelatedUserGroupIDs(userGroups []*pbUser.RelatedUserGroup, db *pg.Tx) ([]uuid.UUID, error) {
	relatedUserGroups := make([]*UserGroup, len(userGroups))
	relatedUserGroupIDs := make([]uuid.UUID, len(userGroups))
	for i, userGroup := range userGroups {
		id, twerr := uuid.Parse(userGroup.Id)
		if twerr != nil {
			return nil, twerr.(error)
		}
		relatedUserGroups[i] = &UserGroup{IDRecord: IDRecord{ID: id}}
		pgerr := db.Model(relatedUserGroups[i]).
			WherePK().
			Returning("id", "display_name", "avatar").
			Select()
		if pgerr != nil {
			return nil, pgerr
		}
		userGroup.DisplayName = relatedUserGroups[i].DisplayName
		// userGroup.Avatar = relatedUserGroups[i].Avatar
		relatedUserGroupIDs[i] = relatedUserGroups[i].ID
	}
	return relatedUserGroupIDs, nil
}

func getLinkIDs(l []*pbUser.Link, db *pg.Tx) ([]uuid.UUID, error) {
	links := make([]*Link, len(l))
	linkIDs := make([]uuid.UUID, len(l))
	for i, link := range l {
		if link.ID == "" {
			links[i] = &Link{Platform: link.Platform, URI: link.Uri}
			_, pgerr := db.Model(links[i]).Returning("*").Insert()
			if pgerr != nil {
				return nil, pgerr
			}
			linkIDs[i] = links[i].ID
			link.ID = links[i].ID.String()
		} else {
			linkID, twerr := uuid.Parse(link.ID)
			if twerr != nil {
				return nil, twerr.(error)
			}
			linkIDs[i] = linkID
		}
	}
	return linkIDs, nil
}

/*type TrackAnalytics struct {
  ID uuid.UUID
  Title string
  PaidPlays int32
  FreePlays int32
  TotalCredits float32
}

// DEPRECATED - moved to Payment API
func (u *UserGroup) GetUserGroupTrackAnalytics(db *pg.DB) ([]*pbUser.TrackAnalytics, twirp.Error) {
  pgerr := db.Model(u).
    Column("OwnerOfTracks").
    WherePK().
    Select()
  if pgerr != nil {
    return nil, errorpkg.CheckError(pgerr, "user_group")
  }
  tracks := make([]TrackAnalytics, len(u.OwnerOfTracks))
  trackIDs := make([]uuid.UUID, len(u.OwnerOfTracks))
  for i, track := range(u.OwnerOfTracks) {
    tracks[i] = TrackAnalytics{
      Title: track.Title,
    }
    trackIDs[i] = track.ID
  }
  artistTrackAnalytics := make([]*pbUser.TrackAnalytics, len(tracks))

  if len(u.OwnerOfTracks) > 0 {
    _, pgerr := db.Query(&tracks, `
      SELECT play.track_id AS id,
        count(case when play.type = 'paid' then 1 else null end) AS paid_plays,
        count(case when play.type = 'free' then 1 else null end) AS free_plays,
        SUM(play.credits) AS total_credits
      FROM plays AS play
      WHERE play.track_id IN (?)
      GROUP BY play.track_id
    `, pg.In(trackIDs))
    if pgerr != nil {
      return nil, errorpkg.CheckError(pgerr, "play")
    }
    for i, track := range(tracks) {
      artistTrackAnalytics[i] = &pbUser.TrackAnalytics{
        ID: track.ID.String(),
        Title: track.Title,
        TotalPlays: track.PaidPlays + track.FreePlays,
        PaidPlays: track.PaidPlays,
        FreePlays: track.FreePlays,
        TotalCredits: float32(track.TotalCredits),
        UserGroupCredits: 0.7*float32(track.TotalCredits),
        ResonateCredits: 0.3*float32(track.TotalCredits),
      }
    }
  }

  return artistTrackAnalytics, nil
}*/
