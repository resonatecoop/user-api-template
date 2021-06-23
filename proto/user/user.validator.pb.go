// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: user/user.proto

//package example;
//package resonate.api.user;

package user

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UserRequest) Validate() error {
	return nil
}
func (this *UserOptionalRequest) Validate() error {
	return nil
}
func (this *ResetUserPasswordRequest) Validate() error {
	return nil
}
func (this *UserUpdateRequest) Validate() error {
	return nil
}
func (this *UserUpdateRestrictedRequest) Validate() error {
	return nil
}
func (this *UserPrivateResponse) Validate() error {
	return nil
}
func (this *UserPublicResponse) Validate() error {
	return nil
}
func (this *UserAddRequest) Validate() error {
	return nil
}
func (this *UserListResponse) Validate() error {
	for _, item := range this.User {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("User", err)
			}
		}
	}
	return nil
}
