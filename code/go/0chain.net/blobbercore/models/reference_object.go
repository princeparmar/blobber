package models

import (
	"gorm.io/datatypes"
)

const TableNameReferenceObject = "reference_objects"

type ReferenceObject struct {
	ID                  int64          `gorm:"column:id;primary_key"`
	Type                string         `gorm:"column:type"`
	AllocationID        string         `gorm:"column:allocation_id"`
	LookupHash          string         `gorm:"column:lookup_hash"`
	Name                string         `gorm:"column:name"`
	Path                string         `gorm:"column:path"`
	Hash                string         `gorm:"column:hash"`
	NumBlocks           int64          `gorm:"column:num_of_blocks"`
	PathHash            string         `gorm:"column:path_hash"`
	ParentPath          string         `gorm:"column:parent_path"`
	PathLevel           int            `gorm:"column:level"`
	CustomMeta          string         `gorm:"column:custom_meta"`
	ContentHash         string         `gorm:"column:content_hash"`
	Size                int64          `gorm:"column:size"`
	MerkleRoot          string         `gorm:"column:merkle_root"`
	ActualFileSize      int64          `gorm:"column:actual_file_size"`
	ActualFileHash      string         `gorm:"column:actual_file_hash"`
	MimeType            string         `gorm:"column:mimetype"`
	WriteMarker         string         `gorm:"column:write_marker"`
	ThumbnailSize       int64          `gorm:"column:thumbnail_size"`
	ThumbnailHash       string         `gorm:"column:thumbnail_hash"`
	ActualThumbnailSize int64          `gorm:"column:actual_thumbnail_size"`
	ActualThumbnailHash string         `gorm:"column:actual_thumbnail_hash"`
	EncryptedKey        string         `gorm:"column:encrypted_key"`
	Attributes          datatypes.JSON `gorm:"column:attributes"`

	OnCloud bool `gorm:"column:on_cloud"`

	ChunkSize int64 `gorm:"column:chunk_size"`

	ModelWithTS
}

func (ReferenceObject) TableName() string {
	return TableNameReferenceObject
}

// ChallengeMeta basic information for challenge
type ChallengeMeta struct {
	Hash      string `gorm:"column:hash"`
	NumBlocks int64  `gorm:"column:num_of_blocks" `
}

// ChallengeStorage storage object for challenge
type ChallengeStorage struct {
	ID          int64  `gorm:"column:id;primary_key" json:"id,omitempty"`
	Name        string `gorm:"column:name" json:"name,omitempty"`
	Size        int64  `gorm:"column:size" json:"size,omitempty"`
	Type        string `gorm:"column:type" json:"type,omitempty"`
	Path        string `gorm:"column:path" json:"path,omitempty"`
	PathLevel   int    `gorm:"column:level" json:"path_level,omitempty"`
	ParentPath  string `gorm:"column:parent_path" json:"parent_path,omitempty"`
	NumBlocks   int64  `gorm:"column:num_of_blocks" json:"num_blocks,omitempty"`
	ContentHash string `gorm:"column:content_hash" json:"content_hash,omitempty"`
	ChunkSize   int64  `gorm:"column:chunk_size" json:"chunk_size,omitempty"`
}

func (ChallengeStorage) TableName() string {
	return TableNameReferenceObject
}
