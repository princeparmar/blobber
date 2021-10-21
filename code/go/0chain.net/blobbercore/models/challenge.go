package models

import (
	"time"

	"gorm.io/datatypes"
)

type (
	ChallengeStatus int
	ChallengeResult int
)

const (
	ChallengeStatusAccepted ChallengeStatus = iota + 1
	ChallengeStatusProcessed
	ChallengeStatusCommitted
)

const (
	ChallengeResultUnknown ChallengeResult = iota
	ChallengeResultSuccess
	ChallengeResultFailure
)

const (
	TableNameChallenge = "challenges"
)

type Challenge struct {
	ChallengeID string `gorm:"column:challenge_id;primary_key"`

	AllocationID            string `gorm:"column:allocation_id"`
	AllocationRoot          string `gorm:"column:allocation_root"`
	RespondedAllocationRoot string `gorm:"column:responded_allocation_root"`

	Status        ChallengeStatus `gorm:"column:status"`
	StatusMessage string          `gorm:"column:status_message"`

	Result ChallengeResult `gorm:"column:result"`

	Seed     int64 `gorm:"column:seed"`
	BlockNum int64 `gorm:"column:block_num"`

	//Sequence int64 `gorm:"sequence"` // it is used https://www.postgresql.org/docs/9.4/functions-sequence.html. can't be inserted/updated by gorm

	CommitTxnID       string         `gorm:"column:commit_txn_id"`
	LastCommitTxnList datatypes.JSON `gorm:"column:last_commit_txn_ids"`

	ValidationTickets datatypes.JSON `gorm:"column:validation_tickets"`
	Validators        datatypes.JSON `gorm:"column:validators"`

	ObjectPath datatypes.JSON `gorm:"column:object_path"`

	RefID int64 `gorm:"column:ref_id"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Challenge) TableName() string {
	return TableNameChallenge
}
