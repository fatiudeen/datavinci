// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: api/proto/visualization.proto

package grpc

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

type CreateVisualizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataSourceId      string            `protobuf:"bytes,1,opt,name=data_source_id,json=dataSourceId,proto3" json:"data_source_id,omitempty"`
	VisualizationType string            `protobuf:"bytes,2,opt,name=visualization_type,json=visualizationType,proto3" json:"visualization_type,omitempty"`
	Dimensions        []string          `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
	Measures          []string          `protobuf:"bytes,4,rep,name=measures,proto3" json:"measures,omitempty"`
	Filters           map[string]string `protobuf:"bytes,5,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Options           map[string]string `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateVisualizationRequest) Reset() {
	*x = CreateVisualizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVisualizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVisualizationRequest) ProtoMessage() {}

func (x *CreateVisualizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVisualizationRequest.ProtoReflect.Descriptor instead.
func (*CreateVisualizationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{0}
}

func (x *CreateVisualizationRequest) GetDataSourceId() string {
	if x != nil {
		return x.DataSourceId
	}
	return ""
}

func (x *CreateVisualizationRequest) GetVisualizationType() string {
	if x != nil {
		return x.VisualizationType
	}
	return ""
}

func (x *CreateVisualizationRequest) GetDimensions() []string {
	if x != nil {
		return x.Dimensions
	}
	return nil
}

func (x *CreateVisualizationRequest) GetMeasures() []string {
	if x != nil {
		return x.Measures
	}
	return nil
}

func (x *CreateVisualizationRequest) GetFilters() map[string]string {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *CreateVisualizationRequest) GetOptions() map[string]string {
	if x != nil {
		return x.Options
	}
	return nil
}

type UpdateVisualizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisualizationId string                      `protobuf:"bytes,1,opt,name=visualization_id,json=visualizationId,proto3" json:"visualization_id,omitempty"`
	UpdateData      *CreateVisualizationRequest `protobuf:"bytes,2,opt,name=update_data,json=updateData,proto3" json:"update_data,omitempty"`
}

func (x *UpdateVisualizationRequest) Reset() {
	*x = UpdateVisualizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVisualizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVisualizationRequest) ProtoMessage() {}

func (x *UpdateVisualizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVisualizationRequest.ProtoReflect.Descriptor instead.
func (*UpdateVisualizationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateVisualizationRequest) GetVisualizationId() string {
	if x != nil {
		return x.VisualizationId
	}
	return ""
}

func (x *UpdateVisualizationRequest) GetUpdateData() *CreateVisualizationRequest {
	if x != nil {
		return x.UpdateData
	}
	return nil
}

type GetVisualizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisualizationId string `protobuf:"bytes,1,opt,name=visualization_id,json=visualizationId,proto3" json:"visualization_id,omitempty"`
}

func (x *GetVisualizationRequest) Reset() {
	*x = GetVisualizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVisualizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVisualizationRequest) ProtoMessage() {}

func (x *GetVisualizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVisualizationRequest.ProtoReflect.Descriptor instead.
func (*GetVisualizationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{2}
}

func (x *GetVisualizationRequest) GetVisualizationId() string {
	if x != nil {
		return x.VisualizationId
	}
	return ""
}

type ListVisualizationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListVisualizationsRequest) Reset() {
	*x = ListVisualizationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVisualizationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVisualizationsRequest) ProtoMessage() {}

