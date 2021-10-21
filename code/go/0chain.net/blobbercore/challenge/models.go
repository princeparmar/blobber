package challenge

import "github.com/0chain/blobber/code/go/0chain.net/core/common"

type BlobberChallenge struct {
	BlobberID  string              `json:"blobber_id"`
	Challenges []*StorageChallenge `json:"challenges"`
}

type StorageChallenge struct {
	Created        common.Timestamp  `json:"created"`
	ID             string            `json:"id"`
	PrevID         string            `json:"prev_id"`
	Validators     []*ValidationNode `json:"validators"`
	RandomNumber   int64             `json:"seed"`
	AllocationID   string            `json:"allocation_id"`
	AllocationRoot string            `json:"allocation_root"`
}
