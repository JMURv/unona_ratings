// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/pb/rating.proto

package rating

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_pb_rating_proto_rawDescGZIP(), []int{0}
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingId  uint32                 `protobuf:"varint,1,opt,name=rating_id,json=ratingId,proto3" json:"rating_id,omitempty"`
	UserUuid  string                 `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Rating    uint32                 `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Text      string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_api_pb_rating_proto_rawDescGZIP(), []int{1}
}

func (x *Rating) GetRatingId() uint32 {
	if x != nil {
		return x.RatingId
	}
	return 0
}

func (x *Rating) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *Rating) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Rating) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Rating) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetUserRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
}

func (x *GetUserRatingRequest) Reset() {
	*x = GetUserRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_rating_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRatingRequest) ProtoMessage() {}

func (x *GetUserRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_rating_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRatingRequest.ProtoReflect.Descriptor instead.
func (*GetUserRatingRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_rating_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRatingRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

type CreateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUuid string `protobuf:"bytes,1,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Rating   uint32 `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateReportRequest) Reset() {
	*x = CreateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_rating_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReportRequest) ProtoMessage() {}

func (x *CreateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_rating_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReportRequest.ProtoReflect.Descriptor instead.
func (*CreateReportRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_rating_proto_rawDescGZIP(), []int{3}
}

func (x *CreateReportRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *CreateReportRequest) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *CreateReportRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type UpdateReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingId uint32 `protobuf:"varint,1,opt,name=rating_id,json=ratingId,proto3" json:"rating_id,omitempty"`
	Rating   uint32 `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *UpdateReportRequest) Reset() {
	*x = UpdateReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_rating_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReportRequest) ProtoMessage() {}

func (x *UpdateReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_rating_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReportRequest.ProtoReflect.Descriptor instead.
func (*UpdateReportRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_rating_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReportRequest) GetRatingId() uint32 {
	if x != nil {
		return x.RatingId
	}
	return 0
}

func (x *UpdateReportRequest) GetRating() uint32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *UpdateReportRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type DeleteReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatingId uint32 `protobuf:"varint,1,opt,name=rating_id,json=ratingId,proto3" json:"rating_id,omitempty"`
}

func (x *DeleteReportRequest) Reset() {
	*x = DeleteReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_rating_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReportRequest) ProtoMessage() {}

func (x *DeleteReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_rating_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReportRequest.ProtoReflect.Descriptor instead.
func (*DeleteReportRequest) Descriptor() ([]byte, []int) {
	return file_api_pb_rating_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteReportRequest) GetRatingId() uint32 {
	if x != nil {
		return x.RatingId
	}
	return 0
}

type GetUserRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating float64 `protobuf:"fixed64,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *GetUserRatingResponse) Reset() {
	*x = GetUserRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_pb_rating_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRatingResponse) ProtoMessage() {}

func (x *GetUserRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_pb_rating_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRatingResponse.ProtoReflect.Descriptor instead.
func (*GetUserRatingResponse) Descriptor() ([]byte, []int) {
	return file_api_pb_rating_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserRatingResponse) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

var File_api_pb_rating_proto protoreflect.FileDescriptor

var file_api_pb_rating_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xa9, 0x01, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x33, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x5e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x32, 0x8b, 0x02, 0x0a, 0x0d,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x4d, 0x55, 0x52, 0x76, 0x2f, 0x75, 0x6e,
	0x6f, 0x6e, 0x61, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_pb_rating_proto_rawDescOnce sync.Once
	file_api_pb_rating_proto_rawDescData = file_api_pb_rating_proto_rawDesc
)

func file_api_pb_rating_proto_rawDescGZIP() []byte {
	file_api_pb_rating_proto_rawDescOnce.Do(func() {
		file_api_pb_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_pb_rating_proto_rawDescData)
	})
	return file_api_pb_rating_proto_rawDescData
}

var file_api_pb_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_pb_rating_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: user.Empty
	(*Rating)(nil),                // 1: user.Rating
	(*GetUserRatingRequest)(nil),  // 2: user.GetUserRatingRequest
	(*CreateReportRequest)(nil),   // 3: user.CreateReportRequest
	(*UpdateReportRequest)(nil),   // 4: user.UpdateReportRequest
	(*DeleteReportRequest)(nil),   // 5: user.DeleteReportRequest
	(*GetUserRatingResponse)(nil), // 6: user.GetUserRatingResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_pb_rating_proto_depIdxs = []int32{
	7, // 0: user.Rating.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: user.RatingService.GetUserRating:input_type -> user.GetUserRatingRequest
	3, // 2: user.RatingService.CreateReport:input_type -> user.CreateReportRequest
	4, // 3: user.RatingService.UpdateReport:input_type -> user.UpdateReportRequest
	5, // 4: user.RatingService.DeleteReport:input_type -> user.DeleteReportRequest
	6, // 5: user.RatingService.GetUserRating:output_type -> user.GetUserRatingResponse
	1, // 6: user.RatingService.CreateReport:output_type -> user.Rating
	1, // 7: user.RatingService.UpdateReport:output_type -> user.Rating
	0, // 8: user.RatingService.DeleteReport:output_type -> user.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_pb_rating_proto_init() }
func file_api_pb_rating_proto_init() {
	if File_api_pb_rating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_pb_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_pb_rating_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_api_pb_rating_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRatingRequest); i {
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
		file_api_pb_rating_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReportRequest); i {
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
		file_api_pb_rating_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReportRequest); i {
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
		file_api_pb_rating_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReportRequest); i {
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
		file_api_pb_rating_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRatingResponse); i {
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
			RawDescriptor: file_api_pb_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pb_rating_proto_goTypes,
		DependencyIndexes: file_api_pb_rating_proto_depIdxs,
		MessageInfos:      file_api_pb_rating_proto_msgTypes,
	}.Build()
	File_api_pb_rating_proto = out.File
	file_api_pb_rating_proto_rawDesc = nil
	file_api_pb_rating_proto_goTypes = nil
	file_api_pb_rating_proto_depIdxs = nil
}
