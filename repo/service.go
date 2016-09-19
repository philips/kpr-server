package repo

import (
	"errors"

	"golang.org/x/net/context"

	pb "github.com/philips/kpr-server/kprpb"
)

type RepoService struct{}

func (m *RepoService) List(c context.Context, s *pb.ListRequest) (*pb.ListResponse, error) {
	println("list")
	repos := []*pb.ListResponse_Repos{
		{Name: "quay.io/foo/bar"},
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

func (m *RepoService) PutTag(c context.Context, s *pb.TagRequest) (*pb.Descriptor, error) {
	println("putTag")
	d := s.GetDesc()
	println(d.MediaType)
	return d, nil
}

func (m *RepoService) GetTag(c context.Context, s *pb.TagRequest) (*pb.PackageManifest, error) {
	println("getTag")
	switch s.Tag {
	case "v1.0.0":
		return &pb.PackageManifest{
			SchemaVersion: 1,
			MediaType:     "application/vnd.kpr.image.manifest.v1+json",
			Package: &pb.Descriptor{
				MediaType: "application/vnd.helm.package.tar+gzip",
				Size:      12,
				Digest:    "sha256:cafecafecafe",
			},
		}, nil
	case "v1.0.1":
		return &pb.PackageManifest{
			SchemaVersion: 1,
			MediaType:     "application/vnd.kpr.image.manifest.list.v1+json",
			Packages: []*pb.Descriptor{
				{
					MediaType: "application/vnd.helm.package.tar+gzip",
					Size:      12,
					Digest:    "sha256:cafecafecafe",
				},
			},
		}, nil
	}
	return nil, errors.New("no version")
}

func (m *RepoService) DeleteTag(c context.Context, s *pb.TagRequest) (*pb.Descriptor, error) {
	println("deleteTag")
	d := s.GetDesc()
	return d, nil
}

func NewServer() *RepoService {
	return new(RepoService)
}
