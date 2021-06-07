// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package error

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

type SearchError int32

const (
	//初始值, 无意义
	SearchError_INIT          SearchError = 0
	SearchError_SEARCH_FAILED SearchError = 10001
)

var SearchError_name = map[int32]string{
	0:     "INIT",
	10001: "SEARCH_FAILED",
}

var SearchError_value = map[string]int32{
	"INIT":          0,
	"SEARCH_FAILED": 10001,
}

func (x SearchError) String() string {
	return proto.EnumName(SearchError_name, int32(x))
}

func (SearchError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func init() {
	proto.RegisterEnum("error.SearchError", SearchError_name, SearchError_value)
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 90 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2d, 0x2a, 0xca, 0x2f, 0xd2,
	0xd2, 0xe6, 0xe2, 0x0e, 0x06, 0x0b, 0xbb, 0x82, 0xb8, 0x42, 0x1c, 0x5c, 0x2c, 0x9e, 0x7e, 0x9e,
	0x21, 0x02, 0x0c, 0x42, 0x42, 0x5c, 0xbc, 0xc1, 0xae, 0x8e, 0x41, 0xce, 0x1e, 0xf1, 0x6e, 0x8e,
	0x9e, 0x3e, 0xae, 0x2e, 0x02, 0x13, 0xfd, 0x92, 0xd8, 0xc0, 0x5a, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x74, 0xce, 0xd5, 0x0c, 0x4a, 0x00, 0x00, 0x00,
}
