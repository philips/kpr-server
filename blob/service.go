package blob

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/net/context"

	pb "github.com/philips/kpr-server/kprpb"
)

type BlobService struct {
	BaseDir string
}

func generateDigest(b []byte) string {
	digest := sha256.Sum256(b)
	return "sha256:" + hex.EncodeToString(digest[:])
}

func (m *BlobService) PutBlob(c context.Context, s *pb.BlobRequest) (*pb.Descriptor, error) {
	digestDir := filepath.Join(m.BaseDir, s.Name, "blobs", "sha256")
	err := os.MkdirAll(digestDir, 0755)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}

	blobFile := filepath.Join(m.BaseDir, s.Name, "blobs", "sha256", s.Digest)
	err = ioutil.WriteFile(blobFile, s.Blob, 0755)
	if err != nil {
		return nil, err
	}

	// TODO get the blob
	return &pb.Descriptor{Size: int64(len(s.Blob)), Digest: generateDigest(s.Blob)}, nil
}

func (m *BlobService) GetBlob(c context.Context, s *pb.BlobRequest) (*pb.BlobResponse, error) {
	blobFile := filepath.Join(m.BaseDir, s.Name, "blobs", "sha256", s.Digest)
	b, err := ioutil.ReadFile(blobFile)
	if err != nil {
		return nil, err
	}

	return &pb.BlobResponse{b}, nil
}

func NewServer(baseDir string) *BlobService {
	return &BlobService{filepath.Join(baseDir, "repos")}
}
