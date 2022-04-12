// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: controller/api/services/v1/managed_group_service.proto

package services

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	managedgroups "github.com/hashicorp/boundary/sdk/pbs/controller/api/resources/managedgroups"
	_ "github.com/hashicorp/boundary/sdk/pbs/controller/protooptions"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetManagedGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public"` // @gotags: `class:"public"`
}

func (x *GetManagedGroupRequest) Reset() {
	*x = GetManagedGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManagedGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManagedGroupRequest) ProtoMessage() {}

func (x *GetManagedGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManagedGroupRequest.ProtoReflect.Descriptor instead.
func (*GetManagedGroupRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetManagedGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetManagedGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *managedgroups.ManagedGroup `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetManagedGroupResponse) Reset() {
	*x = GetManagedGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManagedGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManagedGroupResponse) ProtoMessage() {}

func (x *GetManagedGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManagedGroupResponse.ProtoReflect.Descriptor instead.
func (*GetManagedGroupResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetManagedGroupResponse) GetItem() *managedgroups.ManagedGroup {
	if x != nil {
		return x.Item
	}
	return nil
}

type ListManagedGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthMethodId string `protobuf:"bytes,1,opt,name=auth_method_id,proto3" json:"auth_method_id,omitempty" class:"public"` // @gotags: `class:"public"`
	Filter       string `protobuf:"bytes,30,opt,name=filter,proto3" json:"filter,omitempty" class:"public"`                // @gotags: `class:"public"`
}

func (x *ListManagedGroupsRequest) Reset() {
	*x = ListManagedGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListManagedGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListManagedGroupsRequest) ProtoMessage() {}

func (x *ListManagedGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListManagedGroupsRequest.ProtoReflect.Descriptor instead.
func (*ListManagedGroupsRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListManagedGroupsRequest) GetAuthMethodId() string {
	if x != nil {
		return x.AuthMethodId
	}
	return ""
}

func (x *ListManagedGroupsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListManagedGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*managedgroups.ManagedGroup `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListManagedGroupsResponse) Reset() {
	*x = ListManagedGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListManagedGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListManagedGroupsResponse) ProtoMessage() {}

func (x *ListManagedGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListManagedGroupsResponse.ProtoReflect.Descriptor instead.
func (*ListManagedGroupsResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListManagedGroupsResponse) GetItems() []*managedgroups.ManagedGroup {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateManagedGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *managedgroups.ManagedGroup `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateManagedGroupRequest) Reset() {
	*x = CreateManagedGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManagedGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManagedGroupRequest) ProtoMessage() {}

func (x *CreateManagedGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManagedGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateManagedGroupRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateManagedGroupRequest) GetItem() *managedgroups.ManagedGroup {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateManagedGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri  string                      `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty" class:"public"` // @gotags: `class:"public"`
	Item *managedgroups.ManagedGroup `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *CreateManagedGroupResponse) Reset() {
	*x = CreateManagedGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateManagedGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateManagedGroupResponse) ProtoMessage() {}

func (x *CreateManagedGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateManagedGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateManagedGroupResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateManagedGroupResponse) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *CreateManagedGroupResponse) GetItem() *managedgroups.ManagedGroup {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateManagedGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public"` // @gotags: `class:"public"`
	Item       *managedgroups.ManagedGroup `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask      `protobuf:"bytes,3,opt,name=update_mask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateManagedGroupRequest) Reset() {
	*x = UpdateManagedGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManagedGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManagedGroupRequest) ProtoMessage() {}

func (x *UpdateManagedGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManagedGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateManagedGroupRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateManagedGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateManagedGroupRequest) GetItem() *managedgroups.ManagedGroup {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *UpdateManagedGroupRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateManagedGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *managedgroups.ManagedGroup `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *UpdateManagedGroupResponse) Reset() {
	*x = UpdateManagedGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateManagedGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateManagedGroupResponse) ProtoMessage() {}

func (x *UpdateManagedGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateManagedGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateManagedGroupResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateManagedGroupResponse) GetItem() *managedgroups.ManagedGroup {
	if x != nil {
		return x.Item
	}
	return nil
}

type DeleteManagedGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" class:"public"` // @gotags: `class:"public"`
}

func (x *DeleteManagedGroupRequest) Reset() {
	*x = DeleteManagedGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteManagedGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteManagedGroupRequest) ProtoMessage() {}

func (x *DeleteManagedGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteManagedGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteManagedGroupRequest) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteManagedGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteManagedGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteManagedGroupResponse) Reset() {
	*x = DeleteManagedGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteManagedGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteManagedGroupResponse) ProtoMessage() {}

func (x *DeleteManagedGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_services_v1_managed_group_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteManagedGroupResponse.ProtoReflect.Descriptor instead.
func (*DeleteManagedGroupResponse) Descriptor() ([]byte, []int) {
	return file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP(), []int{9}
}

var File_controller_api_services_v1_managed_group_service_proto protoreflect.FileDescriptor

var file_controller_api_services_v1_managed_group_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0x5a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x6a, 0x0a,
	0x19, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x68, 0x0a, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x22, 0x7b, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x12, 0x4b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d,
	0x22, 0xb6, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4b,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x3c, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x22, 0x69, 0x0a, 0x1a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xae, 0x08, 0x0a, 0x13, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc1, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x32, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x92, 0x41, 0x1d, 0x12, 0x1b, 0x47, 0x65, 0x74, 0x73, 0x20,
	0x61, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0xd3, 0x01, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0x34, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51,
	0x92, 0x41, 0x34, 0x12, 0x32, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x20, 0x69, 0x6e, 0x20,
	0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x41, 0x75, 0x74, 0x68, 0x20,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x12, 0xea, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x65, 0x92, 0x41, 0x3c, 0x12, 0x3a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x73, 0x20, 0x61, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x20, 0x69, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x41, 0x75, 0x74, 0x68,
	0x20, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0xcc,
	0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x92, 0x41, 0x19, 0x12, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x20, 0x61, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x32, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x3a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x62, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0xc0, 0x01,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x3b, 0x92, 0x41, 0x19, 0x12, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x73, 0x20, 0x61, 0x20, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x64, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x55, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72,
	0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2,
	0xe3, 0x29, 0x04, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_services_v1_managed_group_service_proto_rawDescOnce sync.Once
	file_controller_api_services_v1_managed_group_service_proto_rawDescData = file_controller_api_services_v1_managed_group_service_proto_rawDesc
)

func file_controller_api_services_v1_managed_group_service_proto_rawDescGZIP() []byte {
	file_controller_api_services_v1_managed_group_service_proto_rawDescOnce.Do(func() {
		file_controller_api_services_v1_managed_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_services_v1_managed_group_service_proto_rawDescData)
	})
	return file_controller_api_services_v1_managed_group_service_proto_rawDescData
}

