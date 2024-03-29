// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: pkg/musician/musician.proto

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

type SongData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Duration   int32  `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	MusicianID int64  `protobuf:"varint,4,opt,name=musicianID,proto3" json:"musicianID,omitempty"`
}

func (x *SongData) Reset() {
	*x = SongData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongData) ProtoMessage() {}

func (x *SongData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongData.ProtoReflect.Descriptor instead.
func (*SongData) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{0}
}

func (x *SongData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SongData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SongData) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *SongData) GetMusicianID() int64 {
	if x != nil {
		return x.MusicianID
	}
	return 0
}

type MusicianData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MusicianName string      `protobuf:"bytes,2,opt,name=musician_name,json=musicianName,proto3" json:"musician_name,omitempty"`
	Name         string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Surname      string      `protobuf:"bytes,4,opt,name=surname,proto3" json:"surname,omitempty"`
	Songs        []*SongData `protobuf:"bytes,5,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *MusicianData) Reset() {
	*x = MusicianData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MusicianData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MusicianData) ProtoMessage() {}

func (x *MusicianData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MusicianData.ProtoReflect.Descriptor instead.
func (*MusicianData) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{1}
}

func (x *MusicianData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MusicianData) GetMusicianName() string {
	if x != nil {
		return x.MusicianName
	}
	return ""
}

func (x *MusicianData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MusicianData) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *MusicianData) GetSongs() []*SongData {
	if x != nil {
		return x.Songs
	}
	return nil
}

type FindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneRequest) Reset() {
	*x = FindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneRequest) ProtoMessage() {}

func (x *FindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneRequest.ProtoReflect.Descriptor instead.
func (*FindOneRequest) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{2}
}

func (x *FindOneRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64         `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data   *MusicianData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FindOneResponse) Reset() {
	*x = FindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneResponse) ProtoMessage() {}

func (x *FindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneResponse.ProtoReflect.Descriptor instead.
func (*FindOneResponse) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *FindOneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *FindOneResponse) GetData() *MusicianData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{4}
}

func (x *SearchRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Musicians []*MusicianData `protobuf:"bytes,1,rep,name=musicians,proto3" json:"musicians,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{5}
}

func (x *SearchResponse) GetMusicians() []*MusicianData {
	if x != nil {
		return x.Musicians
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_musician_musician_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_musician_musician_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_pkg_musician_musician_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_musician_musician_proto protoreflect.FileDescriptor

var file_pkg_musician_musician_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x2f, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x22, 0x6a, 0x0a, 0x08, 0x53, 0x6f, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x49, 0x44, 0x22,
	0x97, 0x01, 0x0a, 0x0c, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x0f, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x42, 0x0a, 0x0e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x09, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x73, 0x22,
	0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0xc7, 0x02, 0x0a, 0x0f, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12,
	0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e,
	0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x12,
	0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x12, 0x13,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x12, 0x13, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_musician_musician_proto_rawDescOnce sync.Once
	file_pkg_musician_musician_proto_rawDescData = file_pkg_musician_musician_proto_rawDesc
)

func file_pkg_musician_musician_proto_rawDescGZIP() []byte {
	file_pkg_musician_musician_proto_rawDescOnce.Do(func() {
		file_pkg_musician_musician_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_musician_musician_proto_rawDescData)
	})
	return file_pkg_musician_musician_proto_rawDescData
}

var file_pkg_musician_musician_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_musician_musician_proto_goTypes = []interface{}{
	(*SongData)(nil),        // 0: auth.SongData
	(*MusicianData)(nil),    // 1: auth.MusicianData
	(*FindOneRequest)(nil),  // 2: auth.FindOneRequest
	(*FindOneResponse)(nil), // 3: auth.FindOneResponse
	(*SearchRequest)(nil),   // 4: auth.SearchRequest
	(*SearchResponse)(nil),  // 5: auth.SearchResponse
	(*DeleteRequest)(nil),   // 6: auth.DeleteRequest
	(*DeleteResponse)(nil),  // 7: auth.DeleteResponse
}
var file_pkg_musician_musician_proto_depIdxs = []int32{
	0, // 0: auth.MusicianData.songs:type_name -> auth.SongData
	1, // 1: auth.FindOneResponse.data:type_name -> auth.MusicianData
	1, // 2: auth.SearchResponse.musicians:type_name -> auth.MusicianData
	2, // 3: auth.MusicianService.FindOne:input_type -> auth.FindOneRequest
	1, // 4: auth.MusicianService.CreateMusician:input_type -> auth.MusicianData
	1, // 5: auth.MusicianService.UpdateMusician:input_type -> auth.MusicianData
	6, // 6: auth.MusicianService.DeleteMusician:input_type -> auth.DeleteRequest
	4, // 7: auth.MusicianService.SearchMusician:input_type -> auth.SearchRequest
	3, // 8: auth.MusicianService.FindOne:output_type -> auth.FindOneResponse
	3, // 9: auth.MusicianService.CreateMusician:output_type -> auth.FindOneResponse
	3, // 10: auth.MusicianService.UpdateMusician:output_type -> auth.FindOneResponse
	7, // 11: auth.MusicianService.DeleteMusician:output_type -> auth.DeleteResponse
	5, // 12: auth.MusicianService.SearchMusician:output_type -> auth.SearchResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_musician_musician_proto_init() }
func file_pkg_musician_musician_proto_init() {
	if File_pkg_musician_musician_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_musician_musician_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongData); i {
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
		file_pkg_musician_musician_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MusicianData); i {
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
		file_pkg_musician_musician_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneRequest); i {
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
		file_pkg_musician_musician_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneResponse); i {
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
		file_pkg_musician_musician_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_pkg_musician_musician_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_pkg_musician_musician_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_pkg_musician_musician_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_pkg_musician_musician_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_musician_musician_proto_goTypes,
		DependencyIndexes: file_pkg_musician_musician_proto_depIdxs,
		MessageInfos:      file_pkg_musician_musician_proto_msgTypes,
	}.Build()
	File_pkg_musician_musician_proto = out.File
	file_pkg_musician_musician_proto_rawDesc = nil
	file_pkg_musician_musician_proto_goTypes = nil
	file_pkg_musician_musician_proto_depIdxs = nil
}
