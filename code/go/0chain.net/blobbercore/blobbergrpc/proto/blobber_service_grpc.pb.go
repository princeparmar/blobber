// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package blobbergrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BlobberServiceClient is the client API for BlobberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlobberServiceClient interface {
	GetAllocation(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*GetAllocationResponse, error)
	GetFileMetaData(ctx context.Context, in *GetFileMetaDataRequest, opts ...grpc.CallOption) (*GetFileMetaDataResponse, error)
	GetFileStats(ctx context.Context, in *GetFileStatsRequest, opts ...grpc.CallOption) (*GetFileStatsResponse, error)
	ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error)
	GetObjectPath(ctx context.Context, in *GetObjectPathRequest, opts ...grpc.CallOption) (*GetObjectPathResponse, error)
	GetReferencePath(ctx context.Context, in *GetReferencePathRequest, opts ...grpc.CallOption) (*GetReferencePathResponse, error)
	GetObjectTree(ctx context.Context, in *GetObjectTreeRequest, opts ...grpc.CallOption) (*GetObjectTreeResponse, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error)
	RenameObject(ctx context.Context, in *RenameObjectRequest, opts ...grpc.CallOption) (*RenameObjectResponse, error)
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error)
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error)
	CalculateHash(ctx context.Context, in *CalculateHashRequest, opts ...grpc.CallOption) (*CalculateHashResponse, error)
	CommitMetaTxn(ctx context.Context, in *CommitMetaTxnRequest, opts ...grpc.CallOption) (*CommitMetaTxnResponse, error)
	UpdateObjectAttributes(ctx context.Context, in *UpdateObjectAttributesRequest, opts ...grpc.CallOption) (*UpdateObjectAttributesResponse, error)
	CopyObject(ctx context.Context, in *CopyObjectRequest, opts ...grpc.CallOption) (*CopyObjectResponse, error)
	Collaborator(ctx context.Context, in *CollaboratorRequest, opts ...grpc.CallOption) (*CollaboratorResponse, error)
	MarketplaceShareInfo(ctx context.Context, in *MarketplaceShareRequest, opts ...grpc.CallOption) (*MarketplaceShareResponse, error)
}

type blobberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlobberServiceClient(cc grpc.ClientConnInterface) BlobberServiceClient {
	return &blobberServiceClient{cc}
}

