// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: server.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoReq) Reset() {
	*x = InfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoReq) ProtoMessage() {}

func (x *InfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoReq.ProtoReflect.Descriptor instead.
func (*InfoReq) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

type InfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey                     string                  `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Host                          *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                          int32                   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	HostVpnIp                     string                  `protobuf:"bytes,4,opt,name=host_vpn_ip,json=hostVpnIp,proto3" json:"host_vpn_ip,omitempty"`
	MetadataEnabled               bool                    `protobuf:"varint,5,opt,name=metadata_enabled,json=metadataEnabled,proto3" json:"metadata_enabled,omitempty"`
	IsAdmin                       bool                    `protobuf:"varint,6,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	AllowedIps                    string                  `protobuf:"bytes,7,opt,name=allowed_ips,json=allowedIps,proto3" json:"allowed_ips,omitempty"`
	DnsEnabled                    bool                    `protobuf:"varint,8,opt,name=dns_enabled,json=dnsEnabled,proto3" json:"dns_enabled,omitempty"`
	DnsAddress                    string                  `protobuf:"bytes,9,opt,name=dns_address,json=dnsAddress,proto3" json:"dns_address,omitempty"`
	Filename                      string                  `protobuf:"bytes,10,opt,name=filename,proto3" json:"filename,omitempty"`
	InactiveDeviceDeletionEnabled bool                    `protobuf:"varint,11,opt,name=inactive_device_deletion_enabled,json=inactiveDeviceDeletionEnabled,proto3" json:"inactive_device_deletion_enabled,omitempty"`
	InactiveDeviceGracePeriod     *durationpb.Duration    `protobuf:"bytes,12,opt,name=inactive_device_grace_period,json=inactiveDeviceGracePeriod,proto3" json:"inactive_device_grace_period,omitempty"`
	ClientConfigDnsServers        string                  `protobuf:"bytes,13,opt,name=client_config_dns_servers,json=clientConfigDnsServers,proto3" json:"client_config_dns_servers,omitempty"`
	ClientConfigDnsSearchDomain   string                  `protobuf:"bytes,14,opt,name=client_config_dns_search_domain,json=clientConfigDnsSearchDomain,proto3" json:"client_config_dns_search_domain,omitempty"`
}

func (x *InfoRes) Reset() {
	*x = InfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRes) ProtoMessage() {}

func (x *InfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRes.ProtoReflect.Descriptor instead.
func (*InfoRes) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *InfoRes) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *InfoRes) GetHost() *wrapperspb.StringValue {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *InfoRes) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *InfoRes) GetHostVpnIp() string {
	if x != nil {
		return x.HostVpnIp
	}
	return ""
}

func (x *InfoRes) GetMetadataEnabled() bool {
	if x != nil {
		return x.MetadataEnabled
	}
	return false
}

func (x *InfoRes) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *InfoRes) GetAllowedIps() string {
	if x != nil {
		return x.AllowedIps
	}
	return ""
}

func (x *InfoRes) GetDnsEnabled() bool {
	if x != nil {
		return x.DnsEnabled
	}
	return false
}

func (x *InfoRes) GetDnsAddress() string {
	if x != nil {
		return x.DnsAddress
	}
	return ""
}

func (x *InfoRes) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *InfoRes) GetInactiveDeviceDeletionEnabled() bool {
	if x != nil {
		return x.InactiveDeviceDeletionEnabled
	}
	return false
}

func (x *InfoRes) GetInactiveDeviceGracePeriod() *durationpb.Duration {
	if x != nil {
		return x.InactiveDeviceGracePeriod
	}
	return nil
}

func (x *InfoRes) GetClientConfigDnsServers() string {
	if x != nil {
		return x.ClientConfigDnsServers
	}
	return ""
}

func (x *InfoRes) GetClientConfigDnsSearchDomain() string {
	if x != nil {
		return x.ClientConfigDnsSearchDomain
	}
	return ""
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x22, 0xf9, 0x04, 0x0a, 0x07, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x76, 0x70, 0x6e, 0x5f, 0x69, 0x70,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x73, 0x74, 0x56, 0x70, 0x6e, 0x49,
	0x70, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x5f, 0x69, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x49, 0x70, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x64,
	0x6e, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6e, 0x73,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x6e, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x20, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x1d, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x5a, 0x0a, 0x1c, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x19, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x47, 0x72, 0x61, 0x63, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x6e, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x1f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x6e, 0x73,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0x32, 0x0a, 0x06,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66,
	0x72, 0x65, 0x69, 0x66, 0x75, 0x6e, 0x6b, 0x4d, 0x55, 0x43, 0x2f, 0x77, 0x67, 0x2d, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_proto_goTypes = []interface{}{
	(*InfoReq)(nil),                // 0: proto.InfoReq
	(*InfoRes)(nil),                // 1: proto.InfoRes
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*durationpb.Duration)(nil),    // 3: google.protobuf.Duration
}
var file_server_proto_depIdxs = []int32{
	2, // 0: proto.InfoRes.host:type_name -> google.protobuf.StringValue
	3, // 1: proto.InfoRes.inactive_device_grace_period:type_name -> google.protobuf.Duration
	0, // 2: proto.Server.Info:input_type -> proto.InfoReq
	1, // 3: proto.Server.Info:output_type -> proto.InfoRes
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoReq); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRes); i {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
