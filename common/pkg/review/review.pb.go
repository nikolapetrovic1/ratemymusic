// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pkg/review/review.proto

package pb

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

type DataType int32

const (
	DataType_album DataType = 0
	DataType_song  DataType = 1
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "album",
		1: "song",
	}
	DataType_value = map[string]int32{
		"album": 0,
		"song":  1,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_review_review_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_pkg_review_review_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_review_review_proto_rawDescGZIP(), []int{0}
}

type ReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error   string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Reviews []*ReviewData `protobuf:"bytes,3,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *ReviewResponse) Reset() {
	*x = ReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_review_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewResponse) ProtoMessage() {}

func (x *ReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_review_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewResponse.ProtoReflect.Descriptor instead.
func (*ReviewResponse) Descriptor() ([]byte, []int) {
	return file_pkg_review_review_proto_rawDescGZIP(), []int{0}
}

func (x *ReviewResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ReviewResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ReviewResponse) GetReviews() []*ReviewData {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type ReviewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TypeId int64  `protobuf:"varint,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Text   string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Type   string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ReviewData) Reset() {
	*x = ReviewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_review_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewData) ProtoMessage() {}

func (x *ReviewData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_review_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewData.ProtoReflect.Descriptor instead.
func (*ReviewData) Descriptor() ([]byte, []int) {
	return file_pkg_review_review_proto_rawDescGZIP(), []int{1}
}

func (x *ReviewData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReviewData) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ReviewData) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *ReviewData) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ReviewData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_review_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_review_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_pkg_review_review_proto_rawDescGZIP(), []int{2}
}

func (x *IdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SongId int64 `protobuf:"varint,2,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
}

func (x *UserSongRequest) Reset() {
	*x = UserSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_review_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSongRequest) ProtoMessage() {}

func (x *UserSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_review_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSongRequest.ProtoReflect.Descriptor instead.
func (*UserSongRequest) Descriptor() ([]byte, []int) {
	return file_pkg_review_review_proto_rawDescGZIP(), []int{3}
}

func (x *UserSongRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserSongRequest) GetSongId() int64 {
	if x != nil {
		return x.SongId
	}
	return 0
}

type UserAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AlbumId int64 `protobuf:"varint,2,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *UserAlbumRequest) Reset() {
	*x = UserAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_review_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAlbumRequest) ProtoMessage() {}

func (x *UserAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_review_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAlbumRequest.ProtoReflect.Descriptor instead.
func (*UserAlbumRequest) Descriptor() ([]byte, []int) {
	return file_pkg_review_review_proto_rawDescGZIP(), []int{4}
}

func (x *UserAlbumRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserAlbumRequest) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

var File_pkg_review_review_proto protoreflect.FileDescriptor

var file_pkg_review_review_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x6c, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22,
	0x76, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6f, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x46, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49,
	0x64, 0x2a, 0x1f, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67,
	0x10, 0x01, 0x32, 0xe6, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x11, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x12, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x12,
	0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_review_review_proto_rawDescOnce sync.Once
	file_pkg_review_review_proto_rawDescData = file_pkg_review_review_proto_rawDesc
)

func file_pkg_review_review_proto_rawDescGZIP() []byte {
	file_pkg_review_review_proto_rawDescOnce.Do(func() {
		file_pkg_review_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_review_review_proto_rawDescData)
	})
	return file_pkg_review_review_proto_rawDescData
}

var file_pkg_review_review_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_review_review_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_review_review_proto_goTypes = []interface{}{
	(DataType)(0),            // 0: review.DataType
	(*ReviewResponse)(nil),   // 1: review.ReviewResponse
	(*ReviewData)(nil),       // 2: review.ReviewData
	(*IdRequest)(nil),        // 3: review.IdRequest
	(*UserSongRequest)(nil),  // 4: review.UserSongRequest
	(*UserAlbumRequest)(nil), // 5: review.UserAlbumRequest
}
var file_pkg_review_review_proto_depIdxs = []int32{
	2, // 0: review.ReviewResponse.reviews:type_name -> review.ReviewData
	3, // 1: review.ReviewService.FindByUser:input_type -> review.IdRequest
	3, // 2: review.ReviewService.FindBySong:input_type -> review.IdRequest
	3, // 3: review.ReviewService.FindByAlbum:input_type -> review.IdRequest
	4, // 4: review.ReviewService.FindByUserSong:input_type -> review.UserSongRequest
	5, // 5: review.ReviewService.FindByUserAlbum:input_type -> review.UserAlbumRequest
	2, // 6: review.ReviewService.CreateReview:input_type -> review.ReviewData
	2, // 7: review.ReviewService.Update:input_type -> review.ReviewData
	3, // 8: review.ReviewService.Delete:input_type -> review.IdRequest
	1, // 9: review.ReviewService.FindByUser:output_type -> review.ReviewResponse
	1, // 10: review.ReviewService.FindBySong:output_type -> review.ReviewResponse
	1, // 11: review.ReviewService.FindByAlbum:output_type -> review.ReviewResponse
	2, // 12: review.ReviewService.FindByUserSong:output_type -> review.ReviewData
	2, // 13: review.ReviewService.FindByUserAlbum:output_type -> review.ReviewData
	2, // 14: review.ReviewService.CreateReview:output_type -> review.ReviewData
	2, // 15: review.ReviewService.Update:output_type -> review.ReviewData
	2, // 16: review.ReviewService.Delete:output_type -> review.ReviewData
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_review_review_proto_init() }
func file_pkg_review_review_proto_init() {
	if File_pkg_review_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_review_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewResponse); i {
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
		file_pkg_review_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewData); i {
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
		file_pkg_review_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_pkg_review_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSongRequest); i {
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
		file_pkg_review_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAlbumRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_review_review_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_review_review_proto_goTypes,
		DependencyIndexes: file_pkg_review_review_proto_depIdxs,
		EnumInfos:         file_pkg_review_review_proto_enumTypes,
		MessageInfos:      file_pkg_review_review_proto_msgTypes,
	}.Build()
	File_pkg_review_review_proto = out.File
	file_pkg_review_review_proto_rawDesc = nil
	file_pkg_review_review_proto_goTypes = nil
	file_pkg_review_review_proto_depIdxs = nil
}