func (c *blobberServiceClient) GetAllocation(ctx context.Context, in *GetAllocationRequest, opts ...grpc.CallOption) (*GetAllocationResponse, error) {
	out := new(GetAllocationResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/GetAllocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) GetFileMetaData(ctx context.Context, in *GetFileMetaDataRequest, opts ...grpc.CallOption) (*GetFileMetaDataResponse, error) {
	out := new(GetFileMetaDataResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/GetFileMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) GetFileStats(ctx context.Context, in *GetFileStatsRequest, opts ...grpc.CallOption) (*GetFileStatsResponse, error) {
	out := new(GetFileStatsResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/GetFileStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) ListEntities(ctx context.Context, in *ListEntitiesRequest, opts ...grpc.CallOption) (*ListEntitiesResponse, error) {
	out := new(ListEntitiesResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/ListEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) GetObjectPath(ctx context.Context, in *GetObjectPathRequest, opts ...grpc.CallOption) (*GetObjectPathResponse, error) {
	out := new(GetObjectPathResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/GetObjectPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) GetReferencePath(ctx context.Context, in *GetReferencePathRequest, opts ...grpc.CallOption) (*GetReferencePathResponse, error) {
	out := new(GetReferencePathResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/GetReferencePath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) GetObjectTree(ctx context.Context, in *GetObjectTreeRequest, opts ...grpc.CallOption) (*GetObjectTreeResponse, error) {
	out := new(GetObjectTreeResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/GetObjectTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	out := new(DownloadFileResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/DownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) RenameObject(ctx context.Context, in *RenameObjectRequest, opts ...grpc.CallOption) (*RenameObjectResponse, error) {
	out := new(RenameObjectResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/RenameObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileResponse, error) {
	out := new(UploadFileResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) CalculateHash(ctx context.Context, in *CalculateHashRequest, opts ...grpc.CallOption) (*CalculateHashResponse, error) {
	out := new(CalculateHashResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/CalculateHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) CommitMetaTxn(ctx context.Context, in *CommitMetaTxnRequest, opts ...grpc.CallOption) (*CommitMetaTxnResponse, error) {
	out := new(CommitMetaTxnResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/CommitMetaTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) UpdateObjectAttributes(ctx context.Context, in *UpdateObjectAttributesRequest, opts ...grpc.CallOption) (*UpdateObjectAttributesResponse, error) {
	out := new(UpdateObjectAttributesResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/UpdateObjectAttributes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) CopyObject(ctx context.Context, in *CopyObjectRequest, opts ...grpc.CallOption) (*CopyObjectResponse, error) {
	out := new(CopyObjectResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/CopyObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) Collaborator(ctx context.Context, in *CollaboratorRequest, opts ...grpc.CallOption) (*CollaboratorResponse, error) {
	out := new(CollaboratorResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/Collaborator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobberServiceClient) MarketplaceShareInfo(ctx context.Context, in *MarketplaceShareRequest, opts ...grpc.CallOption) (*MarketplaceShareResponse, error) {
	out := new(MarketplaceShareResponse)
	err := c.cc.Invoke(ctx, "/blobber.BlobberService/MarketplaceShareInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlobberServiceServer is the server API for BlobberService service.
// All implementations should embed UnimplementedBlobberServiceServer
// for forward compatibility
type BlobberServiceServer interface {
	GetAllocation(context.Context, *GetAllocationRequest) (*GetAllocationResponse, error)
	GetFileMetaData(context.Context, *GetFileMetaDataRequest) (*GetFileMetaDataResponse, error)
	GetFileStats(context.Context, *GetFileStatsRequest) (*GetFileStatsResponse, error)
	ListEntities(context.Context, *ListEntitiesRequest) (*ListEntitiesResponse, error)
	GetObjectPath(context.Context, *GetObjectPathRequest) (*GetObjectPathResponse, error)
	GetReferencePath(context.Context, *GetReferencePathRequest) (*GetReferencePathResponse, error)
	GetObjectTree(context.Context, *GetObjectTreeRequest) (*GetObjectTreeResponse, error)
	DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error)
	RenameObject(context.Context, *RenameObjectRequest) (*RenameObjectResponse, error)
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error)
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
	CalculateHash(context.Context, *CalculateHashRequest) (*CalculateHashResponse, error)
	CommitMetaTxn(context.Context, *CommitMetaTxnRequest) (*CommitMetaTxnResponse, error)
	UpdateObjectAttributes(context.Context, *UpdateObjectAttributesRequest) (*UpdateObjectAttributesResponse, error)
	CopyObject(context.Context, *CopyObjectRequest) (*CopyObjectResponse, error)
	Collaborator(context.Context, *CollaboratorRequest) (*CollaboratorResponse, error)
	MarketplaceShareInfo(context.Context, *MarketplaceShareRequest) (*MarketplaceShareResponse, error)
}

// UnimplementedBlobberServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBlobberServiceServer struct {
}

func (UnimplementedBlobberServiceServer) GetAllocation(context.Context, *GetAllocationRequest) (*GetAllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllocation not implemented")
}
func (UnimplementedBlobberServiceServer) GetFileMetaData(context.Context, *GetFileMetaDataRequest) (*GetFileMetaDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileMetaData not implemented")
}
func (UnimplementedBlobberServiceServer) GetFileStats(context.Context, *GetFileStatsRequest) (*GetFileStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileStats not implemented")
}
func (UnimplementedBlobberServiceServer) ListEntities(context.Context, *ListEntitiesRequest) (*ListEntitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEntities not implemented")
}
func (UnimplementedBlobberServiceServer) GetObjectPath(context.Context, *GetObjectPathRequest) (*GetObjectPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectPath not implemented")
}
func (UnimplementedBlobberServiceServer) GetReferencePath(context.Context, *GetReferencePathRequest) (*GetReferencePathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReferencePath not implemented")
}
func (UnimplementedBlobberServiceServer) GetObjectTree(context.Context, *GetObjectTreeRequest) (*GetObjectTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectTree not implemented")
}
func (UnimplementedBlobberServiceServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedBlobberServiceServer) RenameObject(context.Context, *RenameObjectRequest) (*RenameObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameObject not implemented")
}
func (UnimplementedBlobberServiceServer) UploadFile(context.Context, *UploadFileRequest) (*UploadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedBlobberServiceServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedBlobberServiceServer) CalculateHash(context.Context, *CalculateHashRequest) (*CalculateHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateHash not implemented")
}
func (UnimplementedBlobberServiceServer) CommitMetaTxn(context.Context, *CommitMetaTxnRequest) (*CommitMetaTxnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitMetaTxn not implemented")
}
func (UnimplementedBlobberServiceServer) UpdateObjectAttributes(context.Context, *UpdateObjectAttributesRequest) (*UpdateObjectAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObjectAttributes not implemented")
}
func (UnimplementedBlobberServiceServer) CopyObject(context.Context, *CopyObjectRequest) (*CopyObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyObject not implemented")
}
func (UnimplementedBlobberServiceServer) Collaborator(context.Context, *CollaboratorRequest) (*CollaboratorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collaborator not implemented")
}
func (UnimplementedBlobberServiceServer) MarketplaceShareInfo(context.Context, *MarketplaceShareRequest) (*MarketplaceShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketplaceShareInfo not implemented")
}

// UnsafeBlobberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlobberServiceServer will
// result in compilation errors.
type UnsafeBlobberServiceServer interface {
	mustEmbedUnimplementedBlobberServiceServer()
}

func RegisterBlobberServiceServer(s *grpc.Server, srv BlobberServiceServer) {
	s.RegisterService(&_BlobberService_serviceDesc, srv)
}

func _BlobberService_GetAllocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).GetAllocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/GetAllocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).GetAllocation(ctx, req.(*GetAllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_GetFileMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileMetaDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).GetFileMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/GetFileMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).GetFileMetaData(ctx, req.(*GetFileMetaDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_GetFileStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).GetFileStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/GetFileStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).GetFileStats(ctx, req.(*GetFileStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_ListEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).ListEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/ListEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).ListEntities(ctx, req.(*ListEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_GetObjectPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).GetObjectPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/GetObjectPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).GetObjectPath(ctx, req.(*GetObjectPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_GetReferencePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReferencePathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).GetReferencePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/GetReferencePath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).GetReferencePath(ctx, req.(*GetReferencePathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_GetObjectTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).GetObjectTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/GetObjectTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).GetObjectTree(ctx, req.(*GetObjectTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/DownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).DownloadFile(ctx, req.(*DownloadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_RenameObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).RenameObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/RenameObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).RenameObject(ctx, req.(*RenameObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_CalculateHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).CalculateHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/CalculateHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).CalculateHash(ctx, req.(*CalculateHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_CommitMetaTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitMetaTxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).CommitMetaTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/CommitMetaTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).CommitMetaTxn(ctx, req.(*CommitMetaTxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_UpdateObjectAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObjectAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).UpdateObjectAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/UpdateObjectAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).UpdateObjectAttributes(ctx, req.(*UpdateObjectAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_CopyObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).CopyObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/CopyObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).CopyObject(ctx, req.(*CopyObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_Collaborator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollaboratorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).Collaborator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/Collaborator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).Collaborator(ctx, req.(*CollaboratorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobberService_MarketplaceShareInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketplaceShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobberServiceServer).MarketplaceShareInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blobber.BlobberService/MarketplaceShareInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobberServiceServer).MarketplaceShareInfo(ctx, req.(*MarketplaceShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlobberService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blobber.BlobberService",
	HandlerType: (*BlobberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllocation",
			Handler:    _BlobberService_GetAllocation_Handler,
		},
		{
			MethodName: "GetFileMetaData",
			Handler:    _BlobberService_GetFileMetaData_Handler,
		},
		{
			MethodName: "GetFileStats",
			Handler:    _BlobberService_GetFileStats_Handler,
		},
		{
			MethodName: "ListEntities",
			Handler:    _BlobberService_ListEntities_Handler,
		},
		{
			MethodName: "GetObjectPath",
			Handler:    _BlobberService_GetObjectPath_Handler,
		},
		{
			MethodName: "GetReferencePath",
			Handler:    _BlobberService_GetReferencePath_Handler,
		},
		{
			MethodName: "GetObjectTree",
			Handler:    _BlobberService_GetObjectTree_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _BlobberService_DownloadFile_Handler,
		},
		{
			MethodName: "RenameObject",
			Handler:    _BlobberService_RenameObject_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _BlobberService_UploadFile_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _BlobberService_Commit_Handler,
		},
		{
			MethodName: "CalculateHash",
			Handler:    _BlobberService_CalculateHash_Handler,
		},
		{
			MethodName: "CommitMetaTxn",
			Handler:    _BlobberService_CommitMetaTxn_Handler,
		},
		{
			MethodName: "UpdateObjectAttributes",
			Handler:    _BlobberService_UpdateObjectAttributes_Handler,
		},
		{
			MethodName: "CopyObject",
			Handler:    _BlobberService_CopyObject_Handler,
		},
		{
			MethodName: "Collaborator",
			Handler:    _BlobberService_Collaborator_Handler,
		},
		{
			MethodName: "MarketplaceShareInfo",
			Handler:    _BlobberService_MarketplaceShareInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blobber_service.proto",
}
