// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/account/proto/account/account.proto

package srv_account

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	// gorm.types.UUID id                      = 1  [(gorm.field).tag = { type: "uuid", primary_key: true, not_null: true, unique: true }]; // primary key
	// string id                               = 1  [(gorm.field).tag = { type: "uuid", primary_key: true, default: "uuid_generate_v4()" }]; // primary key
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	Username             string               `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            string               `protobuf:"bytes,8,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string               `protobuf:"bytes,9,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string               `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Profile              *Profile             `protobuf:"bytes,11,opt,name=Profile,proto3" json:"Profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *User) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

// FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223
// with Workaround in .override.go
type UserRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *UserRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *UserRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *UserRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type UserResponse struct {
	Result               *User    `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetResult() *User {
	if m != nil {
		return m.Result
	}
	return nil
}

type UserExistResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserExistResponse) Reset()         { *m = UserExistResponse{} }
func (m *UserExistResponse) String() string { return proto.CompactTextString(m) }
func (*UserExistResponse) ProtoMessage()    {}
func (*UserExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{3}
}

func (m *UserExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserExistResponse.Unmarshal(m, b)
}
func (m *UserExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserExistResponse.Marshal(b, m, deterministic)
}
func (m *UserExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserExistResponse.Merge(m, src)
}
func (m *UserExistResponse) XXX_Size() int {
	return xxx_messageInfo_UserExistResponse.Size(m)
}
func (m *UserExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserExistResponse proto.InternalMessageInfo

func (m *UserExistResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223
// Workaround in .override.go
type UserListQuery struct {
	Limit                *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Username             *wrappers.StringValue `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            *wrappers.StringValue `protobuf:"bytes,5,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             *wrappers.StringValue `protobuf:"bytes,6,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserListQuery) Reset()         { *m = UserListQuery{} }
func (m *UserListQuery) String() string { return proto.CompactTextString(m) }
func (*UserListQuery) ProtoMessage()    {}
func (*UserListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{4}
}

func (m *UserListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListQuery.Unmarshal(m, b)
}
func (m *UserListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListQuery.Marshal(b, m, deterministic)
}
func (m *UserListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListQuery.Merge(m, src)
}
func (m *UserListQuery) XXX_Size() int {
	return xxx_messageInfo_UserListQuery.Size(m)
}
func (m *UserListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_UserListQuery proto.InternalMessageInfo

func (m *UserListQuery) GetLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *UserListQuery) GetPage() *wrappers.UInt32Value {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *UserListQuery) GetSort() *wrappers.StringValue {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *UserListQuery) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *UserListQuery) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *UserListQuery) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *UserListQuery) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type UserListResponse struct {
	Results              []*User  `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{5}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetResults() []*User {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *UserListResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

// Profile Entity
type Profile struct {
	// gorm.types.UUID id                      = 1  [(gorm.field).tag = { type: "uuid", primary_key: true, not_null: true, unique: true }]; // primary key
	// string id                               = 1  [(gorm.field).tag = { type: "uuid", primary_key: true, default: "uuid_generate_v4()" }]; // primary key
	Id        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	Tz        string               `protobuf:"bytes,5,opt,name=tz,proto3" json:"tz,omitempty"`
	Avatar    string               `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	// FIXME: https://github.com/jinzhu/gorm/issues/1978
	Gender string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	// GenderType gender                       = 7; //    GenderType `gorm:"not null;type:ENUM('M', 'F')"`
	Birthday             *timestamp.Timestamp `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Age                  uint32               `protobuf:"varint,9,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{6}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Profile) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Profile) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Profile) GetTz() string {
	if m != nil {
		return m.Tz
	}
	return ""
}

func (m *Profile) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Profile) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Profile) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *Profile) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

// FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223
// Workaround in .override.go
type ProfileRequest struct {
	Id     *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId *wrappers.StringValue `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Tz     *wrappers.StringValue `protobuf:"bytes,3,opt,name=tz,proto3" json:"tz,omitempty"`
	Avatar *wrappers.StringValue `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender *wrappers.StringValue `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	// GenderType gender = 5; // enum // [(validate.rules).enum.defined_only = true]; // FIXME
	Birthday             *timestamp.Timestamp `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProfileRequest) Reset()         { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()    {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{7}
}

func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileRequest.Unmarshal(m, b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
}
func (m *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(m, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ProfileRequest.Size(m)
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ProfileRequest) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *ProfileRequest) GetTz() *wrappers.StringValue {
	if m != nil {
		return m.Tz
	}
	return nil
}

func (m *ProfileRequest) GetAvatar() *wrappers.StringValue {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *ProfileRequest) GetGender() *wrappers.StringValue {
	if m != nil {
		return m.Gender
	}
	return nil
}

func (m *ProfileRequest) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

type ProfileResponse struct {
	Result               *Profile `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileResponse) Reset()         { *m = ProfileResponse{} }
func (m *ProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ProfileResponse) ProtoMessage()    {}
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{8}
}

func (m *ProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileResponse.Unmarshal(m, b)
}
func (m *ProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileResponse.Marshal(b, m, deterministic)
}
func (m *ProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileResponse.Merge(m, src)
}
func (m *ProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ProfileResponse.Size(m)
}
func (m *ProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileResponse proto.InternalMessageInfo

func (m *ProfileResponse) GetResult() *Profile {
	if m != nil {
		return m.Result
	}
	return nil
}

type ProfileListQuery struct {
	Limit                *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	UserId               *wrappers.StringValue `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	Gender               *wrappers.StringValue `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProfileListQuery) Reset()         { *m = ProfileListQuery{} }
func (m *ProfileListQuery) String() string { return proto.CompactTextString(m) }
func (*ProfileListQuery) ProtoMessage()    {}
func (*ProfileListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{9}
}

func (m *ProfileListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileListQuery.Unmarshal(m, b)
}
func (m *ProfileListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileListQuery.Marshal(b, m, deterministic)
}
func (m *ProfileListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileListQuery.Merge(m, src)
}
func (m *ProfileListQuery) XXX_Size() int {
	return xxx_messageInfo_ProfileListQuery.Size(m)
}
func (m *ProfileListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileListQuery proto.InternalMessageInfo

func (m *ProfileListQuery) GetLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *ProfileListQuery) GetPage() *wrappers.UInt32Value {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ProfileListQuery) GetSort() *wrappers.StringValue {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *ProfileListQuery) GetUserId() *wrappers.StringValue {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *ProfileListQuery) GetGender() *wrappers.StringValue {
	if m != nil {
		return m.Gender
	}
	return nil
}

type ProfileListResponse struct {
	Results              []*Profile `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProfileListResponse) Reset()         { *m = ProfileListResponse{} }
func (m *ProfileListResponse) String() string { return proto.CompactTextString(m) }
func (*ProfileListResponse) ProtoMessage()    {}
func (*ProfileListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{10}
}

func (m *ProfileListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileListResponse.Unmarshal(m, b)
}
func (m *ProfileListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileListResponse.Marshal(b, m, deterministic)
}
func (m *ProfileListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileListResponse.Merge(m, src)
}
func (m *ProfileListResponse) XXX_Size() int {
	return xxx_messageInfo_ProfileListResponse.Size(m)
}
func (m *ProfileListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileListResponse proto.InternalMessageInfo

func (m *ProfileListResponse) GetResults() []*Profile {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ProfileListResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "srv.account.User")
	proto.RegisterType((*UserRequest)(nil), "srv.account.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "srv.account.UserResponse")
	proto.RegisterType((*UserExistResponse)(nil), "srv.account.UserExistResponse")
	proto.RegisterType((*UserListQuery)(nil), "srv.account.UserListQuery")
	proto.RegisterType((*UserListResponse)(nil), "srv.account.UserListResponse")
	proto.RegisterType((*Profile)(nil), "srv.account.Profile")
	proto.RegisterType((*ProfileRequest)(nil), "srv.account.ProfileRequest")
	proto.RegisterType((*ProfileResponse)(nil), "srv.account.ProfileResponse")
	proto.RegisterType((*ProfileListQuery)(nil), "srv.account.ProfileListQuery")
	proto.RegisterType((*ProfileListResponse)(nil), "srv.account.ProfileListResponse")
}

func init() {
	proto.RegisterFile("srv/account/proto/account/account.proto", fileDescriptor_f36848fb0b6d28e2)
}

var fileDescriptor_f36848fb0b6d28e2 = []byte{
	// 1068 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x3f, 0x71, 0x9c, 0x13, 0x5a, 0xb2, 0x43, 0xb5, 0x38, 0xe9, 0xcf, 0x46, 0xd1, 0x0a,
	0xc2, 0xd2, 0x38, 0x25, 0x15, 0x88, 0x2d, 0x50, 0x5a, 0x77, 0x17, 0x54, 0x09, 0x10, 0x78, 0xe9,
	0x4a, 0x2c, 0x82, 0xae, 0x13, 0x4f, 0xb3, 0x96, 0x1c, 0xdb, 0x8c, 0xc7, 0xa1, 0x2d, 0x42, 0xe2,
	0x11, 0xb8, 0xe4, 0x31, 0x10, 0x5c, 0x25, 0x77, 0xbc, 0x06, 0xe2, 0x7a, 0x9f, 0x61, 0x15, 0x2e,
	0x40, 0x63, 0x3b, 0x8e, 0xd3, 0x24, 0x34, 0x6d, 0xe0, 0x62, 0xaf, 0xda, 0x99, 0x73, 0xbe, 0x93,
	0x99, 0xef, 0x3b, 0xe7, 0x1b, 0xc3, 0x6b, 0x3e, 0xe9, 0xd6, 0x8d, 0x56, 0xcb, 0x0d, 0x1c, 0x5a,
	0xf7, 0x88, 0x4b, 0xdd, 0x64, 0x15, 0xff, 0x55, 0xc3, 0x5d, 0x94, 0xf7, 0x49, 0x57, 0x8d, 0xb7,
	0x4a, 0x1b, 0x6d, 0xd7, 0x6d, 0xdb, 0x38, 0x02, 0x34, 0x83, 0x93, 0xfa, 0x77, 0xc4, 0xf0, 0x3c,
	0x4c, 0xfc, 0x28, 0xb9, 0x74, 0xeb, 0x62, 0x9c, 0x5a, 0x1d, 0xec, 0x53, 0xa3, 0xe3, 0xc5, 0x09,
	0x3b, 0x6d, 0x8b, 0x3e, 0x09, 0x9a, 0x6a, 0xcb, 0xed, 0xd4, 0x2d, 0xe7, 0xc4, 0x6d, 0xda, 0xee,
	0xa9, 0xeb, 0x61, 0x27, 0x42, 0xb4, 0x6a, 0x6d, 0xec, 0xd4, 0xda, 0x2e, 0xe9, 0xd4, 0x5d, 0x8f,
	0x5a, 0xae, 0xe3, 0xd7, 0xd9, 0x22, 0xc6, 0xee, 0xa7, 0xb0, 0xd8, 0xe9, 0xba, 0x67, 0x1e, 0x71,
	0x4f, 0xcf, 0xd2, 0xc8, 0xae, 0x61, 0x5b, 0xa6, 0x41, 0x71, 0x7d, 0xe2, 0x9f, 0xa8, 0x44, 0xe5,
	0x0f, 0x01, 0xc4, 0x23, 0x1f, 0x13, 0xb4, 0x01, 0xbc, 0x65, 0x2a, 0x5c, 0x99, 0xab, 0xe6, 0xb4,
	0xe5, 0x7e, 0xaf, 0x08, 0x20, 0x23, 0x31, 0x08, 0x2c, 0xb3, 0xca, 0xe9, 0xbc, 0x65, 0x22, 0x0d,
	0x72, 0x2d, 0x82, 0x0d, 0x8a, 0xcd, 0x7d, 0xaa, 0xf0, 0x65, 0xae, 0x9a, 0x6f, 0x94, 0xd4, 0xe8,
	0x72, 0xea, 0xf0, 0x72, 0xea, 0x17, 0xc3, 0xcb, 0x69, 0x72, 0xbf, 0x57, 0x14, 0x81, 0xdf, 0xe3,
	0xf4, 0x11, 0x8c, 0xd5, 0x08, 0x3c, 0x33, 0xae, 0x21, 0x5c, 0xa5, 0x46, 0x02, 0x43, 0x0f, 0x21,
	0x67, 0x62, 0x1b, 0x47, 0x35, 0xc4, 0x4b, 0x6b, 0xac, 0xf5, 0x7b, 0x45, 0x05, 0x6e, 0xea, 0x2b,
	0x96, 0x79, 0x7a, 0x1c, 0xf8, 0x98, 0xf8, 0xc7, 0x31, 0xfc, 0xd8, 0xa0, 0xfa, 0xa8, 0x14, 0x7a,
	0x15, 0x64, 0x16, 0x76, 0x8c, 0x0e, 0x56, 0xb2, 0x21, 0x0b, 0xd0, 0xef, 0x15, 0x25, 0x10, 0x15,
	0x73, 0x8f, 0xd3, 0x93, 0x18, 0x7a, 0x1d, 0x72, 0x27, 0x16, 0xf1, 0xe9, 0xa7, 0x2c, 0x51, 0x0e,
	0x13, 0xf3, 0xfd, 0x5e, 0x31, 0x0b, 0x19, 0xe5, 0x6f, 0x8e, 0x1d, 0x35, 0x89, 0xa2, 0xdb, 0x20,
	0xdb, 0x46, 0x9c, 0x99, 0x0b, 0x33, 0x47, 0x37, 0x4a, 0x22, 0x68, 0x03, 0x32, 0xb8, 0x63, 0x58,
	0xb6, 0x02, 0x17, 0x52, 0xa2, 0x6d, 0xa4, 0x42, 0xf6, 0x33, 0xe2, 0x9e, 0x58, 0x36, 0x56, 0xf2,
	0xe1, 0x75, 0x57, 0xd4, 0x54, 0x03, 0xaa, 0x71, 0x4c, 0x1f, 0x26, 0xed, 0xe4, 0x9e, 0x69, 0x5c,
	0xbf, 0x57, 0xe4, 0x65, 0xae, 0xf2, 0x94, 0x87, 0x3c, 0x13, 0x57, 0xc7, 0xdf, 0x06, 0xd8, 0xa7,
	0xe8, 0xed, 0x44, 0xe3, 0x7c, 0x63, 0x6d, 0x82, 0xb4, 0x07, 0x94, 0x58, 0x4e, 0xfb, 0xa1, 0x61,
	0x07, 0x58, 0x93, 0x07, 0x5a, 0x86, 0x08, 0xbf, 0x70, 0x91, 0xf6, 0x8f, 0x52, 0xdc, 0xf0, 0x73,
	0xa0, 0xcb, 0x03, 0x6d, 0x9d, 0xac, 0x16, 0x44, 0xa5, 0x50, 0xfd, 0x91, 0x6f, 0xa0, 0x6f, 0xbe,
	0x32, 0x6a, 0xe7, 0x5b, 0xb5, 0xbb, 0xc7, 0xb5, 0xaf, 0xbf, 0xdf, 0xde, 0x7c, 0xf3, 0xad, 0x1f,
	0x6e, 0xa7, 0xf8, 0x3c, 0x48, 0xf3, 0x29, 0xcc, 0x51, 0x3c, 0x3b, 0xd0, 0x44, 0xc2, 0x17, 0x84,
	0x34, 0xd3, 0xfb, 0x29, 0xa6, 0xc5, 0xab, 0xd4, 0x18, 0xc9, 0xf0, 0xee, 0x50, 0x86, 0xcc, 0xdc,
	0xf8, 0xc7, 0x43, 0x8d, 0x2a, 0xbb, 0xf0, 0x62, 0xc4, 0xb3, 0xef, 0xb9, 0x8e, 0xcf, 0x9a, 0x44,
	0x22, 0xd8, 0x0f, 0x6c, 0x1a, 0x93, 0x7d, 0x63, 0x4c, 0xb2, 0x30, 0x35, 0x4e, 0xd8, 0x11, 0x9e,
	0x69, 0x5c, 0x65, 0x0b, 0x6e, 0xb0, 0xcd, 0xfb, 0xa7, 0x96, 0x4f, 0x93, 0x22, 0x37, 0xc7, 0x8a,
	0xc8, 0xe3, 0x88, 0x3f, 0x05, 0x58, 0x62, 0x90, 0x8f, 0x2d, 0x9f, 0x7e, 0x1e, 0x60, 0x72, 0x86,
	0xde, 0x87, 0x8c, 0x6d, 0x75, 0x2c, 0x3a, 0x53, 0xdf, 0xa3, 0x43, 0x87, 0x6e, 0x37, 0xa2, 0x0b,
	0xe4, 0x06, 0x9a, 0x74, 0x47, 0x54, 0xd8, 0x70, 0x47, 0x28, 0x74, 0x17, 0x44, 0xcf, 0x68, 0xcf,
	0xd6, 0x37, 0x8d, 0x66, 0xd7, 0xbf, 0xc3, 0x57, 0x39, 0x3d, 0x84, 0xa0, 0x2d, 0x10, 0x7d, 0x97,
	0xd0, 0x79, 0xd4, 0xd3, 0xc3, 0xcc, 0xb1, 0x86, 0x12, 0xff, 0xcf, 0x86, 0xca, 0xfc, 0x07, 0x0d,
	0x25, 0x2d, 0xd8, 0x50, 0xd9, 0x6b, 0x34, 0xd4, 0x23, 0x28, 0x0c, 0xd5, 0x4d, 0xfa, 0xe1, 0x0d,
	0xc8, 0x46, 0x1d, 0xe0, 0x2b, 0x5c, 0x59, 0x98, 0xde, 0x55, 0xc3, 0x0c, 0xb4, 0x02, 0x19, 0xea,
	0x52, 0xc3, 0x0e, 0xf5, 0x5c, 0xd2, 0xa3, 0x45, 0xd4, 0x3a, 0xbf, 0x09, 0x89, 0xa3, 0x3c, 0x37,
	0xae, 0xff, 0xe5, 0xd5, 0x5c, 0xff, 0x56, 0xbf, 0x57, 0x5c, 0x85, 0xa2, 0xfe, 0x0a, 0x73, 0x7d,
	0x2f, 0xba, 0xeb, 0x2c, 0xe3, 0x5f, 0x06, 0x9e, 0x9e, 0x87, 0x8d, 0x92, 0xd3, 0x79, 0x7a, 0xce,
	0xc6, 0xce, 0xe8, 0x1a, 0xd4, 0x20, 0xa1, 0xf0, 0x39, 0x3d, 0x5e, 0xb1, 0xfd, 0x36, 0x76, 0x4c,
	0x4c, 0xa2, 0xe7, 0x41, 0x8f, 0x57, 0x68, 0x0f, 0xe4, 0xa6, 0x45, 0xe8, 0x13, 0xd3, 0x38, 0x0b,
	0xdf, 0x83, 0x79, 0x6f, 0x97, 0xa0, 0x90, 0x02, 0x02, 0x9b, 0x3c, 0xf6, 0x44, 0x2c, 0x69, 0x12,
	0xf3, 0xee, 0x02, 0xa7, 0xb3, 0xad, 0xb4, 0x97, 0xff, 0xc5, 0xc3, 0xf2, 0xd0, 0xeb, 0x17, 0xb4,
	0xf3, 0x5d, 0x90, 0xd8, 0xb4, 0x1c, 0x9a, 0x73, 0x99, 0xf9, 0x08, 0x1b, 0xa3, 0xd0, 0x66, 0xc8,
	0xd8, 0x3c, 0xd3, 0xce, 0xf8, 0xdc, 0x4d, 0xf8, 0x14, 0xe7, 0xfe, 0xb5, 0x9f, 0xd8, 0xaf, 0xc5,
	0xbc, 0xef, 0x27, 0xbc, 0xcf, 0x33, 0xcc, 0xf9, 0x81, 0x26, 0x13, 0x49, 0xe7, 0x3e, 0xd1, 0xb9,
	0x0f, 0xa7, 0x4a, 0x24, 0x5d, 0x2e, 0xd1, 0x40, 0xcb, 0xfc, 0xca, 0xf1, 0xef, 0xa4, 0x24, 0xaa,
	0xdc, 0x83, 0x97, 0x12, 0xf2, 0xe3, 0x71, 0xdc, 0xbc, 0xe0, 0xf1, 0xd3, 0x9f, 0xe5, 0x31, 0xd3,
	0xfe, 0x9d, 0x87, 0x42, 0x1c, 0x78, 0x3e, 0x7d, 0x7b, 0xd4, 0x39, 0xe2, 0xb5, 0x3a, 0x67, 0x71,
	0x2d, 0x2b, 0x8f, 0xe1, 0xe5, 0x14, 0x85, 0x89, 0x1a, 0xea, 0x45, 0x73, 0x9c, 0xf1, 0x95, 0x74,
	0xb9, 0x3f, 0x36, 0x7e, 0x16, 0xa2, 0xaf, 0xa6, 0x07, 0x98, 0x74, 0xad, 0x16, 0x7b, 0x50, 0x32,
	0xe1, 0xc3, 0x8c, 0x94, 0x49, 0xbf, 0x8d, 0x26, 0xb1, 0xb4, 0x31, 0x11, 0x19, 0x7b, 0xca, 0x2b,
	0x2f, 0xa0, 0x03, 0x10, 0xd9, 0x79, 0x51, 0x69, 0x22, 0x33, 0xe9, 0x84, 0xd2, 0xfa, 0xd4, 0x58,
	0xaa, 0xc8, 0x7b, 0x20, 0x7c, 0x84, 0xff, 0xed, 0x1c, 0xc5, 0x29, 0x91, 0x04, 0xfd, 0x01, 0x48,
	0x07, 0xa1, 0x29, 0x2f, 0x50, 0xe0, 0x28, 0x74, 0xe4, 0x05, 0x0a, 0xdc, 0x0b, 0x7d, 0xf7, 0x9a,
	0x05, 0x1a, 0x4f, 0xb9, 0xc4, 0x04, 0x87, 0xea, 0x1c, 0xc6, 0xc4, 0xae, 0x4f, 0xd3, 0x7b, 0xc4,
	0x6d, 0x79, 0x56, 0x38, 0x75, 0x3c, 0x2d, 0xa2, 0x77, 0x75, 0x6a, 0xe7, 0xc4, 0xc7, 0x5b, 0x9b,
	0x1e, 0x4c, 0x6a, 0xdc, 0x4f, 0x48, 0x5e, 0xa4, 0x4c, 0x53, 0x0a, 0x07, 0x62, 0xfb, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x58, 0xb7, 0xaa, 0x1c, 0x95, 0x0e, 0x00, 0x00,
}
