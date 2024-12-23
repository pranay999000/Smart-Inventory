// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: user-service/proto/user/user.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserStatus int32

const (
	UserStatus_ACTIVE   UserStatus = 0
	UserStatus_INACTIVE UserStatus = 1
	UserStatus_BANNED   UserStatus = 2
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "ACTIVE",
		1: "INACTIVE",
		2: "BANNED",
	}
	UserStatus_value = map[string]int32{
		"ACTIVE":   0,
		"INACTIVE": 1,
		"BANNED":   2,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_proto_user_user_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_user_service_proto_user_user_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{0}
}

type UserRole int32

const (
	UserRole_UNKNOWN   UserRole = 0
	UserRole_ADMIN     UserRole = 1
	UserRole_CONSUMER  UserRole = 2
	UserRole_MANAGER   UserRole = 3
	UserRole_VENDOR    UserRole = 4
	UserRole_ASSISTANT UserRole = 5
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "UNKNOWN",
		1: "ADMIN",
		2: "CONSUMER",
		3: "MANAGER",
		4: "VENDOR",
		5: "ASSISTANT",
	}
	UserRole_value = map[string]int32{
		"UNKNOWN":   0,
		"ADMIN":     1,
		"CONSUMER":  2,
		"MANAGER":   3,
		"VENDOR":    4,
		"ASSISTANT": 5,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_user_service_proto_user_user_proto_enumTypes[1].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_user_service_proto_user_user_proto_enumTypes[1]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{1}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string     `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password    string     `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNumber string     `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	FirstName   string     `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName  string     `protobuf:"bytes,5,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName    string     `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Status      UserStatus `protobuf:"varint,7,opt,name=status,proto3,enum=user.UserStatus" json:"status,omitempty"`
	Role        UserRole   `protobuf:"varint,8,opt,name=role,proto3,enum=user.UserRole" json:"role,omitempty"`
	Avatar      string     `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Location    string     `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	BusinessId  int32      `protobuf:"varint,11,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	Id          string     `protobuf:"bytes,12,opt,name=id,proto3" json:"id,omitempty"`
	VendorId    int32      `protobuf:"varint,13,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_ACTIVE
}

func (x *User) GetRole() UserRole {
	if x != nil {
		return x.Role
	}
	return UserRole_UNKNOWN
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *User) GetBusinessId() int32 {
	if x != nil {
		return x.BusinessId
	}
	return 0
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetVendorId() int32 {
	if x != nil {
		return x.VendorId
	}
	return 0
}

type SuccessMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AuthToken  string `protobuf:"bytes,2,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	FirstName  string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName string `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName   string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Avatar     string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SuccessMessage) Reset() {
	*x = SuccessMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessMessage) ProtoMessage() {}

func (x *SuccessMessage) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessMessage.ProtoReflect.Descriptor instead.
func (*SuccessMessage) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *SuccessMessage) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SuccessMessage) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *SuccessMessage) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SuccessMessage) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *SuccessMessage) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SuccessMessage) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *SuccessMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SignUpUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SignUpUserRequest) Reset() {
	*x = SignUpUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpUserRequest) ProtoMessage() {}

func (x *SignUpUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpUserRequest.ProtoReflect.Descriptor instead.
func (*SignUpUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *SignUpUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SignUpUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ErrMessage *ErrorMessage `protobuf:"bytes,2,opt,name=err_message,json=errMessage,proto3" json:"err_message,omitempty"`
}

func (x *SignUpUserResponse) Reset() {
	*x = SignUpUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpUserResponse) ProtoMessage() {}

func (x *SignUpUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpUserResponse.ProtoReflect.Descriptor instead.
func (*SignUpUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *SignUpUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SignUpUserResponse) GetErrMessage() *ErrorMessage {
	if x != nil {
		return x.ErrMessage
	}
	return nil
}

type SignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthTK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthTK) Reset() {
	*x = AuthTK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthTK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTK) ProtoMessage() {}

func (x *AuthTK) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTK.ProtoReflect.Descriptor instead.
func (*AuthTK) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *AuthTK) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AuthTK) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are assignable to Result:
	//
	//	*SignInResponse_ErrMessage
	//	*SignInResponse_AuthTk
	Result isSignInResponse_Result `protobuf_oneof:"result"`
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *SignInResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *SignInResponse) GetResult() isSignInResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *SignInResponse) GetErrMessage() *ErrorMessage {
	if x, ok := x.GetResult().(*SignInResponse_ErrMessage); ok {
		return x.ErrMessage
	}
	return nil
}

