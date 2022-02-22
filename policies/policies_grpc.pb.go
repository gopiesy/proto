// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: policies/policies.proto

package policles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PolicyServiceClient is the client API for PolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolicyServiceClient interface {
	// StreamSnapshots implemented by GT to continuesly stream snapshots and get
	// the status
	StreamSnapshots(ctx context.Context, opts ...grpc.CallOption) (PolicyService_StreamSnapshotsClient, error)
}

type policyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolicyServiceClient(cc grpc.ClientConnInterface) PolicyServiceClient {
	return &policyServiceClient{cc}
}

func (c *policyServiceClient) StreamSnapshots(ctx context.Context, opts ...grpc.CallOption) (PolicyService_StreamSnapshotsClient, error) {
	stream, err := c.cc.NewStream(ctx, &PolicyService_ServiceDesc.Streams[0], "/policies.PolicyService/StreamSnapshots", opts...)
	if err != nil {
		return nil, err
	}
	x := &policyServiceStreamSnapshotsClient{stream}
	return x, nil
}

type PolicyService_StreamSnapshotsClient interface {
	Send(*SnapshotStatus) error
	Recv() (*Snapshot, error)
	grpc.ClientStream
}

type policyServiceStreamSnapshotsClient struct {
	grpc.ClientStream
}

func (x *policyServiceStreamSnapshotsClient) Send(m *SnapshotStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *policyServiceStreamSnapshotsClient) Recv() (*Snapshot, error) {
	m := new(Snapshot)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PolicyServiceServer is the server API for PolicyService service.
// All implementations must embed UnimplementedPolicyServiceServer
// for forward compatibility
type PolicyServiceServer interface {
	// StreamSnapshots implemented by GT to continuesly stream snapshots and get
	// the status
	StreamSnapshots(PolicyService_StreamSnapshotsServer) error
	mustEmbedUnimplementedPolicyServiceServer()
}

// UnimplementedPolicyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolicyServiceServer struct {
}

func (UnimplementedPolicyServiceServer) StreamSnapshots(PolicyService_StreamSnapshotsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamSnapshots not implemented")
}
func (UnimplementedPolicyServiceServer) mustEmbedUnimplementedPolicyServiceServer() {}

// UnsafePolicyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolicyServiceServer will
// result in compilation errors.
type UnsafePolicyServiceServer interface {
	mustEmbedUnimplementedPolicyServiceServer()
}

func RegisterPolicyServiceServer(s grpc.ServiceRegistrar, srv PolicyServiceServer) {
	s.RegisterService(&PolicyService_ServiceDesc, srv)
}

func _PolicyService_StreamSnapshots_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PolicyServiceServer).StreamSnapshots(&policyServiceStreamSnapshotsServer{stream})
}

type PolicyService_StreamSnapshotsServer interface {
	Send(*Snapshot) error
	Recv() (*SnapshotStatus, error)
	grpc.ServerStream
}

type policyServiceStreamSnapshotsServer struct {
	grpc.ServerStream
}

func (x *policyServiceStreamSnapshotsServer) Send(m *Snapshot) error {
	return x.ServerStream.SendMsg(m)
}

func (x *policyServiceStreamSnapshotsServer) Recv() (*SnapshotStatus, error) {
	m := new(SnapshotStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PolicyService_ServiceDesc is the grpc.ServiceDesc for PolicyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolicyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "policies.PolicyService",
	HandlerType: (*PolicyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSnapshots",
			Handler:       _PolicyService_StreamSnapshots_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "policies/policies.proto",
}
