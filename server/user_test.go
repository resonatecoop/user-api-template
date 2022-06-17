package server_test

import (
	"fmt"

	"github.com/resonatecoop/user-api-template/model"

	_ "github.com/jackc/pgx/v4/stdlib"
	pbUser "github.com/resonatecoop/user-api-template/proto/user"
	"github.com/stretchr/testify/assert"
)

// type UserApiTestSuite struct {
// 	suite.Suite
// 	db  *bun.DB
// 	ctx context.Context
// }

// func (suite *UserApiTestSuite) SetupTest() {
// 	//TODO
// }

func getIntPointer(value int32) *int32 {
	return &value
}

func (suite *UserApiTestSuite) TestAddUser() {
	ctx := suite.ctx

	// When attempting to add a user with admin access,
	// the user is created but limited to default user role
	newuser := pbUser.UserAddRequest{
		Username: "joe@bloggs.com",
		FullName: "Joe Bloggs",
		RoleId:   getIntPointer(1),
	}

	_, err := suite.server.AddUser(ctx, &newuser)
	if err != nil {
		panic(err)
	}

	user := new(model.User)

	err = suite.db.NewSelect().
		Model(user).
		Where("username = ?", "joe@bloggs.com").
		Limit(1).
		Scan(ctx)

	// the record exists
	assert.Nil(suite.T(), err)

	// but the role has been limited to default
	assert.Equal(suite.T(), int32(model.UserRole), user.RoleID)
}

func (suite *UserApiTestSuite) TestDeleteUser() {

	ctx := suite.ctx

	type cases []pbUser.UserRequest

	users := cases{
		{Id: "243b4178-6f98-4bf1-bbb1-46b57a901816"},
		{Id: "5253747c-2b8c-40e2-8a70-bab91348a9bd"},
		{Id: "86a72db9-eb16-45ae-8003-ac37d8dc5cbb"},
	}

	var response *pbUser.UserListResponse
	var empty *pbUser.Empty

	response, err := suite.server.ListUsers(ctx, empty)
	if err != nil {
		panic(err)
	}

	assert.Equal(suite.T(), 5, len(response.User))

	_, err = suite.server.DeleteUser(ctx, &users[0])
	if err != nil {
		panic(err)
	}

	response, err = suite.server.ListUsers(ctx, empty)
	if err != nil {
		panic(err)
	}
	assert.Equal(suite.T(), 4, len(response.User))

	newuser := pbUser.UserAddRequest{
		Username: "joe@bloggs.com",
		FullName: "Joe Bloggs",
	}

	_, err = suite.server.AddUser(ctx, &newuser)
	if err != nil {
		panic(err)
	}

	response, err = suite.server.ListUsers(ctx, empty)
	if err != nil {
		panic(err)
	}
	assert.Equal(suite.T(), 5, len(response.User))

	user := new(model.User)

	if err = suite.db.NewSelect().
		Model(user).
		Where("id = ?", "80b26113-37e0-456a-9f75-01db0eb550f8").
		Limit(1).
		Scan(ctx); err != nil {
		fmt.Printf("No such user returned!")
	}
}

// func (db *bun.DB, ctx context.Context) RunUserTests() {
// 	testrun := new(UserApiTestSuite)
// 	testrun.db = db
// 	testrun.ctx = ctx

// 	suite.Run(nil, testrun)
// }

// func (suite *OauthTestSuite) TestUserExistsFindsValidUser() {
// 	validUsername := suite.users[0].Username
// 	assert.True(suite.T(), suite.service.UserExists(validUsername))
// }

// func (suite *OauthTestSuite) TestUserExistsDoesntFindInvalidUser() {
// 	invalidUsername := "bogus_name"
// 	assert.False(suite.T(), suite.service.UserExists(invalidUsername))
// }

// func (suite *OauthTestSuite) TestUpdateUsernameWorksWithValidEntry() {
// 	user, err := suite.service.CreateUser(
// 		model.UserRole,  // role ID
// 		"test@newuser",  // username
// 		"test_password", // password
// 	)

// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), user)
// 	assert.Equal(suite.T(), "test@newuser", user.Username)

// 	newUsername := "mynew@email"

// 	err = suite.service.UpdateUsername(user, newUsername)

