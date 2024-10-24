//  protoc --proto_path=. --go_out=. --go-grpc_out=. stockpb/stock.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: stockpb/stock.proto

package stockpb

import (
	orderpb "github.com/lingjun0314/goder/common/genproto/orderpb"
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

type GetItemsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemIDs []string `protobuf:"bytes,1,rep,name=ItemIDs,proto3" json:"ItemIDs,omitempty"`
}

func (x *GetItemsRequest) Reset() {
	*x = GetItemsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockpb_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsRequest) ProtoMessage() {}

func (x *GetItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsRequest.ProtoReflect.Descriptor instead.
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemsRequest) GetItemIDs() []string {
	if x != nil {
		return x.ItemIDs
	}
	return nil
}

type GetItemsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*orderpb.Item `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *GetItemsResponse) Reset() {
	*x = GetItemsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockpb_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsResponse) ProtoMessage() {}

func (x *GetItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsResponse.ProtoReflect.Descriptor instead.
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemsResponse) GetItems() []*orderpb.Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckIfItemsInSrockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*orderpb.ItemWithQuantity `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *CheckIfItemsInSrockRequest) Reset() {
	*x = CheckIfItemsInSrockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockpb_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfItemsInSrockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemsInSrockRequest) ProtoMessage() {}

func (x *CheckIfItemsInSrockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemsInSrockRequest.ProtoReflect.Descriptor instead.
func (*CheckIfItemsInSrockRequest) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{2}
}

func (x *CheckIfItemsInSrockRequest) GetItems() []*orderpb.ItemWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckIfItemsInSrockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InStock int32           `protobuf:"varint,1,opt,name=InStock,proto3" json:"InStock,omitempty"`
	Items   []*orderpb.Item `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *CheckIfItemsInSrockResponse) Reset() {
	*x = CheckIfItemsInSrockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stockpb_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfItemsInSrockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemsInSrockResponse) ProtoMessage() {}

func (x *CheckIfItemsInSrockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stockpb_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemsInSrockResponse.ProtoReflect.Descriptor instead.
func (*CheckIfItemsInSrockResponse) Descriptor() ([]byte, []int) {
	return file_stockpb_stock_proto_rawDescGZIP(), []int{3}
}

func (x *CheckIfItemsInSrockResponse) GetInStock() int32 {
	if x != nil {
		return x.InStock
	}
	return 0
}

func (x *CheckIfItemsInSrockResponse) GetItems() []*orderpb.Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_stockpb_stock_proto protoreflect.FileDescriptor

var file_stockpb_stock_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x1a, 0x13,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x73,
	0x22, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x4d, 0x0a, 0x1a, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x53, 0x72, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5c, 0x0a, 0x1b, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x53, 0x72, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6e, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x49, 0x6e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x23, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xb1, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x53, 0x72, 0x6f, 0x63, 0x6b, 0x12,
	0x23, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x53, 0x72, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x6e, 0x53, 0x72, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stockpb_stock_proto_rawDescOnce sync.Once
	file_stockpb_stock_proto_rawDescData = file_stockpb_stock_proto_rawDesc
)

func file_stockpb_stock_proto_rawDescGZIP() []byte {
	file_stockpb_stock_proto_rawDescOnce.Do(func() {
		file_stockpb_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_stockpb_stock_proto_rawDescData)
	})
	return file_stockpb_stock_proto_rawDescData
}

var file_stockpb_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_stockpb_stock_proto_goTypes = []any{
	(*GetItemsRequest)(nil),             // 0: stockpb.GetItemsRequest
	(*GetItemsResponse)(nil),            // 1: stockpb.GetItemsResponse
	(*CheckIfItemsInSrockRequest)(nil),  // 2: stockpb.CheckIfItemsInSrockRequest
	(*CheckIfItemsInSrockResponse)(nil), // 3: stockpb.CheckIfItemsInSrockResponse
	(*orderpb.Item)(nil),                // 4: orderpb.Item
	(*orderpb.ItemWithQuantity)(nil),    // 5: orderpb.ItemWithQuantity
}
var file_stockpb_stock_proto_depIdxs = []int32{
	4, // 0: stockpb.GetItemsResponse.Items:type_name -> orderpb.Item
	5, // 1: stockpb.CheckIfItemsInSrockRequest.Items:type_name -> orderpb.ItemWithQuantity
	4, // 2: stockpb.CheckIfItemsInSrockResponse.Items:type_name -> orderpb.Item
	0, // 3: stockpb.StockService.GetItems:input_type -> stockpb.GetItemsRequest
	2, // 4: stockpb.StockService.CheckIfItemsInSrock:input_type -> stockpb.CheckIfItemsInSrockRequest
	1, // 5: stockpb.StockService.GetItems:output_type -> stockpb.GetItemsResponse
	3, // 6: stockpb.StockService.CheckIfItemsInSrock:output_type -> stockpb.CheckIfItemsInSrockResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_stockpb_stock_proto_init() }
func file_stockpb_stock_proto_init() {
	if File_stockpb_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stockpb_stock_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetItemsRequest); i {
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
		file_stockpb_stock_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetItemsResponse); i {
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
		file_stockpb_stock_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CheckIfItemsInSrockRequest); i {
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
		file_stockpb_stock_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CheckIfItemsInSrockResponse); i {
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
			RawDescriptor: file_stockpb_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stockpb_stock_proto_goTypes,
		DependencyIndexes: file_stockpb_stock_proto_depIdxs,
		MessageInfos:      file_stockpb_stock_proto_msgTypes,
	}.Build()
	File_stockpb_stock_proto = out.File
	file_stockpb_stock_proto_rawDesc = nil
	file_stockpb_stock_proto_goTypes = nil
	file_stockpb_stock_proto_depIdxs = nil
}
