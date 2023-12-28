package contracts

import (
	"github.com/mailhedgehog/smtpMessage"
	"time"
)

// MessageSharing represents interface of manipulation of shared messages
type MessageSharing interface {
	// Add new sharing record to storage.
	Add(emailSharingRecord *SharedMessageRecord) (*SharedMessageRecord, error)

	// Find single record by ID.
	Find(id string) (*SharedMessageRecord, error)

	// DeleteExpired will iterate over all shared records and delete all expired records.
	DeleteExpired() (bool, error)
}

// MessageSharingExpiredAtFormat consider always be in UTC
var MessageSharingExpiredAtFormat = "2006-01-02 15:04:05"

// SharedMessageRecord represents one record of shared message.
type SharedMessageRecord struct {
	Id        string
	Room      Room
	MessageId smtpMessage.MessageID
	ExpiredAt time.Time
}

// Exists checks is record exists in storage ()is created identification.
func (record *SharedMessageRecord) Exists() bool {
	return record.Id != ""
}

// IsExpired checks is record expired
func (record *SharedMessageRecord) IsExpired() bool {
	return time.Now().After(record.ExpiredAt)
}

// SetExpirationInHours setting expiration time future from now in passed hours.
func (record *SharedMessageRecord) SetExpirationInHours(hours int) *SharedMessageRecord {
	record.ExpiredAt = time.Now().UTC().Add(time.Duration(hours) * time.Hour)

	return record
}

// GetExpiredAtString string representation of expired at time.
func (record *SharedMessageRecord) GetExpiredAtString() string {
	return record.ExpiredAt.Format(MessageSharingExpiredAtFormat)
}

// NewSharedMessageRecord creates new object of SharedMessageRecord
func NewSharedMessageRecord(room Room, messageID smtpMessage.MessageID) *SharedMessageRecord {
	return &SharedMessageRecord{
		Room:      room,
		MessageId: messageID,
		ExpiredAt: time.Now().Add(time.Duration(1) * time.Hour),
	}
}