func (x *SignInResponse) GetAuthTk() *AuthTK {
	if x, ok := x.GetResult().(*SignInResponse_AuthTk); ok {
		return x.AuthTk
	}
	return nil
}

type isSignInResponse_Result interface {
	isSignInResponse_Result()
}

type SignInResponse_ErrMessage struct {
	ErrMessage *ErrorMessage `protobuf:"bytes,2,opt,name=err_message,json=errMessage,proto3,oneof"`
}

type SignInResponse_AuthTk struct {
	AuthTk *AuthTK `protobuf:"bytes,3,opt,name=auth_tk,json=authTk,proto3,oneof"`
}

func (*SignInResponse_ErrMessage) isSignInResponse_Result() {}

func (*SignInResponse_AuthTk) isSignInResponse_Result() {}

type GetBusinessIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetBusinessIdRequest) Reset() {
	*x = GetBusinessIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessIdRequest) ProtoMessage() {}

func (x *GetBusinessIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessIdRequest.ProtoReflect.Descriptor instead.
func (*GetBusinessIdRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetBusinessIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetBusinessIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are assignable to Result:
	//
	//	*GetBusinessIdResponse_ErrMessage
	//	*GetBusinessIdResponse_BusinessId
	Result isGetBusinessIdResponse_Result `protobuf_oneof:"result"`
}

func (x *GetBusinessIdResponse) Reset() {
	*x = GetBusinessIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessIdResponse) ProtoMessage() {}

func (x *GetBusinessIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessIdResponse.ProtoReflect.Descriptor instead.
func (*GetBusinessIdResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetBusinessIdResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *GetBusinessIdResponse) GetResult() isGetBusinessIdResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *GetBusinessIdResponse) GetErrMessage() string {
	if x, ok := x.GetResult().(*GetBusinessIdResponse_ErrMessage); ok {
		return x.ErrMessage
	}
	return ""
}

func (x *GetBusinessIdResponse) GetBusinessId() int32 {
	if x, ok := x.GetResult().(*GetBusinessIdResponse_BusinessId); ok {
		return x.BusinessId
	}
	return 0
}

type isGetBusinessIdResponse_Result interface {
	isGetBusinessIdResponse_Result()
}

type GetBusinessIdResponse_ErrMessage struct {
	ErrMessage string `protobuf:"bytes,2,opt,name=err_message,json=errMessage,proto3,oneof"`
}

type GetBusinessIdResponse_BusinessId struct {
	BusinessId int32 `protobuf:"varint,3,opt,name=business_id,json=businessId,proto3,oneof"`
}

func (*GetBusinessIdResponse_ErrMessage) isGetBusinessIdResponse_Result() {}

func (*GetBusinessIdResponse_BusinessId) isGetBusinessIdResponse_Result() {}

var File_user_service_proto_user_user_proto protoreflect.FileDescriptor

var file_user_service_proto_user_user_proto_rawDesc = []byte{
	0x0a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x88, 0x03, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x22, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x49, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x0e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x28, 0x0a, 0x0c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x63, 0x0a, 0x12, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x65, 0x72,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x41, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x37, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x54, 0x4b, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x0e,
	0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x27, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x4b, 0x48, 0x00,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6b, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x2f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a,
	0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x42, 0x08, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x32, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x42, 0x41, 0x4e, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x58, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x45,
	0x4e, 0x44, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54,
	0x41, 0x4e, 0x54, 0x10, 0x05, 0x32, 0xc9, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_service_proto_user_user_proto_rawDescOnce sync.Once
	file_user_service_proto_user_user_proto_rawDescData = file_user_service_proto_user_user_proto_rawDesc
)

func file_user_service_proto_user_user_proto_rawDescGZIP() []byte {
	file_user_service_proto_user_user_proto_rawDescOnce.Do(func() {
		file_user_service_proto_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_user_user_proto_rawDescData)
	})
	return file_user_service_proto_user_user_proto_rawDescData
}