func (x *ListVisualizationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVisualizationsRequest.ProtoReflect.Descriptor instead.
func (*ListVisualizationsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{3}
}

func (x *ListVisualizationsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListVisualizationsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type DeleteVisualizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisualizationId string `protobuf:"bytes,1,opt,name=visualization_id,json=visualizationId,proto3" json:"visualization_id,omitempty"`
}

func (x *DeleteVisualizationRequest) Reset() {
	*x = DeleteVisualizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVisualizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVisualizationRequest) ProtoMessage() {}

func (x *DeleteVisualizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVisualizationRequest.ProtoReflect.Descriptor instead.
func (*DeleteVisualizationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteVisualizationRequest) GetVisualizationId() string {
	if x != nil {
		return x.VisualizationId
	}
	return ""
}

type ExportVisualizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisualizationId string `protobuf:"bytes,1,opt,name=visualization_id,json=visualizationId,proto3" json:"visualization_id,omitempty"`
	Format          string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *ExportVisualizationRequest) Reset() {
	*x = ExportVisualizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportVisualizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportVisualizationRequest) ProtoMessage() {}

func (x *ExportVisualizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportVisualizationRequest.ProtoReflect.Descriptor instead.
func (*ExportVisualizationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{5}
}

func (x *ExportVisualizationRequest) GetVisualizationId() string {
	if x != nil {
		return x.VisualizationId
	}
	return ""
}

func (x *ExportVisualizationRequest) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

type VisualizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VisualizationId   string            `protobuf:"bytes,1,opt,name=visualization_id,json=visualizationId,proto3" json:"visualization_id,omitempty"`
	VisualizationData []byte            `protobuf:"bytes,2,opt,name=visualization_data,json=visualizationData,proto3" json:"visualization_data,omitempty"`
	Metadata          map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VisualizationResponse) Reset() {
	*x = VisualizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VisualizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisualizationResponse) ProtoMessage() {}

func (x *VisualizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisualizationResponse.ProtoReflect.Descriptor instead.
func (*VisualizationResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{6}
}

func (x *VisualizationResponse) GetVisualizationId() string {
	if x != nil {
		return x.VisualizationId
	}
	return ""
}

func (x *VisualizationResponse) GetVisualizationData() []byte {
	if x != nil {
		return x.VisualizationData
	}
	return nil
}

func (x *VisualizationResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type ListVisualizationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Visualizations []*VisualizationResponse `protobuf:"bytes,1,rep,name=visualizations,proto3" json:"visualizations,omitempty"`
	TotalCount     int32                    `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *ListVisualizationsResponse) Reset() {
	*x = ListVisualizationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVisualizationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVisualizationsResponse) ProtoMessage() {}

func (x *ListVisualizationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVisualizationsResponse.ProtoReflect.Descriptor instead.
func (*ListVisualizationsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{7}
}

func (x *ListVisualizationsResponse) GetVisualizations() []*VisualizationResponse {
	if x != nil {
		return x.Visualizations
	}
	return nil
}

func (x *ListVisualizationsResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

type DeleteVisualizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteVisualizationResponse) Reset() {
	*x = DeleteVisualizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVisualizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVisualizationResponse) ProtoMessage() {}

func (x *DeleteVisualizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVisualizationResponse.ProtoReflect.Descriptor instead.
func (*DeleteVisualizationResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteVisualizationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ExportVisualizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExportedData []byte `protobuf:"bytes,1,opt,name=exported_data,json=exportedData,proto3" json:"exported_data,omitempty"`
	Format       string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
}

func (x *ExportVisualizationResponse) Reset() {
	*x = ExportVisualizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_visualization_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportVisualizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportVisualizationResponse) ProtoMessage() {}

func (x *ExportVisualizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_visualization_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportVisualizationResponse.ProtoReflect.Descriptor instead.
func (*ExportVisualizationResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_visualization_proto_rawDescGZIP(), []int{9}
}

func (x *ExportVisualizationResponse) GetExportedData() []byte {
	if x != nil {
		return x.ExportedData
	}
	return nil
}

func (x *ExportVisualizationResponse) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

var File_api_proto_visualization_proto protoreflect.FileDescriptor

var file_api_proto_visualization_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc9,
	0x03, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x73, 0x12, 0x50,
	0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x50, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x36, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a,
	0x0a, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x93, 0x01, 0x0a, 0x1a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x69, 0x73,
	0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x44, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x76,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x47, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x5f, 0x0a,
	0x1a, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x76,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xfe,
	0x01, 0x0a, 0x15, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x76, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x11, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x4e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x8b, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c,
	0x0a, 0x0e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x76, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x37, 0x0a,
	0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5a, 0x0a, 0x1b, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x32, 0x9b, 0x05, 0x0a, 0x14, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x76,
	0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x62, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x2e, 0x76, 0x69, 0x73, 0x75,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6e, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6e, 0x0a, 0x13, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x69,
	0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x27, 0x5a, 0x25, 0x64, 0x61, 0x74, 0x61, 0x76, 0x69, 0x6e, 0x63, 0x69, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x69, 0x73, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_visualization_proto_rawDescOnce sync.Once
	file_api_proto_visualization_proto_rawDescData = file_api_proto_visualization_proto_rawDesc
)

func file_api_proto_visualization_proto_rawDescGZIP() []byte {
	file_api_proto_visualization_proto_rawDescOnce.Do(func() {
		file_api_proto_visualization_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_visualization_proto_rawDescData)
	})
	return file_api_proto_visualization_proto_rawDescData
}

var file_api_proto_visualization_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_proto_visualization_proto_goTypes = []any{
	(*CreateVisualizationRequest)(nil),  // 0: visualization.CreateVisualizationRequest
	(*UpdateVisualizationRequest)(nil),  // 1: visualization.UpdateVisualizationRequest
	(*GetVisualizationRequest)(nil),     // 2: visualization.GetVisualizationRequest
	(*ListVisualizationsRequest)(nil),   // 3: visualization.ListVisualizationsRequest
	(*DeleteVisualizationRequest)(nil),  // 4: visualization.DeleteVisualizationRequest
	(*ExportVisualizationRequest)(nil),  // 5: visualization.ExportVisualizationRequest
	(*VisualizationResponse)(nil),       // 6: visualization.VisualizationResponse
	(*ListVisualizationsResponse)(nil),  // 7: visualization.ListVisualizationsResponse
	(*DeleteVisualizationResponse)(nil), // 8: visualization.DeleteVisualizationResponse
	(*ExportVisualizationResponse)(nil), // 9: visualization.ExportVisualizationResponse
	nil,                                 // 10: visualization.CreateVisualizationRequest.FiltersEntry
	nil,                                 // 11: visualization.CreateVisualizationRequest.OptionsEntry
	nil,                                 // 12: visualization.VisualizationResponse.MetadataEntry
}
var file_api_proto_visualization_proto_depIdxs = []int32{
	10, // 0: visualization.CreateVisualizationRequest.filters:type_name -> visualization.CreateVisualizationRequest.FiltersEntry
	11, // 1: visualization.CreateVisualizationRequest.options:type_name -> visualization.CreateVisualizationRequest.OptionsEntry
	0,  // 2: visualization.UpdateVisualizationRequest.update_data:type_name -> visualization.CreateVisualizationRequest
	12, // 3: visualization.VisualizationResponse.metadata:type_name -> visualization.VisualizationResponse.MetadataEntry
	6,  // 4: visualization.ListVisualizationsResponse.visualizations:type_name -> visualization.VisualizationResponse
	0,  // 5: visualization.VisualizationService.CreateVisualization:input_type -> visualization.CreateVisualizationRequest
	1,  // 6: visualization.VisualizationService.UpdateVisualization:input_type -> visualization.UpdateVisualizationRequest
	2,  // 7: visualization.VisualizationService.GetVisualization:input_type -> visualization.GetVisualizationRequest
	3,  // 8: visualization.VisualizationService.ListVisualizations:input_type -> visualization.ListVisualizationsRequest
	4,  // 9: visualization.VisualizationService.DeleteVisualization:input_type -> visualization.DeleteVisualizationRequest
	5,  // 10: visualization.VisualizationService.ExportVisualization:input_type -> visualization.ExportVisualizationRequest
	6,  // 11: visualization.VisualizationService.CreateVisualization:output_type -> visualization.VisualizationResponse
	6,  // 12: visualization.VisualizationService.UpdateVisualization:output_type -> visualization.VisualizationResponse
	6,  // 13: visualization.VisualizationService.GetVisualization:output_type -> visualization.VisualizationResponse
	7,  // 14: visualization.VisualizationService.ListVisualizations:output_type -> visualization.ListVisualizationsResponse
	8,  // 15: visualization.VisualizationService.DeleteVisualization:output_type -> visualization.DeleteVisualizationResponse
	9,  // 16: visualization.VisualizationService.ExportVisualization:output_type -> visualization.ExportVisualizationResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_visualization_proto_init() }
func file_api_proto_visualization_proto_init() {
	if File_api_proto_visualization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_visualization_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateVisualizationRequest); i {
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
		file_api_proto_visualization_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateVisualizationRequest); i {
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
		file_api_proto_visualization_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetVisualizationRequest); i {
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
		file_api_proto_visualization_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListVisualizationsRequest); i {
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
		file_api_proto_visualization_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteVisualizationRequest); i {
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
		file_api_proto_visualization_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ExportVisualizationRequest); i {
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
		file_api_proto_visualization_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*VisualizationResponse); i {
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
		file_api_proto_visualization_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListVisualizationsResponse); i {
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
		file_api_proto_visualization_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteVisualizationResponse); i {
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
		file_api_proto_visualization_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ExportVisualizationResponse); i {
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
			RawDescriptor: file_api_proto_visualization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_visualization_proto_goTypes,
		DependencyIndexes: file_api_proto_visualization_proto_depIdxs,
		MessageInfos:      file_api_proto_visualization_proto_msgTypes,
	}.Build()
	File_api_proto_visualization_proto = out.File
	file_api_proto_visualization_proto_rawDesc = nil
	file_api_proto_visualization_proto_goTypes = nil
	file_api_proto_visualization_proto_depIdxs = nil
}
