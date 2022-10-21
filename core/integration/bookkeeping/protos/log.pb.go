// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: protos/log.proto

package bkpb

import (
	proto "github.com/golang/protobuf/proto"
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

type LogOrigin int32

const (
	LogOrigin_LOG_ORIGIN_NULL    LogOrigin = 0
	LogOrigin_LOG_ORIGIN_HUMAN   LogOrigin = 1
	LogOrigin_LOG_ORIGIN_PROCESS LogOrigin = 2
)

// Enum value maps for LogOrigin.
var (
	LogOrigin_name = map[int32]string{
		0: "LOG_ORIGIN_NULL",
		1: "LOG_ORIGIN_HUMAN",
		2: "LOG_ORIGIN_PROCESS",
	}
	LogOrigin_value = map[string]int32{
		"LOG_ORIGIN_NULL":    0,
		"LOG_ORIGIN_HUMAN":   1,
		"LOG_ORIGIN_PROCESS": 2,
	}
)

func (x LogOrigin) Enum() *LogOrigin {
	p := new(LogOrigin)
	*p = x
	return p
}

func (x LogOrigin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogOrigin) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_log_proto_enumTypes[0].Descriptor()
}

func (LogOrigin) Type() protoreflect.EnumType {
	return &file_protos_log_proto_enumTypes[0]
}

func (x LogOrigin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogOrigin.Descriptor instead.
func (LogOrigin) EnumDescriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{0}
}

type LogSubtype int32

const (
	LogSubtype_LOG_SUBTYPE_NULL         LogSubtype = 0
	LogSubtype_LOG_SUBTYPE_RUN          LogSubtype = 1
	LogSubtype_LOG_SUBTYPE_SUBSYSTEM    LogSubtype = 2
	LogSubtype_LOG_SUBTYPE_ANNOUNCEMENT LogSubtype = 3
	LogSubtype_LOG_SUBTYPE_INTERVENTION LogSubtype = 4
	LogSubtype_LOG_SUBTYPE_COMMENT      LogSubtype = 5
)

// Enum value maps for LogSubtype.
var (
	LogSubtype_name = map[int32]string{
		0: "LOG_SUBTYPE_NULL",
		1: "LOG_SUBTYPE_RUN",
		2: "LOG_SUBTYPE_SUBSYSTEM",
		3: "LOG_SUBTYPE_ANNOUNCEMENT",
		4: "LOG_SUBTYPE_INTERVENTION",
		5: "LOG_SUBTYPE_COMMENT",
	}
	LogSubtype_value = map[string]int32{
		"LOG_SUBTYPE_NULL":         0,
		"LOG_SUBTYPE_RUN":          1,
		"LOG_SUBTYPE_SUBSYSTEM":    2,
		"LOG_SUBTYPE_ANNOUNCEMENT": 3,
		"LOG_SUBTYPE_INTERVENTION": 4,
		"LOG_SUBTYPE_COMMENT":      5,
	}
)

func (x LogSubtype) Enum() *LogSubtype {
	p := new(LogSubtype)
	*p = x
	return p
}

func (x LogSubtype) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogSubtype) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_log_proto_enumTypes[1].Descriptor()
}

func (LogSubtype) Type() protoreflect.EnumType {
	return &file_protos_log_proto_enumTypes[1]
}

func (x LogSubtype) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogSubtype.Descriptor instead.
func (LogSubtype) EnumDescriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{1}
}

type LogCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Text        string  `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	RunNumbers  []int32 `protobuf:"varint,3,rep,packed,name=runNumbers,proto3" json:"runNumbers,omitempty"`
	ParentLogId int32   `protobuf:"varint,4,opt,name=parentLogId,proto3" json:"parentLogId,omitempty"`
}

func (x *LogCreationRequest) Reset() {
	*x = LogCreationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogCreationRequest) ProtoMessage() {}

func (x *LogCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogCreationRequest.ProtoReflect.Descriptor instead.
func (*LogCreationRequest) Descriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogCreationRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *LogCreationRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *LogCreationRequest) GetRunNumbers() []int32 {
	if x != nil {
		return x.RunNumbers
	}
	return nil
}

func (x *LogCreationRequest) GetParentLogId() int32 {
	if x != nil {
		return x.ParentLogId
	}
	return 0
}

type Attachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp when this entity was created.
	CreatedAt    int64  `protobuf:"varint,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Encoding     string `protobuf:"bytes,2,opt,name=encoding,proto3" json:"encoding,omitempty"`
	FileName     string `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Id           int32  `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	LogId        int32  `protobuf:"varint,5,opt,name=logId,proto3" json:"logId,omitempty"`
	MimeType     string `protobuf:"bytes,6,opt,name=mimeType,proto3" json:"mimeType,omitempty"`
	OriginalName string `protobuf:"bytes,7,opt,name=originalName,proto3" json:"originalName,omitempty"`
	Path         string `protobuf:"bytes,8,opt,name=path,proto3" json:"path,omitempty"`
	Size         int32  `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// Unix timestamp when this entity was created.
	UpdatedAt int64 `protobuf:"varint,10,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Attachment) Reset() {
	*x = Attachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attachment) ProtoMessage() {}

