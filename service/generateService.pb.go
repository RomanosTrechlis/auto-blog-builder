// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generateService.proto

/*
Package generate is a generated protocol buffer package.

It is generated from these files:
	generateService.proto

It has these top-level messages:
	GenerateRequest
	StaticPages
	Theme
	DataSource
	Upload
	GenerateResponse
*/
package generate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GenerateRequest struct {
	Author            string         `protobuf:"bytes,1,opt,name=Author" json:"Author,omitempty"`
	BlogURL           string         `protobuf:"bytes,2,opt,name=BlogURL" json:"BlogURL,omitempty"`
	BlogLanguage      string         `protobuf:"bytes,3,opt,name=BlogLanguage" json:"BlogLanguage,omitempty"`
	BlogDescription   string         `protobuf:"bytes,4,opt,name=BlogDescription" json:"BlogDescription,omitempty"`
	DateFormat        string         `protobuf:"bytes,5,opt,name=DateFormat" json:"DateFormat,omitempty"`
	Theme             *Theme         `protobuf:"bytes,12,opt,name=Theme" json:"Theme,omitempty"`
	ThemeFolder       string         `protobuf:"bytes,6,opt,name=ThemeFolder" json:"ThemeFolder,omitempty"`
	BlogTitle         string         `protobuf:"bytes,7,opt,name=BlogTitle" json:"BlogTitle,omitempty"`
	NumPostsFrontPage int64          `protobuf:"varint,8,opt,name=NumPostsFrontPage" json:"NumPostsFrontPage,omitempty"`
	DataSource        *DataSource    `protobuf:"bytes,13,opt,name=DataSource" json:"DataSource,omitempty"`
	Upload            *Upload        `protobuf:"bytes,14,opt,name=Upload" json:"Upload,omitempty"`
	TempFolder        string         `protobuf:"bytes,9,opt,name=TempFolder" json:"TempFolder,omitempty"`
	DestFolder        string         `protobuf:"bytes,10,opt,name=DestFolder" json:"DestFolder,omitempty"`
	StaticPages       []*StaticPages `protobuf:"bytes,11,rep,name=StaticPages" json:"StaticPages,omitempty"`
}

func (m *GenerateRequest) Reset()                    { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()               {}
func (*GenerateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GenerateRequest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *GenerateRequest) GetBlogURL() string {
	if m != nil {
		return m.BlogURL
	}
	return ""
}

func (m *GenerateRequest) GetBlogLanguage() string {
	if m != nil {
		return m.BlogLanguage
	}
	return ""
}

func (m *GenerateRequest) GetBlogDescription() string {
	if m != nil {
		return m.BlogDescription
	}
	return ""
}

func (m *GenerateRequest) GetDateFormat() string {
	if m != nil {
		return m.DateFormat
	}
	return ""
}

func (m *GenerateRequest) GetTheme() *Theme {
	if m != nil {
		return m.Theme
	}
	return nil
}

func (m *GenerateRequest) GetThemeFolder() string {
	if m != nil {
		return m.ThemeFolder
	}
	return ""
}

func (m *GenerateRequest) GetBlogTitle() string {
	if m != nil {
		return m.BlogTitle
	}
	return ""
}

func (m *GenerateRequest) GetNumPostsFrontPage() int64 {
	if m != nil {
		return m.NumPostsFrontPage
	}
	return 0
}

func (m *GenerateRequest) GetDataSource() *DataSource {
	if m != nil {
		return m.DataSource
	}
	return nil
}

func (m *GenerateRequest) GetUpload() *Upload {
	if m != nil {
		return m.Upload
	}
	return nil
}

func (m *GenerateRequest) GetTempFolder() string {
	if m != nil {
		return m.TempFolder
	}
	return ""
}

func (m *GenerateRequest) GetDestFolder() string {
	if m != nil {
		return m.DestFolder
	}
	return ""
}

func (m *GenerateRequest) GetStaticPages() []*StaticPages {
	if m != nil {
		return m.StaticPages
	}
	return nil
}

type StaticPages struct {
	File       string `protobuf:"bytes,1,opt,name=File" json:"File,omitempty"`
	To         string `protobuf:"bytes,2,opt,name=To" json:"To,omitempty"`
	IsTemplate bool   `protobuf:"varint,3,opt,name=IsTemplate" json:"IsTemplate,omitempty"`
}

func (m *StaticPages) Reset()                    { *m = StaticPages{} }
func (m *StaticPages) String() string            { return proto.CompactTextString(m) }
func (*StaticPages) ProtoMessage()               {}
func (*StaticPages) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StaticPages) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *StaticPages) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *StaticPages) GetIsTemplate() bool {
	if m != nil {
		return m.IsTemplate
	}
	return false
}

