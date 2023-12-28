package contracts

import (
	"github.com/mailhedgehog/smtpMessage"
)

// Room name. Each room contains set of emails. By room application diverge what emails should be displayed for login
// user. Now application expects have room name same as logged username (in case auth disabled - will be used name
// "default").
type Room string

// SearchQuery represents key->value map what describe search params.
type SearchQuery = map[string]string

// SearchParam represents search key.
type SearchParam = string

const (
	SearchParamFrom    SearchParam = "from"
	SearchParamTo                  = "to"
	SearchParamContent             = "content"
)

type MessagesStorageConfiguration struct {
	PerRoomLimit int `yaml:"per_room_limit"`
}

// MessagesStorage interface represents a backend flow to store or retrieve messages
type MessagesStorage interface {
	// RoomsRepo returns room repository
	RoomsRepo() RoomsRepo

	// MessagesRepo returns messages repository related to specific room
	MessagesRepo(room Room) MessagesRepo
}

// RoomsRepo represents repository to manipulate rooms.
type RoomsRepo interface {
	// List of rooms in system.
	List(offset, limit int) ([]Room, error)
	// Count total count rooms in storage.
	Count() int
	// Delete all messages in room from storage.
	Delete(room Room) error
}

// MessagesRepo represents repository to manipulate messages related to specific room.
type MessagesRepo interface {
	// Store `message` to specific `room`.
	Store(message *smtpMessage.SmtpMessage) (smtpMessage.MessageID, error)
	// List retrieve list of messages based on `query` starting with `offset` index and count limited by `limit`.
	// `query` - represents of key->value map, where key is search parameter.
	List(query SearchQuery, offset, limit int) ([]smtpMessage.SmtpMessage, int, error)
	// Count total messages in storage.
	Count() int
	// Delete delete specific message from storage by `messageId`.
	Delete(messageId smtpMessage.MessageID) error
	// Load find specific message from storage by `messageId`.
	Load(messageId smtpMessage.MessageID) (*smtpMessage.SmtpMessage, error)
}
