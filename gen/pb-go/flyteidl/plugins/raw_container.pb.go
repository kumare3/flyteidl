// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/raw_container.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// MetadataFormat decides the encoding format in which the input metadata should be made available to the containers.
// If the user has access to the protocol buffer definitions, it is recommended to use the PROTO format.
// JSON and YAML do not need any protobuf definitions to read it
// All remote references in core.LiteralMap are replaced with local filesystem references (the data is downloaded to local filesystem)
type CoPilot_MetadataFormat int32

const (
	// JSON / YAML are serialized represnetation of a map[string]primitive. primitive -> int, string, bool, double/float, bytes
	CoPilot_JSON CoPilot_MetadataFormat = 0
	CoPilot_YAML CoPilot_MetadataFormat = 1
	// Proto is a serialized binary of `core.LiteralMap` defined in flyteidl/core
	CoPilot_PROTO CoPilot_MetadataFormat = 2
)

var CoPilot_MetadataFormat_name = map[int32]string{
	0: "JSON",
	1: "YAML",
	2: "PROTO",
}

var CoPilot_MetadataFormat_value = map[string]int32{
	"JSON":  0,
	"YAML":  1,
	"PROTO": 2,
}

func (x CoPilot_MetadataFormat) String() string {
	return proto.EnumName(CoPilot_MetadataFormat_name, int32(x))
}

func (CoPilot_MetadataFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_aa9a2ab24b00eebe, []int{0, 0}
}

// This configuration allows executing raw containers in Flyte using the Flyte CoPilot system.
// Flyte CoPilot, eliminates the needs of flytekit or sdk inside the container. Any inputs required by the users container are side-loaded in the input_path
// Any outputs generated by the user container - within output_path are automatically uploaded.
// We are starting this as a plugin, but after the container within Pod ordering stabilizes, we can move this
// into the default task definition.
type CoPilot struct {
	// File system path (start at root). This folder will contain all the inputs exploded to a separate file.
	// Example, if the input interface needs (x: int, y: blob, z: multipart_blob) and the input path is "/var/flyte/inputs", then the file system will look like
	// /var/flyte/inputs/inputs.<metadata format dependent -> .pb .json .yaml> -> Format as defined previously. The Blob and Multipart blob will reference local filesystem instead of remote locations
	// /var/flyte/inputs/x -> X is a file that contains the value of x (integer) in string format
	// /var/flyte/inputs/y -> Y is a file in Binary format
	// /var/flyte/inputs/z/... -> Note Z itself is a directory
	// More information about the protocol - refer to docs #TODO reference docs here
	InputPath string `protobuf:"bytes,1,opt,name=input_path,json=inputPath,proto3" json:"input_path,omitempty"`
	// File system path (start at root). This folder should contain all the outputs for the task as individual files and/or an error text file
	OutputPath uint32 `protobuf:"varint,2,opt,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	// In the inputs folder, there will be an additional summary/metadata file that contains references to all files or inlined primitive values.
	// This format decides the actual encoding for the data. Refer to the encoding to understand the specifics of the contents and the encoding
	Format               CoPilot_MetadataFormat `protobuf:"varint,3,opt,name=format,proto3,enum=flyteidl.plugins.CoPilot_MetadataFormat" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CoPilot) Reset()         { *m = CoPilot{} }
func (m *CoPilot) String() string { return proto.CompactTextString(m) }
func (*CoPilot) ProtoMessage()    {}
func (*CoPilot) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa9a2ab24b00eebe, []int{0}
}

func (m *CoPilot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoPilot.Unmarshal(m, b)
}
func (m *CoPilot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoPilot.Marshal(b, m, deterministic)
}
func (m *CoPilot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoPilot.Merge(m, src)
}
func (m *CoPilot) XXX_Size() int {
	return xxx_messageInfo_CoPilot.Size(m)
}
func (m *CoPilot) XXX_DiscardUnknown() {
	xxx_messageInfo_CoPilot.DiscardUnknown(m)
}

