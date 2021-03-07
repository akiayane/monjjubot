// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.2
// source: databaseproto/db.proto

package databaseproto

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

type BookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId int64  `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Subject  string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *BookRequest) Reset() {
	*x = BookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookRequest) ProtoMessage() {}

func (x *BookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookRequest.ProtoReflect.Descriptor instead.
func (*BookRequest) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{0}
}

func (x *BookRequest) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *BookRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type BigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookPacks []*BookPack `protobuf:"bytes,1,rep,name=bookPacks,proto3" json:"bookPacks,omitempty"`
}

func (x *BigResponse) Reset() {
	*x = BigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigResponse) ProtoMessage() {}

func (x *BigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigResponse.ProtoReflect.Descriptor instead.
func (*BigResponse) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{1}
}

func (x *BigResponse) GetBookPacks() []*BookPack {
	if x != nil {
		return x.BookPacks
	}
	return nil
}

type BookPack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookId   int64  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	CourseId int64  `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Subject  string `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	BookName string `protobuf:"bytes,4,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	BookLink string `protobuf:"bytes,5,opt,name=book_link,json=bookLink,proto3" json:"book_link,omitempty"`
}

func (x *BookPack) Reset() {
	*x = BookPack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_databaseproto_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookPack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookPack) ProtoMessage() {}

func (x *BookPack) ProtoReflect() protoreflect.Message {
	mi := &file_databaseproto_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookPack.ProtoReflect.Descriptor instead.
func (*BookPack) Descriptor() ([]byte, []int) {
	return file_databaseproto_db_proto_rawDescGZIP(), []int{2}
}

func (x *BookPack) GetBookId() int64 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *BookPack) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *BookPack) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *BookPack) GetBookName() string {
	if x != nil {
		return x.BookName
	}
	return ""
}

func (x *BookPack) GetBookLink() string {
	if x != nil {
		return x.BookLink
	}
	return ""
}

var File_databaseproto_db_proto protoreflect.FileDescriptor

var file_databaseproto_db_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x22, 0x44, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3d, 0x0a, 0x0b, 0x62, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x50, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x61,
	0x63, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x6e, 0x6b, 0x32, 0x54, 0x0a, 0x15,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x6d, 0x6f, 0x6e, 0x6a, 0x6a, 0x75, 0x62, 0x6f, 0x74, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_databaseproto_db_proto_rawDescOnce sync.Once
	file_databaseproto_db_proto_rawDescData = file_databaseproto_db_proto_rawDesc
)

func file_databaseproto_db_proto_rawDescGZIP() []byte {
	file_databaseproto_db_proto_rawDescOnce.Do(func() {
		file_databaseproto_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_databaseproto_db_proto_rawDescData)
	})
	return file_databaseproto_db_proto_rawDescData
}

var file_databaseproto_db_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_databaseproto_db_proto_goTypes = []interface{}{
	(*BookRequest)(nil), // 0: mailer.bookRequest
	(*BigResponse)(nil), // 1: mailer.bigResponse
	(*BookPack)(nil),    // 2: mailer.BookPack
}
var file_databaseproto_db_proto_depIdxs = []int32{
	2, // 0: mailer.bigResponse.bookPacks:type_name -> mailer.BookPack
	0, // 1: mailer.DatabaseAccessService.CourseRequest:input_type -> mailer.bookRequest
	1, // 2: mailer.DatabaseAccessService.CourseRequest:output_type -> mailer.bigResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_databaseproto_db_proto_init() }
func file_databaseproto_db_proto_init() {
	if File_databaseproto_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_databaseproto_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookRequest); i {
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
		file_databaseproto_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigResponse); i {
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
		file_databaseproto_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookPack); i {
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
			RawDescriptor: file_databaseproto_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_databaseproto_db_proto_goTypes,
		DependencyIndexes: file_databaseproto_db_proto_depIdxs,
		MessageInfos:      file_databaseproto_db_proto_msgTypes,
	}.Build()
	File_databaseproto_db_proto = out.File
	file_databaseproto_db_proto_rawDesc = nil
	file_databaseproto_db_proto_goTypes = nil
	file_databaseproto_db_proto_depIdxs = nil
}