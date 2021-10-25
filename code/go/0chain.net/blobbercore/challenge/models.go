package challenge

import (
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/models"
	"github.com/0chain/blobber/code/go/0chain.net/core/common"
	"github.com/0chain/gosdk/core/util"
)

// BlobberChallenge challenge object sent from sharders
type BlobberChallenge struct {
	BlobberID  string              `json:"blobber_id"`
	Challenges []*StorageChallenge `json:"challenges"`
}

// StorageChallenge challenge smart contract on blockchain
type StorageChallenge struct {
	Created        common.Timestamp  `json:"created"`
	ID             string            `json:"id"`
	PrevID         string            `json:"prev_id"`
	Validators     []*ValidationNode `json:"validators"`
	RandomNumber   int64             `json:"seed"`
	AllocationID   string            `json:"allocation_id"`
	AllocationRoot string            `json:"allocation_root"`
}

// ChallengeTask challenge task that is selected based on StorageChallenge's random seed
type ChallengeTask struct {
	ChallengeID  string                         `json:"challenge_id,omitempty"`
	RootHash     string                         `json:"root_hash,omitempty"`
	Storages     [][]*models.ChallengeStorage   `json:"storages,omitempty"`
	WriteMarkers []*models.ChallengeWriteMarker `json:"write_marker,omitempty"`

	SelectedFile       *models.ChallengeStorage `json:"selected_file,omitempty"`
	SelectedBlockIndex int64                    `json:"selected_block_index,omitempty"`
	SelectedBlockBytes []byte                   `json:"selected_block_bytes,omitempty"`
	SelectedMerklePath *util.MTPath             `json:"selected_merkle_path,omitempty"`
}