var xxx_messageInfo_CoPilot proto.InternalMessageInfo

func (m *CoPilot) GetInputPath() string {
	if m != nil {
		return m.InputPath
	}
	return ""
}

func (m *CoPilot) GetOutputPath() uint32 {
	if m != nil {
		return m.OutputPath
	}
	return 0
}

func (m *CoPilot) GetFormat() CoPilot_MetadataFormat {
	if m != nil {
		return m.Format
	}
	return CoPilot_JSON
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.CoPilot_MetadataFormat", CoPilot_MetadataFormat_name, CoPilot_MetadataFormat_value)
	proto.RegisterType((*CoPilot)(nil), "flyteidl.plugins.CoPilot")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/raw_container.proto", fileDescriptor_aa9a2ab24b00eebe)
}

var fileDescriptor_aa9a2ab24b00eebe = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xdd, 0xaa, 0xd5, 0x8c, 0x58, 0xc2, 0x9e, 0x7a, 0x11, 0x43, 0xf1, 0x90, 0x8b, 0xbb,
	0x60, 0xf1, 0xee, 0x07, 0x78, 0x10, 0x6b, 0x42, 0xf4, 0xa2, 0x97, 0x32, 0x69, 0xf3, 0xb1, 0xb0,
	0xdd, 0x59, 0xe2, 0x04, 0xe9, 0x7f, 0xf3, 0xc7, 0x89, 0x31, 0x16, 0xcc, 0x6d, 0x78, 0xdf, 0x67,
	0x78, 0xe1, 0x81, 0x8b, 0xd2, 0x6e, 0xb9, 0x30, 0x6b, 0xab, 0xbd, 0x6d, 0x2b, 0xe3, 0x3e, 0x74,
	0x83, 0x9f, 0xcb, 0x15, 0x39, 0x46, 0xe3, 0x8a, 0x46, 0xf9, 0x86, 0x98, 0x64, 0xf8, 0x47, 0xa9,
	0x9e, 0x9a, 0x7d, 0x09, 0x38, 0xba, 0xa7, 0xd4, 0x58, 0x62, 0x79, 0x06, 0x60, 0x9c, 0x6f, 0x79,
	0xe9, 0x91, 0xeb, 0xa9, 0x88, 0x44, 0x1c, 0x64, 0x41, 0x97, 0xa4, 0xc8, 0xb5, 0x3c, 0x87, 0x13,
	0x6a, 0x79, 0xd7, 0x8f, 0x22, 0x11, 0x9f, 0x66, 0xf0, 0x1b, 0x75, 0xc0, 0x0d, 0x8c, 0x4b, 0x6a,
	0x36, 0xc8, 0xd3, 0xfd, 0x48, 0xc4, 0x93, 0xab, 0x58, 0x0d, 0xe7, 0x54, 0x3f, 0xa5, 0x16, 0x05,
	0xe3, 0x1a, 0x19, 0x1f, 0x3a, 0x3e, 0xeb, 0xff, 0x66, 0x1a, 0x26, 0xff, 0x1b, 0x79, 0x0c, 0x07,
	0x8f, 0x2f, 0xc9, 0x73, 0xb8, 0xf7, 0x73, 0xbd, 0xdd, 0x2e, 0x9e, 0x42, 0x21, 0x03, 0x38, 0x4c,
	0xb3, 0xe4, 0x35, 0x09, 0x47, 0x77, 0xd7, 0xef, 0xf3, 0xca, 0x70, 0xdd, 0xe6, 0x6a, 0x45, 0x1b,
	0x6d, 0xb7, 0x25, 0xeb, 0x9d, 0x88, 0xaa, 0x70, 0xda, 0xe7, 0x97, 0x15, 0xe9, 0xa1, 0x9b, 0x7c,
	0xdc, 0xe9, 0x98, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x47, 0x6c, 0xa2, 0x6d, 0x36, 0x01, 0x00,
	0x00,
}