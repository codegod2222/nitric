package services

import (
	"context"

	pb "github.com/nitric-dev/membrane/interfaces/nitric/v1"
	"github.com/nitric-dev/membrane/plugins/sdk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPC Interface for registered Nitric Storage Plugins
type StorageServer struct {
	pb.UnimplementedStorageServer
	storagePlugin sdk.StoragePlugin
}

// Checks that the storage server is registered and returns gRPC Unimplemented error if not.
func (s *StorageServer) checkPluginRegistered() (bool, error) {
	if s.storagePlugin == nil {
		return false, status.Errorf(codes.Unimplemented, "Storage plugin not registered")
	}

	return true, nil
}

func (s *StorageServer) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutReply, error) {
	if ok, err := s.checkPluginRegistered(); ok {
		if err := s.storagePlugin.Put(req.GetBucketName(), req.GetKey(), req.GetBody()); err == nil {
			return &pb.PutReply{
				Success: true,
			}, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (s *StorageServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	if ok, err := s.checkPluginRegistered(); ok {
		if object, err := s.storagePlugin.Get(req.GetBucketName(), req.GetKey()); err == nil {
			return &pb.GetReply{
				Body: object,
			}, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (s *StorageServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	if ok, err := s.checkPluginRegistered(); ok {
		if err := s.storagePlugin.Delete(req.GetBucketName(), req.GetKey()); err == nil {
			return &pb.DeleteReply{
				Success: true,
			}, nil
		} else {
			// TODO: handle specific error codes.
			return nil, err
		}
	} else {
		return nil, err
	}
}

func NewStorageServer(storagePlugin sdk.StoragePlugin) pb.StorageServer {
	return &StorageServer{
		storagePlugin: storagePlugin,
	}
}