func (x *Attachment) ProtoReflect() protoreflect.Message {
	mi := &file_protos_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attachment.ProtoReflect.Descriptor instead.
func (*Attachment) Descriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{1}
}

func (x *Attachment) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Attachment) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

func (x *Attachment) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Attachment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Attachment) GetLogId() int32 {
	if x != nil {
		return x.LogId
	}
	return 0
}

func (x *Attachment) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Attachment) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *Attachment) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Attachment) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Attachment) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attachments []*Attachment `protobuf:"bytes,1,rep,name=attachments,proto3" json:"attachments,omitempty"`
	Author      *User         `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// Unix timestamp when this entity was created.
	CreatedAt   int64       `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Id          int32       `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Origin      []LogOrigin `protobuf:"varint,5,rep,packed,name=origin,proto3,enum=o2.bookkeeping.LogOrigin" json:"origin,omitempty"`
	ParentLogId int32       `protobuf:"varint,6,opt,name=parentLogId,proto3" json:"parentLogId,omitempty"`
	Replies     int32       `protobuf:"varint,7,opt,name=Replies,proto3" json:"Replies,omitempty"`
	RootLogId   int32       `protobuf:"varint,8,opt,name=rootLogId,proto3" json:"rootLogId,omitempty"`
	// Array of minified Run objects.
	Runs       []*MinimalRun `protobuf:"bytes,9,rep,name=runs,proto3" json:"runs,omitempty"`
	Subsystems []*Subsystem  `protobuf:"bytes,10,rep,name=subsystems,proto3" json:"subsystems,omitempty"`
	Subtype    LogSubtype    `protobuf:"varint,11,opt,name=subtype,proto3,enum=o2.bookkeeping.LogSubtype" json:"subtype,omitempty"`
	Tags       []*Tag        `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	Text       string        `protobuf:"bytes,14,opt,name=text,proto3" json:"text,omitempty"`
	Title      string        `protobuf:"bytes,15,opt,name=title,proto3" json:"title,omitempty"`
	// Unix timestamp when this entity was last updated.
	UpdatedAt int64 `protobuf:"varint,16,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_protos_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{2}
}

func (x *Log) GetAttachments() []*Attachment {
	if x != nil {
		return x.Attachments
	}
	return nil
}

func (x *Log) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Log) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Log) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Log) GetOrigin() []LogOrigin {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *Log) GetParentLogId() int32 {
	if x != nil {
		return x.ParentLogId
	}
	return 0
}

func (x *Log) GetReplies() int32 {
	if x != nil {
		return x.Replies
	}
	return 0
}

func (x *Log) GetRootLogId() int32 {
	if x != nil {
		return x.RootLogId
	}
	return 0
}

func (x *Log) GetRuns() []*MinimalRun {
	if x != nil {
		return x.Runs
	}
	return nil
}

func (x *Log) GetSubsystems() []*Subsystem {
	if x != nil {
		return x.Subsystems
	}
	return nil
}

func (x *Log) GetSubtype() LogSubtype {
	if x != nil {
		return x.Subtype
	}
	return LogSubtype_LOG_SUBTYPE_NULL
}

func (x *Log) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Log) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Log) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Log) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type MinimalRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunNumber int32 `protobuf:"varint,1,opt,name=runNumber,proto3" json:"runNumber,omitempty"`
	Id        int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MinimalRun) Reset() {
	*x = MinimalRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MinimalRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MinimalRun) ProtoMessage() {}

func (x *MinimalRun) ProtoReflect() protoreflect.Message {
	mi := &file_protos_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MinimalRun.ProtoReflect.Descriptor instead.
func (*MinimalRun) Descriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{3}
}

func (x *MinimalRun) GetRunNumber() int32 {
	if x != nil {
		return x.RunNumber
	}
	return 0
}

func (x *MinimalRun) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Subsystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp when this entity was created.
	CreatedAt int64 `protobuf:"varint,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Id        int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// The label value of the subsystem.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Unix timestamp when this entity was last updated.
	UpdatedAt int64 `protobuf:"varint,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Subsystem) Reset() {
	*x = Subsystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subsystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subsystem) ProtoMessage() {}

func (x *Subsystem) ProtoReflect() protoreflect.Message {
	mi := &file_protos_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subsystem.ProtoReflect.Descriptor instead.
func (*Subsystem) Descriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{4}
}

func (x *Subsystem) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Subsystem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subsystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subsystem) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix timestamp when this entity was created.
	CreatedAt int64 `protobuf:"varint,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Id        int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// The label value of the tag.
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// The email for the tag.
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// All the mattermost tags for updates
	Mattermost string `protobuf:"bytes,5,opt,name=mattermost,proto3" json:"mattermost,omitempty"`
	// Unix timestamp when this entity was last updated.
	UpdatedAt int64 `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	// The last person that edited the email/mattermost fields
	LastEditedName string `protobuf:"bytes,7,opt,name=lastEditedName,proto3" json:"lastEditedName,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_protos_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_protos_log_proto_rawDescGZIP(), []int{5}
}