// 	assert.NoError(suite.T(), err)

// 	assert.Equal(suite.T(), newUsername, user.Username)
// }

// func (suite *OauthTestSuite) TestUpdateUsernameTxWorksWithValidEntry() {
// 	user, err := suite.service.CreateUser(
// 		roles.User,      // role ID
// 		"test@newuser",  // username
// 		"test_password", // password
// 	)

// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), user)
// 	assert.Equal(suite.T(), "test@newuser", user.Username)

// 	newUsername := "mynew@email"

// 	err = suite.service.UpdateUsernameTx(suite.db, user, newUsername)

// 	assert.NoError(suite.T(), err)

// 	assert.Equal(suite.T(), newUsername, user.Username)
// }

// func (suite *OauthTestSuite) TestUpdateUsernameFailsWithABlankEntry() {
// 	user, err := suite.service.CreateUser(
// 		roles.User,      // role ID
// 		"test@newuser",  // username
// 		"test_password", // password
// 	)

// 	assert.NoError(suite.T(), err)
// 	assert.NotNil(suite.T(), user)
// 	assert.Equal(suite.T(), "test@newuser", user.Username)

// 	newUsername := ""

// 	err = suite.service.UpdateUsername(user, newUsername)

// 	assert.EqualError(suite.T(), err, oauth.ErrCannotSetEmptyUsername.Error())

// 	assert.NotEqual(suite.T(), newUsername, user.Username)
// }

// func (suite *OauthTestSuite) TestFindUserByUsername() {
// 	var (
// 		user *model.User
// 		err  error
// 	)

// 	// When we try to find a user with a bogus username
// 	user, err = suite.service.FindUserByUsername("bogus")

// 	// User object should be nil
// 	assert.Nil(suite.T(), user)

// 	// Correct error should be returned
// 	if assert.NotNil(suite.T(), err) {
// 		assert.Equal(suite.T(), oauth.ErrUserNotFound, err)
// 	}

// 	// When we try to find a user with a valid username
// 	user, err = suite.service.FindUserByUsername("test@user")

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// Correct user object should be returned
// 	if assert.NotNil(suite.T(), user) {
// 		assert.Equal(suite.T(), "test@user", user.Username)
// 	}

// 	// Test username case insensitiviness
// 	user, err = suite.service.FindUserByUsername("TeSt@UsEr")

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// Correct user object should be returned
// 	if assert.NotNil(suite.T(), user) {
// 		assert.Equal(suite.T(), "test@user", user.Username)
// 	}
// }

// func (suite *OauthTestSuite) TestCreateUser() {
// 	var (
// 		user *model.User
// 		err  error
// 	)

// 	// We try to insert a non unique user
// 	user, err = suite.service.CreateUser(
// 		roles.User,      // role ID
// 		"test@user",     // username
// 		"test_password", // password
// 	)

// 	// User object should be nil
// 	assert.Nil(suite.T(), user)

// 	// Correct error should be returned
// 	if assert.NotNil(suite.T(), err) {
// 		assert.Equal(suite.T(), oauth.ErrUsernameTaken.Error(), err.Error())
// 	}

// 	// We try to insert a unique user
// 	user, err = suite.service.CreateUser(
// 		roles.User,      // role ID
// 		"test@newuser",  // username
// 		"test_password", // password
// 	)

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// Correct user object should be returned
// 	if assert.NotNil(suite.T(), user) {
// 		assert.Equal(suite.T(), "test@newuser", user.Username)
// 	}

// 	// Test username case insensitivity
// 	user, err = suite.service.CreateUser(
// 		roles.User,      // role ID
// 		"TeSt@NeWuSeR2", // username
// 		"test_password", // password
// 	)

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// Correct user object should be returned
// 	if assert.NotNil(suite.T(), user) {
// 		assert.Equal(suite.T(), "test@newuser2", user.Username)
// 	}
// }

// func (suite *OauthTestSuite) TestSetPassword() {
// 	var (
// 		user *model.User
// 		err  error
// 	)