type Theme struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository" json:"Repository,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=Type" json:"Type,omitempty"`
}

func (m *Theme) Reset()                    { *m = Theme{} }
func (m *Theme) String() string            { return proto.CompactTextString(m) }
func (*Theme) ProtoMessage()               {}
func (*Theme) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Theme) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Theme) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type DataSource struct {
	Repository string `protobuf:"bytes,1,opt,name=Repository" json:"Repository,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=Type" json:"Type,omitempty"`
}

func (m *DataSource) Reset()                    { *m = DataSource{} }
func (m *DataSource) String() string            { return proto.CompactTextString(m) }
func (*DataSource) ProtoMessage()               {}
func (*DataSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DataSource) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *DataSource) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Upload struct {
	URL      string `protobuf:"bytes,1,opt,name=URL" json:"URL,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password" json:"Password,omitempty"`
}

func (m *Upload) Reset()                    { *m = Upload{} }
func (m *Upload) String() string            { return proto.CompactTextString(m) }
func (*Upload) ProtoMessage()               {}
func (*Upload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Upload) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *Upload) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Upload) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GenerateResponse struct {
	Res string `protobuf:"bytes,1,opt,name=Res" json:"Res,omitempty"`
}

func (m *GenerateResponse) Reset()                    { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string            { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()               {}
func (*GenerateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GenerateResponse) GetRes() string {
	if m != nil {
		return m.Res
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateRequest)(nil), "generate.GenerateRequest")
	proto.RegisterType((*StaticPages)(nil), "generate.StaticPages")
	proto.RegisterType((*Theme)(nil), "generate.Theme")
	proto.RegisterType((*DataSource)(nil), "generate.DataSource")
	proto.RegisterType((*Upload)(nil), "generate.Upload")
	proto.RegisterType((*GenerateResponse)(nil), "generate.GenerateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Generator service

type GeneratorClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type generatorClient struct {
	cc *grpc.ClientConn
}

func NewGeneratorClient(cc *grpc.ClientConn) GeneratorClient {
	return &generatorClient{cc}
}

func (c *generatorClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := grpc.Invoke(ctx, "/generate.Generator/Generate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Generator service

type GeneratorServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
}

func RegisterGeneratorServer(s *grpc.Server, srv GeneratorServer) {
	s.RegisterService(&_Generator_serviceDesc, srv)
}

func _Generator_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneratorServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generate.Generator/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneratorServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generate.Generator",
	HandlerType: (*GeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Generator_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generateService.proto",
}

func init() { proto.RegisterFile("generateService.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x26, 0x4d, 0x9a, 0x3a, 0xe3, 0xd2, 0x84, 0x11, 0x45, 0x4b, 0x84, 0x90, 0x65, 0x81, 0xe4,
	0x03, 0xca, 0x21, 0x20, 0x71, 0xe0, 0xc2, 0x4f, 0x15, 0x84, 0x54, 0xa1, 0xb0, 0x71, 0x1e, 0x60,
	0x49, 0x46, 0xa9, 0x25, 0xdb, 0x6b, 0x76, 0xd7, 0xa0, 0x3e, 0x0d, 0xaf, 0x8a, 0x76, 0xbd, 0x8e,
	0xdd, 0xc2, 0x89, 0xdb, 0xcc, 0xf7, 0xcd, 0xdf, 0x7e, 0x3b, 0x03, 0x97, 0x07, 0x2a, 0x49, 0x09,
	0x43, 0x1b, 0x52, 0x3f, 0xb3, 0x1d, 0x2d, 0x2a, 0x25, 0x8d, 0xc4, 0xa0, 0x85, 0xe3, 0xdf, 0x23,
	0x98, 0x7e, 0xf6, 0x0e, 0xa7, 0x1f, 0x35, 0x69, 0x83, 0x4f, 0x60, 0xfc, 0xa1, 0x36, 0x37, 0x52,
	0xb1, 0x41, 0x34, 0x48, 0x26, 0xdc, 0x7b, 0xc8, 0xe0, 0xec, 0x63, 0x2e, 0x0f, 0x5b, 0x7e, 0xcd,
	0x4e, 0x1c, 0xd1, 0xba, 0x18, 0xc3, 0xb9, 0x35, 0xaf, 0x45, 0x79, 0xa8, 0xc5, 0x81, 0xd8, 0xd0,
	0xd1, 0x77, 0x30, 0x4c, 0x60, 0x6a, 0xfd, 0x2b, 0xd2, 0x3b, 0x95, 0x55, 0x26, 0x93, 0x25, 0x1b,
	0xb9, 0xb0, 0xfb, 0x30, 0x3e, 0x07, 0xb8, 0x12, 0x86, 0x56, 0x52, 0x15, 0xc2, 0xb0, 0x53, 0x17,
	0xd4, 0x43, 0xf0, 0x25, 0x9c, 0xa6, 0x37, 0x54, 0x10, 0x3b, 0x8f, 0x06, 0x49, 0xb8, 0x9c, 0x2e,
	0xda, 0xd7, 0x2c, 0x1c, 0xcc, 0x1b, 0x16, 0x23, 0x08, 0x9d, 0xb1, 0x92, 0xf9, 0x9e, 0x14, 0x1b,
	0xbb, 0x3a, 0x7d, 0x08, 0x9f, 0xc1, 0xc4, 0xf6, 0x4e, 0x33, 0x93, 0x13, 0x3b, 0x73, 0x7c, 0x07,
	0xe0, 0x2b, 0x78, 0xf4, 0xb5, 0x2e, 0xd6, 0x52, 0x1b, 0xbd, 0x52, 0xb2, 0x34, 0x6b, 0xfb, 0xb2,
	0x20, 0x1a, 0x24, 0x43, 0xfe, 0x37, 0x81, 0x6f, 0xdc, 0xd0, 0x62, 0x23, 0x6b, 0xb5, 0x23, 0xf6,
	0xd0, 0x4d, 0xf6, 0xb8, 0x9b, 0xac, 0xe3, 0x78, 0x2f, 0x0e, 0x13, 0x18, 0x6f, 0xab, 0x5c, 0x8a,
	0x3d, 0xbb, 0x70, 0x19, 0xb3, 0x2e, 0xa3, 0xc1, 0xb9, 0xe7, 0xad, 0x28, 0x29, 0x15, 0x95, 0x7f,
	0xcc, 0xa4, 0x11, 0xa5, 0x43, 0x9c, 0x68, 0xa4, 0x8d, 0xe7, 0xc1, 0x8b, 0x76, 0x44, 0xf0, 0x2d,
	0x84, 0x1b, 0x23, 0x4c, 0xb6, 0xb3, 0xd3, 0x6a, 0x16, 0x46, 0xc3, 0x24, 0x5c, 0x5e, 0x76, 0xed,
	0x7a, 0x24, 0xef, 0x47, 0xc6, 0xdf, 0xee, 0x24, 0x22, 0xc2, 0x68, 0x95, 0xe5, 0xe4, 0x57, 0xc3,
	0xd9, 0x78, 0x01, 0x27, 0xa9, 0xf4, 0x3b, 0x71, 0x92, 0x4a, 0x3b, 0xcb, 0x17, 0x6d, 0x67, 0xcb,
	0x85, 0x69, 0x96, 0x21, 0xe0, 0x3d, 0x24, 0x7e, 0xe7, 0x3f, 0xd0, 0x06, 0x72, 0xaa, 0xa4, 0xce,
	0x8c, 0x54, 0xb7, 0xbe, 0x64, 0x0f, 0xb1, 0xcd, 0xd2, 0xdb, 0x8a, 0x7c, 0x69, 0x67, 0xc7, 0xef,
	0xfb, 0x42, 0xff, 0x57, 0x05, 0xde, 0x8a, 0x8e, 0x33, 0x18, 0xda, 0x6d, 0x6e, 0xd2, 0xac, 0x89,
	0x73, 0x08, 0xb6, 0x9a, 0x54, 0x29, 0x8a, 0x36, 0xe7, 0xe8, 0x5b, 0x6e, 0x2d, 0xb4, 0xfe, 0x25,
	0xd5, 0xde, 0x6f, 0xf8, 0xd1, 0x8f, 0x5f, 0xc0, 0xac, 0x3b, 0x23, 0x5d, 0xc9, 0x52, 0x93, 0xad,
	0xce, 0x49, 0xb7, 0xd5, 0x39, 0xe9, 0xe5, 0x1a, 0x26, 0x3e, 0x4a, 0x2a, 0xfc, 0x04, 0x41, 0x9b,
	0x82, 0x4f, 0xbb, 0x8f, 0xb8, 0x77, 0x8d, 0xf3, 0xf9, 0xbf, 0xa8, 0xa6, 0x43, 0xfc, 0xe0, 0xfb,
	0xd8, 0x1d, 0xf4, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x0e, 0x02, 0x86, 0xe9, 0x03,
	0x00, 0x00,
}
