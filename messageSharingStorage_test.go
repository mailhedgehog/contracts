package contracts

import (
	"github.com/mailhedgehog/gounit"
	"testing"
	"time"
)

func TestNewSharedMessageRecord(t *testing.T) {
	record := NewSharedMessageRecord("foo", "bar")

	(*gounit.T)(t).AssertFalse(record.IsExpired())

	(*gounit.T)(t).AssertEqualsString("foo", string(record.Room))
	(*gounit.T)(t).AssertEqualsString("bar", string(record.MessageId))

	(*gounit.T)(t).AssertTrue(time.Now().Before(record.ExpiredAt))
	(*gounit.T)(t).AssertTrue(time.Now().Add(time.Duration(66) * time.Minute).After(record.ExpiredAt))

	record.SetExpirationInHours(2)
	(*gounit.T)(t).AssertFalse(time.Now().Add(time.Duration(66) * time.Minute).After(record.ExpiredAt))
	(*gounit.T)(t).AssertTrue(time.Now().Add(time.Duration(121) * time.Minute).After(record.ExpiredAt))

	(*gounit.T)(t).AssertEqualsString("", string(record.Id))
	(*gounit.T)(t).AssertFalse(record.Exists())

	record.Id = "test"
	(*gounit.T)(t).AssertTrue(record.Exists())

}

func TestSharedMessageRecord_IsExpired(t *testing.T) {
	record := NewSharedMessageRecord("foo", "bar")

	(*gounit.T)(t).AssertFalse(record.IsExpired())

	record.SetExpirationInHours(2)
	(*gounit.T)(t).AssertFalse(record.IsExpired())

	record.SetExpirationInHours(-2)
	(*gounit.T)(t).AssertTrue(record.IsExpired())

}