func (x *Tag) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Tag) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Tag) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Tag) GetMattermost() string {
	if x != nil {
		return x.Mattermost
	}
	return ""
}

func (x *Tag) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Tag) GetLastEditedName() string {
	if x != nil {
		return x.LastEditedName
	}
	return ""
}

var File_protos_log_proto protoreflect.FileDescriptor

var file_protos_log_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x12, 0x4c, 0x6f, 0x67, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x75, 0x6e, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x75,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x8e, 0x02, 0x0a, 0x0a, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xbe, 0x04, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x6f, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x72, 0x6f, 0x6f, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x75,
	0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x52, 0x75, 0x6e, 0x52, 0x04, 0x72, 0x75, 0x6e, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x75,
	0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x53, 0x75, 0x62, 0x74, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x73, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x32, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3a, 0x0a, 0x0a,
	0x4d, 0x69, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x75, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x75,
	0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72,
	0x75, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc3, 0x01, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6d,
	0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6d, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61, 0x73,
	0x74, 0x45, 0x64, 0x69, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x4e, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f,
	0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x4c, 0x4f, 0x47, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49, 0x4e, 0x5f, 0x48, 0x55, 0x4d, 0x41,
	0x4e, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x47, 0x5f, 0x4f, 0x52, 0x49, 0x47, 0x49,
	0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x2a, 0xa7, 0x01, 0x0a, 0x0a,
	0x4c, 0x6f, 0x67, 0x53, 0x75, 0x62, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x4f,
	0x47, 0x5f, 0x53, 0x55, 0x42, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x55, 0x42, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x52, 0x55, 0x4e, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x55, 0x42,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x42, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x02,
	0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x55, 0x42, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x4e, 0x4e, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x1c,
	0x0a, 0x18, 0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x55, 0x42, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x56, 0x45, 0x4e, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13,
	0x4c, 0x4f, 0x47, 0x5f, 0x53, 0x55, 0x42, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x05, 0x32, 0x4f, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e,
	0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x6f, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x6f, 0x32, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x69, 0x63, 0x65, 0x4f, 0x32, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x2f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6b,
	0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x62, 0x6b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_log_proto_rawDescOnce sync.Once
	file_protos_log_proto_rawDescData = file_protos_log_proto_rawDesc
)

func file_protos_log_proto_rawDescGZIP() []byte {
	file_protos_log_proto_rawDescOnce.Do(func() {
		file_protos_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_log_proto_rawDescData)
	})
	return file_protos_log_proto_rawDescData
}

var file_protos_log_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_log_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_log_proto_goTypes = []interface{}{
	(LogOrigin)(0),             // 0: o2.bookkeeping.LogOrigin
	(LogSubtype)(0),            // 1: o2.bookkeeping.LogSubtype
	(*LogCreationRequest)(nil), // 2: o2.bookkeeping.LogCreationRequest
	(*Attachment)(nil),         // 3: o2.bookkeeping.Attachment
	(*Log)(nil),                // 4: o2.bookkeeping.Log
	(*MinimalRun)(nil),         // 5: o2.bookkeeping.MinimalRun
	(*Subsystem)(nil),          // 6: o2.bookkeeping.Subsystem
	(*Tag)(nil),                // 7: o2.bookkeeping.Tag
	(*User)(nil),               // 8: o2.bookkeeping.User
}
var file_protos_log_proto_depIdxs = []int32{
	3, // 0: o2.bookkeeping.Log.attachments:type_name -> o2.bookkeeping.Attachment
	8, // 1: o2.bookkeeping.Log.author:type_name -> o2.bookkeeping.User
	0, // 2: o2.bookkeeping.Log.origin:type_name -> o2.bookkeeping.LogOrigin
	5, // 3: o2.bookkeeping.Log.runs:type_name -> o2.bookkeeping.MinimalRun
	6, // 4: o2.bookkeeping.Log.subsystems:type_name -> o2.bookkeeping.Subsystem
	1, // 5: o2.bookkeeping.Log.subtype:type_name -> o2.bookkeeping.LogSubtype
	7, // 6: o2.bookkeeping.Log.tags:type_name -> o2.bookkeeping.Tag
	2, // 7: o2.bookkeeping.LogService.Create:input_type -> o2.bookkeeping.LogCreationRequest
	4, // 8: o2.bookkeeping.LogService.Create:output_type -> o2.bookkeeping.Log
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_protos_log_proto_init() }
func file_protos_log_proto_init() {
	if File_protos_log_proto != nil {
		return
	}
	file_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogCreationRequest); i {
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
		file_protos_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attachment); i {
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
		file_protos_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_protos_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MinimalRun); i {
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
		file_protos_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subsystem); i {
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
		file_protos_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
			RawDescriptor: file_protos_log_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_log_proto_goTypes,
		DependencyIndexes: file_protos_log_proto_depIdxs,
		EnumInfos:         file_protos_log_proto_enumTypes,
		MessageInfos:      file_protos_log_proto_msgTypes,
	}.Build()
	File_protos_log_proto = out.File
	file_protos_log_proto_rawDesc = nil
	file_protos_log_proto_goTypes = nil
	file_protos_log_proto_depIdxs = nil
}