// 	// Insert a test user without a password
// 	user = &model.User{
// 		MyGormModel: model.MyGormModel{
// 			ID:        uuid.New(),
// 			CreatedAt: time.Now().UTC(),
// 		},
// 		RoleID:   util.StringOrNull(roles.User),
// 		Username: "test@user_nopass",
// 		Password: util.StringOrNull(""),
// 	}
// 	err = suite.suite.db.Create(user).Error
// 	assert.NoError(suite.T(), err, "Inserting test data failed")

// 	// Try to set an empty password
// 	err = suite.service.SetPassword(user, "")

// 	// Correct error should be returned
// 	if assert.NotNil(suite.T(), err) {
// 		assert.Equal(suite.T(), oauth.ErrPasswordTooShort, err)
// 	}

// 	// Try changing the password
// 	err = suite.service.SetPassword(user, "test_password")

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// User object should have been updated
// 	assert.Equal(suite.T(), "test@user_nopass", user.Username)
// 	assert.Nil(suite.T(), pass.VerifyPassword(user.Password.String, "test_password"))
// }

// func (suite *OauthTestSuite) TestAuthUser() {
// 	var (
// 		user *model.User
// 		err  error
// 	)

// 	// Insert a test user without a password
// 	err = suite.suite.db.Create(&model.User{
// 		MyGormModel: model.MyGormModel{
// 			ID:        uuid.New(),
// 			CreatedAt: time.Now().UTC(),
// 		},
// 		RoleID:   util.StringOrNull(roles.User),
// 		Username: "test@user_nopass",
// 		Password: util.StringOrNull(""),
// 	}).Error
// 	assert.NoError(suite.T(), err, "Inserting test data failed")

// 	// When we try to authenticate a user without a password
// 	user, err = suite.service.AuthUser("test@user_nopass", "bogus")

// 	// User object should be nil
// 	assert.Nil(suite.T(), user)

// 	// Correct error should be returned
// 	if assert.NotNil(suite.T(), err) {
// 		assert.Equal(suite.T(), oauth.ErrUserPasswordNotSet, err)
// 	}

// 	// When we try to authenticate with a bogus username
// 	user, err = suite.service.AuthUser("bogus", "test_password")

// 	// User object should be nil
// 	assert.Nil(suite.T(), user)

// 	// Correct error should be returned
// 	if assert.NotNil(suite.T(), err) {
// 		assert.Equal(suite.T(), oauth.ErrUserNotFound, err)
// 	}

// 	// When we try to authenticate with an invalid password
// 	user, err = suite.service.AuthUser("test@user", "bogus")

// 	// User object should be nil
// 	assert.Nil(suite.T(), user)

// 	// Correct error should be returned
// 	if assert.NotNil(suite.T(), err) {
// 		assert.Equal(suite.T(), oauth.ErrInvalidUserPassword, err)
// 	}

// 	// When we try to authenticate with valid username and password
// 	user, err = suite.service.AuthUser("test@user", "test_password")

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// Correct user object should be returned
// 	if assert.NotNil(suite.T(), user) {
// 		assert.Equal(suite.T(), "test@user", user.Username)
// 	}

// 	// Test username case insensitivity
// 	user, err = suite.service.AuthUser("TeSt@UsEr", "test_password")

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// Correct user object should be returned
// 	if assert.NotNil(suite.T(), user) {
// 		assert.Equal(suite.T(), "test@user", user.Username)
// 	}
// }

// func (suite *OauthTestSuite) TestBlankPassword() {
// 	var (
// 		user *model.User
// 		err  error
// 	)

// 	user, err = suite.service.CreateUser(
// 		roles.User,         // role ID
// 		"test@user_nopass", // username
// 		"",                 // password,
// 	)

// 	// Error should be nil
// 	assert.Nil(suite.T(), err)

// 	// Correct user object should be returned
// 	if assert.NotNil(suite.T(), user) {
// 		assert.Equal(suite.T(), "test@user_nopass", user.Username)
// 	}

// 	// When we try to authenticate
// 	user, err = suite.service.AuthUser("test@user_nopass", "")

// 	// User object should be nil
// 	assert.Nil(suite.T(), user)

// 	// Correct error should be returned
// 	if assert.NotNil(suite.T(), err) {
// 		assert.Equal(suite.T(), oauth.ErrUserPasswordNotSet, err)
// 	}
// }
