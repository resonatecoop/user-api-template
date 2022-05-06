// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ResonateUserClient is the client API for ResonateUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResonateUserClient interface {
	//GetUser provides a public level of information about a user
	GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserPublicResponse, error)
	//rpc AddUser(UserAddRequest) returns (User) {
	AddUser(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserRequest, error)
	//rpc UpdateUser(UserUpdateRequest) returns (Empty) {
	UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	//rpc UpdateUserRestricted(UserUpdateRestrictedRequest) returns (Empty) {
	UpdateUserRestricted(ctx context.Context, in *UserUpdateRestrictedRequest, opts ...grpc.CallOption) (*Empty, error)
	//GetUserRestricted provides private level of information about a user
	GetUserRestricted(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserPrivateResponse, error)
	GetUserMembership(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserMembershipResponse, error)
	GetUserCredits(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserCreditResponse, error)
	DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Empty, error)
	//ListUsers returns a list of all Users
	ListUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error)
	//AddUploadSubmission adds a UserUploadSubmission
	AddUploadSubmission(ctx context.Context, in *UploadSubmissionAddRequest, opts ...grpc.CallOption) (*UploadSubmissionRequest, error)
	//UpdateUploadSubmission updates an existing UploadSubmission
	UpdateUploadSubmission(ctx context.Context, in *UploadSubmissionUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	//DeleteUploadSubmission deletes an upload submission
	DeleteUploadSubmission(ctx context.Context, in *UploadSubmissionRequest, opts ...grpc.CallOption) (*Empty, error)
	//AddUserGroup adds a UserGroup based on provided attributes
	AddUserGroup(ctx context.Context, in *UserGroupCreateRequest, opts ...grpc.CallOption) (*UserRequest, error)
	//UpdateUserGroup updates an existing UserGroup
	UpdateUserGroup(ctx context.Context, in *UserGroupUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
	//GetUserGroup provides a public level of information about a user group
	GetUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*UserGroupPublicResponse, error)
	//DeleteUserGroup deletes a UserGroup
	DeleteUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*Empty, error)
	ListUsersUserGroups(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserGroupListResponse, error)
	CreateTrackGroup(ctx context.Context, in *TrackGroupCreateRequest, opts ...grpc.CallOption) (*TrackGroupResponse, error)
	GetTrackGroup(ctx context.Context, in *TrackGroupRequest, opts ...grpc.CallOption) (*TrackGroupResponse, error)
}

type resonateUserClient struct {
	cc grpc.ClientConnInterface
}

func NewResonateUserClient(cc grpc.ClientConnInterface) ResonateUserClient {
	return &resonateUserClient{cc}
}

