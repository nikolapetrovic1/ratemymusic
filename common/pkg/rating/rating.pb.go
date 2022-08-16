// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pkg/rating/rating.proto

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

type RateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingValue float32 `protobuf:"fixed32,1,opt,name=rating_value,json=ratingValue,proto3" json:"rating_value,omitempty"`
	UserId      int64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Types that are assignable to RatingType:
	//	*RateRequest_AlbumId
	//	*RateRequest_SongId
	RatingType isRateRequest_RatingType `protobuf_oneof:"rating_type"`
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{0}
}

func (x *RateRequest) GetRatingValue() float32 {
	if x != nil {
		return x.RatingValue
	}
	return 0
}

func (x *RateRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (m *RateRequest) GetRatingType() isRateRequest_RatingType {
	if m != nil {
		return m.RatingType
	}
	return nil
}

func (x *RateRequest) GetAlbumId() int64 {
	if x, ok := x.GetRatingType().(*RateRequest_AlbumId); ok {
		return x.AlbumId
	}
	return 0
}

func (x *RateRequest) GetSongId() int64 {
	if x, ok := x.GetRatingType().(*RateRequest_SongId); ok {
		return x.SongId
	}
	return 0
}

type isRateRequest_RatingType interface {
	isRateRequest_RatingType()
}

type RateRequest_AlbumId struct {
	AlbumId int64 `protobuf:"varint,3,opt,name=albumId,proto3,oneof"`
}

type RateRequest_SongId struct {
	SongId int64 `protobuf:"varint,4,opt,name=songId,proto3,oneof"`
}

func (*RateRequest_AlbumId) isRateRequest_RatingType() {}

func (*RateRequest_SongId) isRateRequest_RatingType() {}

type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RateResponse) GetError() string {
	if x != nil {
		return x.Error
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
		mi := &file_pkg_rating_rating_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[2]
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
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{2}
}

func (x *IdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error      string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	RatingData []*RatingData `protobuf:"bytes,3,rep,name=rating_data,json=ratingData,proto3" json:"rating_data,omitempty"`
}

func (x *RatingResponse) Reset() {
	*x = RatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingResponse) ProtoMessage() {}

func (x *RatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingResponse.ProtoReflect.Descriptor instead.
func (*RatingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{3}
}

func (x *RatingResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RatingResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *RatingResponse) GetRatingData() []*RatingData {
	if x != nil {
		return x.RatingData
	}
	return nil
}

type RatingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingValue float32 `protobuf:"fixed32,1,opt,name=rating_value,json=ratingValue,proto3" json:"rating_value,omitempty"`
	UserId      int64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Types that are assignable to RatingType:
	//	*RatingData_AlbumId
	//	*RatingData_SongId
	RatingType isRatingData_RatingType `protobuf_oneof:"rating_type"`
}

func (x *RatingData) Reset() {
	*x = RatingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_rating_rating_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingData) ProtoMessage() {}

func (x *RatingData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_rating_rating_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingData.ProtoReflect.Descriptor instead.
func (*RatingData) Descriptor() ([]byte, []int) {
	return file_pkg_rating_rating_proto_rawDescGZIP(), []int{4}
}

func (x *RatingData) GetRatingValue() float32 {
	if x != nil {
		return x.RatingValue
	}
	return 0
}

func (x *RatingData) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (m *RatingData) GetRatingType() isRatingData_RatingType {
	if m != nil {
		return m.RatingType
	}
	return nil
}

func (x *RatingData) GetAlbumId() int64 {
	if x, ok := x.GetRatingType().(*RatingData_AlbumId); ok {
		return x.AlbumId
	}
	return 0
}

func (x *RatingData) GetSongId() int64 {
	if x, ok := x.GetRatingType().(*RatingData_SongId); ok {
		return x.SongId
	}
	return 0
}

type isRatingData_RatingType interface {
	isRatingData_RatingType()
}

type RatingData_AlbumId struct {
	AlbumId int64 `protobuf:"varint,3,opt,name=albumId,proto3,oneof"`
}

type RatingData_SongId struct {
	SongId int64 `protobuf:"varint,4,opt,name=songId,proto3,oneof"`
}

func (*RatingData_AlbumId) isRatingData_RatingType() {}

func (*RatingData_SongId) isRatingData_RatingType() {}

var File_pkg_rating_rating_proto protoreflect.FileDescriptor

var file_pkg_rating_rating_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x6f, 0x6e,
	0x67, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x73, 0x6f, 0x6e,
	0x67, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73, 0x0a,
	0x0e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a,
	0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x22, 0x8d, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x6f, 0x6e,
	0x67, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x73, 0x6f, 0x6e,
	0x67, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x32, 0xb4, 0x02, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x11, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x11, 0x2e,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x11, 0x2e, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x52, 0x61, 0x74, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x12, 0x13, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x13, 0x2e, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_rating_rating_proto_rawDescOnce sync.Once
	file_pkg_rating_rating_proto_rawDescData = file_pkg_rating_rating_proto_rawDesc
)

func file_pkg_rating_rating_proto_rawDescGZIP() []byte {
	file_pkg_rating_rating_proto_rawDescOnce.Do(func() {
		file_pkg_rating_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_rating_rating_proto_rawDescData)
	})
	return file_pkg_rating_rating_proto_rawDescData
}

