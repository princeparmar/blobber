// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: blobber_service.proto

package blobbergrpc

import (
	_ "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_blobber_service_proto protoreflect.FileDescriptor

var file_blobber_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72,
	0x1a, 0x16, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfb, 0x11, 0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x62, 0x62,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x62, 0x6c, 0x6f,
	0x62, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6c, 0x6f, 0x62,
	0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x7b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22,
	0x1a, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x7b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x73,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1c,
	0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62,
	0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x22, 0x1b, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d,
	0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x78, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76,
	0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x70, 0x61, 0x74,
	0x68, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x84,
	0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x20, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25,
	0x12, 0x23, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x12, 0x78, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f,
	0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x74, 0x72,
	0x65, 0x65, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x12,
	0x76, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x0c, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65,
	0x72, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1c, 0x2f, 0x76,
	0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x7b, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xb4, 0x01,
	0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x62,
	0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x67, 0x22, 0x1c, 0x2f,
	0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x7b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x21,
	0x1a, 0x1c, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01,
	0x2a, 0x5a, 0x21, 0x2a, 0x1c, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x68, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x22, 0x22, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2f, 0x7b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x7e,
	0x0a, 0x0d, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x68, 0x61, 0x73, 0x68, 0x2f, 0x7b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x7e,
	0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x54, 0x78, 0x6e, 0x12,
	0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x74, 0x61, 0x74, 0x78, 0x6e, 0x2f, 0x7b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x96,
	0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x62, 0x6c, 0x6f, 0x62,
	0x62, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x25, 0x22, 0x20, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x6c, 0x0a, 0x0a, 0x43, 0x6f, 0x70, 0x79, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x70, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x70, 0x79,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x2f, 0x63, 0x6f, 0x70, 0x79, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0xc9, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x7c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x76, 0x22, 0x22, 0x2f, 0x76, 0x32,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a,
	0x01, 0x2a, 0x5a, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x7b, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x5a, 0x27, 0x2a, 0x22, 0x2f, 0x76, 0x32, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01,
	0x2a, 0x12, 0xc0, 0x01, 0x0a, 0x14, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e, 0x62, 0x6c, 0x6f,
	0x62, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x68, 0x61, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x55, 0x22,
	0x26, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x3a, 0x01, 0x2a, 0x5a, 0x28, 0x2a, 0x26, 0x2f, 0x76,
	0x32, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x42, 0x2c, 0x5a, 0x2a, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x6f, 0x2f,
	0x30, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x62,
	0x65, 0x72, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x62, 0x65, 0x72, 0x67, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_blobber_service_proto_goTypes = []interface{}{
	(*GetAllocationRequest)(nil),           // 0: blobber.GetAllocationRequest
	(*GetFileMetaDataRequest)(nil),         // 1: blobber.GetFileMetaDataRequest
	(*GetFileStatsRequest)(nil),            // 2: blobber.GetFileStatsRequest
	(*ListEntitiesRequest)(nil),            // 3: blobber.ListEntitiesRequest
	(*GetObjectPathRequest)(nil),           // 4: blobber.GetObjectPathRequest
	(*GetReferencePathRequest)(nil),        // 5: blobber.GetReferencePathRequest
	(*GetObjectTreeRequest)(nil),           // 6: blobber.GetObjectTreeRequest
	(*DownloadFileRequest)(nil),            // 7: blobber.DownloadFileRequest
	(*RenameObjectRequest)(nil),            // 8: blobber.RenameObjectRequest
	(*UploadFileRequest)(nil),              // 9: blobber.UploadFileRequest
	(*CommitRequest)(nil),                  // 10: blobber.CommitRequest
	(*CalculateHashRequest)(nil),           // 11: blobber.CalculateHashRequest
	(*CommitMetaTxnRequest)(nil),           // 12: blobber.CommitMetaTxnRequest
	(*UpdateObjectAttributesRequest)(nil),  // 13: blobber.UpdateObjectAttributesRequest
	(*CopyObjectRequest)(nil),              // 14: blobber.CopyObjectRequest
	(*CollaboratorRequest)(nil),            // 15: blobber.CollaboratorRequest
	(*MarketplaceShareInfoRequest)(nil),    // 16: blobber.MarketplaceShareInfoRequest
	(*GetAllocationResponse)(nil),          // 17: blobber.GetAllocationResponse
	(*GetFileMetaDataResponse)(nil),        // 18: blobber.GetFileMetaDataResponse
	(*GetFileStatsResponse)(nil),           // 19: blobber.GetFileStatsResponse
	(*ListEntitiesResponse)(nil),           // 20: blobber.ListEntitiesResponse
	(*GetObjectPathResponse)(nil),          // 21: blobber.GetObjectPathResponse
	(*GetReferencePathResponse)(nil),       // 22: blobber.GetReferencePathResponse
	(*GetObjectTreeResponse)(nil),          // 23: blobber.GetObjectTreeResponse
	(*DownloadFileResponse)(nil),           // 24: blobber.DownloadFileResponse
	(*RenameObjectResponse)(nil),           // 25: blobber.RenameObjectResponse
	(*UploadFileResponse)(nil),             // 26: blobber.UploadFileResponse
	(*CommitResponse)(nil),                 // 27: blobber.CommitResponse
	(*CalculateHashResponse)(nil),          // 28: blobber.CalculateHashResponse
	(*CommitMetaTxnResponse)(nil),          // 29: blobber.CommitMetaTxnResponse
	(*UpdateObjectAttributesResponse)(nil), // 30: blobber.UpdateObjectAttributesResponse
	(*CopyObjectResponse)(nil),             // 31: blobber.CopyObjectResponse
	(*CollaboratorResponse)(nil),           // 32: blobber.CollaboratorResponse
	(*MarketplaceShareInfoResponse)(nil),   // 33: blobber.MarketplaceShareInfoResponse
}
var file_blobber_service_proto_depIdxs = []int32{
	0,  // 0: blobber.BlobberService.GetAllocation:input_type -> blobber.GetAllocationRequest
	1,  // 1: blobber.BlobberService.GetFileMetaData:input_type -> blobber.GetFileMetaDataRequest
	2,  // 2: blobber.BlobberService.GetFileStats:input_type -> blobber.GetFileStatsRequest
	3,  // 3: blobber.BlobberService.ListEntities:input_type -> blobber.ListEntitiesRequest
	4,  // 4: blobber.BlobberService.GetObjectPath:input_type -> blobber.GetObjectPathRequest
	5,  // 5: blobber.BlobberService.GetReferencePath:input_type -> blobber.GetReferencePathRequest
	6,  // 6: blobber.BlobberService.GetObjectTree:input_type -> blobber.GetObjectTreeRequest
	7,  // 7: blobber.BlobberService.DownloadFile:input_type -> blobber.DownloadFileRequest
	8,  // 8: blobber.BlobberService.RenameObject:input_type -> blobber.RenameObjectRequest
	9,  // 9: blobber.BlobberService.UploadFile:input_type -> blobber.UploadFileRequest
	10, // 10: blobber.BlobberService.Commit:input_type -> blobber.CommitRequest
	11, // 11: blobber.BlobberService.CalculateHash:input_type -> blobber.CalculateHashRequest
	12, // 12: blobber.BlobberService.CommitMetaTxn:input_type -> blobber.CommitMetaTxnRequest
	13, // 13: blobber.BlobberService.UpdateObjectAttributes:input_type -> blobber.UpdateObjectAttributesRequest
	14, // 14: blobber.BlobberService.CopyObject:input_type -> blobber.CopyObjectRequest
	15, // 15: blobber.BlobberService.Collaborator:input_type -> blobber.CollaboratorRequest
	16, // 16: blobber.BlobberService.MarketplaceShareInfo:input_type -> blobber.MarketplaceShareInfoRequest
	17, // 17: blobber.BlobberService.GetAllocation:output_type -> blobber.GetAllocationResponse
	18, // 18: blobber.BlobberService.GetFileMetaData:output_type -> blobber.GetFileMetaDataResponse
	19, // 19: blobber.BlobberService.GetFileStats:output_type -> blobber.GetFileStatsResponse
	20, // 20: blobber.BlobberService.ListEntities:output_type -> blobber.ListEntitiesResponse
	21, // 21: blobber.BlobberService.GetObjectPath:output_type -> blobber.GetObjectPathResponse
	22, // 22: blobber.BlobberService.GetReferencePath:output_type -> blobber.GetReferencePathResponse
	23, // 23: blobber.BlobberService.GetObjectTree:output_type -> blobber.GetObjectTreeResponse
	24, // 24: blobber.BlobberService.DownloadFile:output_type -> blobber.DownloadFileResponse
	25, // 25: blobber.BlobberService.RenameObject:output_type -> blobber.RenameObjectResponse
	26, // 26: blobber.BlobberService.UploadFile:output_type -> blobber.UploadFileResponse
	27, // 27: blobber.BlobberService.Commit:output_type -> blobber.CommitResponse
	28, // 28: blobber.BlobberService.CalculateHash:output_type -> blobber.CalculateHashResponse
	29, // 29: blobber.BlobberService.CommitMetaTxn:output_type -> blobber.CommitMetaTxnResponse
	30, // 30: blobber.BlobberService.UpdateObjectAttributes:output_type -> blobber.UpdateObjectAttributesResponse
	31, // 31: blobber.BlobberService.CopyObject:output_type -> blobber.CopyObjectResponse
	32, // 32: blobber.BlobberService.Collaborator:output_type -> blobber.CollaboratorResponse
	33, // 33: blobber.BlobberService.MarketplaceShareInfo:output_type -> blobber.MarketplaceShareInfoResponse
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_blobber_service_proto_init() }
func file_blobber_service_proto_init() {
	if File_blobber_service_proto != nil {
		return
	}
	file_blobber_contract_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_blobber_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blobber_service_proto_goTypes,
		DependencyIndexes: file_blobber_service_proto_depIdxs,
	}.Build()
	File_blobber_service_proto = out.File
	file_blobber_service_proto_rawDesc = nil
	file_blobber_service_proto_goTypes = nil
	file_blobber_service_proto_depIdxs = nil
}