func (c *resonateUserClient) GetUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserPublicResponse, error) {
	out := new(UserPublicResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) AddUser(ctx context.Context, in *UserAddRequest, opts ...grpc.CallOption) (*UserRequest, error) {
	out := new(UserRequest)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) UpdateUser(ctx context.Context, in *UserUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) UpdateUserRestricted(ctx context.Context, in *UserUpdateRestrictedRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/UpdateUserRestricted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) GetUserRestricted(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserPrivateResponse, error) {
	out := new(UserPrivateResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/GetUserRestricted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) GetUserMembership(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserMembershipResponse, error) {
	out := new(UserMembershipResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/GetUserMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) GetUserCredits(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserCreditResponse, error) {
	out := new(UserCreditResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/GetUserCredits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) DeleteUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) ListUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) AddUploadSubmission(ctx context.Context, in *UploadSubmissionAddRequest, opts ...grpc.CallOption) (*UploadSubmissionRequest, error) {
	out := new(UploadSubmissionRequest)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/AddUploadSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) UpdateUploadSubmission(ctx context.Context, in *UploadSubmissionUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/UpdateUploadSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) DeleteUploadSubmission(ctx context.Context, in *UploadSubmissionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/DeleteUploadSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) AddUserGroup(ctx context.Context, in *UserGroupCreateRequest, opts ...grpc.CallOption) (*UserRequest, error) {
	out := new(UserRequest)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/AddUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) UpdateUserGroup(ctx context.Context, in *UserGroupUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/UpdateUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) GetUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*UserGroupPublicResponse, error) {
	out := new(UserGroupPublicResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/GetUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) DeleteUserGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/DeleteUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) ListUsersUserGroups(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserGroupListResponse, error) {
	out := new(UserGroupListResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/ListUsersUserGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) CreateTrackGroup(ctx context.Context, in *TrackGroupCreateRequest, opts ...grpc.CallOption) (*TrackGroupResponse, error) {
	out := new(TrackGroupResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/CreateTrackGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resonateUserClient) GetTrackGroup(ctx context.Context, in *TrackGroupRequest, opts ...grpc.CallOption) (*TrackGroupResponse, error) {
	out := new(TrackGroupResponse)
	err := c.cc.Invoke(ctx, "/user.ResonateUser/GetTrackGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResonateUserServer is the server API for ResonateUser service.
// All implementations should embed UnimplementedResonateUserServer
// for forward compatibility
type ResonateUserServer interface {
	//GetUser provides a public level of information about a user
	GetUser(context.Context, *UserRequest) (*UserPublicResponse, error)
	//rpc AddUser(UserAddRequest) returns (User) {
	AddUser(context.Context, *UserAddRequest) (*UserRequest, error)
	//rpc UpdateUser(UserUpdateRequest) returns (Empty) {
	UpdateUser(context.Context, *UserUpdateRequest) (*Empty, error)
	//rpc UpdateUserRestricted(UserUpdateRestrictedRequest) returns (Empty) {
	UpdateUserRestricted(context.Context, *UserUpdateRestrictedRequest) (*Empty, error)
	//GetUserRestricted provides private level of information about a user
	GetUserRestricted(context.Context, *UserRequest) (*UserPrivateResponse, error)
	GetUserMembership(context.Context, *UserRequest) (*UserMembershipResponse, error)
	GetUserCredits(context.Context, *UserRequest) (*UserCreditResponse, error)
	DeleteUser(context.Context, *UserRequest) (*Empty, error)
	//ListUsers returns a list of all Users
	ListUsers(context.Context, *Empty) (*UserListResponse, error)
	//AddUploadSubmission adds a UserUploadSubmission
	AddUploadSubmission(context.Context, *UploadSubmissionAddRequest) (*UploadSubmissionRequest, error)
	//UpdateUploadSubmission updates an existing UploadSubmission
	UpdateUploadSubmission(context.Context, *UploadSubmissionUpdateRequest) (*Empty, error)
	//DeleteUploadSubmission deletes an upload submission
	DeleteUploadSubmission(context.Context, *UploadSubmissionRequest) (*Empty, error)
	//AddUserGroup adds a UserGroup based on provided attributes
	AddUserGroup(context.Context, *UserGroupCreateRequest) (*UserRequest, error)
	//UpdateUserGroup updates an existing UserGroup
	UpdateUserGroup(context.Context, *UserGroupUpdateRequest) (*Empty, error)
	//GetUserGroup provides a public level of information about a user group
	GetUserGroup(context.Context, *UserGroupRequest) (*UserGroupPublicResponse, error)
	//DeleteUserGroup deletes a UserGroup
	DeleteUserGroup(context.Context, *UserGroupRequest) (*Empty, error)
	ListUsersUserGroups(context.Context, *UserRequest) (*UserGroupListResponse, error)
	CreateTrackGroup(context.Context, *TrackGroupCreateRequest) (*TrackGroupResponse, error)
	GetTrackGroup(context.Context, *TrackGroupRequest) (*TrackGroupResponse, error)
}

// UnimplementedResonateUserServer should be embedded to have forward compatible implementations.
type UnimplementedResonateUserServer struct {
}

func (UnimplementedResonateUserServer) GetUser(context.Context, *UserRequest) (*UserPublicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedResonateUserServer) AddUser(context.Context, *UserAddRequest) (*UserRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedResonateUserServer) UpdateUser(context.Context, *UserUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedResonateUserServer) UpdateUserRestricted(context.Context, *UserUpdateRestrictedRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRestricted not implemented")
}
func (UnimplementedResonateUserServer) GetUserRestricted(context.Context, *UserRequest) (*UserPrivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRestricted not implemented")
}
func (UnimplementedResonateUserServer) GetUserMembership(context.Context, *UserRequest) (*UserMembershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMembership not implemented")
}
func (UnimplementedResonateUserServer) GetUserCredits(context.Context, *UserRequest) (*UserCreditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCredits not implemented")
}
func (UnimplementedResonateUserServer) DeleteUser(context.Context, *UserRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedResonateUserServer) ListUsers(context.Context, *Empty) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedResonateUserServer) AddUploadSubmission(context.Context, *UploadSubmissionAddRequest) (*UploadSubmissionRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUploadSubmission not implemented")
}
func (UnimplementedResonateUserServer) UpdateUploadSubmission(context.Context, *UploadSubmissionUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUploadSubmission not implemented")
}
func (UnimplementedResonateUserServer) DeleteUploadSubmission(context.Context, *UploadSubmissionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUploadSubmission not implemented")
}
func (UnimplementedResonateUserServer) AddUserGroup(context.Context, *UserGroupCreateRequest) (*UserRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserGroup not implemented")
}
func (UnimplementedResonateUserServer) UpdateUserGroup(context.Context, *UserGroupUpdateRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGroup not implemented")
}
func (UnimplementedResonateUserServer) GetUserGroup(context.Context, *UserGroupRequest) (*UserGroupPublicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroup not implemented")
}
func (UnimplementedResonateUserServer) DeleteUserGroup(context.Context, *UserGroupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGroup not implemented")
}
func (UnimplementedResonateUserServer) ListUsersUserGroups(context.Context, *UserRequest) (*UserGroupListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsersUserGroups not implemented")
}
func (UnimplementedResonateUserServer) CreateTrackGroup(context.Context, *TrackGroupCreateRequest) (*TrackGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrackGroup not implemented")
}
func (UnimplementedResonateUserServer) GetTrackGroup(context.Context, *TrackGroupRequest) (*TrackGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrackGroup not implemented")
}

// UnsafeResonateUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResonateUserServer will
// result in compilation errors.
type UnsafeResonateUserServer interface {
	mustEmbedUnimplementedResonateUserServer()
}

func RegisterResonateUserServer(s grpc.ServiceRegistrar, srv ResonateUserServer) {
	s.RegisterService(&_ResonateUser_serviceDesc, srv)
}

func _ResonateUser_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).AddUser(ctx, req.(*UserAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).UpdateUser(ctx, req.(*UserUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_UpdateUserRestricted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateRestrictedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).UpdateUserRestricted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/UpdateUserRestricted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).UpdateUserRestricted(ctx, req.(*UserUpdateRestrictedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_GetUserRestricted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).GetUserRestricted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/GetUserRestricted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).GetUserRestricted(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_GetUserMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).GetUserMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/GetUserMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).GetUserMembership(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_GetUserCredits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).GetUserCredits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/GetUserCredits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).GetUserCredits(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).DeleteUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).ListUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_AddUploadSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSubmissionAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).AddUploadSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/AddUploadSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).AddUploadSubmission(ctx, req.(*UploadSubmissionAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_UpdateUploadSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSubmissionUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).UpdateUploadSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/UpdateUploadSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).UpdateUploadSubmission(ctx, req.(*UploadSubmissionUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_DeleteUploadSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).DeleteUploadSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/DeleteUploadSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).DeleteUploadSubmission(ctx, req.(*UploadSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_AddUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).AddUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/AddUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).AddUserGroup(ctx, req.(*UserGroupCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_UpdateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).UpdateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/UpdateUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).UpdateUserGroup(ctx, req.(*UserGroupUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_GetUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).GetUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/GetUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).GetUserGroup(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_DeleteUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).DeleteUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/DeleteUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).DeleteUserGroup(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_ListUsersUserGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).ListUsersUserGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/ListUsersUserGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).ListUsersUserGroups(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_CreateTrackGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackGroupCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).CreateTrackGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/CreateTrackGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).CreateTrackGroup(ctx, req.(*TrackGroupCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResonateUser_GetTrackGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResonateUserServer).GetTrackGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.ResonateUser/GetTrackGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResonateUserServer).GetTrackGroup(ctx, req.(*TrackGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResonateUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.ResonateUser",
	HandlerType: (*ResonateUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _ResonateUser_GetUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _ResonateUser_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ResonateUser_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserRestricted",
			Handler:    _ResonateUser_UpdateUserRestricted_Handler,
		},
		{
			MethodName: "GetUserRestricted",
			Handler:    _ResonateUser_GetUserRestricted_Handler,
		},
		{
			MethodName: "GetUserMembership",
			Handler:    _ResonateUser_GetUserMembership_Handler,
		},
		{
			MethodName: "GetUserCredits",
			Handler:    _ResonateUser_GetUserCredits_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _ResonateUser_DeleteUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _ResonateUser_ListUsers_Handler,
		},
		{
			MethodName: "AddUploadSubmission",
			Handler:    _ResonateUser_AddUploadSubmission_Handler,
		},
		{
			MethodName: "UpdateUploadSubmission",
			Handler:    _ResonateUser_UpdateUploadSubmission_Handler,
		},
		{
			MethodName: "DeleteUploadSubmission",
			Handler:    _ResonateUser_DeleteUploadSubmission_Handler,
		},
		{
			MethodName: "AddUserGroup",
			Handler:    _ResonateUser_AddUserGroup_Handler,
		},
		{
			MethodName: "UpdateUserGroup",
			Handler:    _ResonateUser_UpdateUserGroup_Handler,
		},
		{
			MethodName: "GetUserGroup",
			Handler:    _ResonateUser_GetUserGroup_Handler,
		},
		{
			MethodName: "DeleteUserGroup",
			Handler:    _ResonateUser_DeleteUserGroup_Handler,
		},
		{
			MethodName: "ListUsersUserGroups",
			Handler:    _ResonateUser_ListUsersUserGroups_Handler,
		},
		{
			MethodName: "CreateTrackGroup",
			Handler:    _ResonateUser_CreateTrackGroup_Handler,
		},
		{
			MethodName: "GetTrackGroup",
			Handler:    _ResonateUser_GetTrackGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