var file_controller_api_services_v1_managed_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_controller_api_services_v1_managed_group_service_proto_goTypes = []interface{}{
	(*GetManagedGroupRequest)(nil),     // 0: controller.api.services.v1.GetManagedGroupRequest
	(*GetManagedGroupResponse)(nil),    // 1: controller.api.services.v1.GetManagedGroupResponse
	(*ListManagedGroupsRequest)(nil),   // 2: controller.api.services.v1.ListManagedGroupsRequest
	(*ListManagedGroupsResponse)(nil),  // 3: controller.api.services.v1.ListManagedGroupsResponse
	(*CreateManagedGroupRequest)(nil),  // 4: controller.api.services.v1.CreateManagedGroupRequest
	(*CreateManagedGroupResponse)(nil), // 5: controller.api.services.v1.CreateManagedGroupResponse
	(*UpdateManagedGroupRequest)(nil),  // 6: controller.api.services.v1.UpdateManagedGroupRequest
	(*UpdateManagedGroupResponse)(nil), // 7: controller.api.services.v1.UpdateManagedGroupResponse
	(*DeleteManagedGroupRequest)(nil),  // 8: controller.api.services.v1.DeleteManagedGroupRequest
	(*DeleteManagedGroupResponse)(nil), // 9: controller.api.services.v1.DeleteManagedGroupResponse
	(*managedgroups.ManagedGroup)(nil), // 10: controller.api.resources.managedgroups.v1.ManagedGroup
	(*fieldmaskpb.FieldMask)(nil),      // 11: google.protobuf.FieldMask
}
var file_controller_api_services_v1_managed_group_service_proto_depIdxs = []int32{
	10, // 0: controller.api.services.v1.GetManagedGroupResponse.item:type_name -> controller.api.resources.managedgroups.v1.ManagedGroup
	10, // 1: controller.api.services.v1.ListManagedGroupsResponse.items:type_name -> controller.api.resources.managedgroups.v1.ManagedGroup
	10, // 2: controller.api.services.v1.CreateManagedGroupRequest.item:type_name -> controller.api.resources.managedgroups.v1.ManagedGroup
	10, // 3: controller.api.services.v1.CreateManagedGroupResponse.item:type_name -> controller.api.resources.managedgroups.v1.ManagedGroup
	10, // 4: controller.api.services.v1.UpdateManagedGroupRequest.item:type_name -> controller.api.resources.managedgroups.v1.ManagedGroup
	11, // 5: controller.api.services.v1.UpdateManagedGroupRequest.update_mask:type_name -> google.protobuf.FieldMask
	10, // 6: controller.api.services.v1.UpdateManagedGroupResponse.item:type_name -> controller.api.resources.managedgroups.v1.ManagedGroup
	0,  // 7: controller.api.services.v1.ManagedGroupService.GetManagedGroup:input_type -> controller.api.services.v1.GetManagedGroupRequest
	2,  // 8: controller.api.services.v1.ManagedGroupService.ListManagedGroups:input_type -> controller.api.services.v1.ListManagedGroupsRequest
	4,  // 9: controller.api.services.v1.ManagedGroupService.CreateManagedGroup:input_type -> controller.api.services.v1.CreateManagedGroupRequest
	6,  // 10: controller.api.services.v1.ManagedGroupService.UpdateManagedGroup:input_type -> controller.api.services.v1.UpdateManagedGroupRequest
	8,  // 11: controller.api.services.v1.ManagedGroupService.DeleteManagedGroup:input_type -> controller.api.services.v1.DeleteManagedGroupRequest
	1,  // 12: controller.api.services.v1.ManagedGroupService.GetManagedGroup:output_type -> controller.api.services.v1.GetManagedGroupResponse
	3,  // 13: controller.api.services.v1.ManagedGroupService.ListManagedGroups:output_type -> controller.api.services.v1.ListManagedGroupsResponse
	5,  // 14: controller.api.services.v1.ManagedGroupService.CreateManagedGroup:output_type -> controller.api.services.v1.CreateManagedGroupResponse
	7,  // 15: controller.api.services.v1.ManagedGroupService.UpdateManagedGroup:output_type -> controller.api.services.v1.UpdateManagedGroupResponse
	9,  // 16: controller.api.services.v1.ManagedGroupService.DeleteManagedGroup:output_type -> controller.api.services.v1.DeleteManagedGroupResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_controller_api_services_v1_managed_group_service_proto_init() }
func file_controller_api_services_v1_managed_group_service_proto_init() {
	if File_controller_api_services_v1_managed_group_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManagedGroupRequest); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManagedGroupResponse); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListManagedGroupsRequest); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListManagedGroupsResponse); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManagedGroupRequest); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateManagedGroupResponse); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManagedGroupRequest); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateManagedGroupResponse); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteManagedGroupRequest); i {
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
		file_controller_api_services_v1_managed_group_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteManagedGroupResponse); i {
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
			RawDescriptor: file_controller_api_services_v1_managed_group_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_api_services_v1_managed_group_service_proto_goTypes,
		DependencyIndexes: file_controller_api_services_v1_managed_group_service_proto_depIdxs,
		MessageInfos:      file_controller_api_services_v1_managed_group_service_proto_msgTypes,
	}.Build()
	File_controller_api_services_v1_managed_group_service_proto = out.File
	file_controller_api_services_v1_managed_group_service_proto_rawDesc = nil
	file_controller_api_services_v1_managed_group_service_proto_goTypes = nil
	file_controller_api_services_v1_managed_group_service_proto_depIdxs = nil
}
