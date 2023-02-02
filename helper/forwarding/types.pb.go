// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: helper/forwarding/types.proto

package forwarding

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Not used right now but reserving in case it turns out that streaming
	// makes things more economical on the gRPC side
	// uint64 id = 1;
	Method           string                  `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Url              *URL                    `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	HeaderEntries    map[string]*HeaderEntry `protobuf:"bytes,4,rep,name=header_entries,json=headerEntries,proto3" json:"header_entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body             []byte                  `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Host             string                  `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	RemoteAddr       string                  `protobuf:"bytes,7,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	PeerCertificates [][]byte                `protobuf:"bytes,8,rep,name=peer_certificates,json=peerCertificates,proto3" json:"peer_certificates,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helper_forwarding_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_helper_forwarding_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_helper_forwarding_types_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetUrl() *URL {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *Request) GetHeaderEntries() map[string]*HeaderEntry {
	if x != nil {
		return x.HeaderEntries
	}
	return nil
}

func (x *Request) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Request) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Request) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *Request) GetPeerCertificates() [][]byte {
	if x != nil {
		return x.PeerCertificates
	}
	return nil
}

type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Opaque string `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
	// This isn't needed now but might be in the future, so we'll skip the
	// number to keep the ordering in net/url
	// UserInfo user = 3;
	Host    string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Path    string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	RawPath string `protobuf:"bytes,6,opt,name=raw_path,json=rawPath,proto3" json:"raw_path,omitempty"`
	// This also isn't needed right now, but we'll reserve the number
	// bool force_query = 7;
	RawQuery string `protobuf:"bytes,8,opt,name=raw_query,json=rawQuery,proto3" json:"raw_query,omitempty"`
	Fragment string `protobuf:"bytes,9,opt,name=fragment,proto3" json:"fragment,omitempty"`
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helper_forwarding_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_helper_forwarding_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URL.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_helper_forwarding_types_proto_rawDescGZIP(), []int{1}
}

func (x *URL) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *URL) GetOpaque() string {
	if x != nil {
		return x.Opaque
	}
	return ""
}

func (x *URL) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *URL) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *URL) GetRawPath() string {
	if x != nil {
		return x.RawPath
	}
	return ""
}

func (x *URL) GetRawQuery() string {
	if x != nil {
		return x.RawQuery
	}
	return ""
}

func (x *URL) GetFragment() string {
	if x != nil {
		return x.Fragment
	}
	return ""
}

type HeaderEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *HeaderEntry) Reset() {
	*x = HeaderEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helper_forwarding_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderEntry) ProtoMessage() {}

func (x *HeaderEntry) ProtoReflect() protoreflect.Message {
	mi := &file_helper_forwarding_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderEntry.ProtoReflect.Descriptor instead.
func (*HeaderEntry) Descriptor() ([]byte, []int) {
	return file_helper_forwarding_types_proto_rawDescGZIP(), []int{2}
}

func (x *HeaderEntry) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Not used right now but reserving in case it turns out that streaming
	// makes things more economical on the gRPC side
	// uint64 id = 1;
	StatusCode uint32 `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Body       []byte `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Added in 0.6.2 to ensure that the content-type is set appropriately, as
	// well as any other information
	HeaderEntries map[string]*HeaderEntry `protobuf:"bytes,4,rep,name=header_entries,json=headerEntries,proto3" json:"header_entries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LastRemoteWal uint64                  `protobuf:"varint,5,opt,name=last_remote_wal,json=lastRemoteWal,proto3" json:"last_remote_wal,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helper_forwarding_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_helper_forwarding_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_helper_forwarding_types_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetStatusCode() uint32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Response) GetHeaderEntries() map[string]*HeaderEntry {
	if x != nil {
		return x.HeaderEntries
	}
	return nil
}

func (x *Response) GetLastRemoteWal() uint64 {
	if x != nil {
		return x.LastRemoteWal
	}
	return 0
}

var File_helper_forwarding_types_proto protoreflect.FileDescriptor

var file_helper_forwarding_types_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22, 0xe4, 0x02, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x21, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x52, 0x4c, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x4d, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x10, 0x70, 0x65, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x1a, 0x59, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xb1, 0x01, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x77, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x61, 0x77, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x92, 0x02,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12,
	0x4e, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x77,
	0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x1a, 0x59, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x2f, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helper_forwarding_types_proto_rawDescOnce sync.Once
	file_helper_forwarding_types_proto_rawDescData = file_helper_forwarding_types_proto_rawDesc
)

func file_helper_forwarding_types_proto_rawDescGZIP() []byte {
	file_helper_forwarding_types_proto_rawDescOnce.Do(func() {
		file_helper_forwarding_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_helper_forwarding_types_proto_rawDescData)
	})
	return file_helper_forwarding_types_proto_rawDescData
}

var file_helper_forwarding_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_helper_forwarding_types_proto_goTypes = []interface{}{
	(*Request)(nil),     // 0: forwarding.Request
	(*URL)(nil),         // 1: forwarding.URL
	(*HeaderEntry)(nil), // 2: forwarding.HeaderEntry
	(*Response)(nil),    // 3: forwarding.Response
	nil,                 // 4: forwarding.Request.HeaderEntriesEntry
	nil,                 // 5: forwarding.Response.HeaderEntriesEntry
}
var file_helper_forwarding_types_proto_depIdxs = []int32{
	1, // 0: forwarding.Request.url:type_name -> forwarding.URL
	4, // 1: forwarding.Request.header_entries:type_name -> forwarding.Request.HeaderEntriesEntry
	5, // 2: forwarding.Response.header_entries:type_name -> forwarding.Response.HeaderEntriesEntry
	2, // 3: forwarding.Request.HeaderEntriesEntry.value:type_name -> forwarding.HeaderEntry
	2, // 4: forwarding.Response.HeaderEntriesEntry.value:type_name -> forwarding.HeaderEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_helper_forwarding_types_proto_init() }
func file_helper_forwarding_types_proto_init() {
	if File_helper_forwarding_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helper_forwarding_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_helper_forwarding_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
		file_helper_forwarding_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderEntry); i {
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
		file_helper_forwarding_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_helper_forwarding_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_helper_forwarding_types_proto_goTypes,
		DependencyIndexes: file_helper_forwarding_types_proto_depIdxs,
		MessageInfos:      file_helper_forwarding_types_proto_msgTypes,
	}.Build()
	File_helper_forwarding_types_proto = out.File
	file_helper_forwarding_types_proto_rawDesc = nil
	file_helper_forwarding_types_proto_goTypes = nil
	file_helper_forwarding_types_proto_depIdxs = nil
}