var file_pkg_rating_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_rating_rating_proto_goTypes = []interface{}{
	(*RateRequest)(nil),    // 0: rating.RateRequest
	(*RateResponse)(nil),   // 1: rating.RateResponse
	(*IdRequest)(nil),      // 2: rating.IdRequest
	(*RatingResponse)(nil), // 3: rating.RatingResponse
	(*RatingData)(nil),     // 4: rating.RatingData
}
var file_pkg_rating_rating_proto_depIdxs = []int32{
	4, // 0: rating.RatingResponse.rating_data:type_name -> rating.RatingData
	2, // 1: rating.RatingService.FindByUser:input_type -> rating.IdRequest
	2, // 2: rating.RatingService.FindBySong:input_type -> rating.IdRequest
	2, // 3: rating.RatingService.FindByAlbum:input_type -> rating.IdRequest
	0, // 4: rating.RatingService.RateSong:input_type -> rating.RateRequest
	0, // 5: rating.RatingService.RateAlbum:input_type -> rating.RateRequest
	3, // 6: rating.RatingService.FindByUser:output_type -> rating.RatingResponse
	3, // 7: rating.RatingService.FindBySong:output_type -> rating.RatingResponse
	3, // 8: rating.RatingService.FindByAlbum:output_type -> rating.RatingResponse
	1, // 9: rating.RatingService.RateSong:output_type -> rating.RateResponse
	1, // 10: rating.RatingService.RateAlbum:output_type -> rating.RateResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_rating_rating_proto_init() }
func file_pkg_rating_rating_proto_init() {
	if File_pkg_rating_rating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_rating_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRequest); i {
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
		file_pkg_rating_rating_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
		file_pkg_rating_rating_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_rating_rating_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingResponse); i {
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
		file_pkg_rating_rating_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingData); i {
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
	file_pkg_rating_rating_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RateRequest_AlbumId)(nil),
		(*RateRequest_SongId)(nil),
	}
	file_pkg_rating_rating_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*RatingData_AlbumId)(nil),
		(*RatingData_SongId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_rating_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_rating_rating_proto_goTypes,
		DependencyIndexes: file_pkg_rating_rating_proto_depIdxs,
		MessageInfos:      file_pkg_rating_rating_proto_msgTypes,
	}.Build()
	File_pkg_rating_rating_proto = out.File
	file_pkg_rating_rating_proto_rawDesc = nil
	file_pkg_rating_rating_proto_goTypes = nil
	file_pkg_rating_rating_proto_depIdxs = nil
}
