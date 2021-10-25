package models

import "github.com/0chain/blobber/code/go/0chain.net/core/common"

// ChallengeWriteMarker writemarker for challenge
type ChallengeWriteMarker struct {
	AllocationRoot         string           `gorm:"column:allocation_root;primary_key" json:"allocation_root,omitempty"`
	PreviousAllocationRoot string           `gorm:"column:prev_allocation_root" json:"prev_allocation_root,omitempty"`
	AllocationID           string           `gorm:"column:allocation_id" json:"allocation_id,omitempty"`
	Size                   int64            `gorm:"column:size" json:"size,omitempty"`
	BlobberID              string           `gorm:"column:blobber_id" json:"blobber_id,omitempty"`
	Timestamp              common.Timestamp `gorm:"column:timestamp" json:"timestamp,omitempty"`
	ClientID               string           `gorm:"column:client_id" json:"client_id,omitempty"`
	ClientPublicKey        string           `gorm:"column:client_key" json:"client_public_key,omitempty"`
	Signature              string           `gorm:"column:signature" json:"signature,omitempty"`
}

const (
	TableNameWriteMarker = "write_markers"
)

func (ChallengeWriteMarker) TableName() string {
	return TableNameWriteMarker
}
