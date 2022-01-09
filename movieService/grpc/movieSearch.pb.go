// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: movieSearch.proto

package __

import (
	context "context"
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

type MovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Page    int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *MovieRequest) Reset() {
	*x = MovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movieSearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRequest) ProtoMessage() {}

func (x *MovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_movieSearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRequest.ProtoReflect.Descriptor instead.
func (*MovieRequest) Descriptor() ([]byte, []int) {
	return file_movieSearch_proto_rawDescGZIP(), []int{0}
}

func (x *MovieRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *MovieRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type MovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movies []*Movie `protobuf:"bytes,1,rep,name=movies,proto3" json:"movies,omitempty"`
	Err    string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *MovieResponse) Reset() {
	*x = MovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movieSearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieResponse) ProtoMessage() {}

func (x *MovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_movieSearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieResponse.ProtoReflect.Descriptor instead.
func (*MovieResponse) Descriptor() ([]byte, []int) {
	return file_movieSearch_proto_rawDescGZIP(), []int{1}
}

func (x *MovieResponse) GetMovies() []*Movie {
	if x != nil {
		return x.Movies
	}
	return nil
}

func (x *MovieResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Type   string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ImdbID string `protobuf:"bytes,4,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_movieSearch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_movieSearch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_movieSearch_proto_rawDescGZIP(), []int{2}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Movie) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Movie) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *Movie) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

var File_movieSearch_proto protoreflect.FileDescriptor

var file_movieSearch_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x3c, 0x0a, 0x0c, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x0d, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x06, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22,
	0x75, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x32, 0x43, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e,
	0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_movieSearch_proto_rawDescOnce sync.Once
	file_movieSearch_proto_rawDescData = file_movieSearch_proto_rawDesc
)

func file_movieSearch_proto_rawDescGZIP() []byte {
	file_movieSearch_proto_rawDescOnce.Do(func() {
		file_movieSearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_movieSearch_proto_rawDescData)
	})
	return file_movieSearch_proto_rawDescData
}

var file_movieSearch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_movieSearch_proto_goTypes = []interface{}{
	(*MovieRequest)(nil),  // 0: grpc.MovieRequest
	(*MovieResponse)(nil), // 1: grpc.MovieResponse
	(*Movie)(nil),         // 2: grpc.Movie
}
var file_movieSearch_proto_depIdxs = []int32{
	2, // 0: grpc.MovieResponse.movies:type_name -> grpc.Movie
	0, // 1: grpc.MovieSearch.GetMovies:input_type -> grpc.MovieRequest
	1, // 2: grpc.MovieSearch.GetMovies:output_type -> grpc.MovieResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_movieSearch_proto_init() }
func file_movieSearch_proto_init() {
	if File_movieSearch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_movieSearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRequest); i {
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
		file_movieSearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieResponse); i {
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
		file_movieSearch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
			RawDescriptor: file_movieSearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_movieSearch_proto_goTypes,
		DependencyIndexes: file_movieSearch_proto_depIdxs,
		MessageInfos:      file_movieSearch_proto_msgTypes,
	}.Build()
	File_movieSearch_proto = out.File
	file_movieSearch_proto_rawDesc = nil
	file_movieSearch_proto_goTypes = nil
	file_movieSearch_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MovieSearchClient is the client API for MovieSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieSearchClient interface {
	GetMovies(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error)
}

type movieSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieSearchClient(cc grpc.ClientConnInterface) MovieSearchClient {
	return &movieSearchClient{cc}
}

func (c *movieSearchClient) GetMovies(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieResponse, error) {
	out := new(MovieResponse)
	err := c.cc.Invoke(ctx, "/grpc.MovieSearch/GetMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieSearchServer is the server API for MovieSearch service.
type MovieSearchServer interface {
	GetMovies(context.Context, *MovieRequest) (*MovieResponse, error)
}

// UnimplementedMovieSearchServer can be embedded to have forward compatible implementations.
type UnimplementedMovieSearchServer struct {
}

func (*UnimplementedMovieSearchServer) GetMovies(context.Context, *MovieRequest) (*MovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}

func RegisterMovieSearchServer(s *grpc.Server, srv MovieSearchServer) {
	s.RegisterService(&_MovieSearch_serviceDesc, srv)
}

func _MovieSearch_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieSearchServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MovieSearch/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieSearchServer).GetMovies(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieSearch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MovieSearch",
	HandlerType: (*MovieSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovies",
			Handler:    _MovieSearch_GetMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movieSearch.proto",
}