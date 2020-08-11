// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: todoitem.proto

package todoitem

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TodoItem_Done int32

const (
	TodoItem_UNDEFINED TodoItem_Done = 0
	TodoItem_DONE      TodoItem_Done = 1
	TodoItem_UNDONE    TodoItem_Done = 2
)

// Enum value maps for TodoItem_Done.
var (
	TodoItem_Done_name = map[int32]string{
		0: "UNDEFINED",
		1: "DONE",
		2: "UNDONE",
	}
	TodoItem_Done_value = map[string]int32{
		"UNDEFINED": 0,
		"DONE":      1,
		"UNDONE":    2,
	}
)

func (x TodoItem_Done) Enum() *TodoItem_Done {
	p := new(TodoItem_Done)
	*p = x
	return p
}

func (x TodoItem_Done) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoItem_Done) Descriptor() protoreflect.EnumDescriptor {
	return file_todoitem_proto_enumTypes[0].Descriptor()
}

func (TodoItem_Done) Type() protoreflect.EnumType {
	return &file_todoitem_proto_enumTypes[0]
}

func (x TodoItem_Done) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoItem_Done.Descriptor instead.
func (TodoItem_Done) EnumDescriptor() ([]byte, []int) {
	return file_todoitem_proto_rawDescGZIP(), []int{0, 0}
}

type TodoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Done        TodoItem_Done        `protobuf:"varint,4,opt,name=done,proto3,enum=TodoItem_Done" json:"done,omitempty"`
	LastUpdated *timestamp.Timestamp `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *TodoItem) Reset() {
	*x = TodoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoitem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItem) ProtoMessage() {}

func (x *TodoItem) ProtoReflect() protoreflect.Message {
	mi := &file_todoitem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItem.ProtoReflect.Descriptor instead.
func (*TodoItem) Descriptor() ([]byte, []int) {
	return file_todoitem_proto_rawDescGZIP(), []int{0}
}

func (x *TodoItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TodoItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TodoItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TodoItem) GetDone() TodoItem_Done {
	if x != nil {
		return x.Done
	}
	return TodoItem_UNDEFINED
}

func (x *TodoItem) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type TodoItemList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todolist []*TodoItem `protobuf:"bytes,1,rep,name=todolist,proto3" json:"todolist,omitempty"`
}

func (x *TodoItemList) Reset() {
	*x = TodoItemList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todoitem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoItemList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoItemList) ProtoMessage() {}

func (x *TodoItemList) ProtoReflect() protoreflect.Message {
	mi := &file_todoitem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoItemList.ProtoReflect.Descriptor instead.
func (*TodoItemList) Descriptor() ([]byte, []int) {
	return file_todoitem_proto_rawDescGZIP(), []int{1}
}

func (x *TodoItemList) GetTodolist() []*TodoItem {
	if x != nil {
		return x.Todolist
	}
	return nil
}

var File_todoitem_proto protoreflect.FileDescriptor

var file_todoitem_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x6f, 0x64, 0x6f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a,
	0x08, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x52,
	0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x22, 0x2b, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x44, 0x4f, 0x4e, 0x45, 0x10,
	0x02, 0x22, 0x35, 0x0a, 0x0c, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x08, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08,
	0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xfa, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x64,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x09, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x2e,
	0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x46, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x09, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f,
	0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x09, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x09, 0x2e, 0x54, 0x6f, 0x64,
	0x6f, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x51, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x54, 0x6f,
	0x64, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x6c, 0x69, 0x73, 0x74,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todoitem_proto_rawDescOnce sync.Once
	file_todoitem_proto_rawDescData = file_todoitem_proto_rawDesc
)

func file_todoitem_proto_rawDescGZIP() []byte {
	file_todoitem_proto_rawDescOnce.Do(func() {
		file_todoitem_proto_rawDescData = protoimpl.X.CompressGZIP(file_todoitem_proto_rawDescData)
	})
	return file_todoitem_proto_rawDescData
}

var file_todoitem_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todoitem_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_todoitem_proto_goTypes = []interface{}{
	(TodoItem_Done)(0),          // 0: TodoItem.Done
	(*TodoItem)(nil),            // 1: TodoItem
	(*TodoItemList)(nil),        // 2: TodoItemList
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 4: google.protobuf.Empty
}
var file_todoitem_proto_depIdxs = []int32{
	0, // 0: TodoItem.done:type_name -> TodoItem.Done
	3, // 1: TodoItem.last_updated:type_name -> google.protobuf.Timestamp
	1, // 2: TodoItemList.todolist:type_name -> TodoItem
	1, // 3: TodoService.CreateTodoItem:input_type -> TodoItem
	1, // 4: TodoService.UpdateTodoItem:input_type -> TodoItem
	1, // 5: TodoService.DeleteTodoItem:input_type -> TodoItem
	1, // 6: TodoService.GetTodoItem:input_type -> TodoItem
	4, // 7: TodoService.ListTodoItems:input_type -> google.protobuf.Empty
	1, // 8: TodoService.CreateTodoItem:output_type -> TodoItem
	1, // 9: TodoService.UpdateTodoItem:output_type -> TodoItem
	1, // 10: TodoService.DeleteTodoItem:output_type -> TodoItem
	1, // 11: TodoService.GetTodoItem:output_type -> TodoItem
	2, // 12: TodoService.ListTodoItems:output_type -> TodoItemList
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_todoitem_proto_init() }
func file_todoitem_proto_init() {
	if File_todoitem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todoitem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItem); i {
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
		file_todoitem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoItemList); i {
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
			RawDescriptor: file_todoitem_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todoitem_proto_goTypes,
		DependencyIndexes: file_todoitem_proto_depIdxs,
		EnumInfos:         file_todoitem_proto_enumTypes,
		MessageInfos:      file_todoitem_proto_msgTypes,
	}.Build()
	File_todoitem_proto = out.File
	file_todoitem_proto_rawDesc = nil
	file_todoitem_proto_goTypes = nil
	file_todoitem_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error)
	UpdateTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error)
	DeleteTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error)
	GetTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error)
	ListTodoItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoItemList, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error) {
	out := new(TodoItem)
	err := c.cc.Invoke(ctx, "/TodoService/CreateTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error) {
	out := new(TodoItem)
	err := c.cc.Invoke(ctx, "/TodoService/UpdateTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error) {
	out := new(TodoItem)
	err := c.cc.Invoke(ctx, "/TodoService/DeleteTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodoItem(ctx context.Context, in *TodoItem, opts ...grpc.CallOption) (*TodoItem, error) {
	out := new(TodoItem)
	err := c.cc.Invoke(ctx, "/TodoService/GetTodoItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) ListTodoItems(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TodoItemList, error) {
	out := new(TodoItemList)
	err := c.cc.Invoke(ctx, "/TodoService/ListTodoItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
type TodoServiceServer interface {
	CreateTodoItem(context.Context, *TodoItem) (*TodoItem, error)
	UpdateTodoItem(context.Context, *TodoItem) (*TodoItem, error)
	DeleteTodoItem(context.Context, *TodoItem) (*TodoItem, error)
	GetTodoItem(context.Context, *TodoItem) (*TodoItem, error)
	ListTodoItems(context.Context, *empty.Empty) (*TodoItemList, error)
}

// UnimplementedTodoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (*UnimplementedTodoServiceServer) CreateTodoItem(context.Context, *TodoItem) (*TodoItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodoItem not implemented")
}
func (*UnimplementedTodoServiceServer) UpdateTodoItem(context.Context, *TodoItem) (*TodoItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodoItem not implemented")
}
func (*UnimplementedTodoServiceServer) DeleteTodoItem(context.Context, *TodoItem) (*TodoItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodoItem not implemented")
}
func (*UnimplementedTodoServiceServer) GetTodoItem(context.Context, *TodoItem) (*TodoItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoItem not implemented")
}
func (*UnimplementedTodoServiceServer) ListTodoItems(context.Context, *empty.Empty) (*TodoItemList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodoItems not implemented")
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/CreateTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodoItem(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/UpdateTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodoItem(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/DeleteTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodoItem(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodoItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodoItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/GetTodoItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodoItem(ctx, req.(*TodoItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_ListTodoItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).ListTodoItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TodoService/ListTodoItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).ListTodoItems(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodoItem",
			Handler:    _TodoService_CreateTodoItem_Handler,
		},
		{
			MethodName: "UpdateTodoItem",
			Handler:    _TodoService_UpdateTodoItem_Handler,
		},
		{
			MethodName: "DeleteTodoItem",
			Handler:    _TodoService_DeleteTodoItem_Handler,
		},
		{
			MethodName: "GetTodoItem",
			Handler:    _TodoService_GetTodoItem_Handler,
		},
		{
			MethodName: "ListTodoItems",
			Handler:    _TodoService_ListTodoItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todoitem.proto",
}