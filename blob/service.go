package blob

import (
	"golang.org/x/net/context"

	pb "github.com/philips/kpr-server/kprpb"
)

type BlobService struct {
	BaseDir string
}

func (m *BlobService) PutBlob(c context.Context, s *pb.BlobRequest) (*pb.Descriptor, error) {
	println("putBlob")
	return &pb.Descriptor{}, nil
}

func (m *BlobService) GetBlob(c context.Context, s *pb.BlobRequest) (*pb.BlobResponse, error) {
	println("getBlob")
	return &pb.BlobResponse{[]byte("hello")}, nil
}

func NewServer(baseDir string) *BlobService {
	return &BlobService{baseDir}
}
