package repo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/net/context"

	pb "github.com/philips/kpr-server/kprpb"
)

type RepoService struct {
	BaseDir string
}

func (m *RepoService) List(c context.Context, s *pb.ListRequest) (*pb.ListResponse, error) {
	orgs, err := ioutil.ReadDir(m.BaseDir)
	if err != nil {
		return nil, err
	}
	repoList := []*pb.ListResponse_Repos{}
	for _, org := range orgs {
		repos, err := ioutil.ReadDir(filepath.Join(m.BaseDir, org.Name()))
		if err != nil {
			log.Fatal(err)
		}
		for _, repo := range repos {
			repoList = append(repoList, &pb.ListResponse_Repos{Name: filepath.Join(org.Name(), repo.Name())})
		}
	}

	return &pb.ListResponse{repoList}, nil
}

func (m *RepoService) ListTags(c context.Context, s *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	tags, err := ioutil.ReadDir(filepath.Join(m.BaseDir, s.Name, "tags"))
	if err != nil {
		return nil, err
	}
	tagsList := []*pb.ListTagsResponse_Tags{}
	for _, tag := range tags {
		tagsList = append(tagsList, &pb.ListTagsResponse_Tags{Name: filepath.Join(tag.Name())})
	}

	return &pb.ListTagsResponse{Name: s.Name, Tags: tagsList}, nil
}

func (m *RepoService) PutTag(c context.Context, s *pb.TagRequest) (*pb.Descriptor, error) {
	tagDir := filepath.Join(m.BaseDir, s.Name, "tags")
	err := os.MkdirAll(tagDir, 0755)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	d := s.GetDesc()
	b, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	ioutil.WriteFile(filepath.Join(tagDir, s.Tag), b, 0644)
	return d, nil
}

func (m *RepoService) GetTag(c context.Context, s *pb.TagRequest) (*pb.PackageManifest, error) {
	tagFile := filepath.Join(m.BaseDir, s.Name, "tags", s.Tag)
	b, err := ioutil.ReadFile(tagFile)
	if err != nil {
		return nil, err
	}

	d := pb.Descriptor{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}

	// TODO get the blob
	return &pb.PackageManifest{}, nil
}

func (m *RepoService) DeleteTag(c context.Context, s *pb.TagRequest) (*pb.Descriptor, error) {
	tagFile := filepath.Join(m.BaseDir, s.Name, "tags", s.Tag)
	b, err := ioutil.ReadFile(tagFile)
	if err != nil {
		return nil, err
	}

	d := pb.Descriptor{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return nil, err
	}

	os.Remove(tagFile)
	return &d, nil
}

func NewServer(baseDir string) *RepoService {
	return &RepoService{filepath.Join(baseDir, "repos")}
}
