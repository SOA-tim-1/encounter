// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: encounter_execution_service.proto

package encounter_execution

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

// Definicija enumeracije EncounterExecutionStatus
type EncounterExecutionStatus int32

const (
	EncounterExecutionStatus_EXECUTION_ACTIVE    EncounterExecutionStatus = 0
	EncounterExecutionStatus_EXECUTION_COMPLETED EncounterExecutionStatus = 1
	EncounterExecutionStatus_EXECUTION_ABANDONED EncounterExecutionStatus = 2
)

// Enum value maps for EncounterExecutionStatus.
var (
	EncounterExecutionStatus_name = map[int32]string{
		0: "EXECUTION_ACTIVE",
		1: "EXECUTION_COMPLETED",
		2: "EXECUTION_ABANDONED",
	}
	EncounterExecutionStatus_value = map[string]int32{
		"EXECUTION_ACTIVE":    0,
		"EXECUTION_COMPLETED": 1,
		"EXECUTION_ABANDONED": 2,
	}
)

func (x EncounterExecutionStatus) Enum() *EncounterExecutionStatus {
	p := new(EncounterExecutionStatus)
	*p = x
	return p
}

func (x EncounterExecutionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncounterExecutionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_encounter_execution_service_proto_enumTypes[0].Descriptor()
}

func (EncounterExecutionStatus) Type() protoreflect.EnumType {
	return &file_encounter_execution_service_proto_enumTypes[0]
}

func (x EncounterExecutionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncounterExecutionStatus.Descriptor instead.
func (EncounterExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return file_encounter_execution_service_proto_rawDescGZIP(), []int{0}
}

// Definicija DTO-a CoordinateDto
type CoordinateDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *CoordinateDto) Reset() {
	*x = CoordinateDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_execution_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinateDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinateDto) ProtoMessage() {}

func (x *CoordinateDto) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_execution_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinateDto.ProtoReflect.Descriptor instead.
func (*CoordinateDto) Descriptor() ([]byte, []int) {
	return file_encounter_execution_service_proto_rawDescGZIP(), []int{0}
}

func (x *CoordinateDto) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CoordinateDto) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// Definicija DTO-a EncounterExecutionDto
type EncounterExecutionDtoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EncounterId            string                   `protobuf:"bytes,2,opt,name=encounter_id,json=encounterId,proto3" json:"encounter_id,omitempty"`
	TouristId              int64                    `protobuf:"varint,3,opt,name=tourist_id,json=touristId,proto3" json:"tourist_id,omitempty"`
	Status                 EncounterExecutionStatus `protobuf:"varint,4,opt,name=status,proto3,enum=EncounterExecutionStatus" json:"status,omitempty"`
	LastActivity           *timestamppb.Timestamp   `protobuf:"bytes,5,opt,name=last_activity,json=lastActivity,proto3" json:"last_activity,omitempty"`
	LocationEntryTimestamp *timestamppb.Timestamp   `protobuf:"bytes,6,opt,name=location_entry_timestamp,json=locationEntryTimestamp,proto3" json:"location_entry_timestamp,omitempty"`
	LastPosition           *CoordinateDto           `protobuf:"bytes,7,opt,name=last_position,json=lastPosition,proto3" json:"last_position,omitempty"`
}

func (x *EncounterExecutionDtoResponse) Reset() {
	*x = EncounterExecutionDtoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_execution_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterExecutionDtoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterExecutionDtoResponse) ProtoMessage() {}

func (x *EncounterExecutionDtoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_execution_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterExecutionDtoResponse.ProtoReflect.Descriptor instead.
func (*EncounterExecutionDtoResponse) Descriptor() ([]byte, []int) {
	return file_encounter_execution_service_proto_rawDescGZIP(), []int{1}
}

func (x *EncounterExecutionDtoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EncounterExecutionDtoResponse) GetEncounterId() string {
	if x != nil {
		return x.EncounterId
	}
	return ""
}

func (x *EncounterExecutionDtoResponse) GetTouristId() int64 {
	if x != nil {
		return x.TouristId
	}
	return 0
}

func (x *EncounterExecutionDtoResponse) GetStatus() EncounterExecutionStatus {
	if x != nil {
		return x.Status
	}
	return EncounterExecutionStatus_EXECUTION_ACTIVE
}

func (x *EncounterExecutionDtoResponse) GetLastActivity() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActivity
	}
	return nil
}

func (x *EncounterExecutionDtoResponse) GetLocationEntryTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.LocationEntryTimestamp
	}
	return nil
}

func (x *EncounterExecutionDtoResponse) GetLastPosition() *CoordinateDto {
	if x != nil {
		return x.LastPosition
	}
	return nil
}

type ActivateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncounterId     string         `protobuf:"bytes,1,opt,name=encounter_id,json=encounterId,proto3" json:"encounter_id,omitempty"`
	TouristId       int64          `protobuf:"varint,2,opt,name=tourist_id,json=touristId,proto3" json:"tourist_id,omitempty"`
	CurrentPosition *CoordinateDto `protobuf:"bytes,3,opt,name=current_position,json=currentPosition,proto3" json:"current_position,omitempty"`
}

func (x *ActivateRequest) Reset() {
	*x = ActivateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_execution_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRequest) ProtoMessage() {}

func (x *ActivateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_execution_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRequest.ProtoReflect.Descriptor instead.
func (*ActivateRequest) Descriptor() ([]byte, []int) {
	return file_encounter_execution_service_proto_rawDescGZIP(), []int{2}
}

func (x *ActivateRequest) GetEncounterId() string {
	if x != nil {
		return x.EncounterId
	}
	return ""
}

func (x *ActivateRequest) GetTouristId() int64 {
	if x != nil {
		return x.TouristId
	}
	return 0
}

func (x *ActivateRequest) GetCurrentPosition() *CoordinateDto {
	if x != nil {
		return x.CurrentPosition
	}
	return nil
}

type AbandonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionId string `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	TouristId   int64  `protobuf:"varint,2,opt,name=tourist_id,json=touristId,proto3" json:"tourist_id,omitempty"`
}

func (x *AbandonRequest) Reset() {
	*x = AbandonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_execution_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbandonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbandonRequest) ProtoMessage() {}

func (x *AbandonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_execution_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbandonRequest.ProtoReflect.Descriptor instead.
func (*AbandonRequest) Descriptor() ([]byte, []int) {
	return file_encounter_execution_service_proto_rawDescGZIP(), []int{3}
}

func (x *AbandonRequest) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *AbandonRequest) GetTouristId() int64 {
	if x != nil {
		return x.TouristId
	}
	return 0
}

type CheckIfCompletedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionId     string         `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	TouristId       int64          `protobuf:"varint,2,opt,name=tourist_id,json=touristId,proto3" json:"tourist_id,omitempty"`
	CurrentPosition *CoordinateDto `protobuf:"bytes,3,opt,name=current_position,json=currentPosition,proto3" json:"current_position,omitempty"`
}

func (x *CheckIfCompletedRequest) Reset() {
	*x = CheckIfCompletedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_execution_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfCompletedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfCompletedRequest) ProtoMessage() {}

func (x *CheckIfCompletedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_execution_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfCompletedRequest.ProtoReflect.Descriptor instead.
func (*CheckIfCompletedRequest) Descriptor() ([]byte, []int) {
	return file_encounter_execution_service_proto_rawDescGZIP(), []int{4}
}

func (x *CheckIfCompletedRequest) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *CheckIfCompletedRequest) GetTouristId() int64 {
	if x != nil {
		return x.TouristId
	}
	return 0
}

func (x *CheckIfCompletedRequest) GetCurrentPosition() *CoordinateDto {
	if x != nil {
		return x.CurrentPosition
	}
	return nil
}

type CompleteMiscEncounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionId string `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	TouristId   int64  `protobuf:"varint,2,opt,name=tourist_id,json=touristId,proto3" json:"tourist_id,omitempty"`
}

func (x *CompleteMiscEncounterRequest) Reset() {
	*x = CompleteMiscEncounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_execution_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteMiscEncounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteMiscEncounterRequest) ProtoMessage() {}

func (x *CompleteMiscEncounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_execution_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteMiscEncounterRequest.ProtoReflect.Descriptor instead.
func (*CompleteMiscEncounterRequest) Descriptor() ([]byte, []int) {
	return file_encounter_execution_service_proto_rawDescGZIP(), []int{5}
}

func (x *CompleteMiscEncounterRequest) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *CompleteMiscEncounterRequest) GetTouristId() int64 {
	if x != nil {
		return x.TouristId
	}
	return 0
}

var File_encounter_execution_service_proto protoreflect.FileDescriptor

var file_encounter_execution_service_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0d, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22,
	0xf0, 0x02, 0x0a, 0x1d, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x54, 0x0a, 0x18, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x16, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x33, 0x0a,
	0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x44, 0x74, 0x6f, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x75,
	0x72, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x44,
	0x74, 0x6f, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0e, 0x41, 0x62, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x75, 0x72,
	0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f,
	0x75, 0x72, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x75, 0x72,
	0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x44, 0x74, 0x6f, 0x52,
	0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x60, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x63,
	0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x75, 0x72, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x2a, 0x62, 0x0a, 0x18, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x10, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x42, 0x41, 0x4e, 0x44,
	0x4f, 0x4e, 0x45, 0x44, 0x10, 0x02, 0x32, 0xf3, 0x02, 0x0a, 0x19, 0x45, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x14, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x50, 0x43, 0x12, 0x10, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x13, 0x41, 0x62, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x45, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x50, 0x43, 0x12, 0x0f, 0x2e, 0x41, 0x62, 0x61, 0x6e, 0x64, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x1c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x50, 0x43, 0x12, 0x18, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x66, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x21, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x69, 0x73, 0x63, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45,
	0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x69, 0x73, 0x63, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x45, 0x6e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2d,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_encounter_execution_service_proto_rawDescOnce sync.Once
	file_encounter_execution_service_proto_rawDescData = file_encounter_execution_service_proto_rawDesc
)

func file_encounter_execution_service_proto_rawDescGZIP() []byte {
	file_encounter_execution_service_proto_rawDescOnce.Do(func() {
		file_encounter_execution_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_encounter_execution_service_proto_rawDescData)
	})
	return file_encounter_execution_service_proto_rawDescData
}

var file_encounter_execution_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_encounter_execution_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_encounter_execution_service_proto_goTypes = []interface{}{
	(EncounterExecutionStatus)(0),         // 0: EncounterExecutionStatus
	(*CoordinateDto)(nil),                 // 1: CoordinateDto
	(*EncounterExecutionDtoResponse)(nil), // 2: EncounterExecutionDtoResponse
	(*ActivateRequest)(nil),               // 3: ActivateRequest
	(*AbandonRequest)(nil),                // 4: AbandonRequest
	(*CheckIfCompletedRequest)(nil),       // 5: CheckIfCompletedRequest
	(*CompleteMiscEncounterRequest)(nil),  // 6: CompleteMiscEncounterRequest
	(*timestamppb.Timestamp)(nil),         // 7: google.protobuf.Timestamp
}
var file_encounter_execution_service_proto_depIdxs = []int32{
	0,  // 0: EncounterExecutionDtoResponse.status:type_name -> EncounterExecutionStatus
	7,  // 1: EncounterExecutionDtoResponse.last_activity:type_name -> google.protobuf.Timestamp
	7,  // 2: EncounterExecutionDtoResponse.location_entry_timestamp:type_name -> google.protobuf.Timestamp
	1,  // 3: EncounterExecutionDtoResponse.last_position:type_name -> CoordinateDto
	1,  // 4: ActivateRequest.current_position:type_name -> CoordinateDto
	1,  // 5: CheckIfCompletedRequest.current_position:type_name -> CoordinateDto
	3,  // 6: EncounterExecutionService.ActivateEncounterRPC:input_type -> ActivateRequest
	4,  // 7: EncounterExecutionService.AbandonEncounterRPC:input_type -> AbandonRequest
	5,  // 8: EncounterExecutionService.CheckIfCompletedEncounterRPC:input_type -> CheckIfCompletedRequest
	6,  // 9: EncounterExecutionService.CompleteMiscEncounterEncounterRPC:input_type -> CompleteMiscEncounterRequest
	2,  // 10: EncounterExecutionService.ActivateEncounterRPC:output_type -> EncounterExecutionDtoResponse
	2,  // 11: EncounterExecutionService.AbandonEncounterRPC:output_type -> EncounterExecutionDtoResponse
	2,  // 12: EncounterExecutionService.CheckIfCompletedEncounterRPC:output_type -> EncounterExecutionDtoResponse
	2,  // 13: EncounterExecutionService.CompleteMiscEncounterEncounterRPC:output_type -> EncounterExecutionDtoResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_encounter_execution_service_proto_init() }
func file_encounter_execution_service_proto_init() {
	if File_encounter_execution_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encounter_execution_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinateDto); i {
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
		file_encounter_execution_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterExecutionDtoResponse); i {
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
		file_encounter_execution_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateRequest); i {
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
		file_encounter_execution_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbandonRequest); i {
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
		file_encounter_execution_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIfCompletedRequest); i {
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
		file_encounter_execution_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteMiscEncounterRequest); i {
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
			RawDescriptor: file_encounter_execution_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_encounter_execution_service_proto_goTypes,
		DependencyIndexes: file_encounter_execution_service_proto_depIdxs,
		EnumInfos:         file_encounter_execution_service_proto_enumTypes,
		MessageInfos:      file_encounter_execution_service_proto_msgTypes,
	}.Build()
	File_encounter_execution_service_proto = out.File
	file_encounter_execution_service_proto_rawDesc = nil
	file_encounter_execution_service_proto_goTypes = nil
	file_encounter_execution_service_proto_depIdxs = nil
}