var file_user_service_proto_user_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_service_proto_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_service_proto_user_user_proto_goTypes = []interface{}{
	(UserStatus)(0),               // 0: user.UserStatus
	(UserRole)(0),                 // 1: user.UserRole
	(*User)(nil),                  // 2: user.User
	(*SuccessMessage)(nil),        // 3: user.SuccessMessage
	(*ErrorMessage)(nil),          // 4: user.ErrorMessage
	(*SignUpUserRequest)(nil),     // 5: user.SignUpUserRequest
	(*SignUpUserResponse)(nil),    // 6: user.SignUpUserResponse
	(*SignInRequest)(nil),         // 7: user.SignInRequest
	(*AuthTK)(nil),                // 8: user.AuthTK
	(*SignInResponse)(nil),        // 9: user.SignInResponse
	(*GetBusinessIdRequest)(nil),  // 10: user.GetBusinessIdRequest
	(*GetBusinessIdResponse)(nil), // 11: user.GetBusinessIdResponse
}
var file_user_service_proto_user_user_proto_depIdxs = []int32{
	0,  // 0: user.User.status:type_name -> user.UserStatus
	1,  // 1: user.User.role:type_name -> user.UserRole
	2,  // 2: user.SignUpUserRequest.user:type_name -> user.User
	4,  // 3: user.SignUpUserResponse.err_message:type_name -> user.ErrorMessage
	4,  // 4: user.SignInResponse.err_message:type_name -> user.ErrorMessage
	8,  // 5: user.SignInResponse.auth_tk:type_name -> user.AuthTK
	5,  // 6: user.UserService.SignUp:input_type -> user.SignUpUserRequest
	7,  // 7: user.UserService.SignIn:input_type -> user.SignInRequest
	10, // 8: user.UserService.GetBusinessId:input_type -> user.GetBusinessIdRequest
	6,  // 9: user.UserService.SignUp:output_type -> user.SignUpUserResponse
	9,  // 10: user.UserService.SignIn:output_type -> user.SignInResponse
	11, // 11: user.UserService.GetBusinessId:output_type -> user.GetBusinessIdResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_user_service_proto_user_user_proto_init() }
func file_user_service_proto_user_user_proto_init() {
	if File_user_service_proto_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthTK); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_service_proto_user_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_user_service_proto_user_user_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*SignInResponse_ErrMessage)(nil),
		(*SignInResponse_AuthTk)(nil),
	}
	file_user_service_proto_user_user_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*GetBusinessIdResponse_ErrMessage)(nil),
		(*GetBusinessIdResponse_BusinessId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_proto_user_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_user_user_proto_goTypes,
		DependencyIndexes: file_user_service_proto_user_user_proto_depIdxs,
		EnumInfos:         file_user_service_proto_user_user_proto_enumTypes,
		MessageInfos:      file_user_service_proto_user_user_proto_msgTypes,
	}.Build()
	File_user_service_proto_user_user_proto = out.File
	file_user_service_proto_user_user_proto_rawDesc = nil
	file_user_service_proto_user_user_proto_goTypes = nil
	file_user_service_proto_user_user_proto_depIdxs = nil
}
