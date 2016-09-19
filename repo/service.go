package repo

import (
	"golang.org/x/net/context"

	pb "github.com/philips/kpr-server/kprpb"
)

type RepoService struct{}

func (m *RepoService) List(c context.Context, s *pb.ListRequest) (*pb.ListResponse, error) {
	println("list")
	repos := []*pb.ListResponse_Repos{
		{Name: "quay.io/foo"},
	}
	return &pb.ListResponse{repos}, nil
}

func (m *RepoService) ListTags(c context.Context, s *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	println("listTags")
	tags := []*pb.ListTagsResponse_Tags{
		{Name: "v1.0.0"},
		{Name: "v1.0.1"},
	}
	return &pb.ListTagsResponse{Name: "quay.io/foo/bar", Tags: tags}, nil
}

func NewServer() *RepoService {
	return new(RepoService)
}